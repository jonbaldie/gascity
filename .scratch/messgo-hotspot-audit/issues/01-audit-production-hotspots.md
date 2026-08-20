# 01 — Audit Messgo production-code hotspots

**What to build:** Deliver a reproducible, read-only Markdown report ranking production-code files by Messgo `design`, `codesize`, and `unused` violation concentration, correlating the findings with file-level commit concentration and identifying the commits relevant to the leading hotspots.

**Blocked by:** None — can start immediately.

**Status:** ready-for-agent

- [ ] The report covers each requested Messgo ruleset, or explicitly records why an unavailable ruleset could not be analyzed.
- [ ] The report excludes non-production code and records its inclusion criteria.
- [ ] The report ranks the leading violation and commit-concentration hotspots with traceable supporting evidence.
- [ ] The report identifies the commits associated with the leading hotspots without unsupported causal claims.
- [ ] The report records reproducible commands or method details and makes no production-code changes.

## Comments

- Worker completed the audit on `outpost/messgo-hotspot-audit-01` and pushed commit `085d7cd5a9af0f3f80c64d298ee1c8283e67c7b8`. The report is `engdocs/audits/messgo-production-hotspots.md`; `go test ./test/docsync` passed. Integration is pending a splicer.
- Splicer integrated and pushed the report as `outpost/splice-messgo-hotspot-audit-01` at `a260035aa1e93b3aeb6e7eaa190ee4b8e7901e87`; `go test ./test/docsync` passed.
- Runner is blocked before shipping: the integration checkout's `AGENTS.md` points to `docs/agents/issue-tracker.md`, but that tracker document is not committed on the branch, so the required claim/update workflow cannot run. No shipping changes were made.
