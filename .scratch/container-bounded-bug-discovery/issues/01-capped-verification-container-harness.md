# 01 — Build a capped verification container harness

**What to build:** A repeatable, resource-bounded Docker execution path for
bug-discovery commands that records enough invocation and artifact data to
replay an exploratory failure without running unbounded work on the host.

**Blocked by:** None — can start immediately.

**Status:** complete

- [x] Exploratory verification commands execute in Docker with explicit CPU,
  memory, process-count, temporary-storage, and wall-clock bounds.
- [x] The harness emits or retains a replayable bounded command and a clear
  location for generated-input, seed, or schedule artifacts.
- [x] Normal project test commands remain unaffected and the harness itself has
  deterministic coverage for its argument/resource handling.

## Completion

- Landed on `outpost/container-bounded-bug-discovery-01` as `0b29ad6aa`
  (`add bounded verification container harness`).
- The replay artifact is written as `replay.sh` in the selected artifact
  directory; the container receives that directory at `/artifacts` through
  `GC_VERIFICATION_ARTIFACTS`.
