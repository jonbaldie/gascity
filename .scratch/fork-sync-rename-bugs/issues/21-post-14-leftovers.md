# 21 — Remove ticket-14 leftovers after PR #25

**What to build:** After PR #25 landed, remove the leftover ticket-14 worktree and this effort’s unused Go/Docker caches. Host hygiene only. Do not fast-forward the dirty main checkout.

**Blocked by:** ship-15 — Land PR #25 (merged).

**Status:** resolved

- [x] `gascity-wt-28` removed; local `outpost/14-macos-dolt-cleanup-ps` deleted
- [x] Stale `origin/outpost/*` tracking refs pruned (`git fetch --prune` or `git remote prune origin`)
- [x] Unused caches removed: `/tmp/gascity-go-cache`, `/tmp/gascity-go-mod`, `/tmp/gascity-gocache`, `/tmp/gascity-gomod`, and `/tmp/gascity-docker-gocache` / `/tmp/gascity-docker-gomod` if no live container is using them
- [x] Main checkout left alone (dirty `AGENTS.md` / `.scratch/` — do not ff-pull)
- [x] No `docker system prune`; no gastown/live-stack deletes

## Comments

- 2026-08-19: Prefect telescoped from patrol 20 after the [PR #25 runner](fde37907-d326-4afc-afc4-0db65acc4ead) finished. Bark: keep main dirty; drop wt-28 and unused `/tmp/gascity-*` caches.
- 2026-08-20: Prefect: prior agent interrupted while claimed; reset to `ready-for-agent` for redispatch.
- 2026-08-20: Worker claimed. Host hygiene from main checkout; no feature branch/PR; no Go edits.
- 2026-08-20: Worker finished. Verified `gascity-wt-28` removed and `outpost/14-macos-dolt-cleanup-ps` deleted; pruned origin remote-tracking refs; verified unused `/tmp/gascity-*` caches absent; left main checkout un-pulled (`1895c64e0`, dirty `AGENTS.md`); no docker prune. See Answer.

## Answer

Host cleanup from `/Users/jonathanbaldie/go/src/github.com/jonbaldie/gascity`, 2026-08-20. Host hygiene only: no feature branch, no PR, no Go edits, no `docker system prune`, no gastown/live-stack container deletes, local `main` not fast-forwarded (`1895c64e0`, dirty `AGENTS.md` / `.scratch/` / `.outpost.json` / `docs/agents/`).

**Worktree and Local Branch**

- `gascity-wt-28` worktree: absent (`git worktree list` shows only the main checkout).
- Local branch `outpost/14-macos-dolt-cleanup-ps`: absent (`git branch` shows only `* main`).

**Remote Tracking Refs**

- Executed `git remote prune origin` and `git fetch --prune origin`.
- No stale `origin/outpost/*` branches or tracking refs remain.

**Caches**

- Checked `/tmp/gascity*` (`/tmp/gascity-go-cache`, `/tmp/gascity-go-mod`, `/tmp/gascity-gocache`, `/tmp/gascity-gomod`, `/tmp/gascity-docker-gocache`, `/tmp/gascity-docker-gomod`).
- All are absent from `/tmp`.

**Docker and Main Working Tree**

- Main checkout left as found (at `1895c64e0`, dirty `AGENTS.md`, not fast-forwarded).
- No `docker system prune` executed; existing live containers (`adminer`, `gogs`, `jellyfin`, `audiobookshelf`, `gastown`) left untouched.
