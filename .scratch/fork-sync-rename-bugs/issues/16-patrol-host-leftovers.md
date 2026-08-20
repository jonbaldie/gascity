# 16 — Patrol host leftovers (worktrees, merged branches, Docker)

**What to build:** A bark of host cruft from this Outpost effort (and the prior constitution effort) that can be cleaned without touching live worker trees. Patrol dog only — do not delete, prune, or `git worktree remove`. Prefect turns the bark into follow-up cleanup.

**Blocked by:** None — can start immediately (patrol; `maxPatrolDogs: 1`).

**Status:** resolved

**Type:** task

- [x] Every `gascity-wt-*` worktree is listed with branch, HEAD, and whether that branch is merged to `origin/main` or still live (ticket 11 = `gascity-wt-24`, ticket 15 = `gascity-wt-25` are in-use — do not recommend removing those)
- [x] Local `outpost/*` branches whose PRs are merged (or which have no remote and no ticket) are listed with the smallest next action (`worktree remove` + `branch -d`, leave alone, ask human)
- [x] Docker leftovers (`golang:1.26*`, `/tmp/gascity-docker-gocache`, `/tmp/gascity-docker-gomod`, dangling containers) are listed with keep-vs-prune recommendation; do not prune
- [x] Bark recorded on this ticket under `## Answer` with ticket identity + smallest next prefect action per finding

## Comments

- 2026-08-19: Prefect queued from human “queue up Patrol Dog tickets, it's been a while since any cleanup has happened.” ~25 local worktrees (`gascity-wt-01` … `gascity-wt-25`) from two efforts. Bark only.
- 2026-08-19: Patrol dog claimed, inspected (no mutate), barked, resolved.

## Answer

