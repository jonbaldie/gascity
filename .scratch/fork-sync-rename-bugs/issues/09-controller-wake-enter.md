# 09 — Controller wake delivery must confirm Enter or fail closed

**What to build:** The controller heartbeat / wake path that types into a live pane either confirms the wake submitted (Enter landed) or fails closed / retries. Mid-turn landings must not leave unsubmitted wake text in the input box while the session looks idle and healthy. Autonomous patrol cadence must not depend on a second agent noticing stranded text.

**Blocked by:** 01, 02, 03, 04

**Status:** resolved

**Upstream:** https://github.com/gastownhall/gascity/issues/4935

Distinct from shipped PR #6 / `#5192` (`gc session nudge` submit confirm). Upstream evidence: nudge recovers a dropped wake; the wake/keystroke path is the defect. No upstream fix PR at research time.

- [x] Controller wake/heartbeat delivery confirms submit or returns a non-silent failure (no “success” with stranded input)
- [x] Mid-turn / busy-pane wake landings do not leave exactly-one-copy unsubmitted wake text as the happy path
- [x] Tests cover the wake keystroke/submit seam at the highest existing test layer
- [x] Focused tests ran in resource-capped Docker (`--cpus=2 --memory=4g --pids-limit=256`), not a broad host `make test`
- [x] Dedicated branch + PR against `jonbaldie/gascity` for this ticket only

## Answer

Cause: `SendKeysDebounced` treated a tmux-accepted Enter as delivery. Mid-turn landings typed wake text, dropped Enter, and still returned nil. `NudgeSession`'s later busy poll could also false-confirm the *original* turn as this submit.

Fix on `outpost/09-controller-wake-enter`:

- Eligible panes (readable busy indicator): refuse to type when already busy (`ErrSubmitPaneBusy`); after paste, confirm Enter via `submitEnterAndConfirm` or return `ErrNudgeSubmitUnconfirmed`.
- Ineligible panes (shell / unknown): unchanged best-effort single Enter.
- Same busy-before-type guard on `NudgeSession` and `NudgeNow` so the hidden-attach keystroke path cannot bypass it.

Landed:

- Branch: `outpost/09-controller-wake-enter`
- Commit: `c868eaff2`
- PR: https://github.com/jonbaldie/gascity/pull/19
- Docker: `golang:1.26.6`, `CGO_ENABLED=0`, `go test ./internal/runtime/tmux -count=1` — ok (9.675s)

## Comments

- Worker claimed 2026-08-19. Worktree `/Users/jonathanbaldie/go/src/github.com/jonbaldie/gascity-wt-22`.
