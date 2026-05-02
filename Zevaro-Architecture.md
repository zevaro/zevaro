# Zevaro — Architecture Reference

**Status:** Authoritative. Every task reads this document before writing code. When a prompt conflicts with this document, this document wins.
**Lineage:** Greenfield project. Patterns adapted from Inspector (logging, constants, completion report template, sub-agent rules) and from the user's standing AI-first conventions (single-pass delivery, source-of-truth discipline, test-and-doc-in-same-pass). Stack and approach are inspired by Ollama (single-binary distribution, local-first daemon) and Tailscale (Wails-equivalent native GUI on top of a Go daemon).

---

## 1. Product Overview

Zevaro is a **desktop-first, open-source AI gateway** for individual developers and small teams. It runs as a single native application on macOS, Windows, and Linux. Behind the GUI, it operates as a long-running local HTTP daemon that exposes both an **OpenAI-compatible** and an **Anthropic-compatible** wire-protocol surface. Any AI client that supports a configurable base URL — Claude Code, Cline, Continue, Aider, Cursor (custom mode), Zed, OpenCode, the OpenAI/Anthropic SDKs — can be pointed at Zevaro and treat it as a drop-in replacement for `api.openai.com` or `api.anthropic.com`.

When a request arrives at Zevaro, the routing engine decides which actual provider serves it — Anthropic, OpenAI, xAI, Google, DeepSeek, Mistral, Groq, Together, Cohere, free-tier services, or local backends like Ollama, LM Studio, llama.cpp, vLLM. Routing decisions are driven by user-configured rules, policy bundles, session tags, project-scoped context, fallback chains, budget guardrails, and (optionally) learned cost-aware preferences. Every request is logged locally with full token accounting and cost calculation, exposed through a polished GUI dashboard with spend analytics, full-text history search, multi-model comparison, and budget alerts.

Zevaro is **local-first**: all keys, prompts, history, routing rules, and policies live on the user's machine in a SQLite database. Nothing phones home. There is no telemetry, no cloud-hosted version, no paid tier, no enterprise upsell. The project is licensed Apache 2.0.

---

## 2. Platform & Technology

| Concern                       | Choice                                                                                       |
|-------------------------------|----------------------------------------------------------------------------------------------|
| Daemon language               | Go 1.22+                                                                                     |
| Desktop framework             | Wails v2 (Go backend + HTML/CSS/JS frontend bundled into a small native binary using the OS webview) |
| Frontend stack inside Wails   | React 18 + TypeScript 5 + Tailwind CSS 3                                                     |
| Target platforms              | macOS (universal binary), Windows (x64 + arm64), Linux (x64 + arm64)                         |
| HTTP framework                | `chi` v5 (routing) on top of stdlib `net/http`                                               |
| Streaming                     | Server-Sent Events (SSE) via stdlib `net/http`                                               |
| Storage                       | SQLite (modernc.org/sqlite, pure-Go driver) via GORM v2                                      |
| Schema management             | GORM `AutoMigrate` during development; versioned migrations (golang-migrate) for production  |
| Configuration                 | Custom config layer (defaults → file → env → GUI overrides) reading `config.yaml` via `gopkg.in/yaml.v3` |
| Logging                       | Standard library `log/slog` with JSON handler in production, text handler in dev             |
| Validation                    | `go-playground/validator/v10`                                                                |
| Testing                       | stdlib `testing` + `stretchr/testify` + `httptest` + in-memory SQLite                        |
| HTTP test recording           | `dnaeon/go-vcr/v3` for provider fixture replay                                               |
| Linting                       | `golangci-lint` with strict ruleset (errcheck, govet, staticcheck, revive, gosec, gocyclo, gocritic, misspell, unparam, unconvert, ineffassign, prealloc) |
| Frontend state                | Zustand for global state; React Query for server state                                       |
| Frontend routing              | React Router v6                                                                              |
| Frontend charts               | Recharts                                                                                     |
| Frontend testing              | Vitest + React Testing Library                                                               |
| E2E testing                   | Playwright (Wails-supported via Chromium)                                                    |
| Build orchestration           | `Makefile` + Wails CLI; cross-compilation via `goreleaser`                                   |
| Code signing                  | macOS: Developer ID + notarytool; Windows: EV cert via SignTool; Linux: GPG-signed packages  |
| Auto-update                   | Custom updater using signed manifest + atomic binary replacement                             |
| CI / CD                       | GitHub Actions                                                                               |
| License                       | Apache 2.0                                                                                   |
| Repo layout                   | Monorepo at `github.com/zevaro/zevaro`                                                       |

Minimum window size: 1024×680. Default window size: 1280×800. Dark mode is the default theme.

---

## 3. Technical Architecture

### 3.1 Project Directory Structure

