# 08 — Supervisor stop must not hang on dolt-health bd list

**What to build:** `gc supervisor stop` (and city teardown) completes within a bounded timeout even when Dolt is under memory pressure or `bd list` for dolt-health would block. Operators can always force the city down without an indefinite hang.

**Blocked by:** None — can start immediately.

**Status:** resolved

**Upstream:** https://github.com/gastownhall/gascity/issues/5256 · related https://github.com/gastownhall/gascity/issues/5343

- [x] Stop path does not wait indefinitely on a blocking `bd list` / dolt-health probe
- [x] Forced-stop / grace timeout can fire while shutdown is in flight
- [x] Test or reproduction script covers the hang class (blocked child / unresponsive store)
- [x] Dedicated branch + PR against `jonbaldie/gascity` for this ticket only

## Comments

### 2026-08-19 — worker done

Root cause: `stopManagedCity` called `CityRuntime.shutdown()` synchronously after the grace period. When the city's deferred shutdown was already inside `shutdownOnce` and wedged on an unresponsive backend (`ListRunning` / `bd list`), the force path blocked forever and the second timeout select never ran.

Fix: `forceShutdownAsync` + async force path in `stopManagedCity` so the forced-stop deadline can fire; bead-provider cleanup still proceeds afterward.

- Branch: `outpost/08-supervisor-stop-dolt-hang`
- Commit: `4c868462`
- PR: https://github.com/jonbaldie/gascity/pull/7
- Regression: `TestStopManagedCityBoundsAForcedStopWhileShutdownIsAlreadyRunning`
