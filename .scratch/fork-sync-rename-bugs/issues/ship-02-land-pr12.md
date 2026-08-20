# ship-02 — Land PR #12 (upstream merge)

**What to build:** Merge https://github.com/jonbaldie/gascity/pull/12 into `main` on `jonbaldie/gascity`. Follow `/ship-pr`'s SKILL.md rigorously and `.scratch/fork-constitution-upstream/ship-main-protection.md` (no `--admin`, no direct push to `main`).

**Blocked by:** Ticket 02 completed (PR open).

**Status:** resolved

- [x] PR #12 merged under protected-main ruleset (or blocked with clear reason)

## Answer

Squash-merged https://github.com/jonbaldie/gascity/pull/12 into `jonbaldie/gascity` `main` via `gh pr merge --squash --match-head-commit a209d41c9b46c7faba74775e82ea3e8bc4a47790 --delete-branch` (no `--admin`).

- Merge commit: `6e8178493e9ea4c4b4d2221ddbbc375dd02bb636` (`Merge gastownhall/gascity main into this fork (#12)`)
- Target: `origin/main` (tip is that commit; merge commit is an ancestor)
- Head branch `outpost/02-merge-upstream` deleted on origin (git ref 404)
- No GitHub closing-issue links on the PR; none to close
- No `map.md` in this effort directory to update

## Comments

- 2026-08-19 runner: claimed. Shipping https://github.com/jonbaldie/gascity/pull/12 via squash per repo convention (PRs #1–#11) and `/ship-pr`. Recorded: repo `jonbaldie/gascity`, base `main`, head `outpost/02-merge-upstream` @ `a209d41c9b46c7faba74775e82ea3e8bc4a47790`, not draft, MERGEABLE, reviewDecision empty (ruleset required_approving_review_count=0), no host-native closing issue links, ruleset `21046899` has no required status checks (classic branch protection 404). Not using `--admin`.
- 2026-08-19 runner: merged and verified. PR state MERGED at 2026-08-19T16:33:45Z. `git fetch origin main`; `git merge-base --is-ancestor 6e8178493e9ea4c4b4d2221ddbbc375dd02bb636 origin/main` succeeded; remote head branch gone.
