# Spec: Fork identity, upstream sync, and another upstream-bug pass

Status: ready-for-agent

## Problem Statement

This clone is `jonbaldie/gascity`, but the tree still presents itself as `gastownhall/gascity`: the Go module path, clone URLs, Homebrew tap, and release-download instructions all point at upstream. Those install paths fail for people using this fork (no tap, no matching GitHub releases). Meanwhile upstream has moved, so this fork is drifting. Operators also want another look at the upstream issue board for bugs this fork has not already fixed.

## Solution

Merge current `gastownhall/gascity` into this fork first (preserving fork-only constitution, standards, and already-landed fixes), then rename project identity from `gastownhall/gascity` to `jonbaldie/gascity` and replace failing install instructions with a `go install` recommendation. Then retarget the Beads CLI/repo dependency from `steveyegge/beads` / `gastownhall/beads` to `jonbaldie/beads`. Independently, scan remaining open upstream issues, skip anything this fork already shipped, and telescope new fix tickets for what is still worth fixing here. Each slice ships as its own PR to `jonbaldie/gascity`.

## User Stories

1. As a fork user, I want `go install github.com/jonbaldie/gascity/cmd/gc@latest` documented as the install path, so that I can get a `gc` binary without a Homebrew tap this fork does not publish.
2. As a fork user, I do not want `brew install gastownhall/gascity/gascity` as the recommended install, so that I am not sent to a tap and formula that are not this fork.
3. As a fork user, I do not want clone/download URLs that pull `gastownhall/gascity` when the intent is to install *this* project, so that I do not silently get upstream instead of the fork.
4. As a Go contributor, I want the module path to be `github.com/jonbaldie/gascity`, so that `go install` and imports match the repository I cloned.
5. As a reader of coding standards, I want the documented module path to match the fork, so that contributors are not told to keep the upstream path.
6. As a contributor, I still want citations of upstream GitHub *issues* to keep pointing at `gastownhall/gascity`, so that bug provenance stays accurate.
7. As a contributor, I want the gascity *module* rename to leave Beads alone, so that ticket 03 does not mix two identity changes in one PR.
8. As a fork operator, I want the Beads CLI this city installs and documents to come from `jonbaldie/beads`, so that we are not pinned to `steveyegge/beads` or `gastownhall/beads`.
9. As a CI job, I want `BD_REPO` (and install/download URLs for `bd`) to point at `jonbaldie/beads`, so that image builds do not fetch the upstream beads org.
10. As a doctor/init user, I want install hints for `bd` to use `go install github.com/jonbaldie/beads/cmd/bd@…`, so that the same go-install approach as `gc` works for the beads binary.
11. As a maintainer, I want upstream `main` merged into this fork, so that we are not maintaining a stale snapshot.
12. As a maintainer, I want fork-only work (constitution symlink, coding standards, already-landed bugfixes) preserved across that merge, so that syncing upstream does not erase why this fork exists.
13. As a splicer, I want conflict resolution to keep both intents where they are compatible, so that neither upstream features nor fork fixes are dropped on the floor.
14. As a human operator, I want splicers to follow the merge-conflict skill rigorously, so that conflict hunks are resolved from primary sources rather than by guesswork.
15. As a fork maintainer, I want another pass over open upstream bugs, so that we catch issues that appeared or stayed open after the first pain report.
16. As a fork maintainer, I want that pass to skip bugs this fork already fixed (PRs #3–#11 and the first pain report), so that we do not re-litigate shipped work.
17. As a prefect, I want new fix tickets to telescope from the research report, so that the queue can grow without guessing the full set up front.
18. As a reviewer, I want each implementation slice to open its own PR against `jonbaldie/gascity`, so that review, revert, and bisect stay simple.
19. As a host operator, I want heavy compile/test work in resource-capped Docker, so that the Mac stays cool.
20. As an agent, I want `gh` defaulting to origin `jonbaldie/gascity`, so that PRs never land on upstream by accident.
21. As a future agent, I want the new research report kept under this feature's `.scratch/` tree, so that later sessions can resume without re-mining GitHub.
22. As a user of docs, I want a source-build clone URL that points at this fork after the rename, so that `git clone` of the documented repo succeeds for this project.
23. As a release reader, I want current (unreleased) changelog compare links to this fork, so that HEAD comparisons are not sent to a different GitHub org.

## Implementation Decisions

- **Sequence (intelligent, not priority):** merge upstream (ticket 02) before the gascity rename (ticket 03) so the rename is applied once to the merged tree. Beads retarget (ticket 04) follows 03 because both touch install docs / `deps.env` / doctor hints. Bug analysis (ticket 01) is independent of the merge. New code fixes (05+) are blocked by 01, 02, 03, and 04 so they land on the renamed, synced, beads-retargeted tree.
- **Upstream remote:** if missing, add `upstream` as `https://github.com/gastownhall/gascity.git`. Fetch `upstream/main` and merge it into a fork branch based on current `origin/main`. Do not merge onto the prefect's dirty checkout.
- **Conflict skill:** every merge/conflict ticket follows `/resolving-merge-conflicts` rigorously: inspect history and primary sources, preserve both intents where possible, never `--abort`, then run the project's checks.
- **Preserve fork intent across the merge:** `AGENTS.md` as constitution, `CLAUDE.md` symlink, root `CODING_STANDARDS.md`, and already-landed fork fixes. Preserve upstream intent for new features and fixes this fork does not have. Where incompatible, prefer the merge's stated goal (integrate upstream into this fork without dropping fork identity or prior fixes) and record the trade-off on the ticket.
- **Rename scope:** every *this-project* identity of `gastownhall/gascity` becomes `jonbaldie/gascity` — Go module path and imports, clone URLs, install/release-download URLs that would fetch the wrong org, contributor links for this repo, coding-standards module-path rule, unreleased changelog compare links.
- **Rename non-scope (ticket 03):** upstream gascity issue/PR citations; Gas Town-the-product names; historical notes that are explicitly about upstream. Beads identity is **ticket 04**, not 03.
- **Beads retarget (ticket 04):** every *dependency/install* identity of `steveyegge/beads` or `gastownhall/beads` becomes `jonbaldie/beads` — `deps.env` `BD_REPO`, CI archive install URLs, doctor/init install hints, README/install/troubleshooting `bd` links, `go install github.com/…/beads/cmd/bd`. Do not rewrite Gas Town (`steveyegge/gastown`) citations. Historical comments that are about Steve's original beads project may stay if they are provenance, not install instructions. The current tree pins `gastownhall/beads`; upstream merge may reintroduce `steveyegge/beads`. Both become `jonbaldie/beads`. Prefer `go install github.com/jonbaldie/beads/cmd/bd@<pin>` when this fork has no GitHub release assets matching the old pin.
- **Install docs:** remove Homebrew tap / upstream release-tarball / `gh attestation --repo gastownhall/gascity` as the recommended ways to install *this* fork. Recommend `go install github.com/jonbaldie/gascity/cmd/gc@latest` (plus still-required runtime deps: tmux, git, jq, and beads/dolt unless using the file provider). Source clone URL becomes this fork. `make install` from a clone of this repo remains valid as a contributor path, secondary to `go install`.
- **Shipping:** PRs against `jonbaldie/gascity` only; one ticket per PR; follow `/ship-pr` (no `--admin`). Main is protected by a GitHub ruleset (old branch-protection API may 404).
- **Thermal:** heavy `go test` / compile in Docker with `--cpus=2 --memory=4g --pids-limit=256`. See `.scratch/fork-constitution-upstream/host-thermal-steer.md`.
- **Tracker:** local markdown under `.scratch/fork-sync-rename-bugs/` per `docs/agents/issue-tracker.md`.
- **Prior research:** first-pass report lives at `.scratch/fork-constitution-upstream/upstream-pain-report.md` (tickets 04–12 already shipped as PRs #3–#11). Ticket 01 must not re-open those as new work unless a regression is proven.

## Testing Decisions

- Good tests assert external behaviour, not the particular search-and-replace tool used.
- Merge ticket: after conflict resolution, focused package tests in Docker still pass; fork-only constitution files still present (`CLAUDE.md` is a symlink to `AGENTS.md`; `CODING_STANDARDS.md` exists).
- Rename ticket: module path is `github.com/jonbaldie/gascity`; `go test` of affected packages compiles against the new import path; user-facing install docs no longer recommend the upstream Homebrew tap or upstream release assets as the way to install this project; `go install` is the documented default.
- Bug-analysis ticket: research artifact + new fix tickets; no code change required.
- Beads retarget: doctor/init install-hint tests and any `BD_REPO` consumers assert `jonbaldie/beads`; focused Docker tests, not a host-wide suite.
- Bug-fix tickets (05+): failing test at the highest existing seam (package test, then integration tag) before the fix; prior art is existing `*_test.go` beside packages and `test/` integration tests. Prefer production behaviour over mock choreography.
- Do not run broad host `make test` / `-race` shards; use capped Docker.

## Out of Scope

- Publishing a Homebrew tap or GitHub release assets for this fork (recommend `go install` instead).
- Opening PRs against `gastownhall/gascity` unless a later ticket explicitly says so.
- Renaming Gas Town (`steveyegge/gastown`) or any GitHub org other than this module (ticket 03) and the beads CLI/repo (ticket 04).
- Replacing Beads as the in-product issue store.
- Re-doing constitution / `CODING_STANDARDS.md` from the previous effort except where the module-path sentence must change.

## Further Notes

- Human said the three goals are in no particular order; the blocking edges above are for merge-safety (rename once after sync), not priority.
- Caps: `.outpost.json` `maxWorkers: 2`, `maxSplicers: 1`, `maxRunners: 1`.
- Ticket 01 telescopes **05+** (04 is the beads retarget). Do not pre-create fix tickets before the report exists.
- Human appended: CI on `main` is red. Ticket **15** makes fork `main` green (Notify Image Rebuilds + remaining red jobs). Dispatch when a worker slot is free (`maxWorkers: 2`).
- Human appended: queue Patrol Dog tickets; cleanup has not happened in a while. Tickets **16** (host leftovers: worktrees/branches/Docker) and **17** (tracker/agent drift). `maxPatrolDogs: 1` — dispatch 16 first. Ticket **18** executes 16’s bark (wt-01…23 only). Ticket **19** finishes what 18 skipped plus shipped wt-24…27. Leave wt-28 while ticket 14 is live.
- After 02 has an open PR, a runner lands it before 03 starts on `origin/main` (or 03 branches from 02's merge branch if the runner has not landed yet — prefer landed `origin/main`).
