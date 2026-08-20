# 05 — Supervisor stop must be durable against launchd restart

**What to build:** On macOS, `gc supervisor stop --wait` either leaves the machine supervisor durably stopped (launchd cannot restart it) or fails visibly. Operators who stop the supervisor must not later find preserved sessions re-adopted and dispatching work.

**Blocked by:** 01, 02, 03, 04

**Status:** resolved

**Upstream:** https://github.com/gastownhall/gascity/issues/5324

Distinct from shipped PR #7 / first-pass ticket 08 (`#5256`): that bounded a wedged `bd list` during city teardown. This ticket is the platform-service postcondition: stop vs launchd KeepAlive.

- [x] After a successful `gc supervisor stop --wait`, launchd does not treat the supervisor job as loaded/restartable
- [x] If the platform cannot prove the job will stay down, the command fails (non-zero) instead of reporting success
- [x] Tests cover the durable-stop vs “reported success but still restartable” class
- [x] Focused tests ran in resource-capped Docker (`--cpus=2 --memory=4g --pids-limit=256`), not a broad host `make test`
- [x] Dedicated branch + PR against `jonbaldie/gascity` for this ticket only

## Answer

Stop now fail-closes on the platform-service postcondition. After the socket ACK, Darwin disable + bootout (unload fallback) runs, then `launchctl print` is polled until the job is absent (or the bounded wait expires). If launchd is still loaded/restartable, `gc supervisor stop` exits non-zero with `platform service did not stop durably` instead of printing `Supervisor stopped.` Launchd is behind `supervisorRuntimeGOOS` / `supervisorLaunchctlRun` / `supervisorLaunchdLoaded` so Docker can cover the class without a live launchd. Install/start re-enable before load.

- Branch: `outpost/05-supervisor-stop-launchd-durable`
- Commit: `d131a053e`
- PR: https://github.com/jonbaldie/gascity/pull/16
