# ZV-002 — Skeleton Diff Report

**Generated against:** Zevaro-Architecture.md §3.1 (patched after ZV-001 to add `internal/constants/` and `frontend/src/assets/fonts/`)
**Repository state at:** `4c48653c594cb0cf8eef59972fb28544205dc37c` (HEAD before ZV-002 started)

---

## Summary

| Metric | Count |
|--------|-------|
| Declared §3.1 entries verified | **149** (102 directories + 47 files) |
| PASS | **130** (99 directories + 31 files) |
| MISSING (created by this task) | **3** (`internal/constants/`, `frontend/src/assets/`, `frontend/src/assets/fonts/`) |
| MISSING (NOT created — future tasks, expected) | **16** (files listed in §3.1 produced by ZV-003 through ZV-059) |
| EXTRA (in repo, not in §3.1) | **14** notable items (see Extras section) |

---

## Directory Verification

All 102 directories declared in §3.1.

| Path (per §3.1) | Status | Notes |
|---|---|---|
| `.github/` | PASS | |
| `.github/ISSUE_TEMPLATE/` | PASS | |
| `.github/workflows/` | REMOVED (ZV-CI-REMOVE) | CI removed by ZV-CI-REMOVE; directory deleted |
| `api/` | PASS | |
| `api/openapi/` | PASS | Contains `.gitkeep` |
| `cmd/` | PASS | |
| `cmd/zevaro/` | PASS | |
| `docs/` | PASS | |
| `docs/public/` | PASS | Contains `.gitkeep` |
| `docs/src/` | PASS | |
| `docs/src/components/` | PASS | Contains `.gitkeep` |
| `docs/src/content/` | PASS | |
| `docs/src/content/architecture/` | PASS | Contains `.gitkeep` |
| `docs/src/content/clients/` | PASS | Contains `.gitkeep` |
| `docs/src/content/faq/` | PASS | Contains `.gitkeep` |
| `docs/src/content/getting-started/` | PASS | Contains `.gitkeep` |
| `docs/src/content/providers/` | PASS | Contains `.gitkeep` |
| `docs/src/content/routing/` | PASS | Contains `.gitkeep` |
| `docs/src/content/troubleshooting/` | PASS | Contains `.gitkeep` |
| `frontend/` | PASS | |
| `frontend/src/` | PASS | |
| `frontend/src/assets/` | MISSING (created) | Added by ZV-002 per post-ZV-001 patch |
| `frontend/src/assets/fonts/` | MISSING (created) | Added by ZV-002 per post-ZV-001 patch; font files bundled by ZV-041 |
| `frontend/src/components/` | PASS | |
| `frontend/src/components/charts/` | PASS | Contains `.gitkeep` |
| `frontend/src/components/dialogs/` | PASS | Contains `.gitkeep` |
| `frontend/src/components/forms/` | PASS | Contains `.gitkeep` |
| `frontend/src/components/layout/` | PASS | Contains `.gitkeep` |
| `frontend/src/components/tables/` | PASS | Contains `.gitkeep` |
| `frontend/src/lib/` | PASS | Contains `.gitkeep` |
| `frontend/src/routes/` | PASS | |
| `frontend/src/routes/budgets/` | PASS | Contains `.gitkeep` |
| `frontend/src/routes/compare/` | PASS | Contains `.gitkeep` |
| `frontend/src/routes/dashboard/` | PASS | Contains `.gitkeep` |
| `frontend/src/routes/history/` | PASS | Contains `.gitkeep` |
| `frontend/src/routes/onboarding/` | PASS | Contains `.gitkeep` |
| `frontend/src/routes/providers/` | PASS | Contains `.gitkeep` |
| `frontend/src/routes/routing/` | PASS | Contains `.gitkeep` |
| `frontend/src/routes/settings/` | PASS | Contains `.gitkeep` |
| `frontend/src/state/` | PASS | Contains `.gitkeep` |
| `frontend/src/styles/` | PASS | Contains `globals.css` |
| `installers/` | PASS | |
| `installers/linux/` | PASS | Contains `.gitkeep` |
| `installers/macos/` | PASS | Contains `.gitkeep` |
| `installers/windows/` | PASS | Contains `.gitkeep` |
| `internal/` | PASS | |
| `internal/api/` | PASS | |
| `internal/api/anthropic/` | PASS | Contains `.gitkeep` |
| `internal/api/management/` | PASS | Contains `.gitkeep` |
| `internal/api/openai/` | PASS | Contains `.gitkeep` |
| `internal/api/streaming/` | PASS | Contains `.gitkeep` |
| `internal/app/` | PASS | Contains `app.go`, `app_test.go` |
| `internal/audit/` | PASS | Contains `.gitkeep` |
| `internal/budgets/` | PASS | Contains `.gitkeep` |
| `internal/cache/` | PASS | Contains `.gitkeep` |
| `internal/compare/` | PASS | Contains `.gitkeep` |
| `internal/config/` | PASS | |
| `internal/config/project/` | PASS | Contains `.gitkeep` |
| `internal/constants/` | MISSING (created) | Added by ZV-002 per post-ZV-001 patch; constants populated as needed by later tasks |
| `internal/discovery/` | PASS | Contains `.gitkeep` |
| `internal/history/` | PASS | |
| `internal/history/search/` | PASS | Contains `.gitkeep` |
| `internal/mcp/` | PASS | Contains `.gitkeep` |
| `internal/normalize/` | PASS | |
| `internal/normalize/errors/` | PASS | Contains `.gitkeep` |
| `internal/normalize/messages/` | PASS | Contains `.gitkeep` |
| `internal/normalize/tools/` | PASS | Contains `.gitkeep` |
| `internal/offline/` | PASS | Contains `.gitkeep` |
| `internal/plugins/` | PASS | Contains `.gitkeep` |
| `internal/privacy/` | PASS | Contains `.gitkeep` |
| `internal/providers/` | PASS | |
| `internal/providers/cloud/` | PASS | Contains `.gitkeep` |
| `internal/providers/freetier/` | PASS | Contains `.gitkeep` |
| `internal/providers/local/` | PASS | Contains `.gitkeep` |
| `internal/routing/` | PASS | |
| `internal/routing/auto/` | PASS | Contains `.gitkeep` |
| `internal/routing/balance/` | PASS | Contains `.gitkeep` |
| `internal/routing/embeddings/` | PASS | Contains `.gitkeep` |
| `internal/routing/fallback/` | PASS | Contains `.gitkeep` |
| `internal/routing/images/` | PASS | Contains `.gitkeep` |
| `internal/routing/policies/` | PASS | Contains `.gitkeep` |
| `internal/routing/rules/` | PASS | Contains `.gitkeep` |
| `internal/routing/tags/` | PASS | Contains `.gitkeep` |
| `internal/server/` | PASS | Contains `.gitkeep`; `server.go` and `healthz.go` added by ZV-005 |
| `internal/server/middleware/` | PASS | Contains `.gitkeep`; content added by ZV-005 |
| `internal/sharing/` | PASS | Contains `.gitkeep` |
| `internal/spend/` | PASS | |
| `internal/spend/pricing/` | PASS | Contains `.gitkeep` |
| `internal/storage/` | PASS | |
| `internal/storage/entities/` | PASS | Contains `.gitkeep` |
| `internal/storage/platform/` | PASS | Contains `.gitkeep` |
| `internal/telemetry/` | PASS | Contains `doc.go` (intentionally empty sentinel package per §8) |
| `internal/updater/` | PASS | Contains `.gitkeep` |
| `pkg/` | PASS | |
| `pkg/pricing-data/` | PASS | Contains `.gitkeep` |
| `pkg/zevaro-plugin/` | PASS | Contains `.gitkeep` |
| `scripts/` | PASS | |
| `test/` | PASS | |
| `test/e2e/` | PASS | Contains `.gitkeep` |
| `test/fixtures/` | PASS | Contains `.gitkeep` |
| `test/helpers/` | PASS | Contains `.gitkeep` |
| `test/integration/` | PASS | Contains `.gitkeep` |