```
zevaro/
├── CLAUDE.md                           # Project persona + execution rules (read on session start)
├── CONVENTIONS.md                      # AI-first conventions (user preferences verbatim)
├── Zevaro-Architecture.md              # This file — single source of truth
├── Zevaro-Audit.md                     # Generated / refreshed after each task
├── openapi.yaml                        # OpenAPI 3.1 spec — every endpoint, every field, every enum
├── README.md
├── LICENSE                             # Apache 2.0
├── NOTICE                              # Apache 2.0 NOTICE file
├── CONTRIBUTING.md
├── go.mod / go.sum                     # Daemon dependencies
├── wails.json                          # Wails v2 project manifest (required for `wails dev` / `wails build`)
├── Makefile                            # build, test, lint, package, release
├── .editorconfig
├── .gitignore
├── .golangci.yml                       # Strict lint config
├── .goreleaser.yaml                    # Cross-platform release config
├── .github/
│   ├── ISSUE_TEMPLATE/
│   ├── PULL_REQUEST_TEMPLATE.md
│   └── workflows/
│       ├── ci.yaml                     # build + test + lint on push/PR
│       ├── release.yaml                # tag-driven cross-platform signed release
│       └── nightly.yaml                # nightly canary build
│
├── cmd/
│   └── zevaro/
│       └── main.go                     # Entry point: bootstrap, Wails, daemon lifecycle
│
├── internal/                           # Implementation; not importable by external code
│   ├── app/                            # Top-level wiring (Wails app, IPC bridge, lifecycle)
│   ├── constants/                      # Centralized magic values (durations, sizes, port defaults, key names)
│   ├── server/                         # HTTP server: chi mux, middleware, graceful shutdown
│   │   ├── server.go
│   │   ├── middleware/                 # request_id, recovery, logging, cors
│   │   └── healthz.go
│   ├── api/
│   │   ├── openai/                     # OpenAI-compatible endpoints
│   │   ├── anthropic/                  # Anthropic-compatible endpoints
│   │   ├── streaming/                  # SSE encoder, chunk normalization
│   │   └── management/                 # /api/v1/... GUI-facing management endpoints
│   ├── normalize/
│   │   ├── tools/                      # OpenAI ↔ Anthropic tool-call translation
│   │   ├── messages/                   # Cross-format message normalization
│   │   └── errors/                     # Provider-specific → canonical error mapping
│   ├── providers/
│   │   ├── provider.go                 # Provider interface + capability descriptor
│   │   ├── registry.go                 # Provider registry (register, lookup, lifecycle)
│   │   ├── cloud/                      # Anthropic, OpenAI, xAI, Google, DeepSeek, Mistral, Groq, Together, Cohere
│   │   ├── local/                      # Ollama, LM Studio, llama.cpp, vLLM
│   │   └── freetier/                   # Big Pickle and other free-tier providers; stability scoring
│   ├── discovery/                      # Local-model auto-discovery (Ollama/LM Studio/etc. on standard ports)
│   ├── routing/
│   │   ├── engine.go                   # Routing decision interface + dispatcher
│   │   ├── rules/                      # Predicate-based rule evaluator
│   │   ├── tags/                       # Session/tag-based routing
│   │   ├── policies/                   # Policy bundles (built-in + user-defined)
│   │   ├── fallback/                   # Multi-provider fallback chains + retry policies
│   │   ├── balance/                    # Multi-key load balancing per provider
│   │   ├── auto/                       # Cost-aware learned auto-routing
│   │   ├── embeddings/                 # Embedding-specific routing
│   │   └── images/                     # Image-generation routing
│   ├── storage/
│   │   ├── db.go                       # GORM setup, AutoMigrate, connection pooling, transaction helpers
│   │   ├── entities/                   # Base entity types (timestamps, soft deletes)
│   │   └── platform/                   # OS-specific data-directory resolution
│   ├── config/
│   │   ├── config.go                   # Typed config struct, precedence, hot-reload
│   │   ├── schema.go                   # Schema validation
│   │   └── project/                    # Per-project (filesystem-path-scoped) config
│   ├── history/
│   │   ├── store.go                    # Prompt-history persistence
│   │   ├── search/                     # FTS5-backed full-text search
│   │   └── retention.go                # Retention policies
│   ├── spend/
│   │   ├── tracker.go                  # Token-count extraction per provider response shape
│   │   ├── pricing/                    # Maintained pricing table (per-model rates)
│   │   ├── aggregate.go                # Aggregation queries (provider/model/project/window)
│   │   └── projection.go               # Spend projection logic
│   ├── budgets/                        # Budget definitions, enforcement, alerts
│   ├── cache/                          # Semantic prompt caching (embedding-based)
│   ├── privacy/                        # PII/secret redaction layer
│   ├── compare/                        # A/B and multi-model comparison
│   ├── mcp/                            # MCP gateway (server registration, tool brokering, transports)
│   ├── plugins/                        # Plugin/extension architecture (HashiCorp go-plugin)
│   ├── offline/                        # Network-state monitoring + offline routing
│   ├── audit/                          # Signed audit-log export
│   ├── sharing/                        # Policy/config bundle export/import (file-based, no SaaS)
│   ├── updater/                        # Auto-update mechanism (manifest verification, atomic replace)
│   └── telemetry/                      # Intentionally empty — sentinel package documenting "no telemetry"
│
├── pkg/                                # Public API for plugin authors / external consumers
│   ├── zevaro-plugin/                  # Plugin SDK (Go module, separately versioned)
│   └── pricing-data/                   # Public pricing data structures (importable by tooling)
│
├── api/                                # API artifacts
│   └── openapi/                        # openapi.yaml + per-surface fragments + generated docs
│
├── frontend/                           # Wails frontend (React + TypeScript + Tailwind)
│   ├── package.json
│   ├── tsconfig.json
│   ├── tailwind.config.js
│   ├── postcss.config.js
│   ├── eslint.config.js
│   ├── vite.config.ts
│   ├── index.html
│   └── src/
│       ├── main.tsx
│       ├── app.tsx                     # Top-level shell: menu, tray, navigation
│       ├── routes/                     # React Router routes
│       │   ├── dashboard/              # Spend dashboard
│       │   ├── providers/              # Provider config screens
│       │   ├── routing/                # Routing rules editor
│       │   ├── history/                # Prompt history search + detail
│       │   ├── budgets/                # Budget configuration
│       │   ├── compare/                # Multi-model comparison view
│       │   ├── settings/               # Settings & preferences
│       │   └── onboarding/             # First-launch onboarding flow
│       ├── components/                 # Reusable UI components
│       │   ├── charts/                 # Recharts wrappers
│       │   ├── forms/                  # Form primitives
│       │   ├── layout/                 # Shell, sidebar, topbar
│       │   ├── tables/                 # Data table primitives
│       │   └── dialogs/                # Modal primitives
│       ├── state/                      # Zustand stores
│       ├── lib/                        # API client, hooks, utilities
│       ├── assets/
│       │   └── fonts/                  # Bundled JetBrains Mono (per §5)
│       └── styles/
│           └── globals.css
│
├── installers/
│   ├── macos/                          # DMG build scripts, entitlements, notarization
│   ├── windows/                        # MSI/EXE build scripts (WiX), code-signing wrapper
│   └── linux/                          # .deb / .rpm / AppImage scripts
│
├── docs/                               # Source for the docs site (zevaro.ai/docs)
│   ├── audit/                          # Audit-task artifacts (skeleton diffs, audit refreshes)
│   ├── astro.config.mjs                # (or docusaurus equivalent)
│   ├── src/
│   │   ├── content/
│   │   │   ├── getting-started/
│   │   │   ├── architecture/
│   │   │   ├── providers/
│   │   │   ├── routing/
│   │   │   ├── clients/                # Client integration guides (Claude Code, Cline, etc.)
│   │   │   ├── troubleshooting/
│   │   │   └── faq/
│   │   └── components/
│   └── public/
│
├── test/
│   ├── integration/                    # Cross-package integration tests
│   ├── e2e/                            # Playwright E2E tests against the built Wails app
│   ├── fixtures/                       # Recorded provider responses (go-vcr cassettes)
│   └── helpers/                        # Shared test utilities
│
└── scripts/
    ├── lint.sh
    ├── release.sh
    ├── notarize-macos.sh
    └── sign-windows.ps1
```

### 3.2 Data Layer

Zevaro persists everything to a single SQLite database file located in the platform-appropriate data directory:

| OS       | Path                                                                          |
|----------|-------------------------------------------------------------------------------|
| macOS    | `~/Library/Application Support/Zevaro/zevaro.db`                              |
| Linux    | `$XDG_DATA_HOME/zevaro/zevaro.db` (default `~/.local/share/zevaro/zevaro.db`) |
| Windows  | `%LOCALAPPDATA%\Zevaro\zevaro.db`                                             |

Configuration is stored alongside the database as `config.yaml` (human-editable). Logs are written to `zevaro.log` in the same directory, rotated nightly.

**Tables (managed by GORM `AutoMigrate` during development; locked migrations for production):**

| Table                  | Purpose                                                                            |
|------------------------|------------------------------------------------------------------------------------|
| `providers`            | Registered provider instances (cloud + local + freetier)                           |
| `provider_keys`        | API keys per provider (encrypted at rest with OS keychain on macOS/Windows)        |
| `models`               | Known models per provider, with capability flags and pricing                       |
| `routing_rules`        | User-defined routing rules with predicates and target selectors                    |
| `policy_bundles`       | Built-in and user-defined policy bundles                                           |
| `project_configs`      | Per-filesystem-path project policies                                               |
| `sessions`             | Logical session groupings (created on a tag-change or client connection)           |
| `requests`             | Every request through the gateway (input, output, model, tokens, cost, latency)   |
| `request_tags`         | Many-to-many tags on requests                                                      |
| `requests_fts`         | FTS5 virtual table for full-text search over prompt content                        |
| `cache_entries`        | Semantic-cache entries (embedding + response)                                      |
| `budgets`              | Budget definitions and current utilization                                         |
| `mcp_servers`          | Registered MCP servers and their transport configs                                 |
| `mcp_tools`            | Cached tool descriptors per MCP server                                             |
| `redaction_patterns`   | User-defined PII patterns (in addition to bundled defaults)                        |
| `audit_log_exports`    | Records of signed audit-log export operations                                      |
| `plugin_registrations` | Installed plugins, their capabilities, and their versioned manifests               |
| `key_health`           | Per-key health stats (success rate, last failure, rate-limit status)               |

Provider keys are **never** logged, never returned through the management API in plaintext, and never included in audit-log exports. They are encrypted at rest using the OS keychain (Keychain on macOS, DPAPI on Windows, `secret-service` on Linux when available; falls back to a derived key file with restrictive permissions when the keyring is unavailable).

### 3.3 Provider Abstraction

