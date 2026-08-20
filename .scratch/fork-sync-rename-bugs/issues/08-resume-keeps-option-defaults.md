# 08 — Resume must keep provider option_defaults

**What to build:** A resumed session is launched with the same provider `option_defaults` as a first spawn (permissions, effort, and any other configured defaults). Resume must not drop those flags so the agent parks on an approval prompt while liveness still reports healthy. This is a regression of the class upstream closed as #799.

**Blocked by:** 01, 02, 03, 04

**Status:** resolved

**Upstream:** https://github.com/gastownhall/gascity/issues/5185

No upstream fix PR at research time. Distinct from shipped nudge fail-closed (PR #6).

- [x] Resume and first-spawn launch commands both include configured provider `option_defaults`
- [x] A session resumed after drain/sleep does not lose the permission / effort defaults that spawn applied
- [x] Tests cover spawn vs resume command construction for at least the builtin claude-shaped provider
- [x] Focused tests ran in resource-capped Docker (`--cpus=2 --memory=4g --pids-limit=256`), not a broad host `make test`
- [x] Dedicated branch + PR against `jonbaldie/gascity` for this ticket only

## Answer

Resume dropped provider `option_defaults` because `BuildProviderResumeCommand` gated on explicit schema overrides only (`hasSchemaOptionOverrides`). A defaults-only resume of a builtin claude-shaped provider therefore emitted a bare `--resume` template and lost `--dangerously-skip-permissions` / `--effort max`. The launch path already counted `EffectiveDefaults` via `hasProviderOptionValues`; resume now uses the same predicate so stored/bare resume templates are strip-and-rebuilt with the merged defaults.

- Branch: `outpost/08-resume-keeps-option-defaults`
- Commit: `5ce8e6d9396f11ddcd1b727b838f5f735a903415`
- PR: https://github.com/jonbaldie/gascity/pull/18
- Docker (`golang:1.26.6`, `--cpus=2 --memory=4g --pids-limit=256`, `CGO_ENABLED=0`):
  - `go test ./internal/config -count=1 -run 'TestBuildProvider(Launch|Resume)Command'`
  - `go test ./cmd/gc -count=1 -run 'TestResolvedWorkerRuntimeKeepsProviderOptionDefaultsOnResume'`
