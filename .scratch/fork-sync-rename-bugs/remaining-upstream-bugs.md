# Remaining upstream bugs worth fixing here

**Date:** 2026-08-19  
**Source:** Open GitHub issues on `gastownhall/gascity` (kind/bug, priority P0/P1, comment volume, recurring ops themes)  
**Ticket:** 01  
**Scope note:** This is an **initial** ranked cut after the first pain report. The 04+ fix queue is allowed to **telescope** — more tickets may be added later. This batch is not a complete inventory of upstream pain.

## Method

1. Listed open P0 issues, open `priority/p1`+`kind/bug` (80), and a 100-issue `kind/bug` sample via `gh issue list`, plus `gh search` sorted by comments.
2. Subtracted the first-pass skip list (PRs #3–#11 on `jonbaldie/gascity` and `.scratch/fork-constitution-upstream/upstream-pain-report.md`) unless a regression was proven. **No regression was proven.**
3. Weighted remaining issues by: P0 > P1 > P2, comment volume, whether the same operational theme still recurs after the shipped slice, whether a concrete unsolved bug can become a fork-scoped vertical ticket, and whether the seam is still unfixed on upstream `main` (open unmerged PRs do **not** arrive via ticket 02).
4. Skipped pure feature wishlists, mega-trackers as *primary* work, beads-upstream-only deadlocks, and pack/provider niches this fork does not operate.

Label snapshot (this pass): **1** open P0 (`#5135`); **~80** open P1 bugs; comment-volume leaders are still first-pass items, epic trackers, or feature workstreams.

## Explicit skip list (already shipped here — do not re-ticket)

| Fork PR | First-pass ticket | Upstream anchors | Why skipped |
|---------|-------------------|------------------|-------------|
| [#3](https://github.com/jonbaldie/gascity/pull/3) | 04 idle Dolt/bd CPU | [#2463](https://github.com/gastownhall/gascity/issues/2463), [#4133](https://github.com/gastownhall/gascity/issues/4133), [#2201](https://github.com/gastownhall/gascity/issues/2201) | Event-trigger self-feed on `order-tracking` beads bounded |
| [#4](https://github.com/jonbaldie/gascity/pull/4) | 05 pool worktree isolation | [#1181](https://github.com/gastownhall/gascity/issues/1181) | Parallel pool agents isolated |
| [#5](https://github.com/jonbaldie/gascity/pull/5) | 06 `gc init` issue_prefix | [#1436](https://github.com/gastownhall/gascity/issues/1436) | Prefix seeded/verified on init |
| [#6](https://github.com/jonbaldie/gascity/pull/6) | 07 nudge Enter fail-closed | [#5192](https://github.com/gastownhall/gascity/issues/5192) | `gc session nudge` fail-closed; **does not** cover controller wake keystrokes |
| [#7](https://github.com/jonbaldie/gascity/pull/7) | 08 supervisor stop hang | [#5256](https://github.com/gastownhall/gascity/issues/5256) | Forced-stop timeout can fire during wedged `bd list`; **does not** cover launchd durability or busy-reload `gc stop` |
| [#8](https://github.com/jonbaldie/gascity/pull/8) | 10 ACP drain process leak | [#5218](https://github.com/gastownhall/gascity/issues/5218) | Provider process terminated on drain/idle |
| [#9](https://github.com/jonbaldie/gascity/pull/9) | 09 pool assignee identity | [#5048](https://github.com/gastownhall/gascity/issues/5048) | Claims match runtime session identity |
| [#10](https://github.com/jonbaldie/gascity/pull/10) | 11 macOS supervisor FD leak | [#4504](https://github.com/gastownhall/gascity/issues/4504) | Config-watcher reload FD leak; **does not** cover [#5385](https://github.com/gastownhall/gascity/issues/5385) PIPE/socket profile |
| [#11](https://github.com/jonbaldie/gascity/pull/11) | 12 adhoc pool origin gate | [#5009](https://github.com/gastownhall/gascity/issues/5009) | Demand-spawned seats receive routed work |

Also skipped as *this* ticket's work: constitution PRs #1–#2 (not bugs). Ticket 13 from the previous effort was host cruft cleanup, not an upstream bug.

## Ranked remaining themes (highest pain first)

| Rank | Theme | Why it hurts here | Anchor upstream issues | Initial fork tickets |
|------|--------|-------------------|------------------------|----------------------|
| 1 | **Stop/start still lies** | After #5256, operators still cannot trust teardown/startup on macOS launchd: stop reports success while the job is restartable, `gc stop` restores a live city when reload is busy, `gc start` treats a reload timeout as fatal before readiness | [#5324](https://github.com/gastownhall/gascity/issues/5324) P1, [#5343](https://github.com/gastownhall/gascity/issues/5343) P2 (4c, updated today), [#5333](https://github.com/gastownhall/gascity/issues/5333) P1 | **05**, **06**, **07** |
| 2 | **Resume drops provider defaults** | Resumed sessions lose `option_defaults` (permissions / effort). The agent parks on an approval prompt while every liveness signal says healthy — regression of upstream #799 | [#5185](https://github.com/gastownhall/gascity/issues/5185) P1 (no upstream PR) | **08** |
| 3 | **Controller wake Enter drop** | Heartbeat/wake types into the pane mid-turn and never submits. Distinct from shipped `gc session nudge` fail-closed. Patrol cadence stalls until a human/other agent notices | [#4935](https://github.com/gastownhall/gascity/issues/4935) P1 (8 live occurrences; no upstream PR) | **09** |
| 4 | **Formula steps born unprimed** | Reconciler-spawned step sessions sit at a bare prompt; `nudge-on-route` only watches `bead.updated`. Autonomous formula runs stall at every step | [#4382](https://github.com/gastownhall/gascity/issues/4382) P1 (3c) | **10** |
| 5 | **Reset-pending wedge + crash-event flood** | Config/provider change leaves `reset-pending` forever; `session.crashed` fires ~2/min with full pane dumps; event log explodes. One vertical slice of the #3872 family | [#5355](https://github.com/gastownhall/gascity/issues/5355) P1 (filed 2026-08-18) | **11** |
| 6 | **Compact quarantines live stores** | Integrity check is strict equality across a flatten window; concurrent inserts quarantine the DB with no alert, blocking GC/backups. First-pass said follow *after* idle CPU — that slice has shipped | [#3341](https://github.com/gastownhall/gascity/issues/3341) P1, related [#2348](https://github.com/gastownhall/gascity/issues/2348), [#2740](https://github.com/gastownhall/gascity/issues/2740) | **12** |
| 7 | **Fresh managed-Dolt schema incomplete** | Boot path residual after prefix seeding: a new managed-Dolt city can come up without `hq` / `config` / migrations, so `bd` workflow never starts | [#5086](https://github.com/gastownhall/gascity/issues/5086) P1 | **13** |
| 8 | **macOS dolt-cleanup reap false-fails** | After ticket 02, this fork gains upstream's `gc dolt-cleanup`. On Darwin, `ps … rss=` hard-fails the reap stage even with zero orphans | [#5201](https://github.com/gastownhall/gascity/issues/5201) P1 | **14** |

Open unmerged upstream PRs exist for 05/06/07/11/12/13/14. They will **not** land via ticket 02. 08 and 09 have **no** upstream PR — highest unique value for this fork. Tracker 04 is the Beads retarget (prefect), not a bug fix.

## Why these matter for this fork

- **Ops trust, round two:** PR #7 unblocked wedged `bd list` shutdown. Operators on this Mac still cannot believe `gc stop` / `gc supervisor stop` / `gc start`. Launchd KeepAlive plus a busy reconcile queue is the production shape here.
- **Lights-out correctness:** Resume without `option_defaults` and controller-wake Enter-drop look “healthy” while work is frozen. That is the same class of silent delivery failure the first pass attacked on `gc session nudge`.
- **Orchestration actually runs:** Unprimed formula steps mean molecules never start without a human. That breaks the “hook has work → run it” contract at the engine layer, not the mail layer.
- **Storage after idle-CPU:** PR #3 stopped the event self-feed. Compaction still races writers and then *disables* GC. Unbounded HQ growth (#2740) stays in force until compact can succeed on a live city.
- **Boot path:** Prefix init shipped; managed-Dolt schema completion did not. First-run still dies with a half store.
- **Merge-then-fix:** `gc dolt-cleanup` is on upstream `main` and absent from this fork today. Ticket 14 is gated on 02 so we do not invent the command, only stop Darwin `ps` from failing a successful reap.

## Considered and not ticketed (with evidence)

| Issue | Why not a 04+ ticket now |
|-------|--------------------------|
| [#5135](https://github.com/gastownhall/gascity/issues/5135) **P0** | Remaining deadlock is [beads#4628](https://github.com/gastownhall/beads/issues/4628) / dirty-tables in `bd`, not a Gas City write path. Reporter's earlier symptoms were version/env specific (bd 1.1.2, root+`--dangerously-skip-permissions`). `gc doctor` already has `BeadsStoreCheck`. Fail-closed city start on hard store-open is a **behavior decision** — telescope only if operators still get a half-working city after a current `bd` pin. |
| [#3872](https://github.com/gastownhall/gascity/issues/3872) family (8c) | Five incidents, not one PR. Ticket **11** takes the reset-pending / priming-loss slice. Adoption-alias, serve-loop store pinning, ghost session beads, scale-from-zero stay telescope seeds. |
| [#2903](https://github.com/gastownhall/gascity/issues/2903) bead-leak tracker (7c) | Stacked-PR mega-tracker. Prefer a single child (wisp prune #3926, sling wrappers, mail wisps) once compact/idle residuals settle. `engdocs/contributors/bead-leak-bookkeeping.md` is **not** in this fork today. |
| [#3926](https://github.com/gastownhall/gascity/issues/3926) order-tracking wisp prune | Real idle-CPU residual after PR #3, but the prune abort is a second storage slice. Follow ticket 12 rather than parallelize two store-churn PRs. |
| [#5317](https://github.com/gastownhall/gascity/issues/5317) / [#4534](https://github.com/gastownhall/gascity/issues/4534) / [#4299](https://github.com/gastownhall/gascity/issues/4299) / [#5278](https://github.com/gastownhall/gascity/issues/5278) | Nudge-queue residuals after fail-closed Enter. Concrete, but 09 already takes the wake/heartbeat seam; do not land four nudge PRs in one batch. |
| [#4448](https://github.com/gastownhall/gascity/issues/4448) / [#4849](https://github.com/gastownhall/gascity/issues/4849) | Pool slot starvation / `wake_mode=fresh` resume. Real P1s; wait until 02 merge so we do not fight upstream pool-desired-state churn twice. |
| [#5385](https://github.com/gastownhall/gascity/issues/5385) | Distinct PIPE/socket FD profile; ticket 11 of the first pass explicitly out-of-scoped it. Telescope. |
| [#5348](https://github.com/gastownhall/gascity/issues/5348) ambient `bd` schema ceiling | Related boot footgun; ticket 13 is the incomplete-schema slice. Pin/ceiling mismatch can follow. |
| [#2713](https://github.com/gastownhall/gascity/issues/2713) serialization storms | Likely belongs in beads retry, not a one-shot gc ticket. |
| [#4308](https://github.com/gastownhall/gascity/issues/4308) PID1 `bd` zombies | Hurts Docker PID1; this host is launchd. Telescope for capped-Docker operators. |
| [#2735](https://github.com/gastownhall/gascity/issues/2735) / [#5042](https://github.com/gastownhall/gascity/issues/5042) | Pack/role-specific. Fork tickets must stay zero-hardcoded-roles. |
| herdr / k8s P1s (`#4914`, `#4891`, `#5143`, …) | This fork's operator path is tmux + macOS. |
| Feature / chore trackers (`#2504`, `#2120`, `#1709`, `#5188`, …) | Wishlists, not bugs. |
| [#5336](https://github.com/gastownhall/gascity/issues/5336) events.jsonl NUL tail | Real event-bus framing bug, P2, no live repro on the reporter's cities. Telescope. |

Ticket 02 merging `upstream/main` will **not** close the 05–14 set: their fix PRs (where they exist) are still **open**. A few older compact/session PRs *are* merged (`#2350`, `#3547`, `#4034`); `#3341` and `#5355` remaining open means those merges did not finish the job.

## Telescope: how to add more 04+ tickets later

This report publishes an **initial** batch only (`issues/05`–`14`). Tracker **04** is the Beads retarget, not a bug. More tickets **will** be warranted. **05+ may telescope.**

**When to add a follow-on:**

1. A theme in the table above still has open upstream children after the first fork PR for that theme ships.
2. Fixing one ticket reveals a new failing acceptance criterion (file a new `NN-*.md`, do not expand the finished ticket).
3. An epic (`#3872`, `#2903`, `#5135`, `#2740`) is ready to split into a **single** vertical slice with clear repro (next likely: `#3926` wisp prune, `#4448` drain reaping, `#5278` dead-letter prune, `#5385` PIPE FDs, `#5348` schema ceiling).
4. Comment volume or a new P0/P1 appears on `gastownhall/gascity` that is not covered here **and** is not on the skip list.

**How to add:**

1. Create `.scratch/fork-sync-rename-bugs/issues/NN-<slug>.md` using the local ticket template (`Status: ready-for-agent`).
2. Until 02, 03, and 04 are resolved, set `Blocked by: 01, 02, 03, 04` so fixes land on the renamed, synced, beads-retargeted tree. After those are resolved, new tickets may use `Blocked by: None` unless they have a real remaining gate.
3. Append a row or bullet under this report’s “Ranked remaining themes” or a new “Follow-ons” section linking `NN` → upstream URL.
4. Prefetch may dispatch when blockers are `resolved`.

**Do not** invent a fake “complete” set. Prefer a short, shippable queue over a speculative backlog.

## Initial fix tickets published

| ID | Title | Upstream anchor |
|----|-------|-----------------|
| 05 | Supervisor stop must be durable against launchd restart | #5324 |
| 06 | `gc stop` must not leave the city running when reload is busy | #5343 |
| 07 | `gc start` must not treat reload timeout as fatal before readiness | #5333 |
| 08 | Resume must keep provider `option_defaults` | #5185 |
| 09 | Controller wake delivery must confirm Enter or fail closed | #4935 |
| 10 | Reconciler-spawned formula steps receive a first prompt | #4382 |
| 11 | Evict wedged reset-pending sessions; bound crash-event payloads | #5355 |
| 12 | Dolt compact must not quarantine on concurrent inserts | #3341 |
| 13 | Fresh managed-Dolt init must finish a usable Beads schema | #5086 |
| 14 | macOS dolt-cleanup reap must not hard-fail on `ps rss=` | #5201 |
