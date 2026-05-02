# Zevaro Roadmap

A flat index of every task required to ship Zevaro v1. v1 contains the entire charter scope — both core capabilities and adjacent capabilities.

This document is a quick-reference overview. The **canonical task descriptions** live in `Zevaro-Architecture.md` §6. **Live status** (which tasks are done, in progress, or blocked) lives in `ZEVARO_TASK_TRACKER.md`. When this document conflicts with either of those, they win — this file is just an at-a-glance index.

---

## How this roadmap works

This is **not** a prioritized backlog, a phased plan, or a sprint roadmap. It is a strict dependency order: each task depends on artifacts produced by earlier tasks. Tasks higher in the list must complete before tasks lower in the list can begin. Within a phase, tasks are independent and could in principle be executed in parallel by separate Claude Code sessions, but the recommended order is top-to-bottom for cleaner audit checkpoints.

Per the project conventions, every task:

- Is a standalone `.md` file artifact (`ZV-NNN.md`).
- Begins with the standard STOP block referencing `CLAUDE.md`, `CONVENTIONS.md`, `openapi.yaml`, `Zevaro-Audit.md`, and `Zevaro-Architecture.md`. ZV-001 has a modified STOP block because the OpenAPI and audit files don't exist yet.
- Ends with `Compile, Run, Test, Commit, Push to Github`.
- Ends with the standard report template including the Git commit hash.
- Ships its scope completely in a single pass — implementation, tests (100% coverage), and documentation comments — with no follow-up prompts for the same scope.

Audit-refresh tasks at the end of each phase regenerate `Zevaro-Audit.md` from the current state of the codebase so subsequent tasks work from accurate source-of-truth.

---

## Stack & architectural decisions (locked in)

- **Daemon language:** Go 1.22+
- **Desktop framework:** Wails v2
- **Frontend:** React 18 + TypeScript 5 + Tailwind CSS 3, **pnpm** as the package manager
- **Storage:** SQLite (modernc.org/sqlite) via GORM with AutoMigrate in dev; golang-migrate for production
- **Logging:** `log/slog`, structured, single project-wide logger
- **HTTP framework:** `chi` v5 on stdlib `net/http`
- **API surface:** OpenAI-compatible and Anthropic-compatible (drop-in) plus a `/api/v1/...` management surface
- **Testing:** stdlib `testing` + `testify` + `httptest` + in-memory SQLite + `go-vcr` cassettes for provider replay
- **Build:** Makefile; `goreleaser` for cross-platform local releases (no CI)
- **License:** Apache 2.0
- **Repo layout:** Monorepo at `github.com/zevaro/zevaro`

Full detail in `Zevaro-Architecture.md` §2.

---

## Phase 1 — Bootstrap

| Task | Description |
|------|-------------|
| **ZV-001** | Project skeleton (Go module, Wails project, directory tree, Makefile) |
| **ZV-002** | Architecture spec verification (diff actual skeleton against §3.1) |
| **ZV-003** | OpenAPI specification — every endpoint, field, enum, error shape |
| **ZV-004** | Initial audit baseline (`Zevaro-Audit.md`) |

## Phase 2 — Foundation

| Task | Description |
|------|-------------|
| **ZV-005** | HTTP server foundation (chi, middleware, graceful shutdown, `/healthz`) |
| **ZV-006** | Configuration system (precedence, hot-reload, platform-specific paths) |
| **ZV-007** | Storage layer (GORM + SQLite, AutoMigrate, transaction helpers) |
| **ZV-008** | Provider abstraction layer (interface, registry, mock provider) |
| **ZV-009** | Audit refresh |

## Phase 3 — Core API Surface

| Task | Description |
|------|-------------|
| **ZV-010** | OpenAI-compatible API endpoints |
| **ZV-011** | Anthropic-compatible API endpoints |
| **ZV-012** | Streaming support (SSE encoder, chunk normalization, disconnect handling) |
| **ZV-013** | Tool/function call normalization (OpenAI ↔ Anthropic) |
| **ZV-014** | Audit refresh |

## Phase 4 — Routing Engine

| Task | Description |
|------|-------------|
| **ZV-015** | Routing engine core, rule evaluator, tag-based routing |
| **ZV-016** | Fallback chains and retry logic |
| **ZV-017** | Policy bundles (built-in + user-defined) |
| **ZV-018** | Audit refresh |

## Phase 5 — Provider Implementations

| Task | Description |
|------|-------------|
| **ZV-019** | Cloud providers (Anthropic, OpenAI, xAI, Google, DeepSeek, Mistral, Groq, Together, Cohere) |
| **ZV-020** | Local providers + auto-discovery (Ollama, LM Studio, llama.cpp, vLLM) |
| **ZV-021** | Free-tier providers + stability scoring |
| **ZV-022** | Audit refresh |

## Phase 6 — Storage, History, and Spend Tracking

| Task | Description |
|------|-------------|
| **ZV-023** | Prompt history storage |
| **ZV-024** | Spend tracking + pricing table + projection |
| **ZV-025** | Full-text search over prompt history (FTS5) |
| **ZV-026** | Audit refresh |

