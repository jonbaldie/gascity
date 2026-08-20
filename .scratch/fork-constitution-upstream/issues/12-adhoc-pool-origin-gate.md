# 12 — Demand-spawned adhoc pool seats receive routed work

**What to build:** Adhoc pool seats created by demand-spawn are eligible for routed work (not origin-gated out). Scale-from-zero / demand spawn results in agents that can actually take the work that caused the spawn.

**Blocked by:** None — can start immediately.

**Status:** resolved

**Upstream:** https://github.com/gastownhall/gascity/issues/5009

- [x] Demand-spawned adhoc seat is not excluded from routed work solely by an origin gate that permanent seats pass
- [x] Minimal repro: spawn-from-demand → work is claimable/routable to that seat
- [x] Regression test at the dispatcher/origin-gate seam
- [x] Dedicated branch + PR against `jonbaldie/gascity` for this ticket only

## Answer

Root cause: city-scoped Huma `POST .../sessions` with `kind=agent` stamped `session_origin=manual`, while the legacy agent create path and controller demand-create stamp `ephemeral`. The default work_query origin gate only admits `ephemeral` or empty, so adhoc pool seats drained without claiming `gc.routed_to=<base route>` work.

Fix: align Huma agent create to `session_origin=ephemeral`. Provider create remains `manual`. Origin-gate regression tests cover ephemeral/empty admit and manual/named exclude.

- Branch: `outpost/12-adhoc-pool-origin-gate`
- PR: https://github.com/jonbaldie/gascity/pull/11
- Commit: `d9fa5648`

## Comments

- 2026-08-19: Claimed and fixed in worktree `gascity-wt-12`. PR opened; not merged (ship policy / prefect).
