# 13 — Fresh managed-Dolt init must finish a usable Beads schema

**What to build:** A new city using managed local Dolt completes Beads database/schema initialization (`hq` and required tables exist) and returns successfully. Operators are not left with a started Dolt server and a missing schema, with only the file provider as a workaround.

Follow `/diagnosing-bugs`'s SKILL.md rigorously.

**Blocked by:** 01, 02, 03, 04

**Status:** resolved

**Upstream:** https://github.com/gastownhall/gascity/issues/5086

Distinct from shipped PR #5 / `#1436` (issue_prefix seeding). Distinct from P0 `#5135` (beads dirty-tables deadlock — not this ticket). Ambient-`bd` schema-ceiling mismatch (`#5348`) is a follow-on, not this ticket.

- [x] Fresh managed-Dolt city init creates the Beads database and schema (not “server up, tables missing”)
- [x] Partial schema creation is retried or rolled back so init does not return “success” with an unusable store
- [x] Tests cover fresh managed-Dolt bootstrap completeness at the highest existing seam
- [x] Focused tests ran in resource-capped Docker (`--cpus=2 --memory=4g --pids-limit=256`), not a broad host `make test`
- [x] Dedicated branch + PR against `jonbaldie/gascity` for this ticket only

## Comments

- 2026-08-19: Worker claimed on `outpost/13-managed-dolt-schema` in worktree `gascity-wt-27` (from `eae28a8e6`).
- 2026-08-19: Prefect: Dispatch notes — this ticket is a bug, so `/diagnosing-bugs` after What to build (not `/implement`).

## Answer

bd 1.2 treats Gas City's managed `.beads/dolt` server root as a legacy workspace unless `.beads/.local_version` is present before `bd init`. Fresh city init pre-creates that layout and the Dolt catalog database, then `bd init` refused with `legacy Dolt server workspace detected`, so `hq` / `config` / migrations never appeared.

Fix: `gc-beads-bd.sh` tracks whether `ensure_database_registered` created the database in this invocation and writes the installed bd version witness only then. Pre-existing databases get no marker (migration guard preserved). Schema still invisible after wait+retry dies (`bd schema not visible`) instead of returning success.

- Branch: `outpost/13-managed-dolt-schema`
- Commit: `dc45320d1f05ca06c7ceab2596dbe482e73d1ac5`
- PR: https://github.com/jonbaldie/gascity/pull/24
- Docker (`golang:1.26.6`, `--cpus=2 --memory=4g --pids-limit=256`, `CGO_ENABLED=0`):
  `go test ./cmd/gc -count=1 -run 'TestGcBeadsBdInitMetadataOnlyFallsThroughToForcedBdInitWithPinnedDatabaseWhenSchemaMissing|TestGcBeadsBdInitDoesNotStampVersionWitnessOnPreexistingDatabase|TestGcBeadsBdInitFailsClosedWhenSchemaNeverBecomesVisible|TestGcBeadsBdInitWaitsForSchemaVisibilityBeforeRuntimeRepair|TestGcBeadsBdInitRetriesPlainInitWhenSchemaStillMissingAfterSuccess|TestGcBeadsBdInitDropsMetadataBeforeRetryingInitAfterForcedFallback|TestGCBeadsBDScript_SeedsVersionWitnessOnlyForCreatedDatabase|TestGCBeadsBDScript_InitForcesReinitOverPreSeededMetadata'` — ok

Red loop (before fix): fake bd 1.2 exited 3 with `legacy Dolt server workspace detected; explicit migration is required`.

### Code review

**Standards:** no documented-standard breaches. Witness write is atomic (`tmp` + `mv`, umask 077). Zero hardcoded roles. Source-inspection test `TestGCBeadsBDScript_SeedsVersionWitnessOnlyForCreatedDatabase` is implementation-coupled (judgement call / duplicated-code smell vs the behavioral lifecycle tests) but matches the existing `TestGCBeadsBDScript_InitForcesReinitOverPreSeededMetadata` pattern in the same file.

**Spec:** all five acceptance boxes covered. Witness gated on `database_created_by_gc` (fresh CREATE only). Fail-closed when schema never appears. Did not expand into `#5348` schema-ceiling or `#5135` dirty-tables.

Standards 0 hard / 1 judgement; Spec 0 missing. Worst Standards: source-inspection test. Worst Spec: none.
