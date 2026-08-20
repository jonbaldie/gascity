package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jonbaldie/gascity/internal/fsys"
)

// darwinPSResourceDeniedMessage is the host ps stderr observed on macOS when
// resource-usage columns require an entitlement the caller does not have
// (gastownhall/gascity#5201).
const darwinPSResourceDeniedMessage = "ps: %mem/vsz/rss/time: requires entitlement"

// installPSDenyingResourceFields puts a fake ps first on PATH that exits
// non-zero with the Darwin entitlement message when asked for rss=/vsz=/%mem=/time=,
// and otherwise prints stdoutLines. Darwin-specific entitlement is simulated;
// the tests never need a live Mac.
func installPSDenyingResourceFields(t *testing.T, stdoutLines string) {
	t.Helper()
	binDir := t.TempDir()
	script := `#!/bin/sh
for arg in "$@"; do
  case "$arg" in
    *rss=*|*%mem=*|*vsz=*|*time=*)
      echo 'ps: %mem/vsz/rss/time: requires entitlement' >&2
      exit 1
      ;;
  esac
done
`
	if stdoutLines != "" {
		script += "cat <<'EOF'\n" + stdoutLines
		if !strings.HasSuffix(stdoutLines, "\n") {
			script += "\n"
		}
		script += "EOF\n"
	}
	if err := os.WriteFile(filepath.Join(binDir, "ps"), []byte(script), 0o755); err != nil {
		t.Fatalf("WriteFile(ps): %v", err)
	}
	t.Setenv("PATH", binDir+string(os.PathListSeparator)+os.Getenv("PATH"))
}

func TestPSLStartCommandLines_ResourceFieldsDeniedDoesNotFailScan(t *testing.T) {
	installPSDenyingResourceFields(t, "  12345 Sun May 17 09:31:24 2026 /usr/bin/sleep 1\n")

	lines, err := psLStartCommandLines()
	if err != nil {
		t.Fatalf("psLStartCommandLines: %v (stderr must not be %q)", err, darwinPSResourceDeniedMessage)
	}
	if len(lines) != 1 {
		t.Fatalf("lines = %v, want 1 process line after rss= denial", lines)
	}
}

func TestDiscoverDoltProcessesFromPS_ResourceFieldsDeniedDoesNotFailScan(t *testing.T) {
	installPSDenyingResourceFields(t, "2222 Sun May 17 09:31:24 2026 dolt sql-server --config /tmp/TestMailRouter9182/config.yaml\n")

	procs, err := discoverDoltProcessesFromPS()
	if err != nil {
		t.Fatalf("discoverDoltProcessesFromPS: %v (rss= denial must not fail discovery)", err)
	}
	if len(procs) != 1 || procs[0].PID != 2222 {
		t.Fatalf("procs = %+v, want one dolt sql-server pid 2222", procs)
	}
	if procs[0].RSSBytes != 0 {
		t.Fatalf("RSSBytes = %d, want 0 (best-effort when rss= is unavailable)", procs[0].RSSBytes)
	}
	if procs[0].StartIdentity != "Sun May 17 09:31:24 2026" {
		t.Fatalf("StartIdentity = %q, want lstart without an rss column", procs[0].StartIdentity)
	}
}

func TestParseDoltPSLine_WithoutRSSStillClassifiesOrphan(t *testing.T) {
	line := "2222 Sun May 17 09:31:24 2026 dolt sql-server --config /tmp/TestMailRouter9182/config.yaml"
	got, ok := parseDoltPSLine(line, nil)
	if !ok {
		t.Fatal("parseDoltPSLine did not recognize dolt sql-server without rss column")
	}
	if got.RSSBytes != 0 {
		t.Fatalf("RSSBytes = %d, want 0 when ps omits rss", got.RSSBytes)
	}
	c := classifyDoltProcess(got, nil, "/home/u", "", nil)
	if c.Action != "reap" {
		t.Fatalf("Action = %q, want reap (orphan decision must not depend on RSS)", c.Action)
	}
}

func TestParseDoltPSLine_WithoutRSSStillProtectsNonTest(t *testing.T) {
	line := "1138290 Sun May 17 09:31:24 2026 dolt sql-server --config /var/lib/dolt/config.yaml"
	got, ok := parseDoltPSLine(line, nil)
	if !ok {
		t.Fatal("parseDoltPSLine did not recognize dolt sql-server without rss column")
	}
	c := classifyDoltProcess(got, nil, "/home/u", "", nil)
	if c.Action != "protect" {
		t.Fatalf("Action = %q, want protect (safety decision must not depend on RSS)", c.Action)
	}
}

func TestRunDoltCleanup_ReapScanSucceedsWhenPSDeniesRSS(t *testing.T) {
	installPSDenyingResourceFields(t, "1138290 Sun May 17 09:31:24 2026 dolt sql-server --config /var/lib/dolt/config.yaml\n")

	var stdout, stderr bytes.Buffer
	opts := cleanupOptions{
		FS:                fsys.NewFake(),
		JSON:              true,
		HomeDir:           "/home/u",
		DiscoverProcesses: discoverDoltProcessesFromPS,
		ActiveTestRoots:   []string{},
	}
	code := runDoltCleanup(opts, &stdout, &stderr)
	if code != 0 {
		t.Fatalf("exit=%d, stderr=%s", code, stderr.String())
	}
	var r CleanupReport
	if err := json.Unmarshal(stdout.Bytes(), &r); err != nil {
		t.Fatalf("Unmarshal: %v\nstdout: %s", err, stdout.String())
	}
	if r.Summary.ErrorsTotal != 0 {
		t.Fatalf("ErrorsTotal = %d, errors=%#v; rss= denial must not fail a zero-orphan reap", r.Summary.ErrorsTotal, r.Errors)
	}
	for _, e := range r.Errors {
		if e.Stage == "reap" {
			t.Fatalf("reap-stage error %q; scan must succeed when ps denies rss=", e.Error)
		}
	}
	if r.Reaped.Count != 0 {
		t.Fatalf("Reaped.Count = %d, want 0 orphans", r.Reaped.Count)
	}
	if len(r.Reaped.ProtectedPIDs) != 1 || r.Reaped.ProtectedPIDs[0] != 1138290 {
		t.Fatalf("ProtectedPIDs = %v, want [1138290] (safety decision without RSS)", r.Reaped.ProtectedPIDs)
	}
	if r.Summary.BytesFreedRSS != 0 {
		t.Fatalf("BytesFreedRSS = %d, want 0 (best-effort when rss= is unavailable)", r.Summary.BytesFreedRSS)
	}
}

func TestPSLStartOutputFormatExcludesRSS(t *testing.T) {
	if strings.Contains(psLStartOutputFormat, "rss") {
		t.Fatalf("psLStartOutputFormat = %q includes rss; Darwin ps hard-fails the whole scan when rss= is requested", psLStartOutputFormat)
	}
}