The `Provider` interface (`internal/providers/provider.go`) is the single contract every concrete provider implements:

| Method                         | Purpose                                                                  |
|--------------------------------|--------------------------------------------------------------------------|
| `ID() string`                  | Stable provider identifier                                               |
| `Capabilities() Capabilities`  | Capability descriptor (chat, embeddings, images, streaming, tools, etc.) |
| `Models(ctx) []Model`          | List models exposed by the provider                                      |
| `Chat(ctx, ChatRequest)`       | Non-streaming chat completion                                            |
| `ChatStream(ctx, ChatRequest)` | Streaming chat completion (returns a `<-chan ChatChunk`)                 |
| `Embed(ctx, EmbedRequest)`     | Embeddings                                                               |
| `Image(ctx, ImageRequest)`     | Image generation                                                         |
| `Health(ctx) HealthStatus`     | Liveness + reachability check                                            |
| `Close() error`                | Cleanup (HTTP client connections, etc.)                                  |

Requests and responses to the `Provider` interface use **canonical types** in `internal/providers/types.go` — independent of any specific provider's wire format. Each concrete provider translates between the canonical types and the provider's actual API. Streaming chunks are normalized into a canonical `ChatChunk` type with consistent fields regardless of source.

A `Capabilities` descriptor lets the routing engine ask "can this provider handle a request that requires X?" without invoking the provider — used for fast filtering before dispatch.

### 3.4 Routing Engine

The routing engine (`internal/routing/engine.go`) takes a normalized request plus context (session tags, project path, user identity, request metadata) and produces a **routing decision**: an ordered list of providers to try, with per-attempt timeout and retry policy.

Strategies, evaluated in this order:

1. **Manual override.** If the session or request has a pinned provider/model, it wins.
2. **Per-project policy.** If the request originates from a configured project path, that project's policy is applied next.
3. **Tag-based routing.** Session tags (`scaffolding`, `architecture`, `qa`, etc.) match policy-bundle entries.
4. **Rule evaluation.** User-defined rules with arbitrary predicates (regex on content, header values, model name, etc.) match in declared order.
5. **Auto-routing (optional).** When enabled, the cost-aware recommender suggests the cheapest provider/model meeting the user's quality threshold based on personal history.
6. **Default policy.** The user's default policy bundle.
7. **Capability filter.** The candidate list is filtered to providers whose `Capabilities` cover the request shape.
8. **Fallback chain.** The dispatcher retries through the candidate list on failure, applying the configured retry policy (exponential backoff, retryable error classes, max attempts).

Every routing decision is logged to the request record with full provenance: which strategy fired, which providers were tried, which succeeded, latency per attempt, the final cost.

**Built-in policy bundles:** Cheap Coding, Privacy First, Quality Over Cost, Balanced, Local Only, Free Tier Only. Each is defined in `internal/routing/policies/builtin/` as a literal YAML embedded via `//go:embed`.

### 3.5 API Surfaces

Zevaro exposes three logical HTTP surfaces, all served from the same daemon on the same port (default `:39237`):

**OpenAI-compatible** (`/v1/...`) — drop-in for clients pointing at `api.openai.com`:

- `POST /v1/chat/completions`
- `POST /v1/completions`
- `POST /v1/embeddings`
- `GET  /v1/models`
- `POST /v1/images/generations`
- `POST /v1/audio/transcriptions`

Wire format matches OpenAI's published shape exactly per `openapi.yaml`. Streaming uses SSE with the same chunk envelope.

**Anthropic-compatible** (`/v1/messages`) — drop-in for clients pointing at `api.anthropic.com`:

- `POST /v1/messages`
- `POST /v1/messages` with `stream: true`
- `POST /v1/messages/count_tokens`

The `anthropic-version` header is parsed; supported values match what `openapi.yaml` declares.

**Management** (`/api/v1/...`) — used by the GUI and by power-user CLI tooling. Authenticated with a local-only bearer token generated on first launch (stored in the keychain; auto-injected by the GUI):

- `GET    /api/v1/providers`
- `POST   /api/v1/providers`
- `PATCH  /api/v1/providers/{id}`
- `DELETE /api/v1/providers/{id}`
- `POST   /api/v1/providers/{id}/test`
- `GET    /api/v1/models`
- `GET    /api/v1/routes` / `POST` / `PATCH` / `DELETE`
- `GET    /api/v1/policies` / `POST` / `PATCH` / `DELETE`
- `GET    /api/v1/projects` / `POST` / `PATCH` / `DELETE`
- `GET    /api/v1/sessions` / `POST` / `PATCH` / `DELETE`
- `GET    /api/v1/history` (with FTS query params)
- `GET    /api/v1/history/{id}`
- `GET    /api/v1/spend` (with aggregation params)
- `GET    /api/v1/spend/projection`
- `GET    /api/v1/budgets` / `POST` / `PATCH` / `DELETE`
- `GET    /api/v1/cache/stats` / `DELETE /api/v1/cache`
- `POST   /api/v1/compare`
- `GET    /api/v1/mcp/servers` / `POST` / `DELETE`
- `GET    /api/v1/mcp/tools`
- `POST   /api/v1/audit/export`
- `POST   /api/v1/sharing/export`
- `POST   /api/v1/sharing/import`
- `GET    /api/v1/discovery/local-models`
- `GET    /api/v1/health/keys`
- `GET    /api/v1/plugins` / `POST /api/v1/plugins/install` / `DELETE /api/v1/plugins/{id}`
- `GET    /api/v1/updater/status` / `POST /api/v1/updater/check`
- `GET    /healthz`

Every endpoint, every field, every enum, every error shape is defined in `openapi.yaml`. The OpenAPI file is the canonical contract — Go handlers and TypeScript clients are validated against it.

### 3.6 Streaming

SSE is implemented in `internal/api/streaming/`. The encoder writes properly framed `data:` lines with a single trailing `\n\n`, flushes after every chunk (`http.Flusher`), and writes a `data: [DONE]\n\n` terminator on stream completion. Client disconnects are detected via `r.Context().Done()` and propagate cancellation upstream so provider HTTP requests are aborted.

When a request arriving on the OpenAI surface is routed to an Anthropic provider (or vice versa), the streaming layer translates chunk-by-chunk between formats. The translation tables live in `internal/normalize/messages/` and `internal/normalize/tools/`, and round-trip equivalence is enforced by the test suite.

### 3.7 Local Model Auto-Discovery

`internal/discovery/` runs a periodic scan (every 30 seconds, configurable) for local model servers on their well-known ports:

| Server     | Port (default) | Detection method                            |
|------------|----------------|---------------------------------------------|
| Ollama     | 11434          | `GET /api/tags` returns valid JSON          |
| LM Studio  | 1234           | `GET /v1/models` returns valid JSON         |
| llama.cpp  | 8080           | `GET /props` returns valid JSON             |
| vLLM       | 8000           | `GET /v1/models` returns valid JSON         |

When a server is detected, it is auto-registered as a provider with its discovered models. The user is shown a non-modal banner ("Detected Ollama with 4 models — Add to Zevaro?") and accepting wires the provider into the registry. If a previously-discovered server stops responding, its provider is marked as unhealthy but **not** removed (preserving rules and history).

### 3.8 Spend Tracking

Every request's response includes provider-reported token counts when available; otherwise Zevaro estimates with a tokenizer matching the provider family. Cost is computed at request-completion time using the pricing table in `internal/spend/pricing/` (versioned, embedded via `//go:embed`, refreshed via auto-update). The pricing table is a YAML keyed by `provider:model` with input/output rates.

Aggregations (`internal/spend/aggregate.go`):

- By provider × time-window (day/week/month/all-time)
- By model × time-window
- By project × time-window
- By session × time-window
- By tag × time-window
- Top-N most expensive requests
- Top-N most cached prompts (with cache savings)

