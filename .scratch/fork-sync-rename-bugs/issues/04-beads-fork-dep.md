# 04 — Retarget Beads from steveyegge/gastownhall to jonbaldie/beads

**What to build:** This fork's Beads CLI dependency (install hints, CI pin, doctor/init, docs) uses `jonbaldie/beads` instead of `steveyegge/beads` or `gastownhall/beads`. Operators can install `bd` with `go install github.com/jonbaldie/beads/cmd/bd@…`. Gas Town (`steveyegge/gastown`) citations stay. Open one PR against `jonbaldie/gascity`.

**Blocked by:** 02 — Merge gastownhall/gascity into this fork; 03 — Rename gastownhall/gascity → jonbaldie/gascity (so go.mod/docs/install are stable before the beads retarget).

**Status:** resolved

- [x] `deps.env` `BD_REPO` is `jonbaldie/beads` (not `gastownhall/beads` / `steveyegge/beads`)
- [x] CI/scripts that download `bd` from GitHub point at `jonbaldie/beads` (or `go install` that module — no remaining upstream beads release URL as the install path)
- [x] Doctor/init install hints use `go install github.com/jonbaldie/beads/cmd/bd@…` (or equivalent pin)
- [x] User-facing docs (README / installation / troubleshooting) that tell you where to get `bd` point at `jonbaldie/beads`
- [x] No leftover *dependency* references to `steveyegge/beads` or `gastownhall/beads`; Gas Town provenance citations unchanged
- [x] Focused tests (doctor hint tests and any `BD_REPO` consumers) ran in resource-capped Docker
- [x] Dedicated PR opened against `jonbaldie/gascity` for this ticket only

## Comments

- 2026-08-19: Appended to the chain by prefect. Current tree pins `BD_REPO=gastownhall/beads` (historical name was `steveyegge/beads`). Target repo `https://github.com/jonbaldie/beads` exists (fork of gastownhall/beads). If that fork has no matching release assets, use `go install` rather than keeping upstream release URLs.
- 2026-08-19: Prefect: tickets 02 and 03 are on `origin/main` (PRs #12, #13). Start from that tip. After the upstream merge, `go.mod` may also require `github.com/steveyegge/beads` — retarget that module require too if present.
- 2026-08-19: Worker claimed on `outpost/04-beads-fork-dep` in worktree `gascity-wt-17` (from `3718a4fe6`).
- 2026-08-19: Worker: `jonbaldie/beads` has zero GitHub releases. Install path is `go install github.com/jonbaldie/beads/cmd/bd@latest`; CI `install-bd-archive.sh` clones that fork at the pinned tag and `go build`s. `go.mod` still requires `github.com/steveyegge/beads` at `bf97b737…` because that commit's module path is still `github.com/steveyegge/beads`; the fork rename is 1.2.1 (`96adeb582f45`) and would be a library upgrade. Provenance `gastownhall/beads#…` comments left alone.

## Answer

Shipped on `outpost/04-beads-fork-dep` as https://github.com/jonbaldie/gascity/pull/14 (`830600124`). Operators and CI install `bd` from `jonbaldie/beads`; the linked Go library pin is unchanged (see Comments).
