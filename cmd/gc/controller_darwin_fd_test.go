package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/gastownhall/gascity/internal/config"
)

// TestWatchConfigTargets_DarwinReloadDoesNotLeakFDs is the regression for
// gastownhall/gascity#4504 / fork ticket 11: on macOS/kqueue, fsnotify
// v1.9.0's Watcher.Close closes the kqueue but leaks watched REG/DIR
// descriptors. Restarting the recursive config watcher must not
// monotonically accumulate FDs.
//
// Measurement matches the upstream repro: unique numeric descriptors via
// `lsof -nP -p <pid> -F f`.
func TestWatchConfigTargets_DarwinReloadDoesNotLeakFDs(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("macOS/kqueue descriptor leak is Darwin-specific")
	}
	if _, err := exec.LookPath("lsof"); err != nil {
		t.Skip("lsof not installed")
	}

	root := t.TempDir()
	// Enough watched directories that a per-path leak is unambiguous across
	// a few restarts (one FD per watched dir on kqueue).
	const dirs = 80
	for i := 0; i < dirs; i++ {
		subdir := filepath.Join(root, "pack", "agents", "a"+strconv.Itoa(i), "overlay")
		if err := os.MkdirAll(subdir, 0o755); err != nil {
			t.Fatalf("MkdirAll: %v", err)
		}
		if err := os.WriteFile(filepath.Join(subdir, "settings.json"), []byte(`{}`), 0o644); err != nil {
			t.Fatalf("WriteFile: %v", err)
		}
	}

	targets := []config.WatchTarget{{Path: root, Recursive: true}}
	var dirty atomic.Bool
	pokeCh := make(chan struct{}, 1)
	var stderr bytes.Buffer

	baseline := countUniqueProcessFDs(t)
	const restarts = 5
	var afterFirst int
	for i := 0; i < restarts; i++ {
		cleanup := watchConfigTargets(targets, &dirty, pokeCh, &stderr)
		cleanup()
		n := countUniqueProcessFDs(t)
		if i == 0 {
			afterFirst = n
		}
	}
	final := countUniqueProcessFDs(t)

	// Allow a small absolute cushion for unrelated test/runtime FDs, but
	// reject the #4504 pattern: each Close leaving ~dirs descriptors behind.
	growth := final - afterFirst
	perRestart := growth / (restarts - 1)
	if perRestart > dirs/4 {
		t.Fatalf("config watcher restart leaked FDs on Darwin: baseline=%d afterFirst=%d final=%d growth=%d (~%d/restart over %d dirs); stderr=%q",
			baseline, afterFirst, final, growth, perRestart, dirs, stderr.String())
	}
	if final-baseline > dirs {
		t.Fatalf("config watcher FD count grew too far from baseline: baseline=%d final=%d delta=%d (dirs=%d); stderr=%q",
			baseline, final, final-baseline, dirs, stderr.String())
	}
}

func countUniqueProcessFDs(t *testing.T) int {
	t.Helper()
	out, err := exec.Command("lsof", "-nP", "-p", strconv.Itoa(os.Getpid()), "-F", "f").Output()
	if err != nil {
		t.Fatalf("lsof: %v", err)
	}
	seen := make(map[string]struct{})
	for _, line := range strings.Split(string(out), "\n") {
		if len(line) > 1 && line[0] == 'f' {
			seen[line[1:]] = struct{}{}
		}
	}
	return len(seen)
}
