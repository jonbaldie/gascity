# 03 — Fix Nightly Ollama Claude Configuration Validation

**What to build:** Fix the Ollama Claude configuration validation failure in `Tier C acceptance tests (Ollama)` in the Nightly workflow. Test fixes in capped Docker container and ship via PR.

**Blocked by:** 01 — Diagnose & Map Root Causes for Main CI Failures.

**Status:** ready-for-agent

- [ ] `Validate Ollama Claude configuration` step updated to gracefully skip when `OLLAMA_API_KEY` secret is not configured on the repository.
- [ ] `Tier C Ollama inference tests` step guarded with `if: env.OLLAMA_API_KEY != ''` or internal credential check.
- [ ] PR opened against `jonbaldie/gascity` with isolated fix
- [ ] All checks pass on PR

## Comments

### Root Cause Details
- `tier-c` job in `.github/workflows/nightly.yml` unconditionally executes `test -n "$OLLAMA_API_KEY" || { echo "Missing OLLAMA_API_KEY GitHub secret" >&2; exit 1; }`.
- On the fork repository, `OLLAMA_API_KEY` is not set, causing the validation step to fail with exit code 1.
- Fix: Guard the validation and inference steps so they skip gracefully when credentials are not present.
