# 20 — Patrol after ticket 19 and aborted ticket 14

**What to build:** A bark of this spec’s perimeter plus host leftovers: claimed 14 with no live agent, completed work still open, ship tickets, leftover worktrees/branches/Docker slag after ticket 19. Patrol dog only — do not claim or close tickets except this patrol’s own resolve.

**Blocked by:** None — cadence stamp after 12, 13, 19 and ships landed.

**Status:** resolved

**Type:** task

- [x] Every open/claimed ticket under `.scratch/fork-sync-rename-bugs/issues/` has one finding (or explicit all-clear) and a smallest next prefect action
- [x] Host leftovers: worktrees, local `outpost/*` branches, Docker slag this effort created
- [x] Bark recorded on this ticket under `## Answer`

## Comments

- 2026-08-19: Prefect cadence stamp. Ticket 19 resolved leftover trees. Ticket 14 claimed but its worker aborted; tests in `gascity-wt-28` were reverted. Leave 14 to the re-dispatched worker. Do not bark “dispatch 13” — that work shipped as PR #24.
- 2026-08-19 ~20:00 BST: Patrol dog claimed, barked, resolving. Did not claim/close any other ticket. Ignored ticket 17’s stale bark (predates PRs #23/#24 and ticket 19).

## Answer

Patrol of `.scratch/fork-sync-rename-bugs/` plus host leftovers for `jonbaldie/gascity`, 2026-08-19 ~20:00 BST. Bark only. Other tickets left as found.

Caps: `maxWorkers=2` (**1/2 used:** ticket 14 worker `c2056da2` live on `gascity-wt-28`), `maxRunners=1` (**0/1**, no open PRs), `maxPatrolDogs=1` (this dog). **Do not dispatch another worker, runner, or patrol.**

`origin/main` = `7a79b3962` (PR [#23](https://github.com/jonbaldie/gascity/pull/23) squash-merged 2026-08-19T18:53:05Z; parent is PR [#24](https://github.com/jonbaldie/gascity/pull/24) `9c7f32871`). Local main checkout still at `1895c64e0`, **behind 15**, dirty `AGENTS.md` + untracked `.outpost.json` / `.scratch/` / `docs/agents/` — **ask human** before ff-pull.

No `map.md` in this effort (nothing to pointer). Open GitHub PRs: none.

Ticket 17 bark ignored (stale; predates #23/#24 and ticket 19).

### Open / claimed — one finding each

| Ticket | Tracker | Live agent? | Finding | Smallest next prefect action |
|---|---|---|---|---|
| **14** macOS `ps rss=` | claimed | **Yes** `c2056da2` (dispatched 19:58, live on `gascity-wt-28`) | Not idle. Previous keeper aborted; this **fresh** worker merged `origin/main` onto `outpost/14-macos-dolt-cleanup-ps` (`7a79b3962`), is rewriting Darwin `ps` denial tests (`cmd/gc/dolt_cleanup_ps_rss_test.go` untracked), and is running capped Docker `go test ./cmd/gc/`. | **Leave.** Do not re-dispatch. Do not bark “14 has no agent.” Keep `gascity-wt-28`. |
| **20** this patrol | claimed → resolved | this dog | Cadence stamp after 12/13/19 and ships. | **Close (this file).** |

### 12 / 13 / ships — all-clear (no runner)

| Ticket | Tracker | Evidence | Next action |
|---|---|---|---|
| **12** | resolved | PR [#23](https://github.com/jonbaldie/gascity/pull/23) **MERGED** `7a79b3962` | **Leave.** No runner. |
| **13** | resolved | PR [#24](https://github.com/jonbaldie/gascity/pull/24) **MERGED** `9c7f32871` | **Leave.** Do not dispatch 13. |
| **ship-13** land PR #23 | resolved | Merge verified on origin | **All-clear.** |
| **ship-13** land PR #24 (`ship-13-land-pr24.md`) | resolved | Duplicate numbering; PR #24 actually landed here | **Leave** (hygiene only). |
| **ship-14** land PR #24 | resolved | Prefect closed as duplicate of `ship-13-land-pr24.md` | **Leave.** |

### Remaining spec tickets — all-clear

01–11, 15–19, ship-02–ship-12: tracker `resolved`, matching shipped PRs #12–#22 plus host-cleanup 18/19. No completed work still open. No blocked tickets. No ready-for-agent work sitting idle.

Optional CI on `origin/main` after #23/#24 is still queued/pending (not a ticket; 15 already shipped CI-green). **Leave.** Do not invent a runner or a new ticket unless those runs go red.

### Host leftovers

| Identity | Finding | Smallest next prefect action |
|---|---|---|
| worktree **main** `/Users/jonathanbaldie/go/src/github.com/jonbaldie/gascity` | Dirty checkout `1895c64e0`, behind origin by 15 | **Ask human** before ff-pull. **Keep.** |
| worktree **gascity-wt-28** `outpost/14-macos-dolt-cleanup-ps` @ `7a79b3962` | Live ticket 14; only leftover tree besides main (ticket 19’s remaining set) | **Keep** until 14 ships. |
| local branch `outpost/14-macos-dolt-cleanup-ps` | Live ticket 14 | **Keep.** |
| stale remote-tracking `origin/outpost/{09,10,11,12,13,15}-*` | Origin heads already gone (`git ls-remote` empty); `git remote prune origin --dry-run` would drop six tracking refs | **Prune** stale tracking (`git fetch --prune` / `git remote prune origin`). Not live branches. |
| Docker `sweet_hypatia` (`golang:1.26.6`, `--rm`, ticket 14 tests) | Live worker slag | **Keep** while 14 runs. |
| `/tmp/gascity-docker-gocache` (7.4G) + `/tmp/gascity-docker-gomod` (983M) | Mounted by the live ticket-14 `docker run` | **Keep** until 14 exits, then prune. |
| `/tmp/gascity-go-cache` (2.0G), `/tmp/gascity-go-mod` (983M), `/tmp/gascity-gocache` (439M), `/tmp/gascity-gomod` (865M) | This effort’s compiler caches; **no live process** (`lsof` empty) | **Prune slag** now (`rm -rf` those four dirs only). Do **not** `docker system prune` (would hit gastown/gogs/adminer/jellyfin). |
| gastown / gogs / adminer / jellyfin / audiobookshelf containers and images | Not this effort | **Keep.** |

**Worktrees after this patrol:** main + `gascity-wt-28` only. **Local `outpost/*`:** `outpost/14-macos-dolt-cleanup-ps` only.

### Caps / prefect sequence

1. **Leave 14.** Worker `c2056da2` is live; slot 1/2 occupied. Do not fill slot 2 — no other ready ticket.
2. **Do not dispatch a runner.** #23/#24 already on `main`; no open PRs.
3. **Ask human** before fast-forwarding dirty local `main` (`1895c64e0`).
4. **Prune** unused `/tmp/gascity-go-{cache,mod}` and `/tmp/gascity-{gocache,gomod}` plus stale `origin/outpost/*` tracking refs. After 14 ships, remove `gascity-wt-28`, delete `outpost/14-macos-dolt-cleanup-ps`, and drop `/tmp/gascity-docker-*`.
5. **Do not dispatch another patrol** (`maxPatrolDogs=1`; this one is done). Next cadence is after 14 (and its ship) complete.
