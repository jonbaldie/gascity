package main

import (
	"path/filepath"
	"testing"
	"time"
)

func TestParseConfig_UsesExplicitBoundsAndCommand(t *testing.T) {
	t.Parallel()

	workspace := t.TempDir()
	artifacts := filepath.Join(workspace, "artifacts")
	config, dockerBinary, err := parseConfig([]string{
		"--image", "golang:1.25.9",
		"--workspace", workspace,
		"--artifact-dir", artifacts,
		"--cpus", "1.5",
		"--memory", "768m",
		"--pids", "64",
		"--tmpfs-size", "256m",
		"--timeout", "7m",
		"--docker", "docker-test",
		"--", "go", "test", "./internal/config",
	}, time.Date(2026, time.August, 20, 10, 30, 0, 0, time.UTC))
	if err != nil {
		t.Fatalf("parseConfig() error = %v", err)
	}

	if dockerBinary != "docker-test" {
		t.Errorf("docker binary = %q, want docker-test", dockerBinary)
	}
	if config.DockerBinary != dockerBinary {
		t.Errorf("replay Docker binary = %q, want %q", config.DockerBinary, dockerBinary)
	}
	if config.Image != "golang:1.25.9" || config.Workspace != workspace || config.ArtifactDir != artifacts {
		t.Errorf("locations = %#v", config)
	}
	if config.ContainerName != "gc-verify-artifacts-1787221800000000000" {
		t.Errorf("container name = %q, want gc-verify-artifacts-1787221800000000000", config.ContainerName)
	}
	if config.CPUs != "1.5" || config.Memory != "768m" || config.PIDs != 64 || config.TmpfsSize != "256m" || config.Timeout != 7*time.Minute {
		t.Errorf("bounds = %#v", config)
	}
	assertStringsEqual(t, config.Command, []string{"go", "test", "./internal/config"})
}

func TestParseConfig_RequiresCommand(t *testing.T) {
	t.Parallel()

	_, _, err := parseConfig(nil, time.Now())
	if err == nil {
		t.Fatal("parseConfig() error = nil, want missing command error")
	}
}

func assertStringsEqual(t *testing.T, got, want []string) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("argument count = %d, want %d", len(got), len(want))
	}
	for index := range want {
		if got[index] != want[index] {
			t.Fatalf("argument %d = %q, want %q", index, got[index], want[index])
		}
	}
}