## Phase 7 — Adjacent Capabilities

| Task | Description |
|------|-------------|
| **ZV-027** | Semantic prompt caching |
| **ZV-028** | Privacy guard rails (PII redaction) |
| **ZV-029** | Per-project configuration |
| **ZV-030** | A/B and multi-model comparison |
| **ZV-031** | Budget guardrails |
| **ZV-032** | MCP gateway support |
| **ZV-033** | Cost-aware auto-routing |
| **ZV-034** | Plugin/extension architecture |
| **ZV-035** | Failover and load balancing |
| **ZV-036** | Offline mode |
| **ZV-037** | Audit log export |
| **ZV-038** | Embedding + image generation routing |
| **ZV-039** | Team sharing (export/import) |
| **ZV-040** | Audit refresh |

## Phase 8 — GUI

| Task | Description |
|------|-------------|
| **ZV-041** | Wails app shell |
| **ZV-042** | Provider configuration screens |
| **ZV-043** | Routing rules editor |
| **ZV-044** | Spend dashboard |
| **ZV-045** | Prompt history search |
| **ZV-046** | Multi-model comparison + budgets screens |
| **ZV-047** | Settings & preferences |
| **ZV-048** | First-launch onboarding |

## Phase 9 — Distribution

| Task | Description |
|------|-------------|
| **ZV-049** | Audit refresh + GUI completion sweep |
| **ZV-050** | macOS installer + Developer ID signing + notarization |
| **ZV-051** | Windows installer + EV cert signing |
| **ZV-052** | Linux packages (`.deb`, `.rpm`, AppImage) |
| **ZV-053** | Auto-update mechanism (signed manifest, atomic replacement, rollback) |

## Phase 10 — Quality and Integration Testing

| Task | Description |
|------|-------------|
| **ZV-054** | Audit refresh |
| **ZV-055** | Cross-provider integration test suite |
| **ZV-056** | End-to-end GUI tests (Playwright) |
| **ZV-057** | Local release script (`make release` + `scripts/release/`) |

## Phase 11 — Documentation and Launch

| Task | Description |
|------|-------------|
| **ZV-058** | Audit refresh |
| **ZV-059** | Documentation site (Astro at zevaro.ai/docs) |
| **ZV-060** | Client integration guides (Claude Code, Cline, Continue, Aider, Cursor, Zed, OpenCode, OpenAI/Anthropic SDKs) |
| **ZV-061** | HN launch post + 90-second demo video script |
| **ZV-062** | `CONTRIBUTING.md` polish + good-first-issue seeding |
| **ZV-063** | Final audit + v1.0.0 release tag |

---

## Standard prompt template

Every prompt is generated against this template. ZV-001 is the single exception (modified STOP block because `openapi.yaml` and `Zevaro-Audit.md` don't exist yet).

```markdown
# ZV-NNN — [Task Title]

STOP: Before writing ANY code, read these files completely:
1. ~/Documents/Github/zevaro/CLAUDE.md
2. ~/Documents/Github/zevaro/CONVENTIONS.md
3. ~/Documents/Github/zevaro/openapi.yaml
4. ~/Documents/Github/zevaro/Zevaro-Audit.md
5. ~/Documents/Github/zevaro/Zevaro-Architecture.md

Do not rely on the descriptions in this prompt alone. If this prompt conflicts
with the source files, the source files win.

If any of these files are missing, STOP, and ask before proceeding.

## Goal

[What this task accomplishes — single declarative statement.]

## Constraints

[Constraints from CLAUDE.md plus task-specific ones.]

## Files to read first

[Specific source files Claude Code must read before producing any output.]

## Acceptance criteria

[Concrete, testable outcomes — checkbox list.]

## Out of scope

[Things this task explicitly does NOT do, with pointers to which task does them.]

Compile, Run, Test, Commit, Push to Github.

## Report template

[Reference to Zevaro-Architecture.md §7 with any task-specific additions.]
```

---

## What's not in this roadmap (intentionally)

- **No timelines.** Per the project conventions, traditional time estimates do not apply. Each task is single-pass at AI-first velocity.
- **No phasing by severity or priority.** v1 contains the entire charter scope. There is no "v1.1 will add X."
- **No priority labels.** The dependency order is the only ordering. Within a phase, tasks are parallelizable.
- **No "may defer if too complex" notes.** Complexity is not a deferral reason.
- **No code in the roadmap.** This document and every prompt tell Claude Code *what* to build and *what to read* before building. Claude Code reads source files and writes the code.

---

## Related documents

- **`Zevaro-Architecture.md`** — Single source of truth for the full system architecture, design language, and the canonical task roadmap with detailed task descriptions in §6. Read this before generating any prompt.
- **`ZEVARO_TASK_TRACKER.md`** — Live status of every task: 🟢 done · 🟡 prompt generated, awaiting execution · 🔵 in progress · ⚪ not started · 🔴 blocked. Updated after every task.
- **`CLAUDE.md`** — Binding behavioral contract Claude Code reads at session start. Defines the absolute rules and the five Critical Execution Rules.
- **`CONVENTIONS.md`** — Engineering conventions (testing, documentation, build tooling, prompt format).
