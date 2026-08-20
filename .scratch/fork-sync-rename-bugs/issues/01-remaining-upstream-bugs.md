# 01 — Analyse remaining upstream bugs worth fixing here

**What to build:** A ranked report of open bugs/frustrations on upstream `gastownhall/gascity` that this fork has **not** already shipped, plus an **initial** set of fix tickets (04+) on this same tracker. Skip anything already covered by `.scratch/fork-constitution-upstream/upstream-pain-report.md` and landed as PRs #3–#11 unless a regression is proven. Telescope: 04+ may grow later. Prefer high-pain open bugs (P0/P1, comment volume, recurring ops themes) that are still fixable in this fork.

**Blocked by:** None — can start immediately.

**Status:** resolved

- [x] Ranked report exists under `.scratch/fork-sync-rename-bugs/` (themes, upstream issue refs, why they matter here, explicit skip list of already-shipped first-pass items)
- [x] Report states that 05+ may telescope and how to add follow-ons (04 is Beads retarget)
- [x] At least one fix ticket (05+) is published **if** any remaining high-pain bug is worth fixing here; if none, the report says so with evidence and no empty tickets are invented
- [x] Published 05+ tickets (if any) are `ready-for-agent` with acceptance criteria and `Blocked by: 01, 02, 03, 04` (spec: 04 is Beads retarget; bug fixes start at 05)
- [x] No code fix in this ticket — research + ticketization only

## Answer

Remaining high-pain upstream bugs exist after PRs #3–#11. Published ranked report and initial telescope batch (05–14). Tracker 04 is the prefect Beads retarget, left untouched. No first-pass regression proven; skip list is explicit in the report.

- Report: `.scratch/fork-sync-rename-bugs/remaining-upstream-bugs.md`
- Fix tickets: `issues/05`–`14` (all `ready-for-agent`, `Blocked by: 01, 02, 03, 04`)

Top remaining themes: stop/start still lies (#5324/#5343/#5333), resume drops `option_defaults` (#5185), controller wake Enter-drop (#4935), unprimed formula steps (#4382), reset-pending wedge (#5355), compact quarantine on concurrent inserts (#3341), managed-Dolt incomplete schema (#5086), Darwin `gc dolt-cleanup` `ps rss=` (#5201). P0 #5135 left as beads-upstream, not a fork ticket.

05+ **may telescope** — how-to is in the report.

## Comments

- 2026-08-19: Claimed and completed by outpost worker on branch `outpost/01-remaining-upstream-bugs` (research + ticketization only; artifacts on main checkout `.scratch/fork-sync-rename-bugs/`). No Go source changes, no PR.
