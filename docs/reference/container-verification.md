---
title: Container verification
description: Run exploratory verification commands with Docker resource limits and replay artifacts.
---

`verify-container` runs an exploratory command in Docker without changing the normal project test targets. It mounts the repository read-only at `/workspace`, provides a bounded writable `/tmp`, and places generated inputs, seeds, and schedules in the host artifact directory mounted at `/artifacts`.

Run it from the repository root:

```sh
go run ./cmd/verify-container -- go test -fuzz=FuzzParse -fuzztime=10m ./internal/config
```

The defaults cap the run at 2 CPUs, 4 GiB of memory, 256 processes, 1 GiB of temporary storage, and 15 minutes of wall-clock time. The command has no network access. Override a limit only for the exploratory run that needs it:

```sh
go run ./cmd/verify-container \
  --cpus 1.5 \
  --memory 768m \
  --pids 64 \
  --tmpfs-size 256m \
  --timeout 7m \
  -- go test ./internal/config -run TestParser
```

The harness prints its artifact directory and writes an executable `replay.sh` there before starting Docker. The command runs with `GC_VERIFICATION_ARTIFACTS=/artifacts`; exploratory tools should write their generated input, seed, or schedule files beneath that location. Re-run the saved script from the same checkout to repeat the bounded invocation.
