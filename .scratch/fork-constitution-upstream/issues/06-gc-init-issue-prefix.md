# 06 — Seed beads issue_prefix on gc init

**What to build:** After a fresh `gc init`, the city's beads store has a usable `issue_prefix` so session attach and other bead writes succeed without a manual `bd init --force --prefix …` recovery (and without creating a parasitic split store).

**Blocked by:** None — can start immediately.

**Status:** resolved

**Upstream:** https://github.com/gastownhall/gascity/issues/1436

- [x] Fresh `gc init` leaves `issue_prefix` set in the city beads config/store
- [x] A post-init bead write path used by operators (e.g. session attach or equivalent) succeeds without manual prefix recovery
- [x] Regression test covers the init → prefix-present path
- [x] Dedicated branch + PR against `jonbaldie/gascity` for this ticket only

## Answer

Managed `gc-beads-bd` already SQL-seeds `issue_prefix` (bd rejects `bd config set issue_prefix`). This ticket hardens the path:

1. Fail-closed verification after managed init (`verifyRuntimeIssuePrefix`)
2. `gc doctor` `issue-prefix:*` checks so doctor can no longer be green while creates fail
3. Stronger regressions for the #1436 init → prefix-present path

- Branch: `outpost/06-gc-init-issue-prefix`
- Commit: `d5e029db`
- PR: https://github.com/jonbaldie/gascity/pull/5

## Comments

- PREFECT STEER: host thermal — use capped Docker for remaining go test/builds (see `.scratch/fork-constitution-upstream/host-thermal-steer.md` and worktree `.outpost-steer.md`).
- Worker: claimed, diagnosed (#1436 — missing Dolt `issue_prefix` row; `bd config set issue_prefix` still rejected on bd 1.2.2), shipped PR #5, resolved.
