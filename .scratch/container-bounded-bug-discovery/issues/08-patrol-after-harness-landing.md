# 08 — Patrol after harness landing

**Type:** patrol stamp

**Spec:** container-bounded-bug-discovery

**Scope:** Reconcile all tickets in this spec plus host leftovers for this
repository: live agents, worktrees, local/remote branches, Docker containers,
images, cache directories, and compiler output created by this effort.

**Cadence:** Due after three completed non-patrol tickets since the previous
patrol stamp (01, 06, and 07).

**Status:** resolved

## Patrol findings

Patrol completed from the provisioned checkout
`/Users/jonathanbaldie/go/src/github.com/jonbaldie/gascity-patrol-container-bounded-08`
on 2026-08-20 11:30 BST. Bark only; no production files, worktrees, branches,
containers, images, or caches were changed.

`.outpost.json` was present and applied: `maxWorkers=2`, `maxSplicers=1`,
`maxRunners=1`, `maxPatrolDogs=1`.

### Tickets

| Ticket | Finding | Smallest next prefect action |
|---|---|---|
| 01 — capped verification container harness | Complete, with landed revision recorded. No live agent. | All-clear; leave ticket evidence. Remove its stale worker tree/branch after this patrol. |
| 02 — fuzz configuration and formula boundaries | Claimed (tracker changed from ready while patrolling); live process `38207` is running `go test ./internal/config ./internal/formula` in `gascity-fuzz-config-formula-02`. | Keep; let the worker finish and record progress/result. |
| 03 — model-check session lifecycle | Claimed; live bounded process `36136` (plus wrapper tests `36133`/`37821`) is running in `gascity-model-session-lifecycle-03`. | Keep; let the worker finish and record progress/result. |
| 04 — replay session mutation concurrency | Ready-for-agent, no agent or worktree. | Queue; do not dispatch until a worker slot frees (`maxWorkers=2`, both slots occupied by 02/03). |
| 05 — audit credential and redaction flows | Ready-for-agent, no agent or worktree. | Queue; do not dispatch until a worker slot frees (`maxWorkers=2`, both slots occupied by 02/03). |
| 06 — splice capped verification container harness | Complete, integrated revision recorded; no live agent. | All-clear; remove its stale splice tree/branch after this patrol. |
| 07 — land capped verification container harness | Complete; PR #26 is squash-merged at `origin/main` `6721f7455`. No live agent. | All-clear; remove the stale runner tree/branch after this patrol. |
| 08 — this patrol | Patrol active; `maxPatrolDogs=1` is occupied by this dog. | Close/retain this resolved patrol record; do not dispatch another patrol. |

No ticket in this spec is blocked without a path. The only open unassigned work
(04/05) is intentionally queued behind the two live workers.

### Worktrees and local branches

All listed worktrees were clean. The relevant local branches have no matching
remote heads (`git ls-remote` returned none for the scoped `outpost/*`
branches).

