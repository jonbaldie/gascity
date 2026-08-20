# 10 — Reconciler-spawned formula steps receive a first prompt

**What to build:** When the session reconciler spawns a session for a ready formula step, that session receives a first work prompt without a human `gc session nudge`. Autonomous formula runs must not stall at every step because routing was stamped at bead creation and nudge-on-route only watches later updates.

**Blocked by:** 01, 02, 03, 04

**Status:** resolved

**Upstream:** https://github.com/gastownhall/gascity/issues/4382

- [x] A reconciler-spawned formula-step session is primed (first prompt enqueued/delivered) without an external sling `--nudge` or human nudge
- [x] Step beads whose `gc.routed_to` is set at creation still get a nudge/prime once the target session exists
- [x] Tests cover the created-already-routed vs `bead.updated`-only trigger gap
- [x] Focused tests ran in resource-capped Docker (`--cpus=2 --memory=4g --pids-limit=256`), not a broad host `make test`
- [x] Dedicated branch + PR against `jonbaldie/gascity` for this ticket only

## Answer

Cause: formula steps stamp `gc.routed_to` at `bead.created`. Core pack `nudge-on-route` watched `bead.updated` only, so those beads never woke a worker. The reconciler later spawned the step session with no first work prompt.

Fix on `outpost/10-formula-step-first-prompt`:

- Event-trigger `on` accepts comma-separated types (OR). `nudge-on-route` is `bead.updated,bead.created`; the script queries both.
- After a successful reconciler spawn with a trigger bead, enqueue a kickoff nudge (`source=spawn`, stable id `kickoff-<session>-<trigger>`). Duplicate enqueue is a no-op. Mechanical: trigger bead present → first prompt; no trigger → silent.

Landed:

- Branch: `outpost/10-formula-step-first-prompt`
- Commit: `63e31b4690053e8cd32c9ae4afc0335e96500a45`
- PR: https://github.com/jonbaldie/gascity/pull/20
- Worktree: `/Users/jonathanbaldie/go/src/github.com/jonbaldie/gascity-wt-23`
- Docker: `golang:1.26.6`, `CGO_ENABLED=0`, `--cpus=2 --memory=4g --pids-limit=256` — focused `go test ./internal/orders/ ./internal/bootstrap/packs/core/ ./cmd/gc/` ok

## Comments

- Worker claimed 2026-08-19. Branch `outpost/10-formula-step-first-prompt` in worktree `gascity-wt-23`. Implementing reconciler spawn kickoff + nudge-on-route `bead.created` (created-already-routed gap).
- 2026-08-19: Resolved. PR #20. Tests cover created-already-routed vs `bead.updated`-only, plus reconciler spawn kickoff enqueue.
