# CI Failures Root Cause Diagnosis & Remediation Report

**Date:** 2026-08-20  
**Target Repository:** `jonbaldie/gascity` (`origin/main` at `ba38067b2`)  
**Scope:** Main CI workflows (`Nightly`, `Mac Regression`, `CI`, `Review Formulas`)

---

## Executive Summary

Continuous Integration on `jonbaldie/gascity` `main` is blocked by five distinct failure modes across four workflows:

1. **Nightly — Tier C Ollama:** Fails in `Validate Ollama Claude configuration` because the `OLLAMA_API_KEY` GitHub secret is not set on the repository.
2. **Nightly — SQLite Coordination Store:** Fails in `Acceptance tests (SQLite coordination store, Tier A)` because the experimental `sqlite` bead store provider was deprecated/removed from the engine, yet `nightly.yml` still schedules a dedicated job forcing `GC_BEADS=sqlite`.
3. **Mac Regression — Quality Gate:** Fails in `Mac / quality` because `.github/actions/setup-gascity-macos` defaulted to Go `1.26.5` while `go.mod` requires `1.26.6`, failing `golangci-lint` with toolchain skew under `GOTOOLCHAIN=local`.
4. **Mac Regression — cmd/gc Process Shards:** Shards fail due to three issues:
   - **Command Census Drift:** New command `gc worktree` (and subcommands `add`, `remove`, `list`) was added to Cobra without updating `cmd/gc/productmetrics_command_census.json`, causing the entire product metrics classifier to fail-closed on every command (`census-mismatch`).
   - **Nil Map Panic in Drift Test:** `TestClassifyProductMetricsCommandRejectsAnnotationDrift` mutates `command.Annotations` directly without nil-map initialization.
   - **Codex Resume Command Expectation:** `TestBuildResumeCommandIncludesWrappedCodexResumeDefaults` in `cmd_session_test.go` asserts a pre-subcommand flag ordering that conflicts with canonical `BuildProviderResumeCommand` flag placement.
