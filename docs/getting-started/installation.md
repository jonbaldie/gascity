---
title: Installation
description: Install Gas City with go install, or build from source.
---

## Which method should I use?

| Method | Best for | Installs deps? | Auto-upgrades? |
|--------|----------|----------------|----------------|
| [go install](#go-install-recommended) | Daily use | No | Re-run `go install` |
| [Source build](#build-from-source) | Contributors, local checkouts | No | Manual |

**Most users should use `go install`.** It builds `gc` from this repository
and puts it on your PATH. Choose source when you already cloned this repo
and want `make install` from that checkout.

This fork does not publish a Homebrew tap or GitHub release tarballs. Do not
install from `gastownhall/gascity` if you want this project.

## Prerequisites

Gas City requires a small set of runtime tools. `go install` installs only
the `gc` binary; install the tools below separately.

| Tool | Required | Min version | macOS | Linux | Notes |
|------|----------|-------------|-------|-------|-------|
| tmux | Yes | — | `brew install tmux` | `apt install tmux` | Session management |
| jq | Yes | — | `brew install jq` | `apt install jq` | JSON processing |
| git | Yes | — | (built-in) | (built-in) | Version control |
| Go 1.26+ | Yes | 1.26 | `brew install go` | [golang.org](https://go.dev/dl/) | Compiler (`go install` and source) |
| dolt | Beads provider `bd` | 2.1.0 or newer | `brew install dolt` | [releases](https://github.com/dolthub/dolt/releases) | Beads data plane |
| bd (Beads CLI) | Beads provider `bd` | 1.0.0 | `go install github.com/jonbaldie/beads/cmd/bd@latest` | `go install github.com/jonbaldie/beads/cmd/bd@latest` | Issue tracking |
| flock | Beads provider `bd` | — | `brew install flock` | (built-in via util-linux) | File locking |
| gh | Optional | — | `brew install gh` | [cli.github.com](https://cli.github.com/) | GitHub gate checks |
| make | Source only | — | (built-in) | `apt install make` (or `build-essential`) | Drives `make install` |

The `bd` (beads) provider is the default. To use a file-based store instead
(no dolt/bd/flock needed), set `GC_BEADS=file` or add `[beads] provider = "file"`
to your `city.toml`.

Use a final Dolt 2.1.0 or newer. Gas City's managed Dolt checks reject older
and pre-release builds because they are below the managed bd/Dolt compatibility
floor; releases before 1.86.2 can also miss the upstream GC/writer deadlock
fix in dolthub/dolt commit `ccf7bde206`, which can hang `dolt_backup sync`
under heavy write load.

The exact versions CI pins are in [`deps.env`](https://github.com/jonbaldie/gascity/blob/main/deps.env).

## go install (recommended)

Requires Go 1.26+ (the version is pinned in `go.mod`).

```bash
go install github.com/jonbaldie/gascity/cmd/gc@latest
gc version
```

`go install` places `gc` in `$(go env GOPATH)/bin` (often `~/go/bin`). Put
that directory on your PATH if `gc version` is not found.

<Warning>
If you use Oh My Zsh with the `git` plugin, `gc` may already be an alias for
`git commit --verbose`. Run `command gc version` once to bypass the alias. For
a persistent fix, add `unalias gc 2>/dev/null` or
`zstyle ':omz:plugins:git' aliases no 'gc'` after Oh My Zsh loads in
`~/.zshrc`, or put that line in a file such as
`~/.oh-my-zsh/custom/gascity.zsh`.
</Warning>

### Upgrading a go-install install

Re-run the same `go install` command. After upgrading, restart any running
city so the supervisor picks up the new binary:

```bash
gc service restart     # restarts the launchd/systemd service
```

`gc start` auto-regenerates the service file on each invocation, so a
reinstall followed by `gc start` always picks up template changes.

## Build from source

Requires `make` and Go 1.26+ (pinned in `go.mod`). This is the contributor
path for a clone of this repository; daily installs should use
`go install` above.

```bash
git clone https://github.com/jonbaldie/gascity.git
cd gascity
make install        # builds and installs to $(GOPATH)/bin/gc
gc version
```

To build without installing globally:

```bash
make build          # outputs bin/gc in the repo root
./bin/gc version
```

On macOS, `make build` signs the binary with a stable local codesigning
identity when one is available, which helps macOS remember local permission
grants across rebuilds. Without a stable identity, the build leaves Go's
linker-produced signature unchanged. Set `GC_SIGN_IDENTITY=<certificate name>`
to choose a specific certificate, `GC_SIGN_IDENTIFIER=<identifier>` to use a
separate local TCC identity, or `GC_ADHOC_SIGN=1` to opt into ad-hoc signing
for a local experiment. Successful local signing also removes stale
`com.apple.provenance` metadata when present.

On macOS, `make build` needs ICU for a transitive Dolt CGO dependency —
`brew install icu4c`. On Linux, `apt install libicu-dev`. On macOS the
Makefile auto-detects the keg-only `icu4c` paths.

### Contributor setup

After building, install the dev toolchain and pre-commit hooks:

```bash
make setup
make check          # runs fmt, lint, vet, and unit tests
```

See [CONTRIBUTING.md](https://github.com/jonbaldie/gascity/blob/main/CONTRIBUTING.md)
for the full contributor workflow.

## Verify your installation

Regardless of install method, confirm everything is working:

```bash
gc version          # should print the installed version and commit
```

If that runs `git commit` instead of Gas City, your shell has a `gc` alias.
Use `command gc version` for this check and see
[Troubleshooting](/getting-started/troubleshooting#oh-my-zsh-git-plugin-hides-gc)
for the permanent fix.

Then create your first city:

```bash
gc init ~/my-city
cd ~/my-city
```

`gc init` registers the city with the supervisor, which then starts it. By the
time the command returns, the city is running.
See the [Quickstart](/getting-started/quickstart) for a complete walkthrough.

Gas City ships a JSONL archive that snapshots every bead database for
disaster recovery. By default it runs in local-only mode and keeps commits
on this host. To enable off-box backup, see
[JSONL archive push failures](/getting-started/troubleshooting#jsonl-archive-push-failures).

## Docs preview

The docs site uses [Mintlify](https://mintlify.com). Preview locally from the
repo root:

```bash
./mint.sh dev
```

Or run a link check without starting the server:

```bash
make check-docs
```
