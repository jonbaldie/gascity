# 19 — Remaining host leftovers after patrol 16/18

**What to build:** Finish host hygiene the first cleanup pass skipped: dirty constitution worktrees `gascity-wt-06`…`13`, shipped leftovers `gascity-wt-24` and `gascity-wt-25` (PRs #21 and #22), plus now-shipped `gascity-wt-26` (PR #23) and `gascity-wt-27` (PR #24). Do not touch `gascity-wt-28` (live ticket 14) or the dirty main checkout.

**Blocked by:** 18 — Execute host leftover cleanup from patrol 16.

**Status:** resolved

- [x] `gascity-wt-06` through `gascity-wt-13` removed. Prefect confirms the untracked `.outpost-steer.md` / `.outpost-ship-policy.md` / constitution scratch files are harness copies — disposable; `git worktree remove --force` is allowed for those eight trees
- [x] `gascity-wt-24` and `gascity-wt-25` removed; local `outpost/11-reset-pending-evict` and `outpost/15-main-ci-green` deleted
- [x] `gascity-wt-26` and `gascity-wt-27` removed; local `outpost/12-dolt-compact-concurrent-inserts` and `outpost/13-managed-dolt-schema` deleted
- [x] `gascity-wt-28` and the main checkout left alone
- [x] No `docker system prune`; no gastown/live-stack deletes; local `main` not fast-forwarded

## Comments

- 2026-08-19: Prefect telescoped from patrol 17. Ticket 18 skipped dirty trees (no `--force`) and kept 24/25 while PRs were open. Those PRs are now merged.
- 2026-08-19: Prefect: PRs #23 and #24 are on `main`. Expand cleanup to wt-26/27. Untracked `.outpost-*` files on wt-06…13 are prefect copies — `--force` is OK. Leave wt-28 (ticket 14 live).
- 2026-08-19: Worker claimed. Host hygiene from main checkout; no feature branch/PR; no Go edits. Aborted a raced ticket-14 attempt first (reverted uncommitted test edits in `gascity-wt-28`; no PR).
- 2026-08-19: Worker finished. 12 worktrees removed; 12 local branches deleted. See Answer.

## Answer

Host cleanup from `/Users/jonathanbaldie/go/src/github.com/jonbaldie/gascity`, 2026-08-19. No feature branch, no PR, no Go edits, no `docker system prune`, no gastown/live-stack deletes, local `main` not fast-forwarded (`1895c64e0`, still dirty `AGENTS.md` / `.scratch/` / `.outpost.json` / `docs/agents/`).

Before this pass, raced ticket-14 edits in `gascity-wt-28` (`cmd/gc/cmd_dolt_cleanup_test.go`, `cmd/gc/dolt_cleanup_discovery_test.go`) were reverted with `git checkout --`. Working tree left clean. No PR for 14.

**Removed worktrees + local branches**

| Worktree | Branch | How |
|---|---|---|
| `gascity-wt-06` | `outpost/06-gc-init-issue-prefix` | `worktree remove --force` + `branch -D` |
| `gascity-wt-07` | `outpost/07-nudge-enter-fail-closed` | `worktree remove --force` + `branch -D` |
| `gascity-wt-08` | `outpost/08-supervisor-stop-dolt-hang` | `worktree remove --force` + `branch -D` |
| `gascity-wt-09` | `outpost/09-pool-assignee-identity` | `worktree remove --force` + `branch -D` |
| `gascity-wt-10` | `outpost/10-acp-drain-process-leak` | `worktree remove --force` + `branch -D` |
| `gascity-wt-11` | `outpost/11-macos-supervisor-fd-leak` | `worktree remove --force` + `branch -D` |
| `gascity-wt-12` | `outpost/12-adhoc-pool-origin-gate` | `worktree remove --force` + `branch -D` |
| `gascity-wt-13` | `outpost/13-host-gc-cruft-cleanup` | `worktree remove --force` + `branch -d` |
| `gascity-wt-24` | `outpost/11-reset-pending-evict` | `worktree remove` + `branch -D` |
| `gascity-wt-25` | `outpost/15-main-ci-green` | `worktree remove` + `branch -D` |
| `gascity-wt-26` | `outpost/12-dolt-compact-concurrent-inserts` | `worktree remove` + `branch -D` |
| `gascity-wt-27` | `outpost/13-managed-dolt-schema` | `worktree remove` + `branch -D` |

**Kept**

- Main checkout `/Users/jonathanbaldie/go/src/github.com/jonbaldie/gascity` (`main` @ `1895c64e0`, behind origin, dirty as found)
- `gascity-wt-28` `outpost/14-macos-dolt-cleanup-ps` @ `eae28a8e6` (live ticket 14)

**Remaining worktrees after this pass:** main + `gascity-wt-28` only. Remaining local `outpost/*` branch: `outpost/14-macos-dolt-cleanup-ps`.
