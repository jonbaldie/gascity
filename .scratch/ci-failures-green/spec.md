# Spec: Fix Main CI, Nightly, and Mac Regression Failures

Status: ready-for-agent

## Problem Statement

Remote origin CI on `jonbaldie/gascity` `main` is not green. Specifically:
- **Nightly** workflow fails on `Tier C acceptance tests (Ollama)` (`Validate Ollama Claude configuration`) and `Integration / SQLite coordination store` (`Acceptance tests (SQLite coordination store, Tier A)`).
- **Mac Regression** workflow fails on `Mac / quality` and multiple `cmd/gc process` test shards (shards 3, 5, 10).
- **CI** and **Review Formulas** workflows on `main` remain stuck pending.

These failures prevent confident continuous delivery and mask regressions.

## Solution

Diagnose the exact failure causes across each failing workflow, reproduce them in resource-capped Docker containers, and implement vertical-slice fixes for each distinct problem area. Ship each fix as its own PR against `jonbaldie/gascity`, keep patrol dogs on cadence, and verify all workflows run to green on `origin/main`.

## User Stories

1. As a contributor, I want `Nightly` CI to pass green, so that nightly verification runs succeed.
2. As a contributor, I want `Mac Regression` CI to pass green, so that platform-specific quality gates and process test shards are reliable.
3. As a maintainer, I want `CI` and `Review Formulas` workflows on `main` to run and succeed promptly without getting stuck pending.
4. As a developer, I want all heavy test/build executions to run in resource-capped Docker, so that the Mac host stays cool.
5. As an operator, I want each fix shipped via an isolated PR and verified on origin `main`, so that progress is traceable and bisectable.
6. As a prefect, I want patrol dogs dispatched on cadence, so that host leftovers and tracker state remain clean.

## Implementation Decisions

- **Diagnosis first (Ticket 01):** Run research/investigation subagent to inspect logs and reproduce each failure mode, producing `.scratch/ci-failures-green/diagnosis-report.md`.
- **Vertical-slice fix PRs (Tickets 02–05):** Each fix is developed in its own worktree and branch, tested in Docker with resource caps (`--cpus=2 --memory=4g --pids-limit=256`), and landed via its own PR.
- **Docker enforcement:** All heavy compilation and test execution runs in Docker containers per `host-thermal-steer.md`.
- **Patrol dog cadence:** Dispatch a patrol dog after every `patrolEvery` completed non-patrol tickets (or after major milestones) to keep slag and tracker clean.
- **Verification (Ticket 06):** Final end-to-end trigger and check of all workflows on `main`.

## Testing Decisions

- Good tests assert external behavior at the highest available seam.
- Reproduce the failing tests in capped Docker before applying fixes.
- Verify fixes both locally in Docker and remotely via GitHub Actions runs.

## Out of Scope

- Upstream synchronization or PRs against `gastownhall/gascity`.
- Unrelated feature work.

## Further Notes

- `.outpost.json` caps: `maxWorkers: 2`, `maxSplicers: 1`, `maxRunners: 1`, `maxPatrolDogs: 1`.
- Follow thermal steering in `.scratch/fork-constitution-upstream/host-thermal-steer.md`.
