# Outpost host thermal steer (human)

**Effective immediately for this effort:** heavy compile/test/integration work MUST run inside Docker with hard resource caps. Keep the Mac host cool — prefer editing on the host, execute builds/tests in containers.

## Caps (defaults)

```bash
DOCKER_CPUS=2
DOCKER_MEMORY=4g
DOCKER_PIDS=256
```

## Pattern

From a ticket worktree (bind-mount the worktree, not the whole machine):

```bash
IMG=golang:1.25  # match go.mod; adjust if module version differs
docker run --rm \
  --cpus="${DOCKER_CPUS:-2}" \
  --memory="${DOCKER_MEMORY:-4g}" \
  --pids-limit="${DOCKER_PIDS:-256}" \
  -v "$PWD":/src -w /src \
  -e GOCACHE=/tmp/go-cache -e GOMODCACHE=/tmp/go-mod \
  "$IMG" \
  go test ./path/under/test -count=1
```

## Rules

- Cap concurrent workers at **2** (see `.outpost.json` `maxWorkers`).
- Do not run broad `make test` / large parallel shards on the host.
- Prefer focused package tests in Docker; skip host `-race` unless inside a capped container.
- `gh` / git / tracker edits may stay on the host (light I/O).
- If Docker is unavailable, `/outpost-blocked` rather than burning the host CPU.
