# 02 — Fix Mac Regression & cmd/gc Process Failures

**What to build:** Fix the failing `Mac / quality` and `cmd/gc process` test shards (shards 3, 5, 10) in the Mac Regression workflow. Test fixes in capped Docker container and ship via PR.

**Blocked by:** 01 — Diagnose & Map Root Causes for Main CI Failures.

**Status:** claimed

- [ ] `Mac / quality` toolchain skew resolved: update `.github/actions/setup-gascity-macos/action.yml` to use `go-version-file: go.mod` (default `go-version: ""`).
- [ ] `cmd/gc` command census drift resolved: add `gc worktree` hierarchy (`gc worktree`, `add`, `remove`, `list`) to `cmd/gc/productmetrics_command_census.json` and regenerate `cmd/gc/metrics_census_gen.go` via `go run ./cmd/gen-command-census`.
- [ ] Map nil-panic in `cmd/gc/metrics_classifier_test.go` fixed by checking/initializing `command.Annotations`.
- [ ] Canonical resume command assertion in `cmd/gc/cmd_session_test.go` (`TestBuildResumeCommandIncludesWrappedCodexResumeDefaults`) updated.
- [ ] Tests verified passing in capped Docker container (`--cpus=2 --memory=4g --pids-limit=256`).
- [ ] PR opened against `jonbaldie/gascity` with isolated fix and verified green on Mac Regression workflow.

## Comments

### Root Cause Details
- **Toolchain skew in `Mac / quality`:** `setup-gascity-macos` defaulted to Go 1.26.5 while `go.mod` required 1.26.6, causing `golangci-lint` to fail under `GOTOOLCHAIN=local`.
- **Command census drift:** `gc worktree` command tree was added without registering in `productmetrics_command_census.json`, causing `validateProductMetricsCommandCensus` to fail closed and classify all commands as `census-mismatch`.
- **Codex resume expectation:** Subcommand-style resume canonicalization places schema flags after the `resume` subcommand token.
