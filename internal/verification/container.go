// Package verification builds bounded Docker invocations for exploratory checks.
package verification

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	// ReplayScriptName is the replay script retained with a verification run's artifacts.
	ReplayScriptName = "replay.sh"

	workspaceMount = "/workspace"
	artifactMount  = "/artifacts"
)

// Config defines a Docker-contained exploratory verification run.
type Config struct {
	ContainerName string
	DockerBinary  string
	Image         string
	Workspace     string
	ArtifactDir   string
	CPUs          string
	Memory        string
	PIDs          int
	TmpfsSize     string
	Timeout       time.Duration
	Command       []string
}

// DockerArgs returns the bounded Docker arguments for Config.
func (c Config) DockerArgs() ([]string, error) {
	if err := c.validate(); err != nil {
		return nil, err
	}

	arguments := []string{
		"run", "--name", c.ContainerName, "--init", "--network", "none",
		"--cpus", c.CPUs,
		"--memory", c.Memory,
		"--pids-limit", fmt.Sprint(c.PIDs),
		"--read-only",
		"--tmpfs", "/tmp:rw,noexec,nosuid,size=" + c.TmpfsSize,
		"--mount", "type=bind,src=" + c.Workspace + ",dst=" + workspaceMount + ",readonly",
		"--mount", "type=bind,src=" + c.ArtifactDir + ",dst=" + artifactMount,
		"--workdir", workspaceMount,
		"--env", "TMPDIR=/tmp",
		"--env", "GOCACHE=/tmp/go-build",
		"--env", "GOMODCACHE=/tmp/go-mod",
		"--env", "GC_VERIFICATION_ARTIFACTS=" + artifactMount,
		c.Image,
	}
	return append(arguments, c.Command...), nil
}

// ReplayScript returns a POSIX shell script that invokes runner with the same bounds.
func (c Config) ReplayScript(runner string) (string, error) {
	if err := c.validate(); err != nil {
		return "", err
	}
	if strings.TrimSpace(runner) == "" {
		return "", fmt.Errorf("replay runner is required")
	}

	arguments := []string{
		"--docker " + shellQuote(c.DockerBinary),
		"--image " + shellQuote(c.Image),
		"--workspace " + shellQuote(c.Workspace),
		"--artifact-dir " + shellQuote(c.ArtifactDir),
		"--cpus " + shellQuote(c.CPUs),
		"--memory " + shellQuote(c.Memory),
		"--pids " + shellQuote(fmt.Sprint(c.PIDs)),
		"--tmpfs-size " + shellQuote(c.TmpfsSize),
		"--timeout " + shellQuote(c.Timeout.String()),
		"--",
	}
	for _, argument := range c.Command {
		arguments = append(arguments, shellQuote(argument))
	}

	return "#!/usr/bin/env sh\nset -eu\ncd " + shellQuote(c.Workspace) + "\nexec " + runner + " " + strings.Join(arguments, " ") + "\n", nil
}

// WriteReplayScript writes the replay script next to the generated artifacts.
func (c Config) WriteReplayScript(runner string) (string, error) {
	script, err := c.ReplayScript(runner)
	if err != nil {
		return "", err
	}
	if err := os.MkdirAll(c.ArtifactDir, 0o755); err != nil {
		return "", fmt.Errorf("creating artifact directory %q: %w", c.ArtifactDir, err)
	}

	temporary, err := os.CreateTemp(c.ArtifactDir, ReplayScriptName+"-*")
	if err != nil {
		return "", fmt.Errorf("creating replay script in %q: %w", c.ArtifactDir, err)
	}
	temporaryName := temporary.Name()
	defer func() { _ = os.Remove(temporaryName) }()

	if _, err := temporary.WriteString(script); err != nil {
		if closeErr := temporary.Close(); closeErr != nil {
			return "", fmt.Errorf("writing and closing replay script: %w", errors.Join(err, closeErr))
		}
		return "", fmt.Errorf("writing replay script: %w", err)
	}
	if err := temporary.Chmod(0o755); err != nil {
		if closeErr := temporary.Close(); closeErr != nil {
			return "", fmt.Errorf("marking and closing replay script: %w", errors.Join(err, closeErr))
		}
		return "", fmt.Errorf("marking replay script executable: %w", err)
	}
	if err := temporary.Close(); err != nil {
		return "", fmt.Errorf("closing replay script: %w", err)
	}

	path := filepath.Join(c.ArtifactDir, ReplayScriptName)
	if err := os.Rename(temporaryName, path); err != nil {
		return "", fmt.Errorf("publishing replay script %q: %w", path, err)
	}
	return path, nil
}

func (c Config) validate() error {
	for _, field := range []struct {
		name  string
		value string
	}{
		{name: "container name", value: c.ContainerName},
		{name: "docker executable", value: c.DockerBinary},
		{name: "image", value: c.Image},
		{name: "workspace", value: c.Workspace},
		{name: "artifact directory", value: c.ArtifactDir},
		{name: "CPU limit", value: c.CPUs},
		{name: "memory limit", value: c.Memory},
		{name: "temporary storage limit", value: c.TmpfsSize},
	} {
		if strings.TrimSpace(field.value) == "" {
			return fmt.Errorf("%s is required", field.name)
		}
	}
	if !filepath.IsAbs(c.Workspace) {
		return fmt.Errorf("workspace must be an absolute path: %q", c.Workspace)
	}
	if !filepath.IsAbs(c.ArtifactDir) {
		return fmt.Errorf("artifact directory must be an absolute path: %q", c.ArtifactDir)
	}
	if c.PIDs <= 0 {
		return fmt.Errorf("process limit must be positive: %d", c.PIDs)
	}
	if c.Timeout <= 0 {
		return fmt.Errorf("wall-clock limit must be positive: %s", c.Timeout)
	}
	if len(c.Command) == 0 {
		return fmt.Errorf("verification command is required after --")
	}
	return nil
}

func shellQuote(value string) string {
	return "'" + strings.ReplaceAll(value, "'", "'\\\"'\\\"'") + "'"
}
