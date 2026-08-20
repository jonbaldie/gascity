# 07 — gc start must not treat reload timeout as fatal before readiness

**What to build:** `gc start` / register may time out waiting for supervisor reload and still succeed if the city becomes ready asynchronously. A reload timeout must not unwind a start that the supervisor is still completing, and must not be the thing that makes later `gc stop` hit a busy queue against a city the CLI thinks never started.

**Blocked by:** 01, 02, 03, 04

**Status:** resolved

**Upstream:** https://github.com/gastownhall/gascity/issues/5333

Pairs with ticket 06 (stop vs busy reload) on the start side. Upstream issue includes a package-test repro shape.

- [x] Supervisor reload timeout during start/register is treated as async reconcile, not an automatic fatal rollback of a start still in flight
- [x] Readiness is still checked; a city that never becomes ready still fails
- [x] Test covers reload-timeout-during-start without requiring a live launchd/systemd install
- [x] Focused tests ran in resource-capped Docker (`--cpus=2 --memory=4g --pids-limit=256`), not a broad host `make test`
- [x] Dedicated branch + PR against `jonbaldie/gascity` for this ticket only

## Answer

Reload timeout during `gc start` / `gc register` is treated as an in-flight async reconcile. The command continues to the existing readiness wait and succeeds if the city becomes ready. Non-timeout reload failures stay fatal. A city that never becomes ready still fails and keeps the registration.

- Branch: `outpost/07-gc-start-reload-timeout`
- Commit: `d0edc2774`
- PR: https://github.com/jonbaldie/gascity/pull/17
- Worktree: `/Users/jonathanbaldie/go/src/github.com/jonbaldie/gascity-wt-20`

## Comments

- 2026-08-19: Worker claimed on `outpost/07-gc-start-reload-timeout` in worktree `gascity-wt-20` (from `58856ae2c`).
- 2026-08-19: Root cause was collapsing supervisor reload `timeout` to exit 1 and unwinding start before the readiness wait. Tests cover timeout-then-ready success and timeout-then-never-ready keep-registration failure, without live launchd/systemd.
