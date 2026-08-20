# 14 — macOS dolt-cleanup reap must not hard-fail on ps rss=

**What to build:** After ticket 02 merges upstream `gc dolt-cleanup` into this fork, a reap/scan with zero orphans on macOS succeeds. `ps` refusing `rss=` (entitlement) must not fail the whole reap stage; RSS in the summary may be zero/best-effort.

Follow `/diagnosing-bugs`'s SKILL.md rigorously.

**Blocked by:** 01, 02, 03, 04

**Status:** resolved

**Upstream:** https://github.com/gastownhall/gascity/issues/5201

The Go `gc dolt-cleanup` command exists on `gastownhall/gascity` `main` and is absent from this fork until ticket 02. Do not invent the command; fix the Darwin `ps` hard-fail on the merged tree.

- [x] Reap/scan does not return a top-level error solely because `ps` denied `rss=` / resource fields
- [x] Safety decisions (which PIDs are orphans) still work without RSS
- [x] Tests cover `ps` failing on rss/resource columns without failing the reap
- [x] Focused tests ran in resource-capped Docker (`--cpus=2 --memory=4g --pids-limit=256`) where applicable; Darwin-specific behavior may use a fake `ps` rather than a live Mac entitlement
- [x] Dedicated branch + PR against `jonbaldie/gascity` for this ticket only

## Comments

- 2026-08-19: Prefect: Dispatch notes — this ticket is a bug, so `/diagnosing-bugs` after What to build (not `/implement`).
- 2026-08-19: Prefect: Two workers raced this ticket on `gascity-wt-28`. Keep **one** worker on 14 in that worktree. The duplicate is redirected to ticket 19 and must stop editing wt-28.
- 2026-08-19: Prefect: Keeper worker aborted mid-red-test. Redirected worker reverted uncommitted tests in wt-28. Re-dispatching 14 on the same worktree. Merge `origin/main` first (behind by PRs #23/#24). Rewrite the Darwin `ps` denial tests; do not assume they are still in the tree.
- 2026-08-19: Worker: claimed on `outpost/14-macos-dolt-cleanup-ps` in worktree `gascity-wt-28`. Merged `origin/main` (PRs #23/#24). Rewrote Darwin `ps` denial tests from scratch.

## Answer

`psLStartCommandLines` requested `rss=` as a `ps -o` column. On Darwin, `ps` exits non-zero for the whole invocation when resource-usage fields require an entitlement (`ps: %mem/vsz/rss/time: requires entitlement`). That error propagated to the reap stage even with zero orphans. RSS from this call is cosmetic (`bytes_freed_rss`); classify/protect uses config path, cwd, and ports.

Fix: drop `rss=` from the Darwin/ps fallback (`pid=,lstart=,command=`), parse six leading fields (pid + lstart), leave `RSSBytes` zero on this path. Linux `/proc` RSS is unchanged.

- Branch: `outpost/14-macos-dolt-cleanup-ps`
- Commit: `e9ee94e078a346fe3a6f892b5dfedb736cc500e2`
- PR: https://github.com/jonbaldie/gascity/pull/25
- Docker (`golang:1.26.6`, `--cpus=2 --memory=4g --pids-limit=256`, `CGO_ENABLED=0`):
  `go test ./cmd/gc -count=1 -run 'TestPSLStartCommandLines_ResourceFieldsDeniedDoesNotFailScan|TestDiscoverDoltProcessesFromPS_ResourceFieldsDeniedDoesNotFailScan|TestParseDoltPSLine|TestRunDoltCleanup_ReapScanSucceedsWhenPSDeniesRSS|TestSameReapProcessIdentity'` — ok; broader `TestRunDoltCleanup|TestLoadRigDoltPorts|TestLooksLikeDoltSQLServer|TestPlanOrphan|TestClassifyDolt` plus `go vet ./cmd/gc` — ok

Red loop (before fix): fake `ps` exited 1 with the Darwin entitlement message; `runDoltCleanup` recorded `ErrorsTotal=1` / `stage:"reap"` / `error:"exit status 1"`.

### Code review

## Standards

No documented-standard breaches. Zero hardcoded roles. Error path still returns `cmd.Output()` errors verbatim (pre-existing; not introduced). Fake `ps` on PATH matches the existing `cmd/gc` test seam (`dolt_process_inspection_test.go`). `TestPSLStartOutputFormatExcludesRSS` is a judgement-call tautological pin of the format constant (the fake-`ps` tests already observe the argv); kept as a cheap re-introduction guard matching upstream #5253.

Smell baseline: no Feature Envy / shotgun surgery. Constants `psLStartOutputFormat` / `psLStartLeadingFieldCount` are used at the exec and parse sites (not speculative).

## Spec

All five acceptance boxes covered. Reap scan no longer errors solely because `ps` would deny `rss=`. `/tmp/Test*` still reaps and `/var/lib/dolt` still protects with no rss column. `BytesFreedRSS` is 0 on the ps path. Did not invent a new command; fixed the merged `gc dolt-cleanup` Darwin `ps` hard-fail. PR is against `jonbaldie/gascity` only (not gastownhall).

Standards 0 hard / 1 judgement; Spec 0 missing. Worst Standards: format-constant pin. Worst Spec: none.
