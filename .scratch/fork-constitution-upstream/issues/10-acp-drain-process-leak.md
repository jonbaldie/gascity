# 10 — Terminate ACP provider process on pool drain-ack

**What to build:** When a pool agent self-completes and drain-acks, the ACP provider OS process is terminated. Idle cities no longer accumulate leaked provider processes that hold RAM and keep sessions falsely “alive.”

**Blocked by:** None — can start immediately.

**Status:** resolved

**Upstream:** https://github.com/gastownhall/gascity/issues/5218

- [x] Drain-ack / self-complete path stops the ACP provider process for that session
- [x] Repeated pool churn does not unbounded-accumulate idle provider processes under a minimal repro
- [x] Test covers terminate-on-drain behaviour at the provider/session seam
- [x] Dedicated branch + PR against `jonbaldie/gascity` for this ticket only

## Answer

Root cause matched upstream #5218 follow-up: drain-ack already calls `Stop`, but ACP `GetLastActivity` was in-memory-only in the reconciler process, so idle/activity-aware reaping could not fire and provider OS processes accumulated.

Fix (port of upstream #4612 shape): durable `gc_last_activity` sidecar + `CanReportActivity`, plus cross-process Stop seam test.

- Branch: `outpost/10-acp-drain-process-leak`
- Commit: `12aab283`
- PR: https://github.com/jonbaldie/gascity/pull/8

## Comments

- 2026-08-19: Worker claimed, diagnosed, shipped PR #8. Not merged (protected main; runner/prefect merges).
