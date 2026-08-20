# Container-bounded bug discovery

**Status:** ready-for-agent

## Problem Statement

Gas City has a large, stateful Go codebase spanning configuration parsing,
workflow/session lifecycles, concurrent mutation, command-line and API input,
and credential handling. Conventional example-based tests provide important
coverage, but they do not systematically search these state spaces for
counterexamples. Bug discovery also needs to be safe and repeatable: test
search must not consume unbounded developer resources or report flaky,
non-actionable failures.

## Solution

Establish a Docker-contained verification harness with explicit resource caps,
then apply the strongest appropriate feedback loop at the existing high-level
seams: coverage-guided fuzzing for external inputs, generated reference-model
properties for lifecycle logic, replayable schedules for concurrency, and
static data-flow analysis for credential-sensitive paths. Every confirmed
counterexample is minimized, deduplicated, fixed, and preserved as a
deterministic regression test. Coverage steers search; it is never treated as
evidence of correctness.

## User Stories

1. As a maintainer, I want automated exploratory checks to run in resource-capped containers, so that a bug hunt cannot destabilize my machine or CI worker.
2. As a maintainer, I want each exploratory run to record its command, seed, and relevant artifact, so that a suspected failure can be replayed.
3. As a maintainer, I want malformed configuration and formula inputs fuzzed at their public boundary, so that parser and validation panics or incorrect acceptance are exposed.
4. As a pack author, I want parsing and validation failures to identify the invalid input deterministically, so that I can correct a configuration without trial and error.
5. As a maintainer, I want lifecycle transitions exercised as generated sequences against a small reference model, so that invalid state changes and missing effects become visible.
6. As an operator, I want concurrent mutation behavior explored under controlled schedules, so that races and lost lifecycle updates can be reproduced rather than merely observed intermittently.
7. As a security reviewer, I want credential-bearing flows traced from sources to observable outputs, so that secrets cannot escape through diagnostics, process arguments, or persisted metadata.
8. As a maintainer, I want static-analysis candidates corroborated by focused tests before they are reported, so that security findings are actionable rather than noisy.
9. As a contributor, I want each verified defect reduced to the smallest practical reproducer, so that its cause is understandable and its regression coverage is durable.
10. As a reviewer, I want duplicate manifestations grouped under one root cause, so that the resulting work remains focused.
11. As a release manager, I want every bug fix to include a deterministic regression test, so that the defect stays fixed after the exploratory harness is removed or evolves.
12. As a CI operator, I want normal fast checks to remain bounded and stable, so that exploratory search does not turn routine verification into an unreliable bottleneck.
13. As a maintainer, I want the final report to contain only reproducible counterexamples, so that I can act on every reported result.

## Implementation Decisions

- A shared container invocation defines explicit CPU, memory, process-count,
  temporary-storage, and wall-clock limits; individual feedback-loop commands
  run through it rather than directly on the host.
- The harness captures enough evidence to replay a failing execution, including
  the generated input or seed, scheduler trace when applicable, tool version,
  and exact bounded command.
- Input-boundary search uses Go coverage-guided fuzzing and assertions at the
  existing public parsers and validation entry points. A failure must be
  minimized before it is treated as a defect.
- Domain state-machine search uses generated command sequences and a deliberately
  simple reference model that describes permitted transitions and observable
  effects. The production system and model are compared at their public seam.
- Concurrency search uses deterministic barriers or a controllable scheduler
  around existing synchronization seams. Schedules are saved in a replayable
  form; the race detector remains corroborating evidence rather than the sole
  oracle.
- Credential/redaction review uses static data-flow analysis targeted to known
  secret sources and external/persistent sinks. Candidates require a focused
  dynamic reproduction before becoming a finding.
- A confirmed issue is fixed in the smallest coherent change and covered by a
  deterministic test beside the affected behavior. Unconfirmed tool output is
  documented only as non-finding evidence, not filed as a bug.

## Testing Decisions

- Tests assert externally observable parser contracts, lifecycle invariants,
  synchronization outcomes, and secret-redaction behavior rather than private
  implementation structure.
- Existing unit, integration, and acceptance tiers remain the destination for
  deterministic regressions; exploratory runs may be longer but are bounded by
  the container harness.
- Existing lifecycle transition/conformance tests and input-validation tests are
  the prior art for reference-model and fuzz assertions. Existing Docker
  session checks are the prior art for isolated execution.
- Each feedback loop has explicit false-positive handling: reproduce, minimize,
  deduplicate, and verify the fixed behavior with a deterministic regression.

## Out of Scope

- Claiming coverage percentages prove correctness.
- Reporting unreproduced crashes, race-detector warnings without an invariant
  failure, or static-analysis warnings without a verified impact.
- Broad architectural refactoring unrelated to a confirmed defect.
- Production load, denial-of-service testing against external services, or
  unbounded fuzzing on a developer machine.
- Symbolic execution or mutation testing unless a later ticket shows that the
  selected feedback loops cannot reach a narrow, important branch or provide a
  credible oracle.

## Further Notes

The exact exploratory techniques are chosen by the state space under test, not
by a blanket coverage target. Findings should identify the feedback loop used,
the bounded reproduction, the minimized input or schedule, the root cause, and
the permanent regression test.
