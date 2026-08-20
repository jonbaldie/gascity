# ship-08 — Land PR #18 (resume keeps option_defaults)

**What to build:** Merge https://github.com/jonbaldie/gascity/pull/18 into `main` on `jonbaldie/gascity`. Follow `/ship-pr`'s SKILL.md rigorously and `.scratch/fork-constitution-upstream/ship-main-protection.md` (no `--admin`, no direct push to `main`).

**Blocked by:** Ticket 08 completed (PR open).

**Status:** resolved

- [x] PR #18 merged under protected-main ruleset (or blocked with clear reason)

## Answer

Squash-merged https://github.com/jonbaldie/gascity/pull/18 into `jonbaldie/gascity` `main` via `gh pr merge --squash --match-head-commit 5ce8e6d9396f11ddcd1b727b838f5f735a903415 --delete-branch` (no `--admin`).

- Merge commit: `604268d1128aba9c2f0012df1ae152d7a78412b1` (`fix(config): keep provider option_defaults when building resume commands (#18)`)
- Target: `origin/main` (tip is that commit; merge commit is an ancestor)
- Head branch `outpost/08-resume-keeps-option-defaults` deleted on origin (git ref 404)
- No GitHub closing-issue links on the PR; none to close
- No `map.md` in this effort directory to update

## Comments

- 2026-08-19 runner: claimed. Shipping https://github.com/jonbaldie/gascity/pull/18 via squash per repo convention (PRs #7–#17) and `/ship-pr`. Recorded: repo `jonbaldie/gascity` (origin; never gastownhall), base `main`, head `outpost/08-resume-keeps-option-defaults` @ `5ce8e6d9396f11ddcd1b727b838f5f735a903415`, not draft, MERGEABLE, mergeStateStatus UNSTABLE (optional checks queued/pending), reviewDecision empty (ruleset required_approving_review_count=0), no host-native closing issue links, ruleset `21046899` has no required status checks (classic branch protection 404; `gh pr checks --required`: none). Optional `scripts/cipolicy` / Runner policy jobs are not required. Not using `--admin`.
- 2026-08-19 runner: merged and verified. PR state MERGED at 2026-08-19T17:51:39Z. `git fetch origin main`; `git merge-base --is-ancestor 604268d1128aba9c2f0012df1ae152d7a78412b1 origin/main` succeeded; remote head branch gone.