**Directory totals:** 99 PASS · 3 MISSING (all created by this task)

---

## File Verification

All 47 files declared by name in §3.1.

### Root-level files

| Path (per §3.1) | Status | Notes |
|---|---|---|
| `.editorconfig` | PASS | |
| `.gitignore` | PASS | |
| `.golangci.yml` | PASS | |
| `.goreleaser.yaml` | PASS | |
| `CLAUDE.md` | PASS | |
| `CONTRIBUTING.md` | PASS | |
| `CONVENTIONS.md` | PASS | |
| `go.mod` | PASS | |
| `go.sum` | PASS | Listed as `go.mod / go.sum` in §3.1 |
| `LICENSE` | PASS | |
| `Makefile` | PASS | |
| `NOTICE` | PASS | Was absent before ZV-001; created by ZV-001 (Observation #1) |
| `openapi.yaml` | MISSING (not created) | Produced by ZV-003; absence expected at this phase |
| `README.md` | PASS | |
| `Zevaro-Architecture.md` | PASS | |
| `Zevaro-Audit.md` | MISSING (not created) | Produced by ZV-004; absence expected at this phase |

### `.github/` files

| Path (per §3.1) | Status | Notes |
|---|---|---|
| `.github/PULL_REQUEST_TEMPLATE.md` | PASS | |
| `.github/workflows/ci.yaml` | REMOVED (ZV-CI-REMOVE) | CI removed; no workflow files in project |
| `.github/workflows/nightly.yaml` | REMOVED (ZV-CI-REMOVE) | CI removed |
| `.github/workflows/release.yaml` | REMOVED (ZV-CI-REMOVE) | CI removed; ZV-057 repurposed to local release script |

### `cmd/zevaro/` files

| Path (per §3.1) | Status | Notes |
|---|---|---|
| `cmd/zevaro/main.go` | PASS | |

### `internal/server/` files

| Path (per §3.1) | Status | Notes |
|---|---|---|
| `internal/server/server.go` | MISSING (not created) | Implemented by ZV-005; directory placeholder exists |
| `internal/server/healthz.go` | MISSING (not created) | Implemented by ZV-005; directory placeholder exists |

### `internal/providers/` files

| Path (per §3.1) | Status | Notes |
|---|---|---|
| `internal/providers/provider.go` | MISSING (not created) | Implemented by ZV-008; directory placeholder exists |
| `internal/providers/registry.go` | MISSING (not created) | Implemented by ZV-008; directory placeholder exists |

### `internal/routing/` files

| Path (per §3.1) | Status | Notes |
|---|---|---|
| `internal/routing/engine.go` | MISSING (not created) | Implemented by ZV-015; directory placeholder exists |

### `internal/storage/` files

| Path (per §3.1) | Status | Notes |
|---|---|---|
| `internal/storage/db.go` | MISSING (not created) | Implemented by ZV-007; directory placeholder exists |

### `internal/config/` files

| Path (per §3.1) | Status | Notes |
|---|---|---|
| `internal/config/config.go` | MISSING (not created) | Implemented by ZV-006; directory placeholder exists |
| `internal/config/schema.go` | MISSING (not created) | Implemented by ZV-006; directory placeholder exists |

### `internal/history/` files

| Path (per §3.1) | Status | Notes |
|---|---|---|
| `internal/history/store.go` | MISSING (not created) | Implemented by ZV-023; directory placeholder exists |
| `internal/history/retention.go` | MISSING (not created) | Implemented by ZV-023; directory placeholder exists |

### `internal/spend/` files

| Path (per §3.1) | Status | Notes |
|---|---|---|
| `internal/spend/tracker.go` | MISSING (not created) | Implemented by ZV-024; directory placeholder exists |
| `internal/spend/aggregate.go` | MISSING (not created) | Implemented by ZV-024; directory placeholder exists |
| `internal/spend/projection.go` | MISSING (not created) | Implemented by ZV-024; directory placeholder exists |

### `frontend/` files

| Path (per §3.1) | Status | Notes |
|---|---|---|
| `frontend/index.html` | PASS | |
| `frontend/package.json` | PASS | |
| `frontend/tailwind.config.js` | PASS | |
| `frontend/tsconfig.json` | PASS | |
| `frontend/vite.config.ts` | PASS | |

### `frontend/src/` files

| Path (per §3.1) | Status | Notes |
|---|---|---|
| `frontend/src/app.tsx` | PASS | Renders "Hello Zevaro" splash; full shell implemented ZV-041 |
| `frontend/src/main.tsx` | PASS | |

### `frontend/src/styles/` files

| Path (per §3.1) | Status | Notes |
|---|---|---|
| `frontend/src/styles/globals.css` | PASS | Full §5 CSS variable palette defined |

### `docs/` files

| Path (per §3.1) | Status | Notes |
|---|---|---|
| `docs/astro.config.mjs` | MISSING (not created) | Implemented by ZV-059 (documentation site task) |

### `scripts/` files

| Path (per §3.1) | Status | Notes |
|---|---|---|
| `scripts/lint.sh` | PASS | |
| `scripts/notarize-macos.sh` | PASS | Stub; full implementation ZV-050 |
| `scripts/release.sh` | PASS | Stub; full implementation ZV-057 |
| `scripts/sign-windows.ps1` | PASS | Stub; full implementation ZV-051 |

**File totals:** 31 PASS · 16 MISSING (none created by this task — all are expected future-task deliverables)

---

## Extras (in repo, not in §3.1)

Items present in the repository that are NOT listed in §3.1. Items excluded per task rules: `.git/`, `node_modules/`, `build/`, `*.lock` files, `npm-placeholder/`. Gitignored artifacts (`.DS_Store`, `coverage.out`, `cmd/zevaro/frontend/dist/`, `frontend/dist/`) are also excluded.

| Path | Classification | Notes |
|---|---|---|
| `wails.json` | EXTRA | Required Wails v2 project config; not listed in §3.1. Recommend adding to §3.1 in ZV-002 or future architecture review. |
| `frontend/postcss.config.js` | EXTRA | Required for Tailwind CSS processing; not listed in §3.1. Standard Vite+Tailwind tooling artifact. |
| `frontend/eslint.config.js` | EXTRA | ESLint flat config; not listed in §3.1. Required by `pnpm lint` target. |
| `cmd/zevaro/embed_desktop.go` | EXTRA | Build-tag conditioned asset embedding for desktop production builds. Not listed in §3.1 (§3.1 only names `main.go`). Expected — required by the `cmd/` layout pattern. |
| `cmd/zevaro/embed_dev.go` | EXTRA | Build-tag conditioned empty embed for dev/test mode. See above. |
| `cmd/zevaro/main_test.go` | EXTRA | Smoke test for main wiring. §3.1 names `main.go` only; test files are expected alongside source. |
| `internal/app/app.go` | EXTRA | Implementation of the Wails App struct. §3.1 lists the `internal/app/` directory but not the files within it at skeleton stage. Expected. |
| `internal/app/app_test.go` | EXTRA | Test for `app.go`. Expected companion to implementation. |
| `internal/telemetry/doc.go` | EXTRA | Intentional sentinel package documenting "no telemetry." Mandated by §8 but not listed as a named file in §3.1. Expected. |
| `frontend/src/app.test.tsx` | EXTRA | Vitest test for the App component. Expected companion to `app.tsx`. |
| `frontend/src/test-setup.ts` | EXTRA | Vitest + Testing Library setup file. Expected tooling artifact. |
| `docs/audit/ZV-002-skeleton-diff.md` | EXTRA | This file. Created by ZV-002 as its primary deliverable. `docs/audit/` is not in §3.1 — recommend adding to §3.1 as a persistent audit directory. |
| `.gitkeep` files (70 instances) | EXTRA | Placeholder files in every empty §3.1 directory to enable git tracking. Not listed in §3.1 individually. Expected housekeeping artifacts. |
| `frontend/.npmrc`, `frontend/.pnpmfile.cjs`, `frontend/.pnpm-scripts-approval.json` | EXTRA | pnpm configuration artifacts written during ZV-001 setup. Should be reviewed — `.npmrc` and `.pnpmfile.cjs` are gitignored per `.gitignore` but appear in the working tree (may not be committed). |

---

## Notes

### On §3.1 file listings vs. directory listings

§3.1 lists specific file names for several packages (`server.go`, `healthz.go`, `provider.go`, `registry.go`, `engine.go`, `db.go`, `config.go`, `schema.go`, `store.go`, `retention.go`, `tracker.go`, `aggregate.go`, `projection.go`). These are named in the architecture to make the intended structure clear, but each is produced by the task that implements that package. All 13 implementation files are MISSING in the expected sense — their parent directories exist with `.gitkeep` placeholders, and the files land when the corresponding ZV task runs.

### On §3.1 directories existing vs. having content

Several directories are classified PASS because they exist on disk, but they contain only `.gitkeep`. This is correct behaviour at the skeleton phase. The table notes which directories are in this state so future task authors know what to expect when they arrive.

### On `frontend/src/assets/` vs. `frontend/src/assets/fonts/`

§3.1 lists `frontend/src/assets/` as a directory and `frontend/src/assets/fonts/` as its subdirectory. The patch created both as missing after ZV-001 ran. ZV-002 creates `frontend/src/assets/fonts/.gitkeep`; the `assets/` directory is implicitly created as its parent. No `.gitkeep` is placed directly in `assets/` since it has a committed subdirectory.

### On `wails.json`

This file is a required Wails v2 project manifest and was created by ZV-001. It is not listed in §3.1. This is a gap in the architecture spec rather than an error in ZV-001. Recommend the architect add `wails.json` to the root-level file list in §3.1 during ZV-002 or the next architecture review.

### On `docs/audit/` and this file

The `docs/audit/` directory holding this report is itself not in §3.1. As audit artefacts accumulate across ZV tasks (ZV-004, ZV-009, ZV-014, etc. produce `Zevaro-Audit.md`), a stable home for diff reports and related audit artefacts is useful. Recommend the architect add `docs/audit/` to §3.1.

### On pnpm config artifacts

Three pnpm-related files (`frontend/.npmrc`, `frontend/.pnpmfile.cjs`, `frontend/.pnpm-scripts-approval.json`) appear in the working tree. The `.gitignore` includes `.npmrc` and `.pnpmfile.cjs` at the root level but may not cover the `frontend/` subdirectory. These should be verified as gitignored and not accidentally committed.

---

*Report produced by ZV-002. All 149 declared §3.1 entries have been classified. No unclassified entries remain.*