Projection (`internal/spend/projection.go`) uses a 14-day rolling window with linear regression on daily totals to project monthly spend, plus a 90% confidence interval. The projection updates after every request.

### 3.9 Privacy Layer

`internal/privacy/` provides pre-send redaction. The default pattern library covers:

- API keys (provider-specific patterns: `sk-`, `sk-ant-`, `xai-`, `gsk_`, etc.)
- AWS access key IDs and secret access keys
- Generic high-entropy 32+-character strings (low-precision fallback)
- Email addresses
- Phone numbers (E.164 + common US formats)
- Credit card numbers (Luhn-validated)
- Filesystem paths matching configured "confidential" prefixes
- IP addresses (configurable on/off)
- UUIDs (configurable on/off)
- User-defined regex patterns

Redaction runs **before** dispatch to a cloud provider. Redaction is configurable per-route — a rule may declare "skip redaction for this route" if the user explicitly trusts the destination. Local providers always skip redaction (the data never leaves the machine). Every redaction is logged with the matched-pattern name and offset (not the redacted value itself), and the full unredacted prompt is retained in local history.

### 3.10 Caching

`internal/cache/` provides semantic prompt caching. On every request:

1. Compute an embedding of the prompt content (using a small bundled embedding model, or a configured embedding provider).
2. Look up the cache by embedding similarity (cosine, threshold default 0.95, configurable).
3. On hit: return the cached response. Mark the request `cache_hit=true`. Record the saved cost.
4. On miss: dispatch to the routing engine; on success, store the (embedding, response) pair.

Cache invalidation:

- TTL (default 7 days, configurable per-policy)
- Manual purge via `DELETE /api/v1/cache`
- Per-tag exclusion (a tag like `realtime` flagged in policy disables cache lookup)

Cache is opt-in per route. Default policy bundles enable it for "Cheap Coding" and "Balanced", disable it for "Quality Over Cost" and "Privacy First".

### 3.11 MCP Gateway

`internal/mcp/` aggregates multiple MCP servers behind Zevaro's single endpoint. Supports `stdio` and `http` transports per the MCP spec. On registration, Zevaro:

1. Spawns or connects to the MCP server.
2. Issues `initialize` and `tools/list`.
3. Caches the tool descriptors with namespace prefix (`<server-id>__<tool-name>`).
4. Forwards tool calls from clients to the appropriate MCP server based on namespace.

Health monitoring re-issues `tools/list` every 5 minutes; servers that fail are marked unhealthy and excluded from tool advertisements until they recover.

### 3.12 Plugin Architecture

`internal/plugins/` uses HashiCorp's `go-plugin` (subprocess RPC) for cross-platform plugin isolation. A plugin is a separate binary that implements the `pkg/zevaro-plugin` SDK interfaces:

- `ProviderPlugin` — adds a new provider
- `RouterPlugin` — adds a new routing strategy
- `RedactorPlugin` — adds new redaction patterns
- `DashboardPlugin` — adds a new dashboard panel
- `ClientGuidePlugin` — adds a new client integration guide

Plugins declare a manifest (name, version, capabilities, permissions). Zevaro enforces capability boundaries: a `RedactorPlugin` cannot make outbound HTTP calls; a `ProviderPlugin` cannot read history. Plugin processes are sandboxed and killed if they violate their capability declarations.

The plugin SDK (`pkg/zevaro-plugin/`) is a separate Go module published under the same repo, semver-versioned independently from the main daemon. Breaking SDK changes bump the major version and require plugin recompilation; minor versions guarantee backward compatibility.

### 3.13 GUI Architecture

The Wails frontend runs in the OS-native webview (WebKit on macOS, WebView2 on Windows, WebKitGTK on Linux). Communication with the Go backend goes through Wails' **bindings layer**: Go methods on the App struct are exposed to the frontend as TypeScript-typed async functions.

**State management:** Zustand stores hold UI state and cached server state. React Query handles server-fetched data with caching and background refresh. The GUI never reaches into SQLite directly — all data flows through `/api/v1/...` endpoints, the same surface a CLI or third-party tool would use.

**Routing:** React Router v6 with these top-level routes:

- `/` → Dashboard (spend overview, recent activity, key health)
- `/providers` → Provider configuration
- `/routing` → Routing rules editor
- `/history` → Prompt history search and detail
- `/budgets` → Budget configuration
- `/compare` → Multi-model comparison
- `/settings` → Settings & preferences
- `/onboarding` → First-launch flow (guards `/` until completed)

**Theme:** Tailwind CSS with a dark default. CSS variables for the palette (defined in `frontend/src/styles/globals.css`) — see §5 for the exact values.

**System tray:** A persistent tray icon (via Wails system-tray API) with: Open Dashboard, Pause Routing, Resume Routing, Daemon Status, Quit Zevaro. Closing the main window minimizes to tray; "Quit" is the only path that fully exits the daemon.

### 3.14 Configuration Precedence

Configuration is resolved in this order (later overrides earlier):

1. **Compiled defaults** — sensible values for every field
2. **Config file** — `config.yaml` in the platform data directory
3. **Environment variables** — `ZEVARO_*` prefix
4. **GUI overrides** — runtime changes from the Settings screen, persisted back to the config file

Hot-reload watches `config.yaml`; changes apply on the next request without a daemon restart, except for the listen address (which requires restart and warns the user accordingly).

### 3.15 Cross-Platform Distribution

| Platform | Format                          | Code signing                                    |
|----------|---------------------------------|-------------------------------------------------|
| macOS    | `.dmg` containing `Zevaro.app`  | Developer ID + notarytool stapled               |
| Windows  | `.msi` (WiX) and portable `.exe`| EV cert via SignTool                            |
| Linux    | `.deb`, `.rpm`, `.AppImage`     | GPG-signed packages; AppImage embeds checksum   |

All artifacts are reproducible via `goreleaser` driven from `.goreleaser.yaml`. Release tags trigger the GitHub Actions workflow which builds, signs, notarizes (macOS), and publishes to a GitHub Release plus the auto-update manifest.

### 3.16 Auto-Update

`internal/updater/` checks an HTTPS-served signed manifest hourly:

- Manifest schema: `version`, `channel` (stable/beta), `platforms[]` (with URL + SHA-256 + signature per platform).
- Signature verification: minisign or Ed25519, public key embedded at build time.
- Download to a temp directory, verify hash and signature, then atomic-replace the running binary on next start.
- On failure (signature mismatch, hash mismatch, write error): rollback to previous binary; log failure; surface to GUI.

User is in control: auto-check is opt-in, channel is selectable (stable or beta), and updates can be deferred or skipped per-version.

### 3.17 Offline Mode

`internal/offline/` polls connectivity to a configurable canary endpoint (default: a small endpoint Zevaro itself hosts at zevaro.ai/connectivity). When connectivity drops:

- Routing is automatically restricted to local providers.
- Cloud-only requests return a clear `503 Network Unavailable — using local fallback` if no local fallback is configured, OR are routed to local fallback per the active policy.
- An optional **queued replay** mode buffers cloud-only requests in a local queue (max queue size configurable, default 100); on connectivity restoration, queued requests are replayed in order with the user's permission via a GUI banner.

### 3.18 Logging

All Go code uses `log/slog` with structured key-value output. A single project-wide logger is initialized in `cmd/zevaro/main.go` and propagated via `context.Context`:

- Production handler: JSON to `zevaro.log` (rotated nightly, retained 14 days), plus stderr for fatal-level only.
- Development handler: text to stderr.
- Log level configurable via `config.yaml` and the GUI Settings screen (`debug` / `info` / `warn` / `error`).
- API keys, full prompt content, and full response content are **never** logged — only metadata (provider, model, latency, token counts, request ID).

The frontend uses a `lib/logger.ts` wrapper that mirrors the slog API and forwards to a backend-bound `log` binding so logs from both surfaces land in the same file with a unified `surface` field (`backend` / `frontend`).

