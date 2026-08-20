# 02 — Add root CODING_STANDARDS.md (Gas City–adapted)

**What to build:** The repo has a root `CODING_STANDARDS.md` modelled on `jonbaldie/gastown`'s coding standards (same major sections: tests, comments/docs, footguns, Go), rewritten for Gas City vocabulary and this fork's invariants (ZFC, zero hardcoded roles, existing test/make gates). Open a PR to `jonbaldie/gascity` as soon as ready — do not wait for other tickets.

**Blocked by:** None — can start immediately.

**Status:** completed

- [x] `CODING_STANDARDS.md` exists at the repository root
- [x] Structure clearly follows the gastown reference without copying Gas Town role hardcoding
- [x] Standards align with this repo's testing and Go quality expectations (point at existing make/test targets where useful)
- [x] Dedicated branch + PR opened against `jonbaldie/gascity` for this ticket only

## Landed

- **Branch:** `outpost/02-coding-standards`
- **Commit:** `c4daaf0f`
- **PR:** https://github.com/jonbaldie/gascity/pull/2
- **Artifact:** `CODING_STANDARDS.md` at repo root (worktree: `gascity-wt-02`)

## Comments

- Worker claimed and shipped Gas City–adapted coding standards modelled on jonbaldie/gastown; no Gas Town role hardcoding; gates point at make test / sharded TESTING.md targets.
