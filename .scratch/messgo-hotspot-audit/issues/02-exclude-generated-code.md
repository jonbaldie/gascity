# 02 — Exclude generated code and refresh the audit

**What to build:** Refresh the Messgo production-hotspot report so that generated source code is excluded from the audited population and from every violation, commit-concentration, and prioritization ranking, leaving maintainers with an actionable view of hand-written production code.

**Blocked by:** None — can start immediately.

**Status:** completed

- [x] The report defines and applies a reproducible generated-code exclusion across every result table.
- [x] The refreshed totals and rankings cover only hand-written production code while retaining the requested ruleset treatment.
- [x] The report retains traceable Messgo and Git evidence and records the revised inclusion boundary.
- [x] No production source code is changed.

## Comments

Completed on `outpost/messgo-hotspot-audit-exclude-generated-02` in commit
`0d5d20ea9`. The refreshed artifact is
`engdocs/audits/messgo-production-hotspots.md`; it excludes the generated
oapi-codegen client by canonical header before Messgo and Git rankings run.
