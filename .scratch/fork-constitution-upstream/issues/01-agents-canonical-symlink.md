# 01 — AGENTS.md canonical; CLAUDE.md symlink

**What to build:** Contributors and agent tools see a single constitution. `AGENTS.md` is the only real file; `CLAUDE.md` is a symlink to it so tools that look for `CLAUDE.md` still work. Any unique content that only existed in `CLAUDE.md` is folded into `AGENTS.md` first. Open a PR to this fork (`jonbaldie/gascity`) as soon as the change is ready — do not wait for other tickets.

**Blocked by:** None — can start immediately.

**Status:** completed

- [x] `AGENTS.md` is the sole editable constitution (no divergent duplicate body in `CLAUDE.md`)
- [x] `CLAUDE.md` is a git-friendly symlink to `AGENTS.md`
- [x] Beads / agent-skills content that must survive is present via `AGENTS.md` after the change
- [x] Dedicated branch + PR opened against `jonbaldie/gascity` for this ticket only

## Comments

- Claimed and finished by outpost worker.
- Branch: `outpost/01-agents-canonical`
- Commit: `ab57a74a` — Make AGENTS.md the sole constitution; symlink CLAUDE.md to it.
- PR: https://github.com/jonbaldie/gascity/pull/1
- Note: old `CLAUDE.md` was `@AGENTS.md` plus a Beads block; `AGENTS.md` already had the Beads block (superset, including reconciler-debugging rule). No fold-in edit required before the symlink.
