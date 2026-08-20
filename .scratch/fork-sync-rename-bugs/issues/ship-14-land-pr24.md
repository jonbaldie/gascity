# ship-14 — Land PR #24 (managed-Dolt schema init)

**What to build:** Merge https://github.com/jonbaldie/gascity/pull/24 into `main` on `jonbaldie/gascity`. Follow `/ship-pr`'s SKILL.md rigorously and `.scratch/fork-constitution-upstream/ship-main-protection.md` (no `--admin`, no direct push to `main`).

**Blocked by:** Ticket 13 completed (PR open). Dispatch after ship-13 if a runner is already landing PR #23 (`maxRunners: 1`).

**Status:** resolved

- [x] PR #24 merged under protected-main ruleset (or blocked with clear reason)

## Comments

- 2026-08-19: Prefect closed as duplicate. PR #24 was already squash-merged by `ship-13-land-pr24.md` (numbering collision). Did not dispatch a second runner.
- 2026-08-19 runner (this dispatch): found ticket already `Status: resolved`; did not revert to claimed. Verified live: PR https://github.com/jonbaldie/gascity/pull/24 is `MERGED` at 2026-08-19T18:52:51Z. Merge commit `9c7f328716d2d2daee5c327966740c8c275b0a45` is ancestor of `origin/main` (tip now `ba38067b254a8d18b803ebd563a7266ad70aa471` after later #23/#25). Head `outpost/13-managed-dolt-schema` gone (empty `git ls-remote --heads`; GitHub git ref 404). Ruleset `21046899` active: required_approving_review_count=0, no required status checks, allowed merge methods include squash. Did not merge again, did not use `--admin`, did not push to `main`, did not ship PR #25. No `map.md` in this effort directory.

## Answer

Duplicate of `.scratch/fork-sync-rename-bugs/issues/ship-13-land-pr24.md`. PR https://github.com/jonbaldie/gascity/pull/24 is `MERGED` at 2026-08-19T18:52:51Z (`9c7f328716d2d2daee5c327966740c8c275b0a45`). No second ship.
