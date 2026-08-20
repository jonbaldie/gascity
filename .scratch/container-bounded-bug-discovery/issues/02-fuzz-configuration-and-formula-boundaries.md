# 02 — Fuzz configuration and formula boundaries

**What to build:** Coverage-guided, container-bounded exploration of public
configuration and formula input boundaries, with assertions for parse,
validation, and round-trip contracts; fix and regress any minimized,
reproducible counterexample.

**Blocked by:** 01 — Build a capped verification container harness.

**Status:** resolved

- [x] The selected public formula boundary runs through the capped harness with
  coverage feedback and explicit invariants.
- [x] Any candidate failure is reproduced, minimized, and deduplicated before
  it is fixed or reported.
- [x] Every confirmed defect has a deterministic regression test; runs with no
  confirmed counterexample report that outcome without treating coverage as
  proof.

## Comments

- Added `internal/config/config_fuzz_test.go` for the public `config.Parse` /
  `City.Marshal` boundary and `internal/formula/parser_fuzz_test.go` for
  `Parser.ParseTOML` / `Formula.Validate` / canonical TOML encoding. Accepted
  inputs must parse, validate where applicable, and be byte-stable after a
  second canonical encode; parser and validation errors are expected outcomes
  for arbitrary bytes.
- Local bounded campaigns used `CGO_ENABLED=0`, one worker, and 10 seconds per
  boundary. Config: 20,128 executions, 52 new interesting inputs (55 total).
  Formula: 39,518 executions, 49 new interesting inputs (52 total). Both passed
  with no failing or minimized inputs.
- Formula capped campaign (Go 1.26.6 image, 1 CPU, 1 GiB memory, 128 PIDs,
  512 MiB tmpfs, 2-minute wall clock, 15-second fuzz budget) ran through
  `cmd/verify-container` with `GOTMPDIR=/artifacts/tmp` and an offline module
  cache archive: baseline 3/3, 359,851 executions, 106 new interesting inputs
  (109 total), PASS. Coverage-guided feedback was active.

  Exact command:

  ```text
  CGO_ENABLED=0 go run ./cmd/verify-container --image golang:1.26.6 --workspace "$PWD" --artifact-dir "$PWD/.gc/verification/fuzz-config-02-formula" --cpus 1 --memory 1g --pids 128 --tmpfs-size 512m --timeout 2m -- sh -c 'mkdir -p /artifacts/tmp /tmp/gomod && tar -xf /artifacts/modules.tar -C /tmp/gomod && GOTMPDIR=/artifacts/tmp GOMODCACHE=/tmp/gomod GOTOOLCHAIN=local go test ./internal/formula -run "^$" -fuzz FuzzParseTOMLValidateRoundTrip -fuzztime=15s -parallel=1'
  ```
- Config was also run in the capped harness with a Linux/arm64 test binary and
  the same 15-second budget: 1,135,305 executions, PASS, but Go reported that a
  precompiled fuzz binary lacked coverage guidance. A source-built config run
  was blocked by the network-none container's large Dolt dependency graph even
  after compatible module-cache staging. This is recorded as a harness blind
  spot, not as correctness evidence; the formula run is the coverage-guided
  boundary campaign.

  Exact supplementary command:

  ```text
  CGO_ENABLED=0 go run ./cmd/verify-container --image golang:1.26.6 --workspace "$PWD" --artifact-dir "$PWD/.gc/verification/fuzz-config-02-config" --cpus 1 --memory 1g --pids 128 --tmpfs-size 512m --timeout 2m -- /artifacts/config.test -test.run '^$' -test.fuzz FuzzParseMarshalRoundTrip -test.fuzztime 15s -test.fuzzcachedir /tmp/fuzzcache -test.parallel 1
  ```
- No reproducible counterexample was found, so there are no RED cases to fix or
  deduplicate. Coverage is treated as search telemetry, not proof of
  correctness. The default `golang:1.25.9` image is also stale versus the
  repository's Go 1.26.6 requirement and was bypassed with `--image
  golang:1.26.6`.
