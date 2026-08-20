# 05 — Audit credential and redaction flows

**What to build:** A container-bounded static data-flow review of known
credential sources through logs, process execution, metadata, and API-facing
sinks, with focused dynamic corroboration and deterministic regressions for
verified leaks.

**Blocked by:** 01 — Build a capped verification container harness.

**Status:** ready-for-agent

- [ ] The review names the credential sources and observable/persistent sinks
  being traced, and runs the chosen analysis within the bounded container.
- [ ] A candidate is reported only when a focused dynamic reproduction confirms
  externally visible exposure or an equivalent actionable security defect.
- [ ] Each confirmed defect is minimized, fixed, and protected by a
  deterministic redaction regression test.
