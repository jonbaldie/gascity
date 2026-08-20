# 17 — Patrol tracker and agent drift

**What to build:** A bark of the markdown perimeter: claimed tickets with no live agent, ready-for-agent work sitting idle, completed work still open, ship tickets waiting, and leftover tickets from `.scratch/fork-constitution-upstream` that are already shipped. Patrol dog only — do not claim or close tickets except this patrol’s own resolve.

**Blocked by:** None — can start immediately, but dispatch after 16 if a patrol-dog slot is already filled (`maxPatrolDogs: 1`).

**Status:** resolved

**Type:** task

- [x] Every open/claimed ticket under `.scratch/fork-sync-rename-bugs/issues/` has one finding (or explicit all-clear) and a smallest next action (re-dispatch, ship, leave, ask human)
- [x] Claimed tickets 11 and 15 are checked against live agents/worktrees (11 = `gascity-wt-24`; 15 was interrupted twice — confirm whether a worker is actually running)
- [x] Ready-for-agent 12, 13, 14 are listed with whether they should stay queued or need a worker slot
- [x] Prior effort `.scratch/fork-constitution-upstream/issues/` is scanned for still-open items that are actually shipped (PRs #1–#11) vs truly leftover
- [x] Bark recorded on this ticket under `## Answer`

## Comments

- 2026-08-19: Prefect queued from human cleanup reminder. Tracker hygiene, not code.
- 2026-08-19: Prefect dispatch (patrol dog). Ticket 16 is resolved; this is the free `maxPatrolDogs: 1` slot. Claim this ticket, bark only, do not close others.
- 2026-08-19: Prior patrol (~19:32, `dca45929`) resolved this ticket; that bark is stale.
- 2026-08-19 ~19:34 BST: Parallel dogs `405eba63` and `58642167` both wrote. This file is the reconciled bark from `58642167` after 18 flipped to `resolved`. Other tickets left as found.

## Answer

Patrol of tracker/agent drift, 2026-08-19 ~19:35 BST. Inspected from the main checkout `.scratch/` tracker. Did not claim/close any other ticket. Host leftovers not re-inventoried (ticket 16 / worker 18); worktrees used only as evidence.

Caps: `maxWorkers=2` (**1/2 used: 12 live; 18 just resolved — one slot free**), `maxRunners=1` (two runners raced on ship-12; both exiting — **do not fill with ship-11**), `maxPatrolDogs=1` (**over cap:** `dca45929` done, `405eba63` + this `58642167` both on ticket 17). **Do not dispatch another patrol or runner.**

`origin/main` = `eae28a8e6` (PR [#22](https://github.com/jonbaldie/gascity/pull/22) squash-merged 2026-08-19T18:33:45Z; parent is PR #21 `1388b7ca5`). Local main checkout still at `1895c64e0`, dirty `AGENTS.md` + untracked `.outpost.json` / `.scratch/` / `docs/agents/` — **ask human** before ff-pull.

No `map.md` in this effort (nothing to pointer).

Prefect snapshot corrections: 12 is `claimed` (not ready-for-agent); PR #22 is **MERGED** (not OPEN); ship-12 is already `resolved`; wt-24 is leftover (11 shipped).

### Open / claimed / ready — one finding each

| Ticket | Tracker | Live agent? | Finding | Smallest next prefect action |
|---|---|---|---|---|
| **12** | claimed | **Yes** `41990abd` (dispatched 19:31, still writing on wt-26) | Not idle. Branch `outpost/12-dolt-compact-concurrent-inserts` @ `468493bbe`, **behind `origin/main` by 2** (PRs #21+#22). | **Leave.** Do not re-dispatch. Worker may need to ff/rebase; only re-dispatch if that agent dies. Keep `gascity-wt-26`. |
| **13** managed-Dolt schema | ready-for-agent | none | Unblocked (01–04 resolved). 18 just freed a worker slot. | **Dispatch a worker now** (prefer 13 before 14). |
| **14** macOS `ps rss=` | ready-for-agent | none | Unblocked. Ticket 02 is merged; `gc dolt-cleanup` exists on this fork. | **Stay queued** until 12 or 13 exits (`maxWorkers=2`). |
| **16** host-leftover patrol | resolved | — | Bark was the work list; 18 executed it. | **Leave.** |
| **18** host cleanup | **resolved** | Worker `6b237046` ended success. | 01–05 and 14–23 removed. **wt-06…13 skipped dirty** (no `--force`; untracked `.outpost-*.md` / constitution scratch). wt-24/25/26 kept per AC. | **Leave.** Follow-up leftover: dirty wt-06…13, plus wt-24/25 now that #21/#22 shipped. |
| **ship-11** land PR #21 | **ready-for-agent** (AC unchecked, no `## Answer`) | No live runner. `bb19b21b` already finished and *did* write resolved+Answer; file was overwritten back to a stub. | **Completed work still open.** PR [#21](https://github.com/jonbaldie/gascity/pull/21) **MERGED** 2026-08-19T18:31:55Z (`1388b7ca5`). Origin head `outpost/11-reset-pending-evict` gone. | **Prefect resolve ship-11** (already merged). **Do not dispatch a runner.** |
| **ship-12** land PR #22 | resolved | Duplicate runners `0a953d7b` (success) + `e120a691` (wrapping) | PR [#22](https://github.com/jonbaldie/gascity/pull/22) MERGED `eae28a8e6`; origin `outpost/15-main-ci-green` deleted. | **All-clear.** Leave the wrapping runner. Do not re-dispatch. |

Do **not** dispatch a third worker. Slot 2/2 is free — fill it with **13**, not 14.

### Tickets 11 and 15 (acceptance)

| Ticket | Tracker | Live agent? | Worktree | Next action |
|---|---|---|---|---|
| **11** | **resolved**. Commit `23808bdf5`, PR [#21](https://github.com/jonbaldie/gascity/pull/21) **MERGED**. Worker `72899595` ended success. | No. | `gascity-wt-24` leftover, clean @ `23808bdf5` | **Leave** the resolved ticket. **Remove wt-24** + `git branch -D outpost/11-reset-pending-evict` (18’s AC forbade it; 18 is done). |
| **15** | **resolved**. PR [#22](https://github.com/jonbaldie/gascity/pull/22) **MERGED** (`eae28a8e6`). Worker `1ae62ac2` ended success. | No. Earlier `1c4ec801` + `5d16de1b` aborted; not zombies. | `gascity-wt-25` leftover, clean @ `b2493445f` | **Leave** the resolved ticket. **Remove wt-25** + `git branch -D outpost/15-main-ci-green`. |

### Ship tickets

| Ticket | Status | Next action |
|---|---|---|
| ship-02 … ship-10 | resolved (PRs #12–#20 merged) | **All-clear.** |
| **ship-11** | **ready-for-agent** while PR #21 is already on `origin/main` | **Prefect close it.** Do not occupy `maxRunners`. |
| **ship-12** | **resolved** (PR #22 merged). Duplicate runners both saw already-merged. | **All-clear.** Next ship is whatever 12 opens — do not invent a ship ticket yet. |

No ship-01 in this effort (numbering starts at PR #12). No open GitHub PRs.

### Resolved work (fork-sync) — all-clear except ship-11

01–11, 15, 16, 18, ship-02–ship-10, ship-12: tracker matches shipped PRs #12–#22. The only completed-work-still-open item is **ship-11**.

### Prior effort `.scratch/fork-constitution-upstream/issues/`

All 24 tickets are `resolved` or `completed`. GitHub PRs #1–#11 are MERGED. **No still-open leftover work.** Nothing to close.

Optional hygiene (not open work): `01`, `02`, `ship-01`–`ship-08` use `Status: completed` instead of `resolved`. **Leave** unless the prefect wants vocabulary cleanup.

### Caps / prefect sequence

1. **Stop extra patrol dogs and runners.** Leave 12 and the wrapping ship-12 runner.
2. **Dispatch 13 now** (one worker slot free). Keep 14 queued.
3. **Tracker-only:** restore ship-11 to `resolved` (PR #21 already on `main`). No runner.
4. **Follow-up hygiene** (do not implement here): remove wt-24 and wt-25; after confirming skipped dirt on wt-06…13 is harness/scratch only, remove those too (18’s Answer). Keep wt-26 while 12 is live.
5. When **12** exits → dispatch **14**.
6. Constitution tracker: no action.
7. Ask human before fast-forwarding the dirty main checkout.
