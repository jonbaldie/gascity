package scripts_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestNotifyImageRebuildSkipsWhenTokenMissing(t *testing.T) {
	result := runNotifyImageRebuild(t, notifyRebuildEnv{
		token: "",
	})
	if result.exitCode != 0 {
		t.Fatalf("exit = %d, want 0 (no-op when token is unset)\nstdout:\n%s\nstderr:\n%s", result.exitCode, result.stdout, result.stderr)
	}
	if !strings.Contains(result.stdout+result.stderr, "skipping") {
		t.Fatalf("output = %q, want a skip reason", result.stdout+result.stderr)
	}
	if calls := result.ghCalls(); len(calls) != 0 {
		t.Fatalf("gh calls = %q, want none when token is unset", calls)
	}
}

func TestNotifyImageRebuildSkipsWhenTargetRepoUnavailable(t *testing.T) {
	result := runNotifyImageRebuild(t, notifyRebuildEnv{
		token:           "hosted-token",
		repoUnavailable: true,
	})
	if result.exitCode != 0 {
		t.Fatalf("exit = %d, want 0 (no-op when target repo is unavailable)\nstdout:\n%s\nstderr:\n%s", result.exitCode, result.stdout, result.stderr)
	}
	if !strings.Contains(result.stdout+result.stderr, "skipping") {
		t.Fatalf("output = %q, want a skip reason", result.stdout+result.stderr)
	}
	for _, call := range result.ghCalls() {
		if strings.Contains(call, "--method POST") || strings.Contains(call, "/dispatches") {
			t.Fatalf("dispatched to target despite unavailable repo: %q", call)
		}
	}
}

func TestNotifyImageRebuildSkipsWhenTargetNotWritable(t *testing.T) {
	result := runNotifyImageRebuild(t, notifyRebuildEnv{
		token: "hosted-token",
		push:  "false",
	})
	if result.exitCode != 0 {
		t.Fatalf("exit = %d, want 0 (no-op when this repo cannot write the target)\nstdout:\n%s\nstderr:\n%s", result.exitCode, result.stdout, result.stderr)
	}
	if !strings.Contains(result.stdout+result.stderr, "skipping") {
		t.Fatalf("output = %q, want a skip reason", result.stdout+result.stderr)
	}
	for _, call := range result.ghCalls() {
		if strings.Contains(call, "--method POST") || strings.Contains(call, "/dispatches") {
			t.Fatalf("dispatched without write access: %q", call)
		}
	}
}

func TestNotifyImageRebuildDispatchesWhenTargetIsWritable(t *testing.T) {
	result := runNotifyImageRebuild(t, notifyRebuildEnv{
		token: "hosted-token",
		push:  "true",
	})
	if result.exitCode != 0 {
		t.Fatalf("exit = %d, want 0\nstdout:\n%s\nstderr:\n%s", result.exitCode, result.stdout, result.stderr)
	}
	var dispatched bool
	for _, call := range result.ghCalls() {
		if strings.Contains(call, "--method POST") && strings.Contains(call, "repos/gascity/gasworks-control-plane/dispatches") {
			dispatched = true
			if !strings.Contains(call, "event_type=runtime-dep-updated") {
				t.Fatalf("dispatch call = %q, want runtime-dep-updated", call)
			}
		}
	}
	if !dispatched {
		t.Fatalf("gh calls = %q, want POST to gascity/gasworks-control-plane/dispatches", result.ghCalls())
	}
}

func TestNotifyImageRebuildUsesHostedTokenSecret(t *testing.T) {
	step := notifyImageRebuildStep(t)
	if got := step.Env["GH_TOKEN"]; got != "${{ secrets.GASCITY_HOSTED_TOKEN }}" {
		t.Fatalf("GH_TOKEN = %q, want secrets.GASCITY_HOSTED_TOKEN (do not invent a substitute secret)", got)
	}
	if strings.Contains(step.Run, "github.token") {
		t.Fatal("notify step must not fall back to github.token; that cannot write the image-host repo")
	}
}

type notifyRebuildEnv struct {
	token           string
	push            string
	repoUnavailable bool
}

type notifyRebuildResult struct {
	exitCode int
	stdout   string
	stderr   string
	logPath  string
}

