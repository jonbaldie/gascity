# 01 — Diagnose & Map Root Causes for Main CI Failures

**What to build:** Comprehensive diagnosis of all failing/stuck workflows on `origin/main` (`Nightly`, `Mac Regression`, stuck `CI`/`Review Formulas`), capturing error logs, identifying root causes, and producing a diagnosis artifact at `.scratch/ci-failures-green/diagnosis-report.md`. If additional distinct bugs are found, telescope them as new tickets.

**Blocked by:** None — can start immediately.

**Status:** done

- [x] Detailed failure logs captured and analyzed for `Nightly` failures (Ollama Claude config validation & SQLite Tier A acceptance)
- [x] Detailed failure logs captured and analyzed for `Mac Regression` failures (`Mac / quality` & `cmd/gc process` shards 3, 5, 10)
- [x] Investigation completed for why `CI` and `Review Formulas` workflows are stuck pending on `main`
- [x] Root causes and proposed fixes documented in `.scratch/ci-failures-green/diagnosis-report.md`
- [x] Follow-on tickets updated with precise diagnostic details

## Comments

Diagnosis completed and documented in `.scratch/ci-failures-green/diagnosis-report.md`. Follow-on tickets 02, 03, 04, and 05 updated with exact root causes and fix plans.
