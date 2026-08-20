# 04 — Fix Nightly SQLite Coordination Store Tier A Acceptance Failure

**What to build:** Fix the SQLite coordination store Tier A acceptance test failure in the Nightly workflow (`Acceptance tests (SQLite coordination store, Tier A)`). Test in capped Docker container and ship via PR.

**Blocked by:** 01 — Diagnose & Map Root Causes for Main CI Failures.

**Status:** ready-for-agent

- [ ] Obsolete `integration-sqlite-coordstore` job removed from `.github/workflows/nightly.yml` (or retargeted to a supported provider).
- [ ] PR opened against `jonbaldie/gascity` with isolated fix
- [ ] All checks pass on PR

## Comments

### Root Cause Details
- Upstream removed the experimental `sqlite` bead store provider (`beads provider "sqlite" is no longer supported: the sqlite coordination-store experiment has been removed`).
- `.github/workflows/nightly.yml` still contains a legacy `integration-sqlite-coordstore` job setting `GC_BEADS: sqlite` and `GC_ACCEPTANCE_BEADS_PROVIDER: sqlite` and running `make test-acceptance`.
- All acceptance tests fail immediately on `gc status` / `gc wait` when trying to open the bead store.
- Fix: Delete the obsolete `integration-sqlite-coordstore` job from `nightly.yml`.
