# Messgo production-code hotspot audit

**Status:** ready-for-agent

## Problem Statement

The project needs a clear, evidence-backed way to focus cleanup work where it has the greatest likely payoff. Messgo rule violations alone do not show which production-code areas have accumulated the most change, and commit concentration alone does not show maintainability risk.

## Solution

Produce a reproducible, read-only audit that ranks production-code hotspots by Messgo `design`, `codesize`, and `unused` violations, correlates those findings with file-level commit concentration, and identifies the commits associated with the highest-ranked files. The result is a concise report that lets maintainers choose follow-up work without changing production code.

## User Stories

1. As a maintainer, I want the highest concentrations of Messgo `design` violations identified, so that I can prioritize design cleanup.
2. As a maintainer, I want the largest `codesize` violations identified, so that I can find code that is difficult to understand or modify.
3. As a maintainer, I want `unused` violations identified, so that I can remove dead production-code surface deliberately.
4. As a maintainer, I want findings limited to production code, so that generated, test, and tooling code do not distort priorities.
5. As a maintainer, I want violation concentration compared with file-level commit counts, so that frequently changed risk areas stand out.
6. As a maintainer, I want the commits associated with leading hotspots listed, so that I can inspect the history behind the concentration.
7. As a maintainer, I want the method and commands recorded, so that I can reproduce or refresh the audit.
8. As a maintainer, I want the report to distinguish raw evidence from prioritization guidance, so that follow-up decisions remain transparent.
9. As a maintainer, I want no production behavior changed by this work, so that the audit can be reviewed independently.

## Implementation Decisions

- This is a read-only repository audit; it produces a tracked Markdown report and no production-code changes.
- The analysis will use the repository's available Messgo configuration and executable, reporting any tool or configuration limitation explicitly rather than substituting a different ruleset.
- Generated source files are excluded from the production-code population and from every ranking, even when they are compiled into the application.
- The ranking will show the three requested rulesets separately and provide a clearly labeled combined view where the evidence supports one.
- Commit concentration will be measured consistently at the file level, with the history window and aggregation method recorded in the report.
- The report will identify relevant commits for the leading hotspots without asserting causal authorship unless Git evidence directly supports it.

## Testing Decisions

- A good result is reproducible from the recorded commands and traces every ranked claim to Messgo or Git output.
- The audit will validate that included paths are production code and that each requested ruleset is either represented or explicitly reported unavailable.
- The report will include enough source data or command output references for a maintainer to check the top findings independently.

## Out of Scope

- Fixing, suppressing, or reconfiguring Messgo violations.
- Refactoring production code.
- Attributing intent or defect causality to individual commits or contributors.
- Analyzing test, generated, vendored, or third-party code as production hotspots.

## Further Notes

The requested outcome is prioritization evidence for subsequent tickets, not an automated policy gate.
