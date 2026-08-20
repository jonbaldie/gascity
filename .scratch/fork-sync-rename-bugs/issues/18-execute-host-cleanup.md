# 18 — Execute host leftover cleanup from patrol 16

**What to build:** Remove merged/abandoned `gascity-wt-01`…`23` worktrees and their local `outpost/*` branches, and delete the leftover `origin/outpost/*` remotes listed in ticket 16’s bark. Host hygiene only — no product code, no `docker system prune`, no gastown/live-stack containers.

**Blocked by:** 16 — Patrol host leftovers (bark is the work list).

**Status:** resolved

- [x] `gascity-wt-01`–`05` and `14`–`23` removed (`git worktree remove`; squash-merged `-D`, ancestors `-d`). `06`–`13` skipped (dirty untracked; no `--force`) — see Answer
- [x] Leftover remotes: five deleted on origin; `outpost/09-controller-wake-enter` and `outpost/10-formula-step-first-prompt` already absent
- [x] `gascity-wt-24`, `gascity-wt-25`, and the main checkout were not removed
- [x] No `docker system prune`; no gastown/1.25/live-stack image or container deletes; `docker rm jolly_payne` + anonymous volume done
- [x] Local `main` was not fast-forwarded (still `1895c64e0`, dirty `AGENTS.md`)

## Comments

- 2026-08-19: Prefect telescoped from ticket 16 bark. Follow `.scratch/fork-sync-rename-bugs/issues/16-patrol-host-leftovers.md` `## Answer` tables exactly.
- 2026-08-19: Worker claimed. Host hygiene from main checkout; no feature branch/PR; no Go edits.
- 2026-08-19: Worker finished. 15 worktrees removed; 8 skipped dirty; 5 remotes deleted; 2 remotes already gone; `jolly_payne` removed. See Answer.

## Answer

Host cleanup from the main checkout, 2026-08-19. No feature branch, no PR, no Go edits, no `docker system prune`, no gastown/live-stack deletes, local `main` not fast-forwarded (`1895c64e0`, dirty `AGENTS.md`).

**Removed worktrees + local branches**

| Worktree | Branch | How |
|---|---|---|
| `gascity-wt-01` | `outpost/01-agents-canonical` | `worktree remove` + `branch -D` |
| `gascity-wt-02` | `outpost/02-coding-standards` | `worktree remove` + `branch -D` |
| `gascity-wt-03` | `outpost/03-upstream-pain-analysis` | `worktree remove` + `branch -d` |
| `gascity-wt-04` | `outpost/04-idle-bd-dolt-cpu` | `worktree remove` + `branch -D` |
| `gascity-wt-05` | `outpost/05-pool-worktree-isolation` | `worktree remove` + `branch -D` |
| `gascity-wt-14` | `outpost/01-remaining-upstream-bugs` | `worktree remove` + `branch -d` (git warned: merged to `origin/main`, not to local HEAD) |
| `gascity-wt-15` | `outpost/02-merge-upstream` | `worktree remove` + `branch -D` |
| `gascity-wt-16` | `outpost/03-rename-go-install` | `worktree remove` + `branch -D` |
| `gascity-wt-17` | `outpost/04-beads-fork-dep` | `worktree remove` + `branch -D` |
| `gascity-wt-18` | `outpost/05-supervisor-stop-launchd-durable` | `worktree remove` + `branch -D` |
| `gascity-wt-19` | `outpost/06-gc-stop-busy-reload` | `worktree remove` + `branch -D` |
| `gascity-wt-20` | `outpost/07-gc-start-reload-timeout` | `worktree remove` + `branch -D` |
| `gascity-wt-21` | `outpost/08-resume-keeps-option-defaults` | `worktree remove` + `branch -D` |
| `gascity-wt-22` | `outpost/09-controller-wake-enter` | `worktree remove` + `branch -D` |
| `gascity-wt-23` | `outpost/10-formula-step-first-prompt` | `worktree remove` + `branch -D` |

**Skipped (dirty untracked; no `--force`)**

| Worktree | Branch | Dirt |
|---|---|---|
| `gascity-wt-06` | `outpost/06-gc-init-issue-prefix` | `?? .outpost-steer.md` |
| `gascity-wt-07` | `outpost/07-nudge-enter-fail-closed` | `?? .outpost-ship-policy.md`, `?? .outpost-steer.md` |
| `gascity-wt-08` | `outpost/08-supervisor-stop-dolt-hang` | same two untracked files |
| `gascity-wt-09` | `outpost/09-pool-assignee-identity` | same two untracked files |
| `gascity-wt-10` | `outpost/10-acp-drain-process-leak` | same two untracked files |
| `gascity-wt-11` | `outpost/11-macos-supervisor-fd-leak` | same two untracked files |
| `gascity-wt-12` | `outpost/12-adhoc-pool-origin-gate` | same two untracked files |
| `gascity-wt-13` | `outpost/13-host-gc-cruft-cleanup` | `?? .scratch/fork-constitution-upstream/host-gc-cruft-{before,after}.md` and `issues/13-host-gc-cruft-cleanup.md` |

Prefect follow-up: confirm those untracked files are harness/scratch only, then `worktree remove` + `branch -D` (06–12 squash) / `branch -d` (13 ancestor).

**Kept (do not remove)**

- Main checkout `/Users/jonathanbaldie/go/src/github.com/jonbaldie/gascity` (`main` @ `1895c64e0`)
- `gascity-wt-24` `outpost/11-reset-pending-evict` (PR #21)
- `gascity-wt-25` `outpost/15-main-ci-green` (PR #22)
- `gascity-wt-26` `outpost/12-dolt-compact-concurrent-inserts` — present, not in ticket 16 bark; left alone

**Remotes** (`git push origin --delete`)

- Deleted: `outpost/01-agents-canonical`, `outpost/02-coding-standards`, `outpost/04-idle-bd-dolt-cpu`, `outpost/05-pool-worktree-isolation`, `outpost/06-gc-init-issue-prefix`
- Already gone (`remote ref does not exist`): `outpost/09-controller-wake-enter`, `outpost/10-formula-step-first-prompt`
- Observation: `origin/outpost/15-main-ci-green` was listed by `ls-remote` at claim time and absent at end. This worker did not delete it (not on the ticket list). Local branch + `gascity-wt-25` remain.

**Docker (optional, allowed)**

- `docker rm jolly_payne` — gone
- `docker volume rm 839e2c2dd8b2133fe83f9230d55e5fae89e47d19a6ae6744899caa8371354af0` — gone
- `gastown-uat-builder` / `gastown-weak-tests` left in place (Exited 255)
