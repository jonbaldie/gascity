# ship-06 — Land PR #15 (gc stop busy reload)

**What to build:** Merge https://github.com/jonbaldie/gascity/pull/15 into `main` on `jonbaldie/gascity`. Follow `/ship-pr`'s SKILL.md rigorously and `.scratch/fork-constitution-upstream/ship-main-protection.md` (no `--admin`, no direct push to `main`).

**Blocked by:** Ticket 06 completed (PR open).

**Status:** resolved

- [x] PR #15 merged under protected-main ruleset (or blocked with clear reason)

## Answer

Squash-merged https://github.com/jonbaldie/gascity/pull/15 into `jonbaldie/gascity` `main` via `gh pr merge --squash --match-head-commit 1d1c29711a993fc33e215bc58a7cc7cabcd58b2f --delete-branch` (no `--admin`).

- Merge commit: `be4f9f452c11500716a5408b1d7de8eff30e2eee` (`fix(stop): do not restore a city when reload is busy (#15)`)
- Target: `origin/main` (tip is that commit; merge commit is an ancestor)
- Head branch `outpost/06-gc-stop-busy-reload` deleted on origin (git ref 404)
- No GitHub closing-issue links on the PR; none to close
- No `map.md` in this effort directory to update

## Comments

- 2026-08-19 runner: claimed. Shipping https://github.com/jonbaldie/gascity/pull/15 via squash per repo convention (PRs #7–#14) and `/ship-pr`. Recorded: repo `jonbaldie/gascity` (origin; never gastownhall), base `main`, head `outpost/06-gc-stop-busy-reload` @ `1d1c29711a993fc33e215bc58a7cc7cabcd58b2f`, not draft, MERGEABLE, mergeStateStatus UNSTABLE, reviewDecision empty (ruleset required_approving_review_count=0), no host-native closing issue links, ruleset `21046899` has no required status checks (classic branch protection 404; GraphQL `isRequired` count=0). Optional CI failures on this SHA: Preflight / static checks (`scripts/cipolicy` TestCurrentWorkflowsMatchPolicy / want exact provider field path) and Preflight / generated artifacts (`docs/reference/cli.md` stale). Same optional pair failed on merged PR #14. Waited through queued CI rather than prefetch-merge. Not using `--admin`.
- 2026-08-19 runner: merged and verified. PR state MERGED at 2026-08-19T17:40:40Z. `git fetch origin main`; `git merge-base --is-ancestor be4f9f452c11500716a5408b1d7de8eff30e2eee origin/main` succeeded; remote head branch gone.
