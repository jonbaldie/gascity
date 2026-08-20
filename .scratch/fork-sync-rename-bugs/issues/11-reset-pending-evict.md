# 11 — Evict wedged reset-pending sessions; bound crash-event payloads

**What to build:** After a config/provider change, a session left `reset-pending` because a stale-but-alive runtime still holds the tmux session is repaired (stale runtime evicted, new command started) instead of sitting wedged for hours. `session.crashed` must not fire on a loop with full pane dumps when nothing is actually crashing.

**Blocked by:** 01, 02, 03, 04

**Status:** resolved

**Upstream:** https://github.com/gastownhall/gascity/issues/5355

One vertical slice of the `#3872` durable-vs-runtime family (drift-relaunch / reset). Do not expand this ticket into adoption-alias, ghost beads, or scale-from-zero.

- [x] Reset-pending with a stale live runtime is evicted/replaced so the session is not wedged indefinitely
- [x] Crash detection does not emit unbounded full-pane `session.crashed` events for this wedge (~2/min for hours is the failure class)
- [x] Tests cover reset-pending + stale runtime still occupying the session
- [x] Focused tests ran in resource-capped Docker (`--cpus=2 --memory=4g --pids-limit=256`), not a broad host `make test`
- [x] Dedicated branch + PR against `jonbaldie/gascity` for this ticket only

## Answer

Cause: `recordResetStallIfDue` only logged/recorded the stall once the startup timeout elapsed, so a tmux session that was still running with a dead process stayed put indefinitely. Independently, the zombie branch emitted `session.crashed` on every ~30s tick with the full pane capture in `Message`.

Fix on `outpost/11-reset-pending-evict`:

- When a continuation reset is stalled and the tmux session is still running, evict via `workerKillSessionTargetWithConfig` (once per `markResetStall` dedup mark).
- Dedup zombie `session.crashed` across ticks (`drainTracker.zombieCrashes`, same shape as `resetStalls`); clear the mark when the session is observed alive or gone.
- Bound crash-event pane dumps to 24 lines (first/last 12) via `truncateCrashPaneOutput`. Classifier peeks and telemetry still see the full capture.

Landed:

- Branch: `outpost/11-reset-pending-evict`
- Commit: `23808bdf525279f27a652042d9d841ff6b02803b`
- PR: https://github.com/jonbaldie/gascity/pull/21
- Docker (`golang:1.26.6`, `--cpus=2 --memory=4g --pids-limit=256`, `CGO_ENABLED=0`):
  `go test ./cmd/gc -count=1 -run 'TestReconcileSessionBeads_ResetStallEvictsStaleRuntime|TestReconcileSessionBeads_ZombieCrashDedupesAcrossTicks|TestReconcileSessionBeads_ZombieCrashPayloadTruncated|TestTruncateCrashPaneOutput|TestReconcileSessionBeads_RecordsResetStallDiagnostic'` — ok

Same commit also drops a leftover `Daemon.FormulaV2: true` literal in `internal/workdir/implicit_pool_test.go` (`FormulaV2` is `*bool`; nil is default-on). Needed for `make vet` in pre-commit.

## Comments

- Worker claimed 2026-08-19. Worktree `/Users/jonathanbaldie/go/src/github.com/jonbaldie/gascity-wt-24`. Branch `outpost/11-reset-pending-evict`.
