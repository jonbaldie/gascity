# ship-05 — Land PR #16 (durable launchd supervisor stop)

**What to build:** Merge https://github.com/jonbaldie/gascity/pull/16 into `main` on `jonbaldie/gascity`. Follow `/ship-pr`'s SKILL.md rigorously and `.scratch/fork-constitution-upstream/ship-main-protection.md` (no `--admin`, no direct push to `main`).

**Blocked by:** Ticket 05 completed (PR open).

**Status:** resolved

- [x] PR #16 merged under protected-main ruleset (or blocked with clear reason)

## Answer

Squash-merged https://github.com/jonbaldie/gascity/pull/16 into `jonbaldie/gascity` `main` via `gh pr merge --squash --match-head-commit d131a053eacd91e1a7fd4479d827a28b84ce7616 --delete-branch` (no `--admin`).

- Merge commit: `52ba42c7a76532fb0e564d0b3ec3c02409f39b90` (`fix(supervisor): fail closed when launchd stop is not durable (#16)`)
- Target: `origin/main` (tip is that commit; merge commit is an ancestor)
- Head branch `outpost/05-supervisor-stop-launchd-durable` deleted on origin (git ref 404)
- No GitHub closing-issue links on the PR; none to close
- No `map.md` in this effort directory to update

## Comments

- 2026-08-19: Prefect: runner cap is 1; ship-06 (PR #15) landed. Dispatch this next.
- 2026-08-19 runner: claimed. Shipping https://github.com/jonbaldie/gascity/pull/16 via squash per repo convention (PRs #7–#15) and `/ship-pr`. Recorded: repo `jonbaldie/gascity` (origin; never gastownhall), base `main`, head `outpost/05-supervisor-stop-launchd-durable` @ `d131a053eacd91e1a7fd4479d827a28b84ce7616`, not draft, MERGEABLE, mergeStateStatus UNSTABLE, reviewDecision empty (ruleset required_approving_review_count=0), no host-native closing issue links, ruleset `21046899` has no required status checks (classic branch protection 404; GraphQL `isRequired` count=0). Optional CI failures on this SHA: Preflight / static checks (`scripts/cipolicy` TestCurrentWorkflowsMatchPolicy / want exact provider field path) and cmd/gc product metrics testhook (also failed on merged PR #14; isRequired=false). Preflight / generated artifacts succeeded here. Not using `--admin`.
- 2026-08-19 runner: merged and verified. PR state MERGED at 2026-08-19T17:42:29Z. `git fetch origin main`; `git merge-base --is-ancestor 52ba42c7a76532fb0e564d0b3ec3c02409f39b90 origin/main` succeeded; remote head branch gone.
