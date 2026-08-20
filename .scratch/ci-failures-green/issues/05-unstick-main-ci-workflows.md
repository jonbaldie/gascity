# 05 — Unstick & Fix Pending Workflows on Main (CI & Review Formulas)

**What to build:** Investigate and resolve why `CI` and `Review Formulas` workflow runs are stuck pending on `origin/main`. Re-trigger or fix runner/dispatch issues.

**Blocked by:** 01 — Diagnose & Map Root Causes for Main CI Failures.

**Status: in-progress**

- [ ] Stale historical runs `32265272842` (CI) and `32265272839` (Review Formulas) and other stuck queued runs cancelled with `gh run cancel` to unblock concurrency queues.
- [ ] Concurrency settings in `.github/workflows/ci.yml` and `.github/workflows/review-formulas.yml` checked to ensure `cancel-in-progress: true` on push or safe concurrency keys so orphaned push runs never block subsequent runs.
- [ ] Verified that pushes to `main` trigger and run to completion without wedging.

## Comments

### Root Cause Details
- Concurrency groups `ci-push-refs/heads/main` and `review-formulas-push-refs/heads/main` have `cancel-in-progress: false` on push events.
- Historical push run `32265272842` (PR #1) used an old `runner_policy.py` that selected `blacksmith-*-ubuntu-2404` runners, which do not exist in this fork repo.
- That run has been queued for 17 hours at the head of the concurrency queue, holding all subsequent push runs (including latest commit on main) in `pending`.
- Fix: Cancel wedged runs with `gh run cancel` and ensure workflow concurrency rules and runner policies are resilient.
