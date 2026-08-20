package verification

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestConfigDockerArgs_UsesBoundedIsolatedInvocation(t *testing.T) {
	t.Parallel()

	config := Config{
		ContainerName: "gc-verify-run-42",
		DockerBinary:  "docker-test",
		Image:         "golang:1.25.9",
		Workspace:     "/repo",
		ArtifactDir:   "/artifacts/run-42",
		CPUs:          "1.5",
		Memory:        "768m",
		PIDs:          64,
		TmpfsSize:     "256m",
		Timeout:       7 * time.Minute,
		Command:       []string{"go", "test", "./internal/config", "-run", "TestParser"},
	}

	got, err := config.DockerArgs()
	if err != nil {
		t.Fatalf("DockerArgs() error = %v", err)
	}

	want := []string{
		"run", "--name", "gc-verify-run-42", "--init", "--network", "none",
		"--cpus", "1.5", "--memory", "768m", "--pids-limit", "64",
		"--read-only", "--tmpfs", "/tmp:rw,noexec,nosuid,size=256m",
		"--mount", "type=bind,src=/repo,dst=/workspace,readonly",
		"--mount", "type=bind,src=/artifacts/run-42,dst=/artifacts",
		"--workdir", "/workspace",
		"--env", "TMPDIR=/tmp", "--env", "GOCACHE=/tmp/go-build", "--env", "GOMODCACHE=/tmp/go-mod",
		"--env", "GC_VERIFICATION_ARTIFACTS=/artifacts",
		"golang:1.25.9", "go", "test", "./internal/config", "-run", "TestParser",
	}
	assertStringsEqual(t, got, want)
}

func TestConfigDockerArgs_RejectsMissingCommand(t *testing.T) {
	t.Parallel()

	_, err := (Config{
		ContainerName: "gc-verify-run-42",
		DockerBinary:  "docker-test",
		Image:         "golang:1.25.9",
		Workspace:     "/repo",
		ArtifactDir:   "/artifacts/run-42",
		CPUs:          "2",
		Memory:        "1g",
		PIDs:          128,
		TmpfsSize:     "512m",
		Timeout:       time.Minute,
	}).DockerArgs()
	if err == nil {
		t.Fatal("DockerArgs() error = nil, want error for missing command")
	}
}

func TestConfigReplayScript_ReplaysBoundedInvocation(t *testing.T) {
	t.Parallel()

	config := Config{
		ContainerName: "gc-verify-run-42",
		DockerBinary:  "docker-test",
		Image:         "golang:1.25.9",
		Workspace:     "/repo",
		ArtifactDir:   "/artifacts/run-42",
		CPUs:          "1.5",
		Memory:        "768m",
		PIDs:          64,
		TmpfsSize:     "256m",
		Timeout:       7 * time.Minute,
		Command:       []string{"go", "test", "./..."},
	}

	got, err := config.ReplayScript("go run ./cmd/verify-container")
	if err != nil {
		t.Fatalf("ReplayScript() error = %v", err)
	}

	want := "#!/usr/bin/env sh\nset -eu\ncd '/repo'\nexec go run ./cmd/verify-container --docker 'docker-test' --image 'golang:1.25.9' --workspace '/repo' --artifact-dir '/artifacts/run-42' --cpus '1.5' --memory '768m' --pids '64' --tmpfs-size '256m' --timeout '7m0s' -- 'go' 'test' './...'\n"
	if got != want {
		t.Errorf("ReplayScript() = %q, want %q", got, want)
	}

	if filepath.Base(ReplayScriptName) != ReplayScriptName {
		t.Fatalf("ReplayScriptName = %q must be a file name", ReplayScriptName)
	}
}

func TestConfigWriteReplayScript_RetainsExecutableReplayArtifact(t *testing.T) {
	t.Parallel()

	artifactDir := filepath.Join(t.TempDir(), "run")
	config := Config{
		ContainerName: "gc-verify-run",
		DockerBinary:  "docker-test",
		Image:         "golang:1.25.9",
		Workspace:     "/repo",
		ArtifactDir:   artifactDir,
		CPUs:          "2",
		Memory:        "1g",
		PIDs:          128,
		TmpfsSize:     "512m",
		Timeout:       time.Minute,
		Command:       []string{"go", "test", "./..."},
	}

	path, err := config.WriteReplayScript("go run ./cmd/verify-container")
	if err != nil {
		t.Fatalf("WriteReplayScript() error = %v", err)
	}
	if path != filepath.Join(artifactDir, ReplayScriptName) {
		t.Errorf("replay path = %q, want %q", path, filepath.Join(artifactDir, ReplayScriptName))
	}
	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("stat replay script: %v", err)
	}
	if info.Mode().Perm() != 0o755 {
		t.Errorf("replay mode = %o, want 755", info.Mode().Perm())
	}
}

func assertStringsEqual(t *testing.T, got, want []string) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("argument count = %d, want %d\\ngot:  %#v\\nwant: %#v", len(got), len(want), got, want)
	}
	for index := range want {
		if got[index] != want[index] {
			t.Fatalf("argument %d = %q, want %q\\ngot:  %#v\\nwant: %#v", index, got[index], want[index], got, want)
		}
	}
}
