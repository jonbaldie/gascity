# 09 — Splice model-checked session lifecycle

**What to build:** Integrate the completed session lifecycle model-check branch
into a clean integration branch and verify its bounded campaign from the
integrated revision, preserving its no-finding report and replay artifacts.

**Blocked by:** 03 — Model-check session lifecycle transitions.

**Status:** resolved

- [x] The Ticket 03 branch is merged into the designated integration branch
  with conflicts resolved according to project conventions.
- [x] The bounded lifecycle campaign and focused package checks are green from
  the integrated revision.
- [x] The integration commit and no-finding evidence are recorded for the
  eventual runner.

## Comments

Spliced `outpost/model-session-lifecycle-03` at `752ebcb52` into
`outpost/splice-model-session-lifecycle-03` with integration commit
`45fa76986` (`merge: integrate model-checked session lifecycle`). The source
change was an independent test file, so the merge completed without conflicts.

Checks from the integrated revision:

- `CGO_ENABLED=0 go test ./internal/session` — PASS.
- `CGO_ENABLED=0 go vet ./internal/session` — PASS.
- `gofmt -d internal/session/state_machine_model_test.go` and `git diff --check` — PASS.
- The default host `go test`/`go vet` attempts were blocked before package
  execution by the host's missing `unicode/regex.h` for `go-icu-regex`; the
  documented CGO-disabled hermetic path passed.

The documented Docker-bounded campaign passed from the integrated revision:

```sh
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go test -c -o .gc/verification/session-lifecycle-20260820/session-model.test ./internal/session
go run ./cmd/verify-container --workspace "$PWD" --artifact-dir "$PWD/.gc/verification/session-lifecycle-20260820" --cpus 2 --memory 4g --pids 256 --tmpfs-size 1g --timeout 2m -- /workspace/.gc/verification/session-lifecycle-20260820/session-model.test -test.run='^TestTransitionGeneratedReferenceModel$' -test.count=1 -test.v
```

Result: seed `202608201503`, 2,048 sequences × 48 steps, 98,304 executions,
62,771 legal and 35,533 illegal steps, all 12 states and all 12 command values
covered, with no RED sequence or minimized trace. Preserved artifacts are in
`/Users/jonathanbaldie/go/src/github.com/jonbaldie/gascity-splice-model-session-lifecycle-03/.gc/verification/session-lifecycle-20260820/`:

- `replay.sh` SHA-256 `b21da0ced07367f7abf2096737129b2b7c1b4ca2256c83b3d2c0a53d59c953ac`
- `session-lifecycle-campaign.json` SHA-256 `659f38a8709c3a861d8391ac88e2e421d6ec3515777cfd5a6d53aefd61500fcb`
- `session-model.test` SHA-256 `a40daf116937825a7e2dce062c8f9ca2b9d552f58d2adebbd13ac730a5e16632`