---

## 4. Service Layer Conventions

Every Go package in `internal/`:

- Exposes its public API through interfaces, not concrete structs.
- Constructs concrete implementations via a `New<Name>(deps...) *<Name>` constructor returning the interface.
- Receives dependencies (other interfaces, the logger, the database) via the constructor — no globals beyond the embedded pricing table and the embedded policy bundles.
- Returns typed errors using sentinel errors (`var Err...`) and wrapped errors (`fmt.Errorf("...%w", err)`); error messages never include user content.
- Has matching unit tests at the package level and integration tests in `test/integration/` where a multi-package flow is exercised.
- Has GoDoc comments on every exported package, type, function, method, variable, and constant — except generated code, gRPC stubs, and database entity structs.

`Provider` adapters in `internal/providers/cloud/` and `internal/providers/local/` additionally have integration tests gated behind a `-tags=live` build constraint that exercise real provider APIs when an environment variable like `ZEVARO_LIVE_TEST_OPENAI_KEY` is set; CI runs these only on a nightly schedule, not on every PR.

---

## 5. Design Language

Zevaro's GUI uses a dark-default Tailwind palette tuned for long working sessions and dashboard-heavy screens.

**CSS variables** (in `frontend/src/styles/globals.css`):

| Variable                  | Dark value              | Light value             | Usage                                                |
|---------------------------|-------------------------|-------------------------|------------------------------------------------------|
| `--bg-canvas`             | `#0b0d10`               | `#ffffff`               | App background                                       |
| `--bg-surface`            | `#13171c`               | `#f6f7f9`               | Cards, sidebar, dialogs                              |
| `--bg-surface-elevated`   | `#1a1f26`               | `#ffffff`               | Modals, popovers, hovered cards                      |
| `--border-subtle`         | `#23292f`               | `#e6e8eb`               | Default borders                                      |
| `--border-strong`         | `#323941`               | `#cdd1d6`               | Inputs, primary borders                              |
| `--text-primary`          | `#e6edf3`               | `#0f1115`               | Body text                                            |
| `--text-secondary`        | `#a3acb8`               | `#5b6470`               | Captions, subtitles                                  |
| `--text-muted`            | `#6b737f`               | `#8a929c`               | Placeholders, disabled                               |
| `--accent-primary`        | `#6c5ce7`               | `#5b48d6`               | Primary buttons, focus ring                          |
| `--accent-success`        | `#2ecc71`               | `#1e9e54`               | Success state, healthy keys                          |
| `--accent-warning`        | `#f39c12`               | `#c47800`               | Warning banners, soft budget hits                    |
| `--accent-danger`         | `#e74c3c`               | `#c0392b`               | Errors, hard budget hits, unhealthy keys             |
| `--chart-categorical-1`   | `#6c5ce7`               | `#5b48d6`               | Recharts series color 1                              |
| `--chart-categorical-2`   | `#00cec9`               | `#00a39e`               | Series 2                                             |
| `--chart-categorical-3`   | `#fdcb6e`               | `#d8a23a`               | Series 3                                             |
| `--chart-categorical-4`   | `#ff7675`               | `#e25b5a`               | Series 4                                             |
| `--chart-categorical-5`   | `#74b9ff`               | `#3a8fde`               | Series 5                                             |

**Typography:**

- UI: `system-ui` (San Francisco / Segoe UI / Inter fallback).
- Monospace: `JetBrains Mono`, bundled in `frontend/src/assets/fonts/`, used in code blocks, keys, model names, FTS query syntax.
- Base size 14 px; scale uses Tailwind's default ramp.

**Layout primitives:**

| Concern               | Value                          |
|-----------------------|--------------------------------|
| Min window            | 1024 × 680                     |
| Default window        | 1280 × 800                     |
| Sidebar width         | 240 px (collapsed: 60 px)      |
| Topbar height         | 56 px                          |
| Card corner radius    | 8 px                           |
| Button corner radius  | 6 px                           |
| Default page padding  | 24 px                          |
| Form field height     | 36 px                          |

---

## 6. Task Roadmap

All tasks are single-pass. No deferrals. Tests ship with every task. Documentation ships with every task. Every task ends with `Compile, Run, Test, Commit, Push to Github`. Every task produces a completion report per §7. Audit refresh tasks at the end of each phase regenerate `Zevaro-Audit.md` from the current state of the codebase.

### Phase 1 — Bootstrap (ZV-001 through ZV-004)

| Task | Description |
|------|-------------|
| **ZV-001** | **Project skeleton.** Initialize the Go module (`github.com/zevaro/zevaro`), the Wails v2 project, the directory structure per §3.1, the `.gitignore`, the `.editorconfig`, the `.golangci.yml` (strict ruleset per §2), the `Makefile` with `build / test / lint / package / release / dev` targets, the placeholder `README.md`, the `CONTRIBUTING.md`, the `.github/ISSUE_TEMPLATE/`, the `.github/PULL_REQUEST_TEMPLATE.md`, and the `.github/workflows/ci.yaml` (build + test + lint on push/PR). Confirm `LICENSE`, `NOTICE`, `CONVENTIONS.md`, and `CLAUDE.md` are already present at root from prior steps; do not modify them. Wails app launches and renders an empty "Hello Zevaro" screen. `make build`, `make test`, and `make lint` all succeed against the empty project. |
| **ZV-002** | **Architecture specification finalization.** Confirm `Zevaro-Architecture.md` is current and matches the actual skeleton produced by ZV-001. Update §3.1 if directory layout drifted. This is a verification task — produce a diff report rather than a rewrite. |
| **ZV-003** | **OpenAPI specification.** Produce `openapi.yaml` defining every endpoint surface from §3.5: OpenAI-compatible, Anthropic-compatible, and management. Every field, every enum, every request/response schema, every error shape. Validate the file with `redocly lint` (or equivalent) — zero warnings, zero errors. Generate static HTML docs into `api/openapi/docs/` via `redocly build-docs`. |
| **ZV-004** | **Initial audit baseline.** Produce `Zevaro-Audit.md` at the repo root using the Claude Audit Template applied to the current state. Establishes the baseline format. Subsequent audit refreshes overwrite this file. |

### Phase 2 — Foundation (ZV-005 through ZV-009)

| Task | Description |
|------|-------------|
| **ZV-005** | **HTTP server foundation.** Implement `internal/server/`: graceful startup and shutdown, `chi` v5 routing, structured request logging via `slog`, `request_id` middleware, panic recovery, CORS handling, content-type negotiation, canonical error response shape (matching `openapi.yaml`), `/healthz`. Listen address from config (default `:39237`). Tests cover middleware behavior, graceful shutdown under in-flight requests, error response shape conformance. |
| **ZV-006** | **Configuration system.** Implement `internal/config/`: precedence order from §3.14, platform-specific path resolution from §3.2, schema validation with `validator/v10`, `config.yaml` read/write with formatting preservation where possible, hot-reload via `fsnotify`, typed config struct exposed to the rest of the app. Tests cover precedence, validation failures, platform-specific paths, hot-reload semantics, and the listen-address-restart-required warning. |
| **ZV-007** | **Storage layer.** Implement `internal/storage/`: GORM v2 with `modernc.org/sqlite`, AutoMigrate during dev, connection pool tuning, transaction helpers, base entity types (timestamps, soft deletes), platform-specific database file path. No domain entities yet — those land with their respective features. Tests cover migration, transaction rollback, concurrent access, and pragma settings (WAL, foreign keys, busy timeout). |
| **ZV-008** | **Provider abstraction layer.** Implement `internal/providers/`: the `Provider` interface from §3.3, the canonical request/response types in `internal/providers/types.go`, the provider registry (registration, lookup, lifecycle), the `Capabilities` descriptor model. No concrete providers yet. A `mockProvider` is implemented in `internal/providers/mock/` for use by upstream tests. Tests cover registry behavior, interface contract validation, and mock provider flows. |
| **ZV-009** | **Audit refresh.** Re-run the audit and update `Zevaro-Audit.md` to reflect the foundation layer. |

