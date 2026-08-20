# Host gc cruft — after inventory

Captured: 2026-08-19 ~15:48 local (ticket 13)

## Live processes (gc-related)

None found:

- no `dolt sql-server`
- no `/tmp/gc-wt06 supervisor run`
- no `prefix-repro` / `gc-init-prefix-*` attached agents

## Listeners

- `*:33890` — gone
- `127.0.0.1:8372` — gone

## tmux sockets under `/tmp/tmux-501/`

| Socket | Status |
|--------|--------|
| `default` | preserved (no server running; socket file intact) |

All previously listed test/city sockets removed (`prefix-repro`, `nowt-prefix-*`, `gt-*`, `001-*`).

## `/tmp/gc*` paths

None remaining (`ls /tmp/gc*` → no matches).

## Residuals (with reason)

None for targeted zombies. Intentionally undisturbed:

- `/tmp/tmux-501/default` — personal/default tmux socket
- Git worktrees `gascity-wt-*` (including `gascity-wt-06`) — not host process cruft
- Unrelated Cursor/IDE/agent processes

## Teardown notes

1. `gc stop /tmp/gc-init-prefix-xjnU` via `/tmp/gc-wt06` — stopped city/tmux; reported beads stop script missing, but city marked stopped.
2. Orphaned `supervisor` (47192) and `dolt` (46908) received SIGTERM after stop (tmux already gone).
3. Dead non-default tmux sockets removed after confirming no live server.
4. `/tmp/gc-home-zRBv` needed `chmod -R u+w` before delete (read-only module cache under fake HOME).
