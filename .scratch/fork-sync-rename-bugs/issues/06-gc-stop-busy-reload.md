# 06 — gc stop must not leave the city running when reload is busy

**What to build:** An operator-issued `gc stop` either stops the city or fails in a way that does not restore a still-running city. A busy supervisor reconcile/reload queue must not unregister-then-re-register the city and report a confused half-stop while sessions keep running.

**Blocked by:** 01, 02, 03, 04

**Status:** resolved

**Upstream:** https://github.com/gastownhall/gascity/issues/5343

Distinct from shipped PR #7 / `#5256` (forced-stop timeout during wedged shutdown) and from ticket 05 (launchd durability of `gc supervisor stop`). This is the city `gc stop` path vs “reconcile queue is busy”.

- [x] `gc stop` does not end with the city still `running=true` / sessions alive after a busy-reload rejection
- [x] Busy-reload during stop is either retried until the city is down, queued until the supervisor can apply it, or failed closed without restoring a live registration that keeps the city up
- [x] Test or reproduction covers the busy-reconcile stop class (reload rejected as busy)
- [x] Focused tests ran in resource-capped Docker (`--cpus=2 --memory=4g --pids-limit=256`), not a broad host `make test`
- [x] Dedicated branch + PR against `jonbaldie/gascity` for this ticket only

## Answer

Busy reload during `gc stop` no longer restores a live registration. The stop path waits/retries until the supervisor observes unregistration, or fails closed leaving the city unregistered so it will not auto-restart.

- Branch: `outpost/06-gc-stop-busy-reload`
- Commit: `1d1c29711`
- PR: https://github.com/jonbaldie/gascity/pull/15
- Worktree: `/Users/jonathanbaldie/go/src/github.com/jonbaldie/gascity-wt-19`

## Comments

- 2026-08-19: Worker claimed on `outpost/06-gc-stop-busy-reload` in worktree `gascity-wt-19` (from `58856ae2c`).
- 2026-08-19: Root cause was treating `reconcile queue is busy` as a fatal reload failure that re-registered the city after unregister. Tests cover busy-then-down, busy wait-timeout leave-unregistered, and init-status (`starting_bead_store`) wait.