### Phase 3 — Core API Surface (ZV-010 through ZV-014)

| Task | Description |
|------|-------------|
| **ZV-010** | **OpenAI-compatible API endpoints.** Implement `internal/api/openai/`: `/v1/chat/completions`, `/v1/completions`, `/v1/embeddings`, `/v1/models`, `/v1/images/generations`, `/v1/audio/transcriptions`. Wire-format conformance to `openapi.yaml`. Routes through the provider abstraction; with no concrete providers yet, the `mockProvider` resolves end-to-end. Tests cover request validation, error mapping, and round-trip wire-format compatibility against captured fixtures. |
| **ZV-011** | **Anthropic-compatible API endpoints.** Implement `internal/api/anthropic/`: `/v1/messages` (incl. streaming), `/v1/messages/count_tokens`. Wire-format conformance to `openapi.yaml`. `anthropic-version` header parsing. Tests cover request validation, error mapping, and round-trip wire-format compatibility against captured fixtures. |
| **ZV-012** | **Streaming support.** Implement `internal/api/streaming/`: SSE encoder per §3.6, chunk shape normalization, backpressure handling, client disconnect detection via `r.Context().Done()`, stream termination semantics, `data: [DONE]` terminator on the OpenAI surface. Tests cover full streaming flows, mid-stream client disconnects, error injection, and translation between OpenAI and Anthropic chunk shapes. |
| **ZV-013** | **Tool/function call normalization.** Implement `internal/normalize/tools/`: bidirectional translation between OpenAI tool-call format and Anthropic tool-use format. A request arriving in OpenAI format that gets routed to an Anthropic provider must translate, and vice versa. Tests cover every documented tool-call edge case in both formats and confirm round-trip equivalence with golden fixtures. |
| **ZV-014** | **Audit refresh.** Re-run the audit and update `Zevaro-Audit.md`. |

### Phase 4 — Routing Engine (ZV-015 through ZV-018)

| Task | Description |
|------|-------------|
| **ZV-015** | **Routing engine core.** Implement `internal/routing/engine.go`, `internal/routing/rules/`, and `internal/routing/tags/`: the routing decision interface, the predicate-based rule evaluator (predicates over content, headers, model, request shape), manual override channel (per-session pinned model), tag-based routing, the routing decision log written to the request record. Tests cover every routing strategy independently and in combination with golden decision traces. |
| **ZV-016** | **Fallback chains and retry logic.** Implement `internal/routing/fallback/`: ordered chains, retry policies (exponential backoff, max attempts, retryable error classes per provider), retry budget per request, fallback decision logging. Tests cover provider failure modes (rate limit, network error, 5xx, timeout), retry exhaustion, and chain composition. |
| **ZV-017** | **Policy bundles.** Implement `internal/routing/policies/`: bundle definition format, embedded built-in bundles (Cheap Coding, Privacy First, Quality Over Cost, Balanced, Local Only, Free Tier Only) under `internal/routing/policies/builtin/` via `//go:embed`, bundle loader, user-customizable bundle storage in SQLite. Tests cover bundle composition, override behavior, conflict resolution, and that every built-in bundle parses cleanly. |
| **ZV-018** | **Audit refresh.** Re-run the audit and update `Zevaro-Audit.md`. |

### Phase 5 — Provider Implementations (ZV-019 through ZV-022)

| Task | Description |
|------|-------------|
| **ZV-019** | **Cloud provider implementations.** Implement `internal/providers/cloud/`: Anthropic, OpenAI, xAI (Grok), Google (Gemini), DeepSeek, Mistral, Groq, Together AI, and Cohere. Each adapter handles authentication, request translation, response normalization, streaming, tool-use translation, rate-limit handling, and error mapping per §3.3. Tests cover each provider via recorded `go-vcr` fixtures (no live API calls in CI) plus an opt-in live-test suite gated by `-tags=live` and per-provider env vars. |
| **ZV-020** | **Local provider implementations + auto-discovery.** Implement adapters in `internal/providers/local/` for Ollama, LM Studio, llama.cpp, and vLLM. Implement auto-discovery in `internal/discovery/` per §3.7 (periodic scan of well-known ports, auto-registration with user confirmation, health tracking). Tests cover each adapter and the discovery loop using local fakes. |
| **ZV-021** | **Free-tier provider support.** Implement adapters in `internal/providers/freetier/` for free-tier services (Big Pickle, free-quota DeepSeek, free-quota Gemini) with explicit `unstable: true` flags and a stability scoring mechanism that automatically deprioritizes providers that fail repeatedly within a sliding window. Tests cover stability scoring, sliding-window decay, and degradation behavior. |
| **ZV-022** | **Audit refresh.** Re-run the audit and update `Zevaro-Audit.md`. |

### Phase 6 — Storage, History, and Spend Tracking (ZV-023 through ZV-026)

| Task | Description |
|------|-------------|
| **ZV-023** | **Prompt history storage.** Implement `internal/history/`: schema (request, response, provider, model, latency, token counts, cost, tags, project context), GORM models, retention policies per §3.8, history query API. Every request through the gateway is logged. Tests cover schema migration, query filters, retention enforcement, and high-volume insert performance. |
| **ZV-024** | **Spend tracking.** Implement `internal/spend/`: token-count extraction per provider response shape, cost calculation using the embedded pricing table at `internal/spend/pricing/` per §3.8, aggregation queries, projection logic. Tests cover pricing-table correctness, every aggregation shape, and the projection math (golden values for known inputs). |
| **ZV-025** | **Full-text search over prompt history.** Implement FTS in `internal/history/search/` using SQLite's FTS5 extension. Supports query syntax (`AND`, `OR`, `NOT`, phrase, prefix), filters by model/date/project/cost/outcome/tag, snippet highlighting, ranking. Tests cover correctness, filter combinations, and large-volume performance (10k+ records). |
| **ZV-026** | **Audit refresh.** Re-run the audit and update `Zevaro-Audit.md`. |

### Phase 7 — Adjacent Capabilities (ZV-027 through ZV-040)