| Identity | Finding | Smallest next prefect action |
|---|---|---|
| `gascity-container-bounded-bug-discovery-01` / `outpost/container-bounded-bug-discovery-01` @ `0b29ad6aa` | Ticket 01 worker tree; work is represented by the squash landing, so the branch is not an ancestor of `origin/main` even though the work landed. | Remove worktree, then delete the stale local branch (`-D` for the squash history). |
| `gascity-splice-container-bounded-bug-discovery-01` / `outpost/splice-container-bounded-bug-discovery-01` @ `ba9ebb16d` | Ticket 06 integration tree; no live agent and superseded by ticket 07 landing. | Remove worktree and delete the stale local branch. |
| `gascity-run-container-bounded-bug-discovery-01` / `outpost/run-container-bounded-bug-discovery-01` @ `8dc9a1a41` | Ticket 07 runner tree; landing is complete and no live runner exists. | Remove worktree and delete the stale local branch. |
| `gascity-fuzz-config-formula-02` / `outpost/fuzz-config-formula-02` @ `6721f7455` | Live ticket 02 worker tree; clean and equal to `origin/main`. | Keep until ticket 02 finishes. |
| `gascity-model-session-lifecycle-03` / `outpost/model-session-lifecycle-03` @ `6721f7455` | Live ticket 03 worker tree; clean and equal to `origin/main`. | Keep until ticket 03 finishes. |
| `gascity-patrol-container-bounded-08` / `outpost/patrol-container-bounded-08` @ `6721f7455` | This patrol's active tree; clean and equal to `origin/main`. | Keep until this record is handed off, then remove the patrol tree/branch. |
| Messgo trees: `gascity-messgo-hotspot-audit-01`, `gascity-splice-messgo-hotspot-audit-01`, `gascity-run-messgo-hotspot-audit-01` and branches `outpost/{messgo-hotspot-audit-01,splice-messgo-hotspot-audit-01,run-messgo-hotspot-audit-01}` | Pre-existing effort trees; no live process. Ticket 01 is still `ready-for-agent` despite worker/splicer completion notes and a blocked runner note. | Reconcile the Messgo ticket/runner decision, then remove stale trees/branches (do not redispatch blindly). |
| Messgo trees: `gascity-messgo-hotspot-audit-exclude-generated-02`, `gascity-splice-messgo-exclude-generated-02`, `gascity-run-messgo-exclude-generated-02`, `gascity-close-messgo-upstream-pr-5430` and their matching branches | Pre-existing effort trees; ticket 02 is completed and no live process exists. | Remove stale trees and local branches after preserving the completed ticket evidence. |
| Main checkout `/Users/jonathanbaldie/go/src/github.com/jonbaldie/gascity` @ `1895c64e0` | Dirty (`AGENTS.md`, `.outpost.json`, `.scratch/`, `.serena/`, `docs/agents/`) and 17 commits behind `origin/main`; these are host/user state, not disposable slag. | Keep; ask the human before any fast-forward or cleanup of dirty files. |

### Docker, temporary files, and compiler slag

| Identity | Finding | Smallest next prefect action |
|---|---|---|
| Live bounded run: PID `36136`, `docker run --rm --network none --read-only ... golang:1.25.9 go test ...` | Ticket 03's ephemeral bounded container is active; `--rm` means no persistent container is left. | Keep until the worker exits; do not prune Docker during the run. |
| `golang:1.25.9` | Shared base image used by the live ticket 03 run; pre-existing image, not disposable output. | Keep while ticket 03 runs; ask before later pruning. |
| `/private/tmp/verify-container-help.txt` (0 bytes, created 2026-08-20 10:40) | Empty harness-help artifact from this effort. | Prune this one file after the patrol handoff. |
| `/private/tmp/messgo-verify-{design,unusedcode,codesize}.json` (about 1.4 MiB total) | Slag from the pre-existing Messgo effort, with no live Messgo process. | Prune after Messgo ticket disposition. |
| `/private/tmp/verify-database-evidence` (empty, modified 11:21) | Recent verification scratch directory; ownership is not safely attributable to a finished ticket. | Keep and ask the active worker/prefect before removal. |
| Host `~/Library/Caches/go-build` and Go module caches | Shared compiler caches; no effort-specific compiler output directory was found. | Keep; do not clear shared caches. |
| Docker containers/images/volumes for `beads-dev`, rtorrent, gastown, adminer, Gogs, Jellyfin, Audiobookshelf, BuildKit, and other unrelated stacks | Existing host services or other-project state; no Gas City ownership evidence. | Keep; do not run global Docker prune. |
| Dangling/generic Docker images (including the old builder and PHP image) and other-project Go cache volumes | Not attributable to this effort. | Ask the owner before any targeted prune. |

No effort-specific persistent harness container, artifact directory, or compiler
output was found beyond the explicitly listed empty help file and active
ephemeral run.
