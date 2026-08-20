# 13 — Clean host zombie/stale gc cruft

**What to build:** The Mac host is free of zombie/stale/cruft left by running `gc` (and its tests) outside Docker: orphaned `dolt sql-server` processes, `gc`-owned tmux servers/sessions (non-default sockets only), stray `/tmp/gc-*` test cities, and related listeners. Personal/default tmux and unrelated user processes stay untouched. Prefer `gc stop` / targeted `tmux -L <socket>` teardown over blunt kills. Produce a short before/after inventory under this feature's `.scratch/` tree.

**Blocked by:** None — can start immediately. Coordinate with ticket 06 if a live repro city is still required; do not destroy an in-use worktree mid-PR without checking.

**Status:** resolved

**Known hot spots (prefect survey, may drift):**
- `dolt sql-server` listening on `:33890` under `/tmp/gc-init-prefix-xjnU`
- tmux sockets `prefix-repro`, `nowt-prefix-*`, assorted `gt-*` under `/tmp/tmux-501/`
- `claude` session attached as mayor for `/tmp/gc-init-prefix-xjnU`
- `gc-wt06` listener on `127.0.0.1:8372`
- `/tmp/gc-*` dirs and dolt hang/status logs

- [x] Inventory of gc-related host processes, tmux sockets (non-default), and `/tmp/gc-*` cities written under `.scratch/fork-constitution-upstream/`
- [x] Safe stop/teardown of identified cruft (`gc stop` first where a city root exists; never `tmux kill-server` on the default socket)
- [x] Default/personal tmux and non-gc user processes undisturbed
- [x] After inventory shows no remaining targeted zombies (or residual items listed with reason)
- [x] No code PR required unless a small `scripts/` helper is clearly warranted; prefer cleanup over new tooling

## Answer

Ticket 06 was already **resolved**, so the live `/tmp/gc-init-prefix-xjnU` prefix-repro city was safe to tear down.

**Cleaned:**
- Live city via `/tmp/gc-wt06 stop /tmp/gc-init-prefix-xjnU` (tmux `prefix-repro` + mayor claude gone)
- Orphaned post-stop `supervisor` (PID 47192, `:8372`) and `dolt sql-server` (PID 46908, `:33890`) via SIGTERM
- Stale live test sockets `nowt-prefix-16438` / `nowt-prefix-31351` (`kill-server`)
- Dead non-default sockets (`gt-*`, `001-*`, `prefix-repro`, `nowt-prefix-*`)
- All `/tmp/gc*` cities, logs, bak files, and `/tmp/gc-wt06` binary (`gc-home-zRBv` needed `chmod -R u+w` first)

**Left:**
- `/tmp/tmux-501/default` only
- Git worktrees `gascity-wt-*` (not process cruft)
- No PR (no reusable script added)

**Inventories:**
- `.scratch/fork-constitution-upstream/host-gc-cruft-before.md`
- `.scratch/fork-constitution-upstream/host-gc-cruft-after.md`

## Comments

- 2026-08-19 worker: claimed; confirmed ticket 06 resolved before tearing down prefix-repro city. Host hygiene only; no code PR.