| Task | Description |
|------|-------------|
| **ZV-027** | **Semantic prompt caching.** Implement `internal/cache/` per §3.10: embedding generation (configurable embedding provider), similarity-threshold lookup, invalidation policies, hit/miss tracking, per-policy enable/disable. Tests cover cache correctness, threshold tuning, TTL expiry, manual purge, and tag-based exclusion. |
| **ZV-028** | **Privacy guard rails (PII redaction).** Implement `internal/privacy/` per §3.9: bundled pattern library covering every category in §3.9, pre-send redaction with route-level opt-out, redaction logging (pattern name + offset only), user-defined regex patterns. Tests cover every default pattern (positive + negative cases), custom-pattern injection, route-opt-out semantics, and that local providers always skip redaction. |
| **ZV-029** | **Per-project configuration.** Implement `internal/config/project/`: filesystem-path-to-policy mapping, git-repo detection, project-context injection into routing decisions, nested-project resolution. Tests cover path matching, nested-project handling, policy inheritance, and the routing-decision integration. |
| **ZV-030** | **A/B and multi-model comparison.** Implement `internal/compare/`: parallel dispatch to multiple providers for the same request, response collection with timeouts, side-by-side storage in history, winner-tagging mechanism. Tests cover parallel dispatch, partial-failure handling, comparison query API, and timeout enforcement. |
| **ZV-031** | **Budget guardrails.** Implement `internal/budgets/`: per-provider, per-project, per-day, per-week, per-month definitions; soft-warning and hard-cutoff modes; projection alerts. Tests cover budget enforcement, alert generation, projection accuracy, and the hard-cutoff failure mode (request rejected with explicit error). |
| **ZV-032** | **MCP gateway support.** Implement `internal/mcp/` per §3.11: MCP server registration, tool namespace brokering, `stdio` and `http` transports, access control per server, periodic `tools/list` health refresh. Tests cover MCP protocol compliance, namespace collision handling, transport reliability, and unhealthy-server exclusion. |
| **ZV-033** | **Cost-aware auto-routing.** Implement `internal/routing/auto/`: a learned recommender that uses the user's own prompt/response history (not generic benchmarks) to pick the cheapest model meeting a learned quality threshold. Tests cover recommendation correctness against synthetic history, degradation when history is sparse, and the integration into the routing engine as strategy 5. |
| **ZV-034** | **Plugin/extension architecture.** Implement `internal/plugins/` per §3.12: HashiCorp `go-plugin` subprocess RPC, plugin manifest, capability declarations, capability enforcement, and the plugin SDK in `pkg/zevaro-plugin/` published as a separately-versioned Go module. Tests cover plugin lifecycle (discover, load, unload), capability boundary enforcement, and SDK contract stability. |
| **ZV-035** | **Failover and load balancing.** Implement `internal/routing/balance/`: round-robin and weighted distribution across multiple keys for the same provider, automatic key health tracking in `key_health` table, integration with existing fallback chains. Tests cover distribution correctness, health tracking under failures, and failover behavior. |
| **ZV-036** | **Offline mode.** Implement `internal/offline/` per §3.17: connectivity monitoring against canary endpoint, automatic local-provider routing on disconnect, optional queued replay with user permission, GUI banner integration. Tests cover detection accuracy, queue replay correctness, queue-size enforcement, and graceful degradation. |
| **ZV-037** | **Audit log export.** Implement `internal/audit/`: exportable log format (JSON Lines), Ed25519 signing with a project-level keypair, verification tooling (CLI command), tamper-detection tests. Tests cover export correctness, signature verification, and tamper detection. |
| **ZV-038** | **Embedding routing + Image generation routing.** Implement `internal/routing/embeddings/` and `internal/routing/images/`: provider-agnostic embedding/image requests, dimensional compatibility checks (embeddings), format normalization (image: URL ↔ base64), routing rules tuned for each workload. Tests cover dimensional handling, format normalization, and provider-compatibility filtering. |
| **ZV-039** | **Team sharing (export/import).** Implement `internal/sharing/`: portable policy bundle file format, export/import via file or git URL, sanity validation on import, no central server. Tests cover round-trip export/import and rejection of malformed bundles. |
| **ZV-040** | **Audit refresh.** Re-run the audit and update `Zevaro-Audit.md`. |

### Phase 8 — GUI (ZV-041 through ZV-048)

| Task | Description |
|------|-------------|
| **ZV-041** | **Wails app shell.** Implement `frontend/src/app.tsx` and supporting infrastructure: native window, application menu (File, Edit, View, Help), system tray icon with quick controls per §3.13, keyboard shortcuts, React + TypeScript + Tailwind skeleton, Zustand store scaffolding, React Query setup, React Router routes per §3.13, the Wails IPC bridge with TypeScript types generated from Go bindings. Vitest setup. Tests cover IPC contracts (mocked), router smoke, menu and shortcut behavior. |
| **ZV-042** | **GUI: Provider configuration.** Implement `frontend/src/routes/providers/`: list of supported providers, key input with masked display, connection test, capability summary per provider, key health status, multi-key management for load balancing. Tests cover form validation, key storage flow, connection-test mocking, and table rendering. |
| **ZV-043** | **GUI: Routing rules editor.** Implement `frontend/src/routes/routing/`: visual rule builder (predicates and targets), rule ordering (drag-to-reorder), policy bundle picker, manual override controls, rule testing against sample prompts. Tests cover rule serialization round-trip, the testing harness, and drag-reorder semantics. |
| **ZV-044** | **GUI: Spend dashboard.** Implement `frontend/src/routes/dashboard/`: spend over time (line), spend by provider (bar/pie), spend by project (bar), top expensive prompts (table), projected spend (gauge with confidence interval), key health grid, recent activity. All charts via Recharts using the categorical palette from §5. Tests cover data wiring and chart rendering against fixture data. |
| **ZV-045** | **GUI: Prompt history search.** Implement `frontend/src/routes/history/`: search input with FTS query syntax, filter sidebar (model, date range, project, cost range, outcome, tag), result list with snippet preview, detail view with full request/response, copy-to-clipboard, re-run-prompt action. Tests cover search wiring, filter composition, and detail rendering. |
| **ZV-046** | **GUI: Multi-model comparison + Budgets.** Implement `frontend/src/routes/compare/` and `frontend/src/routes/budgets/`: side-by-side comparison view (prompt input, target selection, response panes, winner-tagging), budget configuration (definitions, thresholds, alert preferences), spend-vs-budget visualization. Tests cover comparison flow, budget CRUD, and alert preview. |
| **ZV-047** | **GUI: Settings & preferences.** Implement `frontend/src/routes/settings/`: general (start at login, system-tray behavior, network port, log level), privacy (redaction patterns, including custom-regex CRUD), update settings (auto-update, channel), import/export of configuration, "About Zevaro" dialog with version + license + GitHub link. The privacy screen prominently states "Zevaro contains no telemetry — there is nothing to opt into or out of." Tests cover settings persistence, validation, and import/export round-trip. |
| **ZV-048** | **GUI: First-launch onboarding.** Implement `frontend/src/routes/onboarding/`: welcome screen, provider quick-setup (paste keys for the providers you use), local model detection display ("we found Ollama running with N models"), policy bundle selection, sample prompt routed end-to-end to confirm everything works. Onboarding guards `/` until completed; completion is persisted. Tests cover the flow and the sample-prompt round-trip. |

### Phase 9 — Distribution (ZV-049 through ZV-053)

| Task | Description |
|------|-------------|
| **ZV-049** | **Audit refresh + GUI completion sweep.** Re-run the audit and update `Zevaro-Audit.md`. Sweep the GUI for accessibility (keyboard navigation, ARIA, contrast against §5 palette) and for visual consistency. |
| **ZV-050** | **macOS installer with code signing.** Implement `installers/macos/`: Wails packaging into `.app`, code signing with Developer ID via env-var-supplied identity, notarization via `notarytool`, stapling, DMG generation, and the GitHub Actions workflow that produces a signed/notarized DMG on release tag. Tests cover the pipeline via CI dry-run mode. |
| **ZV-051** | **Windows installer with code signing.** Implement `installers/windows/`: Wails packaging into `.exe`, code signing via SignTool with EV cert, MSI generation via WiX, portable EXE variant, GitHub Actions workflow. Tests cover the pipeline via CI dry-run mode. |
| **ZV-052** | **Linux packages.** Implement `installers/linux/`: `.deb` for Debian/Ubuntu (with `postinst`/`prerm` for systemd autostart opt-in), `.rpm` for Fedora/RHEL, AppImage with embedded checksum, GPG signing of all artifacts, GitHub Actions workflow. Tests cover the pipeline via CI dry-run mode. |
| **ZV-053** | **Auto-update mechanism.** Implement `internal/updater/` per §3.16: signed manifest format and validation, Ed25519 signature verification, download with hash verification, atomic binary replacement on next start, rollback on failure, GUI integration. Tests cover signature verification (positive + tampered), atomic replacement semantics, and rollback. |

### Phase 10 — Quality, CI, and Integration Testing (ZV-054 through ZV-057)

