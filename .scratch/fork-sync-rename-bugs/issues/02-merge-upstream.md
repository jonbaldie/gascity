# 02 — Merge gastownhall/gascity into this fork

**What to build:** This fork's `main` (via a dedicated merge branch and PR) contains current `gastownhall/gascity` `main`, with conflicts resolved so both intents survive: upstream's new work, and this fork's constitution/standards/already-landed fixes. The splicer follows `/resolving-merge-conflicts` rigorously (primary sources, preserve both intents, never `--abort`, then checks). Open one PR against `jonbaldie/gascity`.

**Blocked by:** None — can start immediately.

**Status:** resolved

- [x] `upstream` remote points at `https://github.com/gastownhall/gascity.git` (added if missing) and `upstream/main` was fetched
- [x] Merge is on a dedicated branch from current `origin/main`, not onto a dirty prefect checkout
- [x] Every conflict hunk was resolved from primary sources; both intents preserved where compatible; trade-offs recorded on this ticket when incompatible
- [x] Fork-only identity still present after the merge: `AGENTS.md` constitution, `CLAUDE.md` symlink to it, root `CODING_STANDARDS.md`, and previously landed fork fixes not silently reverted
- [x] Focused tests after the merge ran in resource-capped Docker (`--cpus=2 --memory=4g --pids-limit=256`), not a broad host `make test`
- [x] Dedicated PR opened against `jonbaldie/gascity` for this merge only

## Answer

- **Branch:** `outpost/02-merge-upstream`
- **Worktree:** `/Users/jonathanbaldie/go/src/github.com/jonbaldie/gascity-wt-15`
- **Merge commit:** `a209d41c9b46c7faba74775e82ea3e8bc4a47790`
- **PR:** https://github.com/jonbaldie/gascity/pull/12 (against `jonbaldie/gascity`, not gastownhall)

Merged `upstream/main` (`a6341f8b1`) into fork `origin/main` (`e9945cb4`). `CLAUDE.md` remains a symlink to `AGENTS.md`. Root `CODING_STANDARDS.md` is still present.

### Conflict trade-offs

- **AGENTS.md:** git auto-merged to upstream's evolved constitution. Fork identity for this file is “AGENTS.md is the constitution”; we kept that role and the `CLAUDE.md` symlink rather than restoring origin's older body.
- **Named-session / pool identity:** kept fork `BEADS_ACTOR` stamp in `setTemplateEnvIdentity` and the EnvIdentityStamped start-merge guard, but skipped that guard when `pool_alias_conflict` is set so upstream's empty-alias scrub still applies.
- **Supervisor stop:** kept fork `forceShutdownAsync` and also set upstream's `forceStopShutdown` flag.
- **Nudge:** took upstream submit-key sequences / poke / fail-closed, and kept fork's C-u clear of pending input before paste.
- **tmux startup:** kept both `disableMouseAndActivity` (upstream) and `setAutoRespawnHook` (fork).
- **Doctor:** kept fork `PoolWorkDirIsolationCheck` and `IssuePrefixCheck` on upstream's `register()` path; added `WarmupEligible()`.
- **Duplicate tests:** dropped fork copies of named-session assignee tests that upstream already shipped; kept wrapped-bead order-tracking coverage.

### Tests (Docker, capped)

`docker run --rm --cpus=2 --memory=4g --pids-limit=256 -e CGO_ENABLED=0 golang:1.26`

Passed: `internal/orders`, `internal/runtime/tmux`, `internal/runtime/acp`, `internal/session` (skipped root-only chmod Pi restore test), `internal/doctor` (skipped BeadsRole git-worktree checks), `internal/config` origin-gate tests, `cmd/gc` identity/stop tests.

Did not run broad host `make test`.

## Comments
