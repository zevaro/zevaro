# Zevaro Makefile
# Targets: build, test, lint, package, release, dev
#
# All targets are runnable from the repo root.
# Requires: Go 1.22+, pnpm, wails, golangci-lint

BINARY     := zevaro
BUILD_DIR  := build/bin
CMD_PKG    := ./cmd/zevaro
FRONTEND   := ./frontend

# Detect the OS for platform-specific commands
GOOS   := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)

.PHONY: build test lint package release dev clean

## build: Compile the frontend and produce the Zevaro binary for the current platform.
build: frontend-build
	@echo "==> Building Zevaro daemon ($(GOOS)/$(GOARCH))"
	@mkdir -p $(CMD_PKG)/frontend
	@cp -r $(FRONTEND)/dist $(CMD_PKG)/frontend/dist
	CGO_ENABLED=1 go build -tags desktop -o $(BUILD_DIR)/$(BINARY) $(CMD_PKG)
	@echo "==> Binary: $(BUILD_DIR)/$(BINARY)"

## frontend-build: Build the React/TypeScript frontend bundle.
frontend-build:
	@echo "==> Building frontend"
	@cd $(FRONTEND) && pnpm install --frozen-lockfile && pnpm run build

## test: Run all Go unit tests and frontend Vitest tests.
test: test-go test-frontend

## test-go: Run Go tests with race detection and coverage.
test-go:
	@echo "==> Running Go tests"
	go test -race -coverprofile=coverage.out ./...
	@go tool cover -func=coverage.out | tail -1

## test-frontend: Run Vitest tests in non-watch mode.
test-frontend:
	@echo "==> Running frontend tests"
	@cd $(FRONTEND) && pnpm test --run

## lint: Run golangci-lint and frontend ESLint + tsc.
lint: lint-go lint-frontend

## lint-go: Run golangci-lint on all Go packages.
lint-go:
	@echo "==> Running golangci-lint"
	golangci-lint run ./...

## lint-frontend: Run ESLint and TypeScript type-check on the frontend.
lint-frontend:
	@echo "==> Running frontend lint"
	@cd $(FRONTEND) && pnpm run lint

## package: Placeholder — packaging implemented in ZV-050/051/052.
package:
	@echo "Packaging is implemented in ZV-050/051/052"

## release: Placeholder — release pipeline implemented in ZV-057.
release:
	@echo "Release pipeline is implemented in ZV-057"

## dev: Start Wails in hot-reload development mode.
dev:
	@echo "==> Starting Wails dev server"
	wails dev

## clean: Remove build artifacts.
clean:
	@rm -rf $(BUILD_DIR) coverage.out $(CMD_PKG)/frontend/dist

## help: List all available targets.
help:
	@grep -E '^##' Makefile | sed 's/## //'
