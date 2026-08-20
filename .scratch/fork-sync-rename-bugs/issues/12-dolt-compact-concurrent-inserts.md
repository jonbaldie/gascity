# 12 — Dolt compact must not quarantine on concurrent inserts

**What to build:** Compaction on a live, continuously-written Beads/Dolt store does not silently quarantine that store because row counts or value hashes diverged from append-only inserts during the flatten window. GC and backups keep running on a busy city; quarantine is reserved for actual data loss, not captured churn.

**Blocked by:** 01, 02, 03, 04

**Status:** resolved

**Upstream:** https://github.com/gastownhall/gascity/issues/3341 · related https://github.com/gastownhall/gascity/issues/2348 · https://github.com/gastownhall/gascity/issues/2740

First-pass report deferred this until idle CPU shipped (PR #3). Do not turn this into unbounded-growth / wisp-prune (#3926) — that is a follow-on.

- [x] Concurrent inserts during flatten do not leave a durable quarantine that disables compaction/GC until a human deletes a marker
- [x] Integrity checking distinguishes row-count *gains* (churn) from *loss*; loss still fails closed
- [x] Tests or a deterministic script cover writer-during-flatten without requiring a multi-day city
- [x] Focused tests ran in resource-capped Docker (`--cpus=2 --memory=4g --pids-limit=256`), not a broad host `make test`
- [x] Dedicated branch + PR against `jonbaldie/gascity` for this ticket only

## Answer

Cause: `verify_counts` / `preflight_counts` called `COUNT(*)` and `DOLT_HASHOF_TABLE` against the live working set. Concurrent inserts during the flatten window changed those values with no HEAD movement, so gain+hash-drift wrote a durable quarantine (`post-flatten table value hash changed with row-count increase`). `HASHOF_TABLE` is unary over the session root and does not honor SQL `AS OF`.

Fix on `outpost/12-dolt-compact-concurrent-inserts`:

- Pin per-table count and hash probes to a Dolt revision database (`db/<commit>`): pre-flight at the stable HEAD, verify at the flatten commit.
- Live working-set inserts no longer look like flatten corruption. Row-count *decrease* at that snapshot still quarantines.

Landed:

- Branch: `outpost/12-dolt-compact-concurrent-inserts`
- Commit: `61d37c7804d013c8ce2dba1e100c725e19e4df3a`
- PR: https://github.com/jonbaldie/gascity/pull/23
- Worktree: `/Users/jonathanbaldie/go/src/github.com/jonbaldie/gascity-wt-26`
- Docker (`golang:1.26.6`, `--cpus=2 --memory=4g --pids-limit=256`, `CGO_ENABLED=0`): `go test ./examples/bd/dolt/ -run TestCompactScript` ok; `sh test/dolt/compact_gain_drift_proof_test.sh` 10/10. New seam: `TestCompactScriptIgnoresLiveInsertsDuringVerify`. Loss controls: `TestCompactScriptFailsOnRowCountDecreaseBeforeGC`, `TestCompactScriptStillQuarantinesRowDecreaseWithStableHead`.

Not in scope: wisp-prune / unbounded-growth (#3926).

## Comments

- 2026-08-19: Prefect dispatched after ticket 11 freed a worker slot. Worktree `/Users/jonathanbaldie/go/src/github.com/jonbaldie/gascity-wt-26`. Branch `outpost/12-dolt-compact-concurrent-inserts` from `origin/main` @ `468493bbe`.
- 2026-08-19: Resolved. PR #23. Snapshot-isolate compact integrity against the flatten commit so concurrent inserts cannot quarantine a live store.
