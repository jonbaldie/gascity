# ship-07 — Land PR #7 (supervisor forced-stop timeout)

**What to build:** Merge https://github.com/jonbaldie/gascity/pull/7 into `main` on `jonbaldie/gascity`. Follow `/ship-pr`'s SKILL.md rigorously and obey `ship-main-protection.md` (no `--admin`, no direct push). Soft-wait behind ship-06 if a runner is already in flight (`maxRunners: 1`).

**Blocked by:** Ticket 08 completed (PR open). Soft-wait: ship-06 ahead.

**Status:** completed

- [x] PR #7 merged under the protected-main ruleset (or blocked with a clear reason)

## Comments

- Claimed by outpost runner for land of https://github.com/jonbaldie/gascity/pull/7
- Resolve record: repo `jonbaldie/gascity`, base `main`, head `outpost/08-supervisor-stop-dolt-hang` @ `4c8684623790498d2656666104200bd4881a78fd`, draft=false, reviewDecision empty (ruleset required_approving_review_count=0), closingIssuesReferences=[]
- Required checks (ruleset 21046899): none — rules are deletion, non_fast_forward, pull_request only. CI rollup jobs are optional/pending; gate only on ruleset.
- Squash-merged via `gh pr merge --squash --match-head-commit 4c8684623790498d2656666104200bd4881a78fd --delete-branch` (no `--admin`) at 2026-08-19T15:00:15Z.
- Merge commit: `0807fe4bb55cd21dbabbdfba53245594d1005c1b` reachable from `origin/main`.
- Head branch `outpost/08-supervisor-stop-dolt-hang` deleted (404 on ref).
- PR: https://github.com/jonbaldie/gascity/pull/7

## Answer

Landed PR #7 on `main` under ruleset `main` (id 21046899). Merge commit `0807fe4bb55cd21dbabbdfba53245594d1005c1b`.
