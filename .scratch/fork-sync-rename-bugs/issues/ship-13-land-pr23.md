# ship-13 — Land PR #23 (dolt compact concurrent inserts)

**What to build:** Merge https://github.com/jonbaldie/gascity/pull/23 into `main` on `jonbaldie/gascity`. Follow `/ship-pr`'s SKILL.md rigorously and `.scratch/fork-constitution-upstream/ship-main-protection.md` (no `--admin`, no direct push to `main`).

**Blocked by:** Ticket 12 completed (PR open).

**Status:** resolved

- [x] PR #23 merged under protected-main ruleset (or blocked with clear reason)

## Comments

- 2026-08-19 runner: claimed. Shipping https://github.com/jonbaldie/gascity/pull/23 via squash per repo convention and `/ship-pr`. Recorded: repo `jonbaldie/gascity`, base `main`, head `outpost/12-dolt-compact-concurrent-inserts` @ `61d37c7804d013c8ce2dba1e100c725e19e4df3a`, not draft, MERGEABLE, mergeStateStatus UNSTABLE (optional checks queued/pending), reviewDecision empty (ruleset required_approving_review_count=0), no host-native closing issue links, ruleset `21046899` has no required status checks (classic branch protection 404; only ruleset; effective rules on `main` are deletion + non_fast_forward + pull_request). Optional CI / Runner policy jobs are not required. Not using `--admin`. Did not ship PR #24.
- 2026-08-19 runner: `gh pr merge 23 --squash --match-head-commit 61d37c7804d013c8ce2dba1e100c725e19e4df3a --delete-branch` succeeded. Verified PR state MERGED at 2026-08-19T18:53:05Z. `git fetch origin main`; `git merge-base --is-ancestor 7a79b3962970211cd3a4f08d53cce6936636ce50 origin/main` succeeded; `origin/main` tip is that squash commit; remote head branch gone (GitHub git ref 404). No `map.md` in this effort directory to update.

## Answer

Squash-merged https://github.com/jonbaldie/gascity/pull/23 into `main` on `jonbaldie/gascity` without `--admin` or a direct push.

- **PR:** https://github.com/jonbaldie/gascity/pull/23 (`MERGED` at 2026-08-19T18:53:05Z)
- **Merge commit:** `7a79b3962970211cd3a4f08d53cce6936636ce50` (`git merge-base --is-ancestor` vs `origin/main` succeeded; `origin/main` is that commit)
- **Target:** `main` (ruleset `21046899`: required_approving_review_count=0, no required status checks; classic branch-protection API 404)
- **Head SHA at merge:** `61d37c7804d013c8ce2dba1e100c725e19e4df3a` (`--match-head-commit`, `--squash`, `--delete-branch`)
- **Branch result:** `outpost/12-dolt-compact-concurrent-inserts` deleted on origin (GitHub git ref 404)
- **GitHub issues:** none linked (`closingIssuesReferences` empty)
- **Not shipped:** PR #24 (ship-14)
