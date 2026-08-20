# Ship policy — protected `main` (ruleset)

**Human notice (2026-08-19):** `main` on `jonbaldie/gascity` is protected via an active repository **ruleset** (`main`, id `21046899`).

## What the ruleset enforces (current)

- No deletion of `main`
- No non-fast-forward history rewrite on `main`
- Changes land only through a **pull request** (merge / squash / rebase allowed)
- Required approving review count: **0** (reviews not required)
- Classic “branch protection” API may still 404 — **rulesets** are the source of truth

## Prefect / runner rules (effective immediately)

- When shipping via a PR/MR, runners **must** follow `/ship-pr`'s SKILL.md **rigorously** (gate on required checks + protections, merge with `--match-head-commit`, verify merge commit on target, delete branch, no admin bypass).
- **Do not** push directly to `main`
- **Do not** use `gh pr merge --admin` (or other bypass flags) to skip protections
- Prefer squash unless the user or repo convention says otherwise (per `/ship-pr`)
- Prefetch merge-while-CI-queued was an emergency for *unprotected* main; that exception is **revoked**
