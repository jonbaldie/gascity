# 03 — Analyse upstream open bugs and frustrations

**What to build:** A ranked analysis of the most common frustrations and reported/unsolved bugs on upstream `gastownhall/gascity`, written under this feature's `.scratch/` tree, plus an **initial** set of fix tickets (04+) published to the same tracker. The fix set may **telescope**: add more tickets later as themes clarify or follow-ons appear — do not pretend the first batch is complete. Prefer high-pain open bugs (P0/P1, high comment volume, recurring operational themes) over feature wishlist tracking issues.

**Blocked by:** None — can start immediately.

**Status:** resolved

- [x] Ranked report exists under `.scratch/fork-constitution-upstream/` (pain themes, issue refs, why they matter here)
- [x] At least one initial fix ticket (04+) is published with clear acceptance criteria and blockers
- [x] Report states explicitly that 04+ may grow (telescope) and how to add follow-ons
- [x] No code fix is required in this ticket — research + ticketization only

## Answer

Published ranked report and initial telescope batch:

- Report: `.scratch/fork-constitution-upstream/upstream-pain-report.md`
- Fix tickets: `issues/04`–`12` (all `ready-for-agent`, `Blocked by: None`)

Top themes: idle Dolt/bd CPU (#2463/#4133), pool shared worktrees (#1181), init prefix (#1436), nudge delivery (#5192), stop hangs (#5256), pool assignee identity (#5048), ACP process leaks (#5218), macOS FD leaks (#4504), adhoc origin gate (#5009). Epics #3872/#2903/#5135 noted as telescope seeds, not first-batch mega-fixes.

## Comments

- 2026-08-19: Claimed and completed by outpost worker on branch `outpost/03-upstream-pain-analysis` (research only; artifacts on main checkout `.scratch/`).
