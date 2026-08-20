# Host gc cruft — before inventory

Captured: 2026-08-19 ~15:46 local (ticket 13)

Ticket 06 status at capture: **resolved** — `/tmp/gc-init-prefix-xjnU` and related
prefix-repro city are not mid-repro; safe to tear down.

## Live processes (gc-related)

| PID   | Command / notes |
|-------|-----------------|
| 46908 | `dolt sql-server --config /tmp/gc-init-prefix-xjnU/.gc/runtime/packs/dolt/dolt-config.yaml` (LISTEN `*:33890`) |
| 47192 | `/tmp/gc-wt06 supervisor run` (LISTEN `127.0.0.1:8372`) |
| 54330 | `tmux -u -L prefix-repro` server for city `/tmp/gc-init-prefix-xjnU` |
| 54331 | `claude` mayor session in that city (`GC_CITY=/tmp/gc-init-prefix-xjnU`) |

## Listeners

- `dolt` PID 46908 → `*:33890`
- `gc-wt06` PID 47192 → `127.0.0.1:8372`

## tmux sockets under `/tmp/tmux-501/`

| Socket | Sessions |
|--------|----------|
| `default` | no server running (socket file present) — **do not kill** |
| `prefix-repro` | `mayor` (live; city `/tmp/gc-init-prefix-xjnU`) |
| `nowt-prefix-16438` | `hq-deacon` (stale test, Aug 18) |
| `nowt-prefix-31351` | `hq-deacon` (stale test, Aug 18) |
| `001-870a57`, `001-8a4019` | dead sockets (no server) |
| `gt-46bfe7`, `gt-blockpane-*`, `gt-fix-shutdown-late-*`, `gt-test-*`, `gt-uat-bugs-17613` | dead sockets (no server) |

## `/tmp/gc*` paths

**Live city (processes attached):**

- `/tmp/gc-init-prefix-xjnU` — full city + runtime; hosts dolt/supervisor/claude
- `/tmp/gc-wt06` — ~97MB Mach-O binary used by supervisor

**Stale cities / scratch (no processes referencing path):**

- `/tmp/gc-deferred-prefix-ZtRQ`
- `/tmp/gc-home-PBPI`
- `/tmp/gc-home-zRBv`
- `/tmp/gc-init-prefix-2bfb` (empty)
- `/tmp/gc-init-prefix2-JyXw`
- `/tmp/gc-prefix-verify-gy0I`

**Logs / misc files:**

- `/tmp/gc-beads-bd.sh.bak`
- `/tmp/gc-init-gitconfig-46782`
- `/tmp/gc-init-out.txt`, `/tmp/gc-init2-out.txt`, `/tmp/gc-start2-out.txt`

## Explicitly left alone (not cruft)

- Default tmux socket `/tmp/tmux-501/default`
- Git worktrees `gascity-wt-*` under the go src tree
- Cursor / IDE / agent processes unrelated to the leftover city
