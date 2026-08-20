# Upstream pain analysis — `gastownhall/gascity`

**Date:** 2026-08-19  
**Source:** Open GitHub issues on `gastownhall/gascity` (kind/bug, priority P0/P1, high comment volume, recurring ops themes)  
**Ticket:** 03  
**Scope note:** This is an **initial** ranked cut. The 04+ fix queue is allowed to **telescope** — more tickets may be added later. This batch is not a complete inventory of upstream pain.

## Method

1. Listed open issues (sample of 100 via `gh issue list`, plus `gh search` sorted by comments for `kind/bug`).
2. Weighted by: priority label (P0 > P1 > P2), comment volume, recurrence of the same operational theme, and whether a concrete unsolved bug (not a multi-year tracker epic) can become a fork-scoped fix ticket.
3. Skipped pure feature wishlists and epic trackers as *primary* work; they remain **telescope seeds** below.

Label snapshot (open sample): ~68 `kind/bug`, ~29 `priority/p1`, 1 `priority/p0`.

## Ranked themes (highest pain first)

| Rank | Theme | Why it hurts here | Anchor upstream issues | Initial fork tickets |
|------|--------|-------------------|------------------------|----------------------|
| 1 | **Idle Dolt / bd CPU & battery** | Idle cities burn laptop battery and starve status/reconcile; operators lose trust before any user work starts | [#2463](https://github.com/gastownhall/gascity/issues/2463) (13c, P1), [#4133](https://github.com/gastownhall/gascity/issues/4133) (7c, P1), [#2201](https://github.com/gastownhall/gascity/issues/2201) (11c, P2), related [#2604](https://github.com/gastownhall/gascity/issues/2604) | **04** |
| 2 | **Pool agents share one worktree** | Parallel dispatch clobbers uncommitted edits; data-loss class failure on the default pool path | [#1181](https://github.com/gastownhall/gascity/issues/1181) (7c, P1) | **05** |
| 3 | **Fresh city unbootable / beads prefix** | `gc init` can leave cities unable to write beads until manual recovery | [#1436](https://github.com/gastownhall/gascity/issues/1436) (5c, P2); related [#5086](https://github.com/gastownhall/gascity/issues/5086), [#5135](https://github.com/gastownhall/gascity/issues/5135) P0 schema | **06** (prefix); P0 schema is a separate telescope candidate |
| 4 | **Nudge / mail delivery silent failure** | Work appears “sent” but never wakes the agent; no health signal | [#5192](https://github.com/gastownhall/gascity/issues/5192) (P1), [#5317](https://github.com/gastownhall/gascity/issues/5317), [#5006](https://github.com/gastownhall/gascity/issues/5006)/[#5007](https://github.com/gastownhall/gascity/issues/5007), [#1543](https://github.com/gastownhall/gascity/issues/1543) | **07** |
| 5 | **Stop / teardown hangs or lies** | Operators cannot shut cities down under Dolt pressure or busy reconcile | [#5256](https://github.com/gastownhall/gascity/issues/5256) (P1), [#5343](https://github.com/gastownhall/gascity/issues/5343), [#5324](https://github.com/gastownhall/gascity/issues/5324) | **08** |
| 6 | **Pool identity / assignee mismatch** | Cron/pool molecules strand with a healthy idle agent because assignee ≠ runtime env | [#5048](https://github.com/gastownhall/gascity/issues/5048) (P1); upstream fix PR [#5372](https://github.com/gastownhall/gascity/pull/5372) | **09** |
| 7 | **Provider / process leaks after drain** | Pool agents accumulate idle OS processes (RAM + false “alive”) | [#5218](https://github.com/gastownhall/gascity/issues/5218) (P1) | **10** |
| 8 | **macOS supervisor FD / pipe leaks** | Reload/idle supervisor exhausts descriptors; city becomes unstable | [#4504](https://github.com/gastownhall/gascity/issues/4504) (P1), [#5385](https://github.com/gastownhall/gascity/issues/5385) | **11** |
| 9 | **Demand-spawn pool seats origin-gated** | Adhoc seats never receive routed work | [#5009](https://github.com/gastownhall/gascity/issues/5009) (5c, P1) | **12** |

## Notable P0 / epic trackers (not in first fix batch)

| Issue | Notes | Telescope guidance |
|-------|--------|-------------------|
| [#5135](https://github.com/gastownhall/gascity/issues/5135) P0 | `bd` provider unusable — dirty Dolt schema migration / broken `ready_issues` | Add a dedicated ticket once fork can reproduce against current `bd`/`dolt` pins |
| [#3872](https://github.com/gastownhall/gascity/issues/3872) P1 | Session durable vs runtime divergence **family** (8c) | Split into one vertical slice at a time (adoption alias, priming on relaunch, ghost beads, …) — do not treat as one PR |
| [#2903](https://github.com/gastownhall/gascity/issues/2903) P1 | Bead-leak **tracker** + idle open-bead stability | Prefer child tickets from `engdocs/contributors/bead-leak-bookkeeping.md` rather than one mega-fix |
| [#3341](https://github.com/gastownhall/gascity/issues/3341) / [#2740](https://github.com/gastownhall/gascity/issues/2740) | Dolt compact / unbounded store growth | Follow after idle CPU (04) — same ops surface |

## Why these matter for this fork

- **Ops trust first:** Idle CPU (#2463/#4133) and stop hangs (#5256) hit every operator before any orchestration value lands.
- **Correctness under parallelism:** Shared worktrees (#1181) and assignee mismatch (#5048) are silent corruption / stranded-work bugs — high severity for multi-agent packs.
- **Boot path:** Prefix / schema init failures turn first-run demos into support tickets.
- **Delivery truth:** Nudge stranding (#5192) breaks GUPP-style “hook has work → run it” because the prompt never lands.

## Telescope: how to add more 04+ tickets later

This report intentionally publishes an **initial** batch only (`issues/04`–`12`). More tickets **will** be warranted.

**When to add a follow-on:**

1. A theme in the table above still has open upstream children after the first fork PR ships.
2. Fixing one ticket reveals a new failing acceptance criterion (file a new `NN-*.md`, do not expand the finished ticket).
3. An epic (#3872, #2903, #5135) is ready to split into a **single** vertical slice with clear repro.
4. Comment volume or a new P0/P1 appears on `gastownhall/gascity` that is not covered here.

**How to add:**

1. Create `.scratch/fork-constitution-upstream/issues/NN-<slug>.md` with the local ticket template (`Status: ready-for-agent`, `Blocked by: None` unless gated).
2. Append a row or bullet under this report’s “Ranked themes” or a new “Follow-ons” section linking `NN` → upstream URL.
3. Prefetch may dispatch immediately when `Blocked by` is None and 03 is resolved.

**Do not** invent a fake “complete” set. Prefer a short, shippable queue over a speculative backlog.

## Initial fix tickets published

| ID | Title | Upstream anchor |
|----|-------|-----------------|
| 04 | Bound idle city bd/Dolt query rate | #2463, #4133, #2201 |
| 05 | Isolate pool-agent worktrees on parallel dispatch | #1181 |
| 06 | Seed beads `issue_prefix` on `gc init` | #1436 |
| 07 | Make `gc session nudge` fail closed if Enter does not land | #5192 |
| 08 | Supervisor stop must not block forever on dolt-health `bd list` | #5256 |
| 09 | Align pool assignee writes with runtime session identity | #5048 |
| 10 | Terminate ACP provider process on pool drain-ack | #5218 |
| 11 | Stop macOS supervisor reload from leaking FDs | #4504 |
| 12 | Demand-spawned adhoc pool seats receive routed work | #5009 |

## Skipped (for now)

- Feature / design-only issues (e.g. worktree enforcement design share-out) unless they encode a concrete bug.
- Pack-specific role names in titles — fork tickets must stay ZFC / zero-hardcoded-roles.
- Opening PRs against upstream; ship on `jonbaldie/gascity` per parent spec.
