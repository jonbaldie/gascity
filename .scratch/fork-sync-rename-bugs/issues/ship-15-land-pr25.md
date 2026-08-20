# ship-15 — Land PR #25 (macOS dolt-cleanup ps rss)

**What to build:** Merge https://github.com/jonbaldie/gascity/pull/25 into `main` on `jonbaldie/gascity`. Follow `/ship-pr`'s SKILL.md rigorously and `.scratch/fork-constitution-upstream/ship-main-protection.md` (no `--admin`, no direct push to `main`).

**Blocked by:** Ticket 14 completed (PR open).

**Status:** resolved

- [x] PR #25 merged under protected-main ruleset (or blocked with clear reason)

## Comments

- 2026-08-19: Prefect: ticket 14 resolved; PR #25 OPEN on `outpost/14-macos-dolt-cleanup-ps` @ `e9ee94e078a346fe3a6f892b5dfedb736cc500e2`.
- 2026-08-19: Runner: claimed. Shipping via squash-merge of PR #25 into `jonbaldie/gascity` `main` (ruleset 21046899; 0 required reviews; no required status checks).
- 2026-08-19 runner: Recorded: repo `jonbaldie/gascity` (origin; never gastownhall), base `main`, head `outpost/14-macos-dolt-cleanup-ps` @ `e9ee94e078a346fe3a6f892b5dfedb736cc500e2`, not draft, MERGEABLE, mergeStateStatus UNSTABLE (optional checks queued/pending), reviewDecision empty (ruleset required_approving_review_count=0), no host-native closing issue links, ruleset `21046899` has no required status checks (classic branch protection 404; only ruleset; effective rules on `main` are deletion + non_fast_forward + pull_request). Optional CI / Runner policy jobs are not required. Not using `--admin`. Did not ff-pull the prefect's dirty local `main`.
- 2026-08-19 runner: `gh pr merge 25 --squash --match-head-commit e9ee94e078a346fe3a6f892b5dfedb736cc500e2 --delete-branch` succeeded. Verified PR state MERGED at 2026-08-19T19:08:03Z. `git fetch origin main`; `git merge-base --is-ancestor ba38067b254a8d18b803ebd563a7266ad70aa471 origin/main` succeeded; `origin/main` tip is that squash commit (single parent `7a79b3962970211cd3a4f08d53cce6936636ce50`); remote head branch gone (`git ls-remote` empty; GitHub git ref 404). No `map.md` in this effort directory to update.

## Answer

Squash-merged https://github.com/jonbaldie/gascity/pull/25 into `main` on `jonbaldie/gascity` without `--admin` or a direct push.

- **PR:** https://github.com/jonbaldie/gascity/pull/25 (`MERGED` at 2026-08-19T19:08:03Z)
- **Merge commit:** `ba38067b254a8d18b803ebd563a7266ad70aa471` (`git merge-base --is-ancestor` vs `origin/main` succeeded; `origin/main` is that commit)
- **Target:** `main` (ruleset `21046899`: required_approving_review_count=0, no required status checks; classic branch-protection API 404)
- **Head SHA at merge:** `e9ee94e078a346fe3a6f892b5dfedb736cc500e2` (`--match-head-commit`, `--squash`, `--delete-branch`)
- **Branch result:** `outpost/14-macos-dolt-cleanup-ps` deleted on origin
- **GitHub issues:** none linked (`closingIssuesReferences` empty)
