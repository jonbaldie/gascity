package main

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"net"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"
)

// TestStopSupervisorWithWaitFailsWhenLaunchdRemainsRestartable is the
// gastownhall/gascity#5324 class: the socket stop ACKs and the CLI reports
// success while launchd still considers the supervisor job loaded, so
// KeepAlive can restart it and re-adopt preserved sessions.
func TestStopSupervisorWithWaitFailsWhenLaunchdRemainsRestartable(t *testing.T) {
	startDarwinSupervisorStopFixture(t)

	oldRun := supervisorLaunchctlRun
	oldLoaded := supervisorLaunchdLoaded
	supervisorLaunchctlRun = func(...string) error { return nil }
	supervisorLaunchdLoaded = func(label string) (bool, string) {
		return true, "state = running\npid = 4242\nlabel = " + label
	}
	t.Cleanup(func() {
		supervisorLaunchctlRun = oldRun
		supervisorLaunchdLoaded = oldLoaded
	})

	var stdout, stderr bytes.Buffer
	code := stopSupervisorWithWait(&stdout, &stderr, true, time.Second)
	if code != 1 {
		t.Fatalf("stopSupervisorWithWait code = %d, want 1 (launchd still restartable); stdout=%q stderr=%q",
			code, stdout.String(), stderr.String())
	}
	if strings.Contains(stdout.String(), "Supervisor stopped.") {
		t.Fatalf("stdout = %q, must not report success while launchd remains restartable", stdout.String())
	}
	got := stderr.String()
	for _, want := range []string{"did not stop durably", "still loaded", supervisorLaunchdLabel()} {
		if !strings.Contains(got, want) {
			t.Fatalf("stderr = %q, want %q", got, want)
		}
	}
}

func TestStopSupervisorWithWaitSucceedsWhenLaunchdJobIsGone(t *testing.T) {
	fixture := startDarwinSupervisorStopFixture(t)

	oldRun := supervisorLaunchctlRun
	oldLoaded := supervisorLaunchdLoaded
	var calls []string
	supervisorLaunchctlRun = func(args ...string) error {
		calls = append(calls, strings.Join(args, " "))
		return nil
	}
	supervisorLaunchdLoaded = func(string) (bool, string) { return false, "" }
	t.Cleanup(func() {
		supervisorLaunchctlRun = oldRun
		supervisorLaunchdLoaded = oldLoaded
	})

	var stdout, stderr bytes.Buffer
	code := stopSupervisorWithWait(&stdout, &stderr, true, time.Second)
	if code != 0 {
		t.Fatalf("stopSupervisorWithWait code = %d, want 0; stderr=%q", code, stderr.String())
	}
	if !strings.Contains(stdout.String(), "Supervisor stopped.") {
		t.Fatalf("stdout = %q, want success confirmation", stdout.String())
	}
	wantCalls := []string{"disable " + fixture.target, "bootout " + fixture.target}
	if strings.Join(calls, "|") != strings.Join(wantCalls, "|") {
		t.Fatalf("launchctl calls = %v, want %v", calls, wantCalls)
	}
	if strings.Contains(strings.Join(calls, "|"), "unload "+fixture.plistPath) {
		t.Fatalf("launchctl calls = %v, should not fall back to unload after bootout succeeds", calls)
	}
}

func TestUnloadSupervisorServiceDarwinWaitsThroughSIGTERMedUntilAbsent(t *testing.T) {
	startDarwinSupervisorStopFixture(t)

	oldRun := supervisorLaunchctlRun
	oldLoaded := supervisorLaunchdLoaded
	var checks int
	supervisorLaunchctlRun = func(...string) error { return nil }
	supervisorLaunchdLoaded = func(label string) (bool, string) {
		checks++
		if checks < 3 {
			return true, "state = SIGTERMed\npid = 4242\nlabel = " + label
		}
		return false, ""
	}
	t.Cleanup(func() {
		supervisorLaunchctlRun = oldRun
		supervisorLaunchdLoaded = oldLoaded
	})

	if err := unloadSupervisorService(); err != nil {
		t.Fatalf("unloadSupervisorService returned error: %v", err)
	}
	if checks < 3 {
		t.Fatalf("launchd loaded checks = %d, want polling through SIGTERMed", checks)
	}
}

func TestUnloadSupervisorServiceDarwinReturnsLaunchctlFailures(t *testing.T) {
	fixture := startDarwinSupervisorStopFixture(t)

	oldRun := supervisorLaunchctlRun
	oldLoaded := supervisorLaunchdLoaded
	var calls []string
	supervisorLaunchctlRun = func(args ...string) error {
		calls = append(calls, strings.Join(args, " "))
		switch args[0] {
		case "disable":
			return errors.New("disable denied")
		case "bootout":
			return errors.New("bootout denied")
		case "unload":
			return errors.New("unload denied")
		default:
			return nil
		}
	}
	supervisorLaunchdLoaded = func(string) (bool, string) { return false, "" }
	t.Cleanup(func() {
		supervisorLaunchctlRun = oldRun
		supervisorLaunchdLoaded = oldLoaded
	})

	err := unloadSupervisorService()
	if err == nil {
		t.Fatal("unloadSupervisorService returned nil, want launchctl failure")
	}
	got := err.Error()
	for _, want := range []string{"launchctl disable", "disable denied", "launchctl bootout", "bootout denied", "launchctl unload", "unload denied"} {
		if !strings.Contains(got, want) {
			t.Fatalf("error = %q, want %q", got, want)
		}
	}
	wantCalls := []string{"disable " + fixture.target, "bootout " + fixture.target, "unload " + fixture.plistPath}
	if strings.Join(calls, "|") != strings.Join(wantCalls, "|") {
		t.Fatalf("launchctl calls = %v, want %v", calls, wantCalls)
	}
}

