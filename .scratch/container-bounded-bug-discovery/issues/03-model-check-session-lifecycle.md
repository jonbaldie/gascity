# 03 — Model-check session lifecycle transitions

**What to build:** Container-bounded generated lifecycle command sequences
compared with a simple reference model at the existing session lifecycle seam;
fix and regress any minimized divergence from the documented transition
contract.

**Blocked by:** 01 — Build a capped verification container harness.

**Status:** resolved

- [x] Generated sequences cover legal and illegal lifecycle commands and assert
  the observable state/effects against a simple reference model.
- [x] A failing sequence records a replayable seed or trace and is minimized
  before being classified as a defect.
- [x] Every confirmed divergence receives a focused fix and deterministic
  regression test; no unconfirmed mismatch is reported as a bug.

## Answer

Added `internal/session/state_machine_model_test.go`, a deterministic bounded
counterexample campaign for the public `session.Transition` seam. The oracle
is an independently spelled-out lifecycle model (it does not read the
production transition table). Each generated command checks legal transitions
for the exact next state; illegal known commands must wrap
`ErrIllegalTransition`, and unknown commands must return an error. A RED path
greedily removes commands while preserving the mismatch and records the seed,
initial state, sequence index, and minimized trace in the artifact report.

Final Docker command:

```sh
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go test -c -o .gc/verification/session-lifecycle-20260820/session-model.test ./internal/session
go run ./cmd/verify-container --workspace "$PWD" --artifact-dir "$PWD/.gc/verification/session-lifecycle-20260820" --cpus 2 --memory 4g --pids 256 --tmpfs-size 1g --timeout 2m -- /workspace/.gc/verification/session-lifecycle-20260820/session-model.test -test.run='^TestTransitionGeneratedReferenceModel$' -test.count=1 -test.v
```

The harness ran with no network, read-only workspace, 2 CPUs, 4 GiB memory,
256 PIDs, 1 GiB tmpfs, and a 2-minute wall-clock cap; it completed in about
5.6 seconds. Budget was 2,048 sequences × 48 steps (98,304 executions):
62,771 legal and 35,533 illegal commands, covering all 12 modeled states and
all 12 generated command values. The saved replay is
`.gc/verification/session-lifecycle-20260820/replay.sh`, and the telemetry is
in `session-lifecycle-campaign.json` in that artifact directory.

Landed on branch `outpost/model-session-lifecycle-03` in commit
`752ebcb52` (`test: model-check session lifecycle transitions`).

Result: bounded no-finding. No minimized RED trace remained after correcting
an initial oracle mistake (the model incorrectly treated an illegal command's
target state as observable even though `Transition` returns an error). No
production divergence was confirmed, so no production fix or regression case
was required.

Blind spots: this campaign does not exercise Manager/store/provider/API side
effects, the reconciler's `StateStartPending` → `StateCreating` pre-start
boundary, concurrency or timing/fault schedules, or states beyond this pure
transition table. The exact returned state value accompanying an error is not
asserted because the documented contract is the error classification; the
model treats the current state as unchanged for subsequent commands.
