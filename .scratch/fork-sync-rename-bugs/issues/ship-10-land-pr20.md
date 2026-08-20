# ship-10 — Land PR #20 (formula-step first prompt)

**What to build:** Merge https://github.com/jonbaldie/gascity/pull/20 into `main` on `jonbaldie/gascity`. Follow `/ship-pr`'s SKILL.md rigorously and `.scratch/fork-constitution-upstream/ship-main-protection.md` (no `--admin`, no direct push to `main`).

**Blocked by:** Ticket 10 completed (PR open).

**Status:** resolved

- [x] PR #20 merged under protected-main ruleset (or blocked with clear reason)

## Answer

Squash-merged https://github.com/jonbaldie/gascity/pull/20 into `jonbaldie/gascity` `main` via `gh pr merge --squash --match-head-commit 63e31b4690053e8cd32c9ae4afc0335e96500a45 --delete-branch` (no `--admin`).

- Merge commit: `468493bbe5836c3e37ce5c99042528654887e358` (`fix(session): prime reconciler-spawned formula-step sessions (#20)`)
- Target: `origin/main` (tip is that commit; merge commit is an ancestor)
- Head branch `outpost/10-formula-step-first-prompt` deleted on origin (git ref 404)
- No GitHub closing-issue links on the PR; none to close
- No `map.md` in this effort directory to update

## Comments

- 2026-08-19 runner: claimed. Shipping https://github.com/jonbaldie/gascity/pull/20 via squash per repo convention (PRs #7–#19) and `/ship-pr`. Recorded: repo `jonbaldie/gascity` (origin; never gastownhall), base `main`, head `outpost/10-formula-step-first-prompt` @ `63e31b4690053e8cd32c9ae4afc0335e96500a45`, not draft, MERGEABLE, mergeStateStatus UNSTABLE (optional checks queued/pending), reviewDecision empty (ruleset required_approving_review_count=0), no host-native closing issue links, ruleset `21046899` has no required status checks (classic branch protection 404; `gh pr checks --required`: none). Optional `scripts/cipolicy` / Runner policy jobs are not required. Not using `--admin`.
- 2026-08-19 runner: merged and verified. PR state MERGED at 2026-08-19T18:17:44Z. `git fetch origin main`; `git merge-base --is-ancestor 468493bbe5836c3e37ce5c99042528654887e358 origin/main` succeeded; remote head branch gone.
