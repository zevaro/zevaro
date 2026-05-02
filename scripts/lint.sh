#!/usr/bin/env bash
# lint.sh — Run all linters for Zevaro.
# Usage: ./scripts/lint.sh
set -euo pipefail

echo "==> Running golangci-lint"
golangci-lint run ./...

echo "==> Running frontend lint"
cd frontend && pnpm run lint

echo "==> All lint checks passed"
