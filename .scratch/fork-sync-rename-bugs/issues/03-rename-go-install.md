# 03 — Rename gastownhall/gascity → jonbaldie/gascity and recommend go install

**What to build:** This project identifies as `jonbaldie/gascity` end to end (Go module path and this-repo clone/install URLs). User-facing install instructions that would fail for the fork (Homebrew tap `gastownhall/gascity`, upstream release tarballs/attestations as the way to install *this* project) are removed or demoted. The documented default install is `go install github.com/jonbaldie/gascity/cmd/gc@latest`. Upstream *issue* citations and non-gascity gastownhall modules stay put. Open one PR against `jonbaldie/gascity`.

**Blocked by:** 02 — Merge gastownhall/gascity into this fork (rename once on the synced tree; start from landed `origin/main` after 02 merges, or from 02's merge branch if that PR is not yet landed).

**Status:** resolved

- [x] `go.mod` module path is `github.com/jonbaldie/gascity` and in-tree imports compile against it
- [x] User-facing install docs recommend `go install github.com/jonbaldie/gascity/cmd/gc@latest` as the default
- [x] Homebrew tap / upstream GitHub release-download / attestation instructions are no longer recommended as the way to install this fork
- [x] Clone/source URLs meant for *this* project point at `jonbaldie/gascity`
- [x] `CODING_STANDARDS.md` no longer tells contributors to keep the upstream module path
- [x] Citations of upstream gascity issues/PRs are unchanged; Beads identity is left for ticket 04 (do not retarget `bd` URLs in this PR)
- [x] Focused compile/tests ran in resource-capped Docker, not a broad host `make test`
- [x] Dedicated PR opened against `jonbaldie/gascity` for this ticket only

## Comments

- 2026-08-19: Prefect: ticket 02 landed on `origin/main` (PR #12). Start from that tip.
- 2026-08-19: Worker: claimed. Working in `/Users/jonathanbaldie/go/src/github.com/jonbaldie/gascity-wt-16` on `outpost/03-rename-go-install` from `6e8178493`.
- 2026-08-19: Worker: done. PR https://github.com/jonbaldie/gascity/pull/13 — commit `f339f4861` on `outpost/03-rename-go-install`. Docker: `go test ./test/docsync` and `go test ./scripts -run TestForkVerifyRunsOnlyInForks`; `go list -m` is `github.com/jonbaldie/gascity`. Beads URLs not retargeted.

## Answer

This fork now identifies as `github.com/jonbaldie/gascity`. Default install is `go install github.com/jonbaldie/gascity/cmd/gc@latest`. PR: https://github.com/jonbaldie/gascity/pull/13
