# ship-09 — Land PR #19 (controller wake Enter fail-closed)

**What to build:** Merge https://github.com/jonbaldie/gascity/pull/19 into `main` on `jonbaldie/gascity`. Follow `/ship-pr`'s SKILL.md rigorously and `.scratch/fork-constitution-upstream/ship-main-protection.md` (no `--admin`, no direct push to `main`).

**Blocked by:** Ticket 09 completed (PR open).

**Status:** resolved

- [x] PR #19 merged under protected-main ruleset (or blocked with clear reason)

## Answer

Squash-merged https://github.com/jonbaldie/gascity/pull/19 into `jonbaldie/gascity` `main` via `gh pr merge --squash --match-head-commit c868eaff298e43e36631e1cf7a63cef3df29ddc6 --delete-branch` (no `--admin`).

- Merge commit: `f7c40d0c99a0714f9face62f62c08fc8ace648ed` (`fix(tmux): confirm wake keystroke submit or fail closed (#19)`)
- Target: `origin/main` (tip is that commit; merge commit is an ancestor)
- Head branch `outpost/09-controller-wake-enter` deleted on origin (git ref 404)
- No GitHub closing-issue links on the PR; none to close
- No `map.md` in this effort directory to update

## Comments

- 2026-08-19 runner: claimed. Shipping https://github.com/jonbaldie/gascity/pull/19 via squash per repo convention (PRs #7–#18) and `/ship-pr`. Recorded: repo `jonbaldie/gascity` (origin; never gastownhall), base `main`, head `outpost/09-controller-wake-enter` @ `c868eaff298e43e36631e1cf7a63cef3df29ddc6`, not draft, MERGEABLE, mergeStateStatus UNSTABLE (optional checks queued/pending), reviewDecision empty (ruleset required_approving_review_count=0), no host-native closing issue links, ruleset `21046899` has no required status checks (classic branch protection 404; `gh pr checks --required`: none). Optional `scripts/cipolicy` / Runner policy jobs are not required. Not using `--admin`.
- 2026-08-19 runner: merged and verified. PR state MERGED at 2026-08-19T17:53:45Z. `git fetch origin main`; `git merge-base --is-ancestor f7c40d0c99a0714f9face62f62c08fc8ace648ed origin/main` succeeded; remote head branch gone.
