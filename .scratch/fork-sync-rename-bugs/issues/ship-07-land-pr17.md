# ship-07 — Land PR #17 (gc start reload timeout)

**What to build:** Merge https://github.com/jonbaldie/gascity/pull/17 into `main` on `jonbaldie/gascity`. Follow `/ship-pr`'s SKILL.md rigorously and `.scratch/fork-constitution-upstream/ship-main-protection.md` (no `--admin`, no direct push to `main`).

**Blocked by:** Ticket 07 completed (PR open).

**Status:** resolved

- [x] PR #17 merged under protected-main ruleset (or blocked with clear reason)

## Answer

Squash-merged https://github.com/jonbaldie/gascity/pull/17 into `jonbaldie/gascity` `main` via `gh pr merge --squash --match-head-commit d0edc2774335fdb59349e3a42966182fda8830c0 --delete-branch` (no `--admin`).

- Merge commit: `748b72c4e7cbfeb4cc59fa8238af99b0aea93c56` (`fix(start): treat reload timeout as async reconcile (#17)`)
- Target: `origin/main` (tip is that commit; merge commit is an ancestor)
- Head branch `outpost/07-gc-start-reload-timeout` deleted on origin (git ref 404)
- No GitHub closing-issue links on the PR; none to close
- No `map.md` in this effort directory to update

## Comments

- 2026-08-19: Prefect: land after ship-05 (PR #16) if a runner is already in flight (maxRunners 1).
- 2026-08-19 runner: claimed. Shipping https://github.com/jonbaldie/gascity/pull/17 via squash per repo convention (PRs #7–#16) and `/ship-pr`. Recorded: repo `jonbaldie/gascity` (origin; never gastownhall), base `main`, head `outpost/07-gc-start-reload-timeout` @ `d0edc2774335fdb59349e3a42966182fda8830c0`, not draft, MERGEABLE, mergeStateStatus UNSTABLE, reviewDecision empty (ruleset required_approving_review_count=0), no host-native closing issue links, ruleset `21046899` has no required status checks (classic branch protection 404; GraphQL `isRequired` count=0). Optional CI failure on this SHA: Preflight / static checks (`CI workflow policy` / `scripts/cipolicy`; isRequired=false). Same optional cipolicy failure landed on merged PRs #14–#16. Docs step skipped after that fail. Not using `--admin`.
- 2026-08-19 runner: merged and verified. PR state MERGED at 2026-08-19T17:44:23Z. `git fetch origin main`; `git merge-base --is-ancestor 748b72c4e7cbfeb4cc59fa8238af99b0aea93c56 origin/main` succeeded; remote head branch gone.
