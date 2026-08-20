// verify-container runs exploratory verification commands in a bounded Docker container.
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/jonbaldie/gascity/internal/verification"
)

const replayRunner = "go run ./cmd/verify-container"

func main() {
	os.Exit(run())
}

func run() int {
	config, dockerBinary, err := parseConfig(os.Args[1:], time.Now().UTC())
	if err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return 0
		}
		fmt.Fprintf(os.Stderr, "verify-container: %v\n", err)
		return 2
	}

	replayPath, err := config.WriteReplayScript(replayRunner)
	if err != nil {
		fmt.Fprintf(os.Stderr, "verify-container: %v\n", err)
		return 1
	}
	arguments, err := config.DockerArgs()
	if err != nil {
		fmt.Fprintf(os.Stderr, "verify-container: %v\n", err)
		return 2
	}

	fmt.Printf("Artifacts: %s\n", config.ArtifactDir)
	fmt.Printf("Replay: %s\n", replayPath)
	fmt.Printf("Wall-clock limit: %s\n", config.Timeout)

	ctx, cancel := context.WithTimeout(context.Background(), config.Timeout)
	defer cancel()
	command := exec.CommandContext(ctx, dockerBinary, arguments...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	runErr := command.Run()
	cleanupErr := removeContainer(dockerBinary, config.ContainerName)
	if cleanupErr != nil {
		fmt.Fprintf(os.Stderr, "verify-container: %v\n", cleanupErr)
	}
	if runErr != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			fmt.Fprintf(os.Stderr, "verify-container: wall-clock limit exceeded after %s; replay with %s\n", config.Timeout, replayPath)
			return 1
		}
		fmt.Fprintf(os.Stderr, "verify-container: Docker execution failed: %v; replay with %s\n", runErr, replayPath)
		return 1
	}
	if cleanupErr != nil {
		return 1
	}
	return 0
}

func parseConfig(arguments []string, now time.Time) (verification.Config, string, error) {
	workingDirectory, err := os.Getwd()
	if err != nil {
		return verification.Config{}, "", fmt.Errorf("finding current directory: %w", err)
	}

	flags := flag.NewFlagSet("verify-container", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	flags.Usage = func() {
		_, _ = fmt.Fprintln(flags.Output(), "Usage: verify-container [flags] -- command [arguments...]")
		flags.PrintDefaults()
	}
	image := flags.String("image", "golang:1.25.9", "Docker image used for the verification command")
	workspace := flags.String("workspace", workingDirectory, "workspace mounted read-only into the container")
	artifactDir := flags.String("artifact-dir", filepath.Join(workingDirectory, ".gc", "verification", now.Format("20060102T150405Z")), "host directory for generated inputs, seeds, and schedules")
	cpus := flags.String("cpus", "2", "Docker CPU limit")
	memory := flags.String("memory", "4g", "Docker memory limit")
	pids := flags.Int("pids", 256, "Docker process-count limit")
	tmpfsSize := flags.String("tmpfs-size", "1g", "size limit for the container temporary filesystem")
	timeout := flags.Duration("timeout", 15*time.Minute, "wall-clock limit for the Docker invocation")
	dockerBinary := flags.String("docker", "docker", "Docker executable")
	if err := flags.Parse(arguments); err != nil {
		return verification.Config{}, "", err
	}
	if len(flags.Args()) == 0 {
		return verification.Config{}, "", fmt.Errorf("verification command is required after --")
	}

	absWorkspace, err := filepath.Abs(*workspace)
	if err != nil {
		return verification.Config{}, "", fmt.Errorf("resolving workspace %q: %w", *workspace, err)
	}
	absArtifactDir, err := filepath.Abs(*artifactDir)
	if err != nil {
		return verification.Config{}, "", fmt.Errorf("resolving artifact directory %q: %w", *artifactDir, err)
	}

	config := verification.Config{
		ContainerName: verificationContainerName(absArtifactDir, now),
		DockerBinary:  *dockerBinary,
		Image:         *image,
		Workspace:     absWorkspace,
		ArtifactDir:   absArtifactDir,
		CPUs:          *cpus,
		Memory:        *memory,
		PIDs:          *pids,
		TmpfsSize:     *tmpfsSize,
		Timeout:       *timeout,
		Command:       flags.Args(),
	}
	if _, err := config.DockerArgs(); err != nil {
		return verification.Config{}, "", err
	}
	return config, *dockerBinary, nil
}

func removeContainer(dockerBinary, containerName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	output, err := exec.CommandContext(ctx, dockerBinary, "rm", "-f", containerName).CombinedOutput()
	if err != nil {
		return fmt.Errorf("removing verification container %q: %w: %s", containerName, err, strings.TrimSpace(string(output)))
	}
	return nil
}

func verificationContainerName(artifactDir string, now time.Time) string {
	base := filepath.Base(artifactDir)
	var sanitized strings.Builder
	for _, character := range base {
		if (character >= 'a' && character <= 'z') || (character >= 'A' && character <= 'Z') || (character >= '0' && character <= '9') || character == '.' || character == '_' || character == '-' {
			sanitized.WriteRune(character)
		}
	}
	name := strings.Trim(sanitized.String(), ".-_")
	if name == "" {
		name = "run"
	}
	if len(name) > 32 {
		name = name[:32]
	}
	return fmt.Sprintf("gc-verify-%s-%d", name, now.UnixNano())
}
