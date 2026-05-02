# Zevaro — the local-first AI gateway

Zevaro is a desktop-first, open-source AI gateway for individual developers and small teams. It runs as a single native application on macOS, Windows, and Linux, and exposes both an OpenAI-compatible and an Anthropic-compatible HTTP surface on localhost. Any AI client that supports a configurable base URL — Claude Code, Cline, Continue, Aider, Cursor, Zed, OpenCode, or the OpenAI/Anthropic SDKs — can point at Zevaro as a drop-in replacement for the real API endpoints.

Behind that surface, Zevaro routes each request to whichever actual provider makes sense: Anthropic, OpenAI, xAI, Google, DeepSeek, Mistral, Groq, Together AI, Cohere, free-tier services, or local backends (Ollama, LM Studio, llama.cpp, vLLM). Routing is driven by user-configured rules, policy bundles, session tags, budget guardrails, and optional cost-aware auto-routing. Every request is logged locally with full token accounting and cost tracking.

All keys, prompts, history, routing rules, and policies live on your machine in a SQLite database. Nothing phones home. There is no telemetry, no cloud-hosted version, no paid tier.

**Status: pre-alpha, in active development.** Not ready for production use.

---

## Building from source

**Prerequisites:**

- Go 1.22+
- Node.js 20+ and [pnpm](https://pnpm.io)
- [Wails CLI v2](https://wails.io/docs/gettingstarted/installation): `go install github.com/wailsapp/wails/v2/cmd/wails@latest`
- [golangci-lint](https://golangci-lint.run): `brew install golangci-lint` (macOS) or see installation docs

**Commands:**

```sh
# Clone
git clone https://github.com/zevaro/zevaro.git
cd zevaro

# Install frontend dependencies
cd frontend && pnpm install && cd ..

# Run all tests
make test

# Run linters
make lint

# Build the native binary
make build

# Start the dev server with hot reload (requires Wails CLI)
make dev
```

The compiled binary lands at `build/bin/zevaro`.

**API reference:** The full OpenAPI 3.1 spec is at [`openapi.yaml`](openapi.yaml). Browsable HTML documentation is at [`api/openapi/docs/index.html`](api/openapi/docs/index.html) — open it in any browser, or run `make docs` to regenerate it.

---

## License

Apache 2.0. See [LICENSE](LICENSE) and [NOTICE](NOTICE).

Project home: [zevaro.ai](https://zevaro.ai)
