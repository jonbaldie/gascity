# 09 — Align pool assignee writes with runtime session identity

**What to build:** Pool/cron-dispatched molecules are assigned with an identity string that the live agent will recognize via its runtime env (`GC_SESSION_ID` / name / alias forms). Healthy idle pool agents no longer leave claimed work stranded because of sanitized `bd__…` assignee mismatch.

**Blocked by:** None — can start immediately.

**Status:** resolved

**Upstream:** https://github.com/gastownhall/gascity/issues/5048 · related PR https://github.com/gastownhall/gascity/pull/5372

- [x] Assignee written at dispatch matches at least one identity form the session exports
- [x] A pool agent with matching runtime identity can claim/see the dispatched work without manual reassignment
- [x] Regression test covers sanitized vs runtime identity forms
- [x] Dedicated branch + PR against `jonbaldie/gascity` for this ticket only

## Answer

Root cause: `setTemplateEnvIdentity` stamped `GC_ALIAS`/`GC_AGENT` with the stable pool identity while leaving `BEADS_ACTOR` on the sanitized session name from `resolveTemplate`. `bd update --claim` writes `BEADS_ACTOR` into `work.Assignee`, so claims landed under `bd__dog-<beadID>` forms that later respawns no longer export as `$GC_ALIAS`.

Fix: align `BEADS_ACTOR` with the alias-first runtime identity at template stamp + `RuntimeEnvWithSessionContext`, preserve stamped identity across start merge, and match named-session demand on runtime session-name assignees.

- Branch: `outpost/09-pool-assignee-identity`
- PR: https://github.com/jonbaldie/gascity/pull/9
- Commit: `f7354f22`

## Comments

- 2026-08-19: outpost worker claimed and completed; PR opened, not merged (ship policy).
