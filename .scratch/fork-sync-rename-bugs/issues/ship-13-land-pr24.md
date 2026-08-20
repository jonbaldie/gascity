# ship-13 — Land PR #24 (managed-Dolt schema init)

**What to build:** Merge https://github.com/jonbaldie/gascity/pull/24 into `main` on `jonbaldie/gascity`. Follow `/ship-pr`'s SKILL.md rigorously and `.scratch/fork-constitution-upstream/ship-main-protection.md` (no `--admin`, no direct push to `main`).

**Blocked by:** Ticket 13 completed (PR open).

**Status:** resolved

- [x] PR #24 merged under protected-main ruleset (or blocked with clear reason)

## Comments

- 2026-08-19 runner: claimed. Shipping https://github.com/jonbaldie/gascity/pull/24 via squash per repo convention and `/ship-pr`. Recorded: repo `jonbaldie/gascity` (origin; never gastownhall), base `main`, head `outpost/13-managed-dolt-schema` @ `dc45320d1f05ca06c7ceab2596dbe482e73d1ac5`, not draft, MERGEABLE, mergeStateStatus UNSTABLE (optional checks queued/pending), reviewDecision empty (ruleset required_approving_review_count=0), no host-native closing issue links, ruleset `21046899` has no required status checks (classic branch protection 404; `gh pr checks --required`: none). Optional `scripts/cipolicy` / Runner policy jobs are not required. Not using `--admin`.
- 2026-08-19 runner: `gh pr merge --squash --match-head-commit dc45320d1f05ca06c7ceab2596dbe482e73d1ac5 --delete-branch` succeeded. Verified PR state MERGED at 2026-08-19T18:52:51Z. `git fetch origin main`; `git merge-base --is-ancestor 9c7f328716d2d2daee5c327966740c8c275b0a45 origin/main` succeeded; `origin/main` tip is that squash commit; remote head branch gone (empty `git ls-remote --heads origin outpost/13-managed-dolt-schema`). No `map.md` in this effort directory to update. Ticket 13 already `resolved`.

## Answer

Squash-merged https://github.com/jonbaldie/gascity/pull/24 into `main` on `jonbaldie/gascity` without `--admin` or a direct push.

- **PR:** https://github.com/jonbaldie/gascity/pull/24 (`MERGED` at 2026-08-19T18:52:51Z)
- **Merge commit:** `9c7f328716d2d2daee5c327966740c8c275b0a45` (`git merge-base --is-ancestor` vs `origin/main` succeeded; `origin/main` is that commit)
- **Target:** `main` (ruleset `21046899`: required_approving_review_count=0, no required status checks; classic branch-protection API 404)
- **Head SHA at merge:** `dc45320d1f05ca06c7ceab2596dbe482e73d1ac5` (`--match-head-commit`, `--squash`, `--delete-branch`)
- **Branch result:** `outpost/13-managed-dolt-schema` deleted on origin (empty `git ls-remote --heads`)
- **GitHub issues:** none linked (local tracker ticket 13 already resolved)
