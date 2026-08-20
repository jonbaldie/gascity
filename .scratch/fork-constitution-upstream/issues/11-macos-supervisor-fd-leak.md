# 11 — Stop macOS supervisor reload from leaking FDs

**What to build:** Supervisor start/reload cycles on macOS do not leak thousands of file descriptors (or unbounded closed sockets/pipes). After a bounded number of reloads, FD count remains stable enough that the city stays operable.

**Blocked by:** None — can start immediately.

**Status:** resolved

**Upstream:** https://github.com/gastownhall/gascity/issues/4504 · related https://github.com/gastownhall/gascity/issues/5385

- [x] Reload/start cycle does not monotonically accumulate FDs across repeated reloads in a repro
- [x] Fix is verified with a measurement approach suitable for macOS (e.g. `lsof`/FD count delta)
- [x] Dedicated branch + PR against `jonbaldie/gascity` for this ticket only

## Answer

Root cause: `fsnotify` v1.9.0 on Darwin closes the kqueue on `Watcher.Close()` but leaks watched REG/DIR descriptors (fsnotify#732). Supervisor config reload recreates the recursive watcher, so each reload retained ~one FD per watched path until EMFILE.

Fix: bump to `fsnotify` v1.10.1 (fsnotify#740) and add `TestWatchConfigTargets_DarwinReloadDoesNotLeakFDs`, which measures unique FDs via `lsof -nP -p <pid> -F f` across recursive watcher restarts.

- Branch: `outpost/11-macos-supervisor-fd-leak`
- PR: https://github.com/jonbaldie/gascity/pull/10
- Commit: `4c9bf34d`

Note: upstream #5385 (PIPE/socket/`/dev/null` retention under session churn) is a distinct leak profile and was not in scope for this ticket.

## Comments

- 2026-08-19: Claimed by outpost worker; Darwin repro went red on v1.9.0 (~244 FDs/restart over 80 dirs) and green after the bump.