| Task | Description |
|------|-------------|
| **ZV-054** | **Audit refresh.** Re-run the audit and update `Zevaro-Audit.md` to reflect the distribution layer. |
| **ZV-055** | **Cross-provider integration test suite.** Implement `test/integration/`: full request → routing → provider → response paths against mock provider servers covering every supported provider's quirks, including tool calls, streaming, error cases, and rate-limit handling. Tests run in CI on every commit. |
| **ZV-056** | **End-to-end GUI tests.** Implement `test/e2e/` using Playwright against the built Wails app: install, onboarding, configure provider, route a prompt, view in dashboard, search history, configure budgets, multi-model comparison flow. Tests run in CI on every commit (nightly for cross-platform matrix). |
| **ZV-057** | **Release pipeline.** Implement `.github/workflows/release.yaml`: triggered by semver tag, runs full test suite, builds for all platforms, signs and notarizes, generates checksums, creates a GitHub Release with all artifacts, updates the auto-update manifest, posts release notes. Includes a dry-run mode against a fake tag. Tests cover the workflow logic in dry-run. |

### Phase 11 — Documentation and Launch (ZV-058 through ZV-063)

| Task | Description |
|------|-------------|
| **ZV-058** | **Audit refresh.** Re-run the audit and update `Zevaro-Audit.md`. |
| **ZV-059** | **Documentation site.** Build the docs site under `docs/` using Astro: getting started, architecture overview (a public-facing distillation of this document), provider setup guides, routing rules tutorial, troubleshooting, FAQ, contributor guide. Deploy via GitHub Pages or Cloudflare Pages mapped to `zevaro.ai/docs`. Tests cover link validity, build success, and search-index generation. |
| **ZV-060** | **Client integration guides.** Write a copy-paste-ready guide for each major AI client: Claude Code, Cline, Continue, Aider, Cursor (custom mode), Zed, OpenCode, the OpenAI Python SDK, the Anthropic Python SDK, and the Anthropic TypeScript SDK. Each guide is a single page in the docs site with the exact base-URL and key configuration the user needs. |
| **ZV-061** | **Launch HN post + demo video script.** Write the Hacker News launch post (concrete walkthrough of one user problem the gateway solves end-to-end, not "we built a thing") and the script for a 90-second demo video showing Claude Code routing through Zevaro with the spend dashboard updating live and a privacy rule keeping sensitive prompts on a local model. Both committed to `docs/launch/`. |
| **ZV-062** | **`CONTRIBUTING.md` polish + `good-first-issue` seeding.** Polish the contributor guide: dev environment setup, running tests, writing a new provider plugin, submitting a PR, review process. Then seed the GitHub issues with a curated set of "good first issue" labeled tasks (new provider integrations, additional redaction patterns, additional client integration guides, dashboard polish items). |
| **ZV-063** | **Final audit + release readiness.** Re-run the audit, update `Zevaro-Audit.md`, and produce a short release-readiness summary at `docs/release-readiness.md` confirming every charter capability is implemented, tested, and documented. Cuts the v1.0.0 tag. |

---

## 7. Task Completion Summary Template

Every task completes with the following report appended to the response. Fields must be filled in literally — placeholders must be replaced with real values, not left as `[...]`.

```
══════════════════════════════════════════════════════════════
                    TASK COMPLETION REPORT
══════════════════════════════════════════════════════════════

TASK:    <ID> — <short description>
PROJECT: Zevaro
BRANCH:  <branch>

FILES CREATED (<count> total, <insertions> insertions)
  <path>
  ...

FILES MODIFIED
  <path> — <what changed>
  ...

TESTS
  Go unit tests:        <passed> / <total> passing
  Go integration tests: <passed> / <total> passing
  Frontend unit tests:  <passed> / <total> passing
  E2E tests (if applicable): <passed> / <total> passing
  Coverage:             <%> line coverage on authored code

DEPENDENCIES
  Added:   <list or "none">
  Removed: <list or "none">

═══════════════════════════════════════════════════════════════
                    VERIFICATION RESULTS
═══════════════════════════════════════════════════════════════

BUILD:
  Command: make build
  Exit:    0
  Status:  PASS

DAEMON LAUNCH (if applicable):
  Command:  ./bin/zevaro
  Health:   /healthz returned 200
  Stable:   process up >= 10s
  Status:   PASS

WAILS GUI LAUNCH (if applicable):
  Command:  wails dev (smoke)
  Window:   appeared
  Console:  no errors
  Status:   PASS

STATIC ANALYSIS:
  Go:        golangci-lint run — "0 issues."
  Frontend:  eslint + tsc --noEmit — "0 issues."
  Status:    PASS

TESTS:
  Command: make test
  Result:  <N> tests passing
  Status:  PASS

═══════════════════════════════════════════════════════════════
OVERALL: ALL CHECKS PASS — TASK VERIFIED COMPLETE
═══════════════════════════════════════════════════════════════

SUB-AGENT USAGE:
  Sub-agents spawned:           <count>
  Full directives propagated:   <YES/NO/N/A>
  Sequential preconditions:     <YES/NO/N/A>

SCOPE ADHERENCE: CONFIRMED
  <one-line confirmation that only the task's scope was touched>

COMMIT:   <full SHA>
BRANCH:   <branch>
PUSHED:   YES / NO — <origin ref range>

OBSERVATIONS (Not Actioned — For Architect Review)
  1. <observation> — actioned: false
  2. <observation> — actioned: false

══════════════════════════════════════════════════════════════
```

---

## 8. Cross-Cutting Conventions

- **Constants.** Every magic number, key name, file extension, duration, or size lives in `internal/constants/` (Go) or `frontend/src/lib/constants.ts` (frontend). No literals in implementation code.
- **Logging.** `slog` is initialized once in `cmd/zevaro/main.go`. The logger is propagated via `context.Context` using `slog.With(...)` to add request-scoped attributes. `log.Println` and `fmt.Println` are banned by the linter (`forbidigo`).
- **Documentation.** GoDoc on every exported package, type, function, method, variable, and constant — except generated code and database entity structs. TSDoc on every exported TypeScript class and public method.
- **Errors.** Use sentinel errors (`var ErrXxx = errors.New(...)`) for distinguishable failure modes. Wrap errors with `fmt.Errorf("...: %w", err)`. Never put user content in error messages.
- **Testing.** Unit tests live alongside source files (`foo.go` ↔ `foo_test.go`). Integration tests live in `test/integration/`. E2E tests live in `test/e2e/`. Shared fixtures in `test/fixtures/`. Recorded provider responses in `test/fixtures/cassettes/` via `go-vcr`. `testify` is the single assertion library. `httptest` is the single HTTP-test harness. In-memory SQLite for storage tests. 100% line coverage on authored code is mandatory.
- **Codegen.** Wails generates frontend bindings; OpenAPI generates docs. Both are checked into the repo and verified up-to-date in CI (a stale-binding check fails the build).
- **Theming.** Tailwind CSS variables per §5. Dark mode is default; light mode is fully supported. Theme is stored in preferences and applied at app startup before the first paint.
- **No telemetry, ever.** Not opt-in, not anonymous, not "to make the product better." None. The `internal/telemetry/` package is intentionally empty; if a future task introduces analytics, that task fails the architecture review.
- **No cloud dependency at runtime.** The daemon must be fully functional with the network disconnected, except for actually reaching cloud providers when the user has configured them. The auto-update check and the connectivity canary are the only outbound calls Zevaro itself makes.
- **No Java, no Maven, no Flyway, no Hibernate, no Gradle.** Go + Wails + SQLite + GORM. The conventional patterns those tools represent (Hibernate → AutoMigrate; Flyway → versioned migrations) are honored using the Go-native equivalents.

---

## 9. Critical Execution Rules

The Critical Execution Rules defined in `CLAUDE.md` — 15-Minute Test Rule, No Silent Spinning, Fail Fast on Tests, Task Time Budget, Build Verification Is Non-Negotiable — apply verbatim to every Zevaro task. These rules override any other guidance, including this architecture document. Read them before every task.
