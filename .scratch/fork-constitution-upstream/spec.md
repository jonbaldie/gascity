# Spec: Fork constitution, coding standards, and upstream pain fixes

Status: ready-for-agent

## Problem Statement

This fork needs a clear agent constitution (`AGENTS.md` only, not a parallel `CLAUDE.md`), coding standards aligned with the maintainer's other work (`jonbaldie/gastown`), and a deliberate program to identify and fix the most common open frustrations and unsolved bugs reported against upstream `gastownhall/gascity` — shipping each slice via its own PR rather than holding everything until the end.

## Solution

Make `AGENTS.md` the single constitution and replace `CLAUDE.md` with a symlink to it. Add a root `CODING_STANDARDS.md` modelled on `jonbaldie/gastown`'s file but adapted to Gas City vocabulary and invariants (including zero hardcoded role names). Analyse upstream open issues for the highest-pain bugs and frustrations, then telescope fix tickets from that analysis and land each fix (and each docs/constitution change) as a separate PR as soon as it is ready.

## User Stories

1. As a contributor, I want one constitution file, so that agent instructions never diverge between `AGENTS.md` and `CLAUDE.md`.
2. As a Claude Code user, I want `CLAUDE.md` to still resolve, so that tools that look for that filename keep working.
3. As a maintainer, I want `CLAUDE.md` to be a symlink to `AGENTS.md`, so that there is only one editable source of truth.
4. As a contributor, I want root coding standards, so that test and Go expectations are explicit for this fork.
5. As a contributor, I want those standards modelled on `jonbaldie/gastown`, so that house style matches across related projects.
6. As a Gas City contributor, I want standards to use city/agent/pack/formula language, so that we do not reintroduce hardcoded Gas Town role names.
7. As a fork maintainer, I want upstream open bugs ranked by pain, so that we fix what users actually hit first.
8. As a fork maintainer, I want unsolved high-priority upstream bugs considered for fixes here, so that the fork earns trust faster than waiting on upstream.
9. As a contributor, I want each constitution/standards/fix slice to open its own PR, so that review and merge are not blocked on the whole program.
10. As a prefect/coordinator, I want fix tickets to telescope from the research artifact, so that the queue can grow without guessing the full fix set up front.
11. As a reviewer, I want each PR scoped to one ticket, so that reverts and bisects stay simple.
12. As an agent worker, I want clear acceptance criteria per ticket, so that done means verifiable and shippable.
13. As a human operator, I want shipping to proceed as tickets finish, so that value lands continuously.
14. As a future agent, I want the research report kept under `.scratch/`, so that later sessions can resume without re-mining upstream.
15. As a contributor, I want Beads session rules preserved inside `AGENTS.md` after the symlink change, so that existing beads workflow text is not lost.

## Implementation Decisions

- Constitution: `AGENTS.md` remains the only real file; `CLAUDE.md` becomes a relative symlink to `AGENTS.md`. Any unique content that lived only in `CLAUDE.md` is folded into `AGENTS.md` first if still needed.
- Coding standards: new root `CODING_STANDARDS.md` adapted from `jonbaldie/gastown`'s file — keep structure (Tests, Comments/docs, Common footguns, Go); rewrite Gas Town-specific terms and paths to Gas City equivalents; honour this repo's ZFC / zero-hardcoded-roles / testing targets from existing agent docs.
- Upstream analysis: use GitHub issues on `gastownhall/gascity` (open bugs, high comment volume, P0/P1 labels, recurring themes such as Dolt CPU, session lifecycle, pool clobbering, bead leaks). Produce a ranked report under `.scratch/fork-constitution-upstream/` and publish fix tickets adaptively (04+ can telescope).
- Shipping: each completed implementation ticket opens a PR against this fork's default repo (`jonbaldie/gascity`); do not batch unrelated tickets into one PR.
- Tracker: local markdown under `.scratch/fork-constitution-upstream/` per `docs/agents/issue-tracker.md`.

## Testing Decisions

- Prefer external behaviour: symlink resolves; `CLAUDE.md` and `AGENTS.md` present the same constitution text to readers/tools.
- `CODING_STANDARDS.md` is documentation — verified by presence, structure, and Gas City vocabulary (no role-name hardcoding), not by Go tests.
- Bug-fix tickets must prove the fix via the highest available seam already used in this repo (package tests, integration tags, or reproduction scripts documented in the ticket). Prefer asserting production behaviour over mock choreography.
- Prior art: existing `*_test.go` beside packages; integration tests under `test/` with build tags; `TESTING.md` shard targets.

## Out of Scope

- Promoting this fork back to upstream or opening PRs against `gastownhall/gascity` unless a later ticket explicitly says so.
- Replacing Beads with the local markdown tracker for all repo work — local markdown is the outpost/skills tracker for this effort; Beads blocks in `AGENTS.md` stay until a separate decision.
- Building a full VS Code extension or PackV2 coordination workstreams from upstream tracking issues.
- Rewriting the entire agent pack ecosystem.

## Further Notes

- Ticket **03** must remain adaptive: publish an initial ranked set of fix tickets, and allow further 04+ tickets to be added as analysis deepens or as fixes uncover follow-ons ("telescope out").
- Prefetch reference for standards authors: `jonbaldie/gastown` `CODING_STANDARDS.md`.
- `gh repo set-default` for this clone is `jonbaldie/gascity` (origin), not the upstream parent.
- **Ticket 13 (added):** clean host zombie/stale cruft from running `gc` raw on the Mac (tmux/dolt/tmp cities). Thermal steer still applies to compile/test; this ticket is host process hygiene.
