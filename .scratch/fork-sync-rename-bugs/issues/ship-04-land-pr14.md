# ship-04 — Land PR #14 (Beads → jonbaldie/beads)

**What to build:** Merge https://github.com/jonbaldie/gascity/pull/14 into `main` on `jonbaldie/gascity`. Follow `/ship-pr`'s SKILL.md rigorously and `.scratch/fork-constitution-upstream/ship-main-protection.md` (no `--admin`, no direct push to `main`).

**Blocked by:** Ticket 04 completed (PR open).

**Status:** resolved

- [x] PR #14 merged under protected-main ruleset (or blocked with clear reason)

## Answer

Squash-merged https://github.com/jonbaldie/gascity/pull/14 into `jonbaldie/gascity` `main` via `gh pr merge --squash --match-head-commit 8306001241d7bb0bb8b33b068134095eff8cb25a --delete-branch` (no `--admin`).

- Merge commit: `58856ae2ce45392710953ec62dc1e82c741b236e` (`Retarget Beads CLI install identity to jonbaldie/beads. (#14)`)
- Target: `origin/main` (tip is that commit; merge commit is an ancestor)
- Head branch `outpost/04-beads-fork-dep` deleted on origin (git ref 404)
- No GitHub closing-issue links on the PR; none to close
- No `map.md` in this effort directory to update

## Comments

- 2026-08-19 runner: claimed. Shipping https://github.com/jonbaldie/gascity/pull/14 via squash per repo convention (PRs #1–#13) and `/ship-pr`. Recorded: repo `jonbaldie/gascity` (origin; never gastownhall), base `main` @ `3718a4fe69e481fadef831f6c1d679586ce5459e`, head `outpost/04-beads-fork-dep` @ `8306001241d7bb0bb8b33b068134095eff8cb25a`, not draft, MERGEABLE, mergeStateStatus UNSTABLE (optional CI still queued/running), reviewDecision empty (ruleset required_approving_review_count=0), no host-native closing issue links, ruleset `21046899` has no required status checks (classic branch protection 404). Not using `--admin`.
- 2026-08-19 runner: merged and verified. PR state MERGED at 2026-08-19T17:12:57Z. `git fetch origin main`; `git merge-base --is-ancestor 58856ae2ce45392710953ec62dc1e82c741b236e origin/main` succeeded; remote head branch gone.