func (r notifyRebuildResult) ghCalls() []string {
	raw, err := os.ReadFile(r.logPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return []string{"read log: " + err.Error()}
	}
	text := strings.TrimSpace(string(raw))
	if text == "" {
		return nil
	}
	return strings.Split(text, "\n")
}

func runNotifyImageRebuild(t *testing.T, env notifyRebuildEnv) notifyRebuildResult {
	t.Helper()
	script := notifyImageRebuildStep(t).Run
	binDir := t.TempDir()
	logPath := filepath.Join(t.TempDir(), "gh.log")
	writeExecutable(t, filepath.Join(binDir, "gh"), `#!/usr/bin/env bash
set -euo pipefail
log="${NOTIFY_GH_LOG:?}"
printf '%s\n' "$*" >> "$log"
if [[ -z "${GH_TOKEN:-}" ]]; then
  echo "gh: To use GitHub CLI in a GitHub Actions workflow, set the GH_TOKEN environment variable. Example:" >&2
  echo "  env:" >&2
  echo "    GH_TOKEN: \${{ github.token }}" >&2
  exit 4
fi
method=GET
endpoint=""
args=("$@")
i=0
while [[ $i -lt ${#args[@]} ]]; do
  case "${args[$i]}" in
    --method)
      i=$((i + 1))
      method="${args[$i]}"
      ;;
    --jq|-f)
      i=$((i + 1))
      ;;
    repos/*)
      endpoint="${args[$i]}"
      ;;
  esac
  i=$((i + 1))
done
if [[ "$method" != "GET" ]]; then
  exit 0
fi
if [[ "${NOTIFY_GH_REPO_UNAVAILABLE:-}" == "1" ]]; then
  echo "gh: Not Found (HTTP 404)" >&2
  exit 1
fi
printf '%s\n' "${NOTIFY_GH_PUSH:-false}"
`)

	cmd := exec.Command("bash", "-c", script)
	cmd.Dir = repoRoot(t)
	cmd.Env = []string{
		"PATH=" + binDir + string(os.PathListSeparator) + "/usr/bin:/bin",
		"HOME=" + t.TempDir(),
		"TMPDIR=" + t.TempDir(),
		"GH_TOKEN=" + env.token,
		"GITHUB_ACTIONS=true",
		"CI=true",
		"NOTIFY_GH_LOG=" + logPath,
		"NOTIFY_GH_PUSH=" + env.push,
	}
	if env.repoUnavailable {
		cmd.Env = append(cmd.Env, "NOTIFY_GH_REPO_UNAVAILABLE=1")
	}
	out, err := cmd.CombinedOutput()
	result := notifyRebuildResult{
		stdout:  string(out),
		stderr:  "",
		logPath: logPath,
	}
	if err == nil {
		return result
	}
	exitErr, ok := err.(*exec.ExitError)
	if !ok {
		t.Fatalf("run notify script: %v\n%s", err, out)
	}
	result.exitCode = exitErr.ExitCode()
	return result
}

type notifyWorkflowFile struct {
	Jobs map[string]struct {
		Steps []struct {
			Name string            `yaml:"name"`
			Run  string            `yaml:"run"`
			Env  map[string]string `yaml:"env"`
		} `yaml:"steps"`
	} `yaml:"jobs"`
}

func notifyImageRebuildStep(t *testing.T) struct {
	Name string            `yaml:"name"`
	Run  string            `yaml:"run"`
	Env  map[string]string `yaml:"env"`
} {
	t.Helper()
	raw, err := os.ReadFile(filepath.Join(repoRoot(t), ".github", "workflows", "notify-image-build.yaml"))
	if err != nil {
		t.Fatalf("read notify-image-build.yaml: %v", err)
	}
	var doc notifyWorkflowFile
	if err := yaml.Unmarshal(raw, &doc); err != nil {
		t.Fatalf("parse notify-image-build.yaml: %v", err)
	}
	job, ok := doc.Jobs["notify"]
	if !ok {
		t.Fatal("notify-image-build.yaml has no notify job")
	}
	for _, step := range job.Steps {
		if strings.TrimSpace(step.Run) != "" {
			return step
		}
	}
	t.Fatal("notify job has no run script")
	return struct {
		Name string            `yaml:"name"`
		Run  string            `yaml:"run"`
		Env  map[string]string `yaml:"env"`
	}{}
}