func TestStopSupervisorWithWaitFailsWhenSystemdStopFails(t *testing.T) {
	oldGOOS := supervisorRuntimeGOOS
	supervisorRuntimeGOOS = "linux"
	t.Cleanup(func() { supervisorRuntimeGOOS = oldGOOS })

	homeDir := t.TempDir()
	gcHome := shortTempDir(t, "gc-home-")
	runtimeDir := shortTempDir(t, "gc-run-")
	t.Setenv("HOME", homeDir)
	t.Setenv("GC_HOME", gcHome)
	t.Setenv("XDG_RUNTIME_DIR", runtimeDir)

	unitPath := supervisorSystemdServicePath()
	if err := os.MkdirAll(filepath.Dir(unitPath), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(unitPath, []byte("unit\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	startDestructiveStopSocket(t, supervisorSocketPath())

	oldRun := supervisorSystemctlRun
	supervisorSystemctlRun = func(args ...string) error {
		if len(args) >= 3 && args[1] == "stop" {
			return errors.New("unit stop denied")
		}
		return nil
	}
	t.Cleanup(func() { supervisorSystemctlRun = oldRun })

	var stdout, stderr bytes.Buffer
	code := stopSupervisorWithWait(&stdout, &stderr, true, time.Second)
	if code != 1 {
		t.Fatalf("stopSupervisorWithWait code = %d, want 1; stdout=%q stderr=%q", code, stdout.String(), stderr.String())
	}
	if strings.Contains(stdout.String(), "Supervisor stopped.") {
		t.Fatalf("stdout = %q, must not report success when systemd stop fails", stdout.String())
	}
	if !strings.Contains(stderr.String(), "did not stop durably") || !strings.Contains(stderr.String(), "unit stop denied") {
		t.Fatalf("stderr = %q, want durable-stop failure with systemctl error", stderr.String())
	}
}

type darwinSupervisorStopFixture struct {
	plistPath string
	target    string
}

func startDarwinSupervisorStopFixture(t *testing.T) darwinSupervisorStopFixture {
	t.Helper()

	oldGOOS := supervisorRuntimeGOOS
	supervisorRuntimeGOOS = "darwin"
	t.Cleanup(func() { supervisorRuntimeGOOS = oldGOOS })

	oldTimeout := supervisorLaunchdStopTimeout
	oldPoll := supervisorLaunchdStopPollInterval
	supervisorLaunchdStopTimeout = 20 * time.Millisecond
	supervisorLaunchdStopPollInterval = time.Millisecond
	t.Cleanup(func() {
		supervisorLaunchdStopTimeout = oldTimeout
		supervisorLaunchdStopPollInterval = oldPoll
	})

	homeDir := t.TempDir()
	gcHome := shortTempDir(t, "gc-home-")
	runtimeDir := shortTempDir(t, "gc-run-")
	t.Setenv("HOME", homeDir)
	t.Setenv("GC_HOME", gcHome)
	t.Setenv("XDG_RUNTIME_DIR", runtimeDir)

	path := supervisorLaunchdPlistPath()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte("<plist/>"), 0o600); err != nil {
		t.Fatal(err)
	}

	startDestructiveStopSocket(t, supervisorSocketPath())

	return darwinSupervisorStopFixture{
		plistPath: path,
		target:    supervisorLaunchdServiceTarget(supervisorLaunchdLabel()),
	}
}

// startDestructiveStopSocket answers ping until stop, ACKs the destructive
// stop, emits done:ok, then stops answering so --wait can observe exit.
func startDestructiveStopSocket(t *testing.T, sockPath string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(sockPath), 0o700); err != nil {
		t.Fatalf("MkdirAll(%q): %v", filepath.Dir(sockPath), err)
	}
	lis, err := net.Listen("unix", sockPath)
	if err != nil {
		t.Fatalf("Listen(unix, %q): %v", sockPath, err)
	}
	t.Cleanup(func() {
		lis.Close()         //nolint:errcheck
		os.Remove(sockPath) //nolint:errcheck
	})

	var (
		mu      sync.Mutex
		stopped bool
	)
	go func() {
		for {
			conn, err := lis.Accept()
			if err != nil {
				return
			}
			go func(conn net.Conn) {
				defer conn.Close() //nolint:errcheck
				r := bufio.NewReader(conn)
				line, err := r.ReadString('\n')
				if err != nil {
					return
				}
				switch strings.TrimSpace(line) {
				case "ping":
					mu.Lock()
					done := stopped
					mu.Unlock()
					if done {
						return
					}
					io.WriteString(conn, "4242\n") //nolint:errcheck
				case "stop":
					mu.Lock()
					stopped = true
					mu.Unlock()
					io.WriteString(conn, "ok\n")      //nolint:errcheck
					io.WriteString(conn, "done:ok\n") //nolint:errcheck
				}
			}(conn)
		}
	}()
}