5. **Main CI & Review Formulas — Stuck Pending on Main:** Workflows on `main` are stuck pending because the concurrency queue for `ci-push-refs/heads/main` and `review-formulas-push-refs/heads/main` is wedged by historical run `32265272842` (PR #1), which requested non-existent `blacksmith-*-ubuntu-2404` runners on push with `cancel-in-progress: false`.

---

## 1. Nightly Failures

### 1.1 Tier C Ollama Claude Configuration Validation

- **Run ID:** `32341195111` (Job ID: `96340489498`)
- **Workflow:** `.github/workflows/nightly.yml` (`tier-c` job)
- **Error Log:**
  ```text
  ##[group]Run test -n "$OLLAMA_API_KEY" || { echo "Missing OLLAMA_API_KEY GitHub secret" >&2; exit 1; }
  test -n "$OLLAMA_API_KEY" || { echo "Missing OLLAMA_API_KEY GitHub secret" >&2; exit 1; }
  ...
  Missing OLLAMA_API_KEY GitHub secret
  ##[error]Process completed with exit code 1.
  ```
- **Root Cause:**
  `nightly.yml` unconditionally runs `Validate Ollama Claude configuration` requiring `OLLAMA_API_KEY` and Anthropic model variables. On the `jonbaldie/gascity` fork, this secret is not configured.
- **Recommended Fix (Ticket 03):**
  Guard the step to gracefully skip when `OLLAMA_API_KEY` is not present (or skip the test step with a notice), allowing Nightly to succeed when optional external provider keys are omitted.

---

### 1.2 SQLite Coordination Store Tier A Acceptance

- **Run ID:** `32341195111` (Job ID: `96340489701`)
- **Workflow:** `.github/workflows/nightly.yml` (`integration-sqlite-coordstore` job)
- **Error Log:**
  ```text
  gc status: opening bead store: beads provider "sqlite" is no longer supported: the sqlite coordination-store experiment has been removed; update provider in city.toml to a supported value such as "doltlite", or remove the setting to use the default
  gc wait list: beads provider "sqlite" is no longer supported: the sqlite coordination-store experiment has been removed; update provider in city.toml to a supported value such as "doltlite", or remove the setting to use the default
  FAIL    github.com/jonbaldie/gascity/test/acceptance    152.436s
  make: *** [Makefile:559: test-acceptance] Error 1
  ```
- **Root Cause:**
  The `sqlite` bead store provider was deprecated/removed upstream. `nightly.yml` still contains a dedicated matrix/job `integration-sqlite-coordstore` that sets `GC_BEADS=sqlite` and `GC_ACCEPTANCE_BEADS_PROVIDER=sqlite` and runs `make test-acceptance`. Every CLI acceptance test attempting to open the bead store exits with status 1.
- **Recommended Fix (Ticket 04):**
  Remove the obsolete `integration-sqlite-coordstore` job from `.github/workflows/nightly.yml` (or update it to a supported provider such as `doltlite` if desired).

---

## 2. Mac Regression Failures

### 2.1 Mac / Quality Gate Toolchain Skew

- **Run ID:** `32330215192` (Job ID: `96309335881`)
- **Action:** `.github/actions/setup-gascity-macos/action.yml`
- **Error Log:**
  ```text
  GOFLAGS="$(go env GOFLAGS | sed -E 's/(^|[[:space:]])-mod=[^[:space:]]+//g') -mod=readonly" /Users/runner/go/bin/golangci-lint run ./...
  level=error msg="Running error: context loading failed: failed to load packages: failed to load packages: failed to load with go/packages: err: exit status 1: stderr: go: go.mod requires go >= 1.26.6 (running go 1.26.5; GOTOOLCHAIN=local)\n"
  make: *** [lint-full] Error 3
  ```
- **Root Cause:**
  In `go.mod`, Go is `1.26.6`. `.github/actions/setup-gascity-macos/action.yml` specified `default: "1.26.5"` rather than reading from `go.mod` (unlike `setup-gascity-ubuntu/action.yml` which uses `go-version-file: go.mod`).
- **Recommended Fix (Ticket 02):**
  Update `.github/actions/setup-gascity-macos/action.yml` to default `go-version: ""` and set `go-version-file: ${{ inputs.go-version == '' && 'go.mod' || '' }}`.

---

### 2.2 cmd/gc Process Shards Failures

- **Run ID:** `32330215192` (Job IDs: 96309336101 [Shard 3], 96309336091 [Shard 5], 96309336068 [Shard 10], etc.)

#### 2.2.1 Product Metrics Command Census Missing `gc worktree`
- **Error Log:**
  ```text
  metrics_census_test.go:59: product-metrics census: live command "gc worktree" is missing
  --- FAIL: TestProductMetricsCommandCensusMatchesProductionBuiltins (0.00s)
  --- FAIL: TestClassifyProductMetricsCommandPolicyMatrix (0.00s)
  ```
- **Root Cause:**
  The `gc worktree` command hierarchy (`gc worktree`, `add`, `remove`, `list`) was added to `cmd/gc/cmd_worktree.go` without registering it in `cmd/gc/productmetrics_command_census.json` or running `cmd/gen-command-census`.
  Because `validateProductMetricsCommandCensus` fails, `applyProductionProductMetricsCommandCensus` fails closed, setting no annotations on `root` and causing all commands to classify as `Exclusion: census-mismatch`.
- **Fix:**
  Add `gc worktree` and its subcommands to `productmetrics_command_census.json` and regenerate `cmd/gc/metrics_census_gen.go`.

#### 2.2.2 Nil Map Panic in `TestClassifyProductMetricsCommandRejectsAnnotationDrift`
- **Error Log:**
  ```text
  --- FAIL: TestClassifyProductMetricsCommandRejectsAnnotationDrift (0.00s)
      --- FAIL: TestClassifyProductMetricsCommandRejectsAnnotationDrift/known_id_swap (0.00s)
  panic: assignment to entry in nil map [recovered, repanicked]
  cmd/gc/metrics_classifier_test.go:255 +0x68
  ```
- **Root Cause:**
  When `command.Annotations` is `nil` (which occurs when census validation failed or when initialized freshly), direct map assignment panics in Go.
- **Fix:**
  In `cmd/gc/metrics_classifier_test.go`, check and initialize `if command.Annotations == nil { command.Annotations = make(map[string]string) }` before mutating.

#### 2.2.3 Codex Resume Flag Placement in `TestBuildResumeCommandIncludesWrappedCodexResumeDefaults`
- **Error Log:**
  ```text
  cmd_session_test.go:1088: resume command = "aimux run codex -- resume --dangerously-bypass-approvals-and-sandbox --model gpt-5.3-codex -c model_reasoning_effort=medium abc-123", want "aimux run codex -- --dangerously-bypass-approvals-and-sandbox -m gpt-5.3-codex resume -c model_reasoning_effort=medium abc-123"
  ```
- **Root Cause:**
  `BuildProviderResumeCommand` canonicalizes schema flag placement for subcommand-style providers (like Codex) after the `resume` subcommand token. The unit test expectation in `cmd_session_test.go` had an older manual flag ordering before `resume`.
- **Fix:**
  Update the expected resume command string in `cmd_session_test.go` to match the canonical schema-derived subcommand resume structure.

---

## 3. Main CI and Review Formulas Stuck Pending

- **Pending Runs:**
  - `CI` on `main`: Run `32291319640` (commit `#25`)
  - `Review Formulas` on `main`: Run `32291319798` (commit `#25`)
- **Stuck Head of Concurrency Queue:**
  - `CI` Run `32265272842` (`Make AGENTS.md the sole constitution; symlink CLAUDE.md to it. (#1)`) — status: `queued` (17h)
  - `Review Formulas` Run `32265272839` (`#1`) — status: `queued` (17h)
- **Root Cause:**
  In `ci.yml` and `review-formulas.yml`:
  ```yaml
  concurrency:
    group: ci-${{ github.event_name }}-${{ github.event.pull_request.number || github.ref || github.run_id }}
    cancel-in-progress: ${{ github.event_name == 'pull_request' }}
  ```
  On `push` to `main`, `cancel-in-progress` is `false`.
  Historical run `32265272842` (from PR #1) used an earlier version of `runner_policy.py` that selected `blacksmith-*-ubuntu-2404` for `push` events. Since Blacksmith runners do not exist on this repository, run `32265272842` entered `queued` status indefinitely.
  Because subsequent `push` events join the same concurrency group (`ci-push-refs/heads/main`) with `cancel-in-progress: false`, all later push runs (such as PR #25) sit in `pending` waiting for run #1 to finish.
- **Recommended Fix (Ticket 05):**
  1. Cancel all stale queued/pending workflow runs via `gh run cancel` (e.g. `32265272842`, `32265272839`, etc.).
  2. In `ci.yml` and `review-formulas.yml`, ensure `cancel-in-progress: true` is enabled or concurrency group properly differentiates runs, and ensure `runner_policy.py` consistently selects GitHub-hosted runners for pushes on non-Blacksmith setups.

---

## 4. Ticket Mapping & Action Plan

| Ticket | Scope | Action Items |
|---|---|---|
| **02** | Mac Regression & `cmd/gc` | 1. Update `setup-gascity-macos/action.yml` to use `go-version-file: go.mod`.<br>2. Add `gc worktree` hierarchy to `productmetrics_command_census.json` & regenerate.<br>3. Fix map nil-check in `metrics_classifier_test.go`.<br>4. Fix expected resume string in `cmd_session_test.go`. |
| **03** | Nightly: Ollama Tier C | Update `tier-c` in `nightly.yml` to gracefully skip validation and test execution when `OLLAMA_API_KEY` secret is absent. |
| **04** | Nightly: SQLite Tier A | Remove obsolete `integration-sqlite-coordstore` job from `nightly.yml`. |
| **05** | Unstick Workflows on Main | Cancel stalled historical runs (`32265272842`, `32265272839`, etc.), verify `concurrency` configuration in `ci.yml` and `review-formulas.yml`. |
| **06** | End-to-End Verification | Trigger and verify all CI workflows (`CI`, `Nightly`, `Mac Regression`, `Review Formulas`) run to green on `main`. |
