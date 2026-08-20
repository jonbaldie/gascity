# 04 — Bound idle city bd/Dolt query rate

**What to build:** A fresh idle city (no user work) no longer saturates CPU via housekeeping / `bd` subprocess storms. Operators can leave a started city idle without melting battery or timing out `gc status` session snapshots.

**Blocked by:** None — can start immediately.

**Status:** resolved

**Upstream:** https://github.com/gastownhall/gascity/issues/2463 · https://github.com/gastownhall/gascity/issues/4133 · https://github.com/gastownhall/gascity/issues/2201

- [x] Documented or measured idle baseline: subprocess/`bd` rate on a fresh idle city is bounded (not tens of procs/s sustaining host saturation)
- [x] Idle-city path no longer makes `gc status` session-snapshot timeout the default experience under normal laptop load
- [x] Tests or a reproducible measurement script prove the bound (prefer production behaviour over mock choreography)
- [x] Dedicated branch + PR against `jonbaldie/gascity` for this ticket only

## Answer

Root cause matched upstream #4133/#3720: event triggers counted lifecycle events from dispatcher `order-tracking` beads, self-feeding on idle cities.

Fix: `internal/orders.checkEvent` skips payloads labeled `order-tracking` (flat + wrapped `{"bead":…}`). Regression tests in `triggers_test.go`. Measurement: `scripts/idle-bd-rate/` aggregates `GC_BD_TRACE` with optional `--max-per-sec`.

- Branch: `outpost/04-idle-bd-dolt-cpu`
- Commit: `0be41451`
- PR: https://github.com/jonbaldie/gascity/pull/3

## Comments

- claimed by outpost worker
- resolved: PR #3 opened; production seam fix + unit regressions + idle bd-rate script