Patrol of host leftovers, 2026-08-19. Inspected from the main repo (`git worktree list`, local `outpost/*` vs `origin/main` at `468493bbe` / PR #20, `docker ps -a`, `docker images golang`, `/tmp/gascity-docker-*`). Did not delete, prune, or `git worktree remove`.

`origin/main` = `468493bbe` (`fix(session): prime reconciler-spawned formula-step sessions (#20)`). No open PRs. Squash-merged branches are **not** git-ancestors of `origin/main` (`ahead=1`); after `worktree remove`, use `git branch -D` (plain `-d` will refuse). True ancestors can use `-d`.

**Keep (in-use — do not remove):**

| Identity | Why | Smallest next prefect action |
|---|---|---|
| `/Users/jonathanbaldie/go/src/github.com/jonbaldie/gascity` (`main` @ `1895c64e0`) | Main checkout. Behind `origin/main`; dirty `AGENTS.md`. | Keep. Ask human before any ff-pull — dirty `AGENTS.md` on stale `main`. |
| `gascity-wt-24` `outpost/11-reset-pending-evict` @ `f7c40d0c9` | Ticket 11 live. Git-ancestor of `origin/main` (behind 1) but worktree has staged edits in `cmd/gc/session_reconciler.go`, `session_reconciler_test.go`, `session_types.go`. | Keep until ticket 11 ships. |
| `gascity-wt-25` `outpost/15-main-ci-green` @ `468493bbe` | Ticket 15 live. Branch **equals** `origin/main`. | Keep until ticket 15 done, then `git worktree remove` + `git branch -d outpost/15-main-ci-green`. |

**Remove now (merged / ticket resolved; not the three in-use trees):**

Constitution effort (`gascity-wt-01`…`13`):

| Worktree | Branch @ HEAD | vs `origin/main` | Smallest next prefect action |
|---|---|---|---|
| `gascity-wt-01` | `outpost/01-agents-canonical` @ `ab57a74ab` | squash-merged PR #1; remote **still on origin** | `git worktree remove` + `git branch -D`; then `git push origin --delete outpost/01-agents-canonical` |
| `gascity-wt-02` | `outpost/02-coding-standards` @ `c4daaf0f7` | squash-merged PR #2; remote **still on origin** | `worktree remove` + `branch -D`; `git push origin --delete outpost/02-coding-standards` |
| `gascity-wt-03` | `outpost/03-upstream-pain-analysis` @ `b77856d1f` | git-ancestor (ahead=0 behind=20); no PR; ticket resolved | `worktree remove` + `branch -d` |
| `gascity-wt-04` | `outpost/04-idle-bd-dolt-cpu` @ `0be41451e` | squash-merged PR #3; remote **still on origin** | `worktree remove` + `branch -D`; `git push origin --delete outpost/04-idle-bd-dolt-cpu` |
| `gascity-wt-05` | `outpost/05-pool-worktree-isolation` @ `a1b17fb8f` | squash-merged PR #4; remote **still on origin** | `worktree remove` + `branch -D`; `git push origin --delete outpost/05-pool-worktree-isolation` |
| `gascity-wt-06` | `outpost/06-gc-init-issue-prefix` @ `d5e029db9` | squash-merged PR #5; remote **still on origin** | `worktree remove` + `branch -D`; `git push origin --delete outpost/06-gc-init-issue-prefix` |
| `gascity-wt-07` | `outpost/07-nudge-enter-fail-closed` @ `f663dcf49` | squash-merged PR #6; remote gone | `worktree remove` + `branch -D` |
| `gascity-wt-08` | `outpost/08-supervisor-stop-dolt-hang` @ `4c8684623` | squash-merged PR #7; remote gone | `worktree remove` + `branch -D` |
| `gascity-wt-09` | `outpost/09-pool-assignee-identity` @ `f7354f22d` | squash-merged PR #9; remote gone | `worktree remove` + `branch -D` |
| `gascity-wt-10` | `outpost/10-acp-drain-process-leak` @ `12aab2832` | squash-merged PR #8; remote gone | `worktree remove` + `branch -D` |
| `gascity-wt-11` | `outpost/11-macos-supervisor-fd-leak` @ `4c9bf34d7` | squash-merged PR #10; remote gone | `worktree remove` + `branch -D` |
| `gascity-wt-12` | `outpost/12-adhoc-pool-origin-gate` @ `d9fa56486` | squash-merged PR #11; remote gone | `worktree remove` + `branch -D` |
| `gascity-wt-13` | `outpost/13-host-gc-cruft-cleanup` @ `ccd1dbbb9` | git-ancestor (at PR #4; behind=16); ticket resolved | `worktree remove` + `branch -d` |

Fork-sync effort (`gascity-wt-14`…`23`; skip 24/25):

| Worktree | Branch @ HEAD | vs `origin/main` | Smallest next prefect action |
|---|---|---|---|
| `gascity-wt-14` | `outpost/01-remaining-upstream-bugs` @ `e9945cb44` | git-ancestor (behind=9); ticket 01 resolved | `worktree remove` + `branch -d` |
| `gascity-wt-15` | `outpost/02-merge-upstream` @ `a209d41c9` | squash-merged PR #12; remote gone; local ahead=2871 (merge commit) | `worktree remove` + `branch -D` |
| `gascity-wt-16` | `outpost/03-rename-go-install` @ `f339f4861` | squash-merged PR #13; remote gone | `worktree remove` + `branch -D` |
| `gascity-wt-17` | `outpost/04-beads-fork-dep` @ `830600124` | squash-merged PR #14; remote gone | `worktree remove` + `branch -D` |
| `gascity-wt-18` | `outpost/05-supervisor-stop-launchd-durable` @ `d131a053e` | squash-merged PR #16; remote gone | `worktree remove` + `branch -D` |
| `gascity-wt-19` | `outpost/06-gc-stop-busy-reload` @ `1d1c29711` | squash-merged PR #15; remote gone | `worktree remove` + `branch -D` |
| `gascity-wt-20` | `outpost/07-gc-start-reload-timeout` @ `d0edc2774` | squash-merged PR #17; remote gone | `worktree remove` + `branch -D` |
| `gascity-wt-21` | `outpost/08-resume-keeps-option-defaults` @ `5ce8e6d93` | squash-merged PR #18; remote gone | `worktree remove` + `branch -D` |
| `gascity-wt-22` | `outpost/09-controller-wake-enter` @ `c868eaff2` | squash-merged PR #19; remote **still on origin** | `worktree remove` + `branch -D`; `git push origin --delete outpost/09-controller-wake-enter` |
| `gascity-wt-23` | `outpost/10-formula-step-first-prompt` @ `63e31b469` | squash-merged PR #20; remote **still on origin** | `worktree remove` + `branch -D`; `git push origin --delete outpost/10-formula-step-first-prompt` |

All 25 local `outpost/*` branches are bound to a worktree; none are orphan local branches. Tickets 12–14 and 17 have no worktree yet — not leftovers.

**Docker (do not prune in this patrol):**

| Identity | Size / state | Keep vs prune | Smallest next prefect action |
|---|---|---|---|
| `/tmp/gascity-docker-gocache` | 1.6G; written 2026-08-19 19:08–19:16 | **Keep** (ticket 15 CI likely hot) | Leave until ticket 15 done; then optional `rm -rf` if unused. |
| `/tmp/gascity-docker-gomod` | 983M; same timestamp | **Keep** | Same as gocache. |
| `golang:1.26` / `golang:1.26.6` (same ID `2800d2462133`, 889MB) | matches `origin/main` `go 1.26.6`; used by fork-sync workers | **Keep** | Leave. |
| `golang:1.26-bookworm` (`00b8e454423d`, 839MB) | unused by named leftover containers | Prune after 15 if unused | `docker image rm golang:1.26-bookworm` |
| `golang:1.26.5-bookworm` (`d4f6d9156212`, 838MB) | superseded patch | Prune | `docker image rm golang:1.26.5-bookworm` |
| `golang:1.26.2-bookworm` (`521832c3ee77`, 838MB) | still referenced by `gastown-weak-tests` | Ask human | Do not rm until gastown container is gone. |
| `golang:1.25` / `golang:1.25-bookworm` (864MB / 814MB) | matches stale local `main` `go 1.25.9`; constitution thermal-steer used 1.25 | Ask human | Keep if gastown still needs 1.25; else `docker image rm` both after ticket 15. |
| `jolly_payne` (`59e4ef5a4e9d`, `dolthub/dolt-sql-server:2.2.0`) | Status=Created (never started), 2026-08-17; testcontainers labels (`org.testcontainers.reap=true`) | Prune (gascity test leftover) | `docker rm jolly_payne` then `docker volume rm 839e2c2dd8b2133fe83f9230d55e5fae89e47d19a6ae6744899caa8371354af0` |
| `gastown-uat-builder`, `gastown-weak-tests` | Exited 255; gastown compose/mounts, not gascity | Leave / ask human | Do not touch (other project). |
| `objective_edison` (`prosie:wolfi` PHP) | Created 2026-08-13; not gascity | Ask human | Leave unless human confirms junk. |
| `adminer-*`, `gogs-*`, `audiobookshelf`, `jellyfin`, `buildx_buildkit_*`, `myrest-mysql-*` | Live or unrelated | **Keep** | Do not `docker system prune`. |
| volumes `gastown-go-build-cache`, `gastown-go-mod-cache`, `gastown_agent-home`, `gastown_dolt-data` | gastown | Ask human | Leave. |

**Prefect sequence (after this bark, not this dog):**

1. Leave wt-24, wt-25, and the main checkout alone.
2. Remove wt-01…wt-23 as tabulated (`-D` for squash, `-d` for ancestors).
3. Delete the seven leftover `origin/outpost/*` remotes listed above.
4. Keep `/tmp/gascity-docker-*` and `golang:1.26`/`1.26.6` until ticket 15 finishes.
5. Optional: `docker rm jolly_payne` + its anonymous volume. Ask human before any gastown/1.25/live-stack prune.
6. Ask human about dirty `AGENTS.md` on local `main` before fast-forwarding the main checkout.
