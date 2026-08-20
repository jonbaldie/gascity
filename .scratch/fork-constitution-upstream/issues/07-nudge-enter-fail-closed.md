# 07 — Fail closed when session nudge does not land Enter

**What to build:** `gc session nudge` no longer silently strands text on an agent's input line. If the submit/Enter step does not land (or the line was not clear), the command fails visibly and/or retries so operators and health surfaces see delivery truth.

**Blocked by:** None — can start immediately.

**Status:** resolved

**Upstream:** https://github.com/gastownhall/gascity/issues/5192

- [x] Nudge delivery either confirms submit landed or returns a non-silent failure
- [x] Mid-turn / busy-pane cases do not leave stranded unsubmitted text as a success path
- [x] Test or integration coverage at the highest available seam for the delivery/confirm path
- [x] Dedicated branch + PR against `jonbaldie/gascity` for this ticket only

## Answer

Landed on branch `outpost/07-nudge-enter-fail-closed` as commit `f663dcf4`.

- `NudgeSession` clears pending input with `C-u` before paste.
- For claude/codex, `submitEnterAndConfirm` re-sends Enter only while idle and requires a busy indicator; otherwise returns `ErrNudgeSubmitUnconfirmed` (fail closed).
- Coverage: unit tests in `nudge_submit_confirm_test.go`; tmux seam tests in `nudge_submit_confirm_integration_test.go`.

**PR:** https://github.com/jonbaldie/gascity/pull/6

## Comments

- 2026-08-19 worker: claimed and implemented fail-closed confirm path adapted from upstream #5012/#5013 shape for this fork.
