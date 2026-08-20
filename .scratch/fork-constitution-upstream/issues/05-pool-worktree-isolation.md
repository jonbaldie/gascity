# 05 — Isolate pool-agent worktrees on parallel dispatch

**What to build:** When more than one pool agent is dispatched in parallel against the same rig, each agent works in an isolated worktree (or equivalent isolation) so git branch switches / resets / stashes cannot clobber another agent's uncommitted edits.

**Blocked by:** None — can start immediately.

**Status:** resolved

**Upstream:** https://github.com/gastownhall/gascity/issues/1181

- [x] Two parallel pool dispatches on one rig do not share a single mutable working tree for concurrent git mutations
- [x] Failure or incomplete isolation is detectable (test or doctor/audit path), not silent
- [x] No hardcoded role names in Go; pool/worktree behaviour stays config-driven
- [x] Dedicated branch + PR against `jonbaldie/gascity` for this ticket only

## Answer

Implicit rig-scoped pool agents now default to per-instance worktrees via `work_dir` + `gc internal ensure-worktree` PreStart. Doctor check `pool-workdir-isolation` warns when multi-session rig agents still share one path.

- Branch: `outpost/05-pool-worktree-isolation`
- Commit: `a1b17fb8`
- PR: https://github.com/jonbaldie/gascity/pull/4

## Comments

- Claimed and finished by outpost worker on `gascity-wt-05`.
