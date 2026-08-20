# 15 — Make fork `main` CI green

**What to build:** Pushes to `jonbaldie/gascity` `main` do not leave a red GitHub Actions status for workflows this fork can actually run. In particular, **Notify Image Rebuilds** must not fail on every `cmd/`/`internal/` push. Other red jobs on the latest `main` CI run that are caused by this fork (identity rename, missing upstream secrets, stale docs) are fixed in the same PR if they are still failing; do not paper over real product-test failures.

**Blocked by:** None — can start immediately (ops-priority; dispatch when a worker slot is free).

**Status:** resolved

**Evidence (2026-08-19):** every recent `main` merge reds **Notify Image Rebuilds** (e.g. https://github.com/jonbaldie/gascity/actions/runs/32283339009). Job `notify` exits 4: `GH_TOKEN` is empty (`secrets.GASCITY_HOSTED_TOKEN` unset on this fork) while the workflow POSTs `repository_dispatch` to an upstream image-host repo this fork does not control. Rapid successive merges also **cancel** in-flight `CI` runs on `main`; that is a symptom of ship cadence, not the primary red X.

- [x] Notify Image Rebuilds does not fail on `jonbaldie/gascity` (skip/no-op when the hosted token or target repo is unavailable; do not invent secrets or dispatch to a repo this fork cannot write)
- [x] Latest `main` (or this ticket’s PR) has no remaining red *fork-caused* jobs; real test failures are fixed, not skipped
- [x] Focused checks ran in resource-capped Docker where they are Go tests; workflow YAML changes are verified by `gh`/actionlint or equivalent if present
- [x] Dedicated branch + PR against `jonbaldie/gascity` for this ticket only

## Answer

Notify Image Rebuilds no-ops when `GASCITY_HOSTED_TOKEN` is unset or `gascity/gasworks-control-plane` is not writable; `scripts/cipolicy` hash pinned for the `jonbaldie/beads` clone URLs. Generated `docs/reference/cli.md` was already fresh.

- Branch: `outpost/15-main-ci-green`
- PR: https://github.com/jonbaldie/gascity/pull/22
- Commit: `b2493445f`

## Comments

- 2026-08-19: Prefect queued from human “CI failure on main”. Workflow: `.github/workflows/notify-image-build.yaml`. Runners already noted optional PR failures (`scripts/cipolicy`, stale `docs/reference/cli.md`) — include those if they still fail this fork’s CI.
- 2026-08-19: Prefect re-dispatch. Prior worker was interrupted mid-diagnosis (no commits). Worktree `gascity-wt-25` / `outpost/15-main-ci-green` fast-forwarded to `origin/main` (includes PR #20). Continue; ticket stays claimed.
- 2026-08-19: Worker finished. Primary red was empty `GH_TOKEN` (Actions exit 4). Remaining fork-caused optional red was cipolicy execution hash after PR #14 beads clone retarget. cli.md not stale. Blacksmith not a current main-push red (runner-policy already falls back).
