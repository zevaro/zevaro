# Contributing to Zevaro

Thank you for your interest in contributing. This document explains how to set up a development environment, run tests, and submit changes.

Zevaro follows the conventions in [CONVENTIONS.md](CONVENTIONS.md). Every pull request must include tests and documentation in the same pass — there are no follow-up tasks.

---

## Development environment setup

### macOS / Linux

1. **Go 1.22+**
   - macOS: `brew install go`
   - Linux: download from [go.dev/dl](https://go.dev/dl/) or use your distro package manager

2. **Node.js 20+ and pnpm**
   - macOS: `brew install node pnpm`
   - Linux: install Node from [nodejs.org](https://nodejs.org), then `npm install -g pnpm`

3. **Wails CLI v2**
   ```sh
   go install github.com/wailsapp/wails/v2/cmd/wails@latest
   ```

4. **golangci-lint**
   - macOS: `brew install golangci-lint`
   - Linux: see [golangci-lint.run/usage/install](https://golangci-lint.run/usage/install/)

5. **Clone and install**
   ```sh
   git clone https://github.com/zevaro/zevaro.git
   cd zevaro
   cd frontend && pnpm install && cd ..
   go mod download
   ```

### Windows

1. Install Go 1.22+ from [go.dev/dl](https://go.dev/dl/)
2. Install Node.js 20+ from [nodejs.org](https://nodejs.org)
3. Install pnpm: `npm install -g pnpm`
4. Install Wails: `go install github.com/wailsapp/wails/v2/cmd/wails@latest`
5. Install golangci-lint from the [releases page](https://github.com/golangci/golangci-lint/releases)
6. Ensure a C compiler is available (Wails requires CGO): install [TDM-GCC](https://jmeubank.github.io/tdm-gcc/) or MSYS2

---

## Running tests

```sh
# All tests (Go + frontend)
make test

# Go tests only
make test-go

# Frontend tests only
make test-frontend
```

All authored code ships with 100% line coverage. Tests are not optional.

---

## Running linters

```sh
make lint
```

This runs `golangci-lint run` on Go code and `eslint + tsc --noEmit` on the frontend. CI enforces zero lint errors.

---

## Building

```sh
make build
```

This builds the frontend, then compiles the Go binary to `build/bin/zevaro`.

For hot-reload development:

```sh
make dev
```

---

## Where to file issues

Use [GitHub Issues](https://github.com/zevaro/zevaro/issues). For bugs, include the output of `zevaro --version`, your OS, and steps to reproduce. For feature requests, describe the problem you're solving, not just the solution you want.

---

## Submitting a pull request

1. Fork the repo and create a branch from `main`.
2. Make your changes. Tests and documentation ship in the same commit.
3. Verify `make build`, `make test`, and `make lint` all pass.
4. Open a PR against `main`. Fill in the pull request template completely.

PRs that are missing tests, missing documentation, or that fail CI will not be merged.

See [CONVENTIONS.md](CONVENTIONS.md) for the full engineering conventions.
