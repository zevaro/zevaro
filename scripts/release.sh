#!/usr/bin/env bash
# release.sh — Trigger a Zevaro release build via goreleaser.
# Full release pipeline implemented in ZV-057.
# Usage: ./scripts/release.sh <version-tag>
set -euo pipefail

VERSION="${1:-}"
if [[ -z "$VERSION" ]]; then
  echo "Usage: $0 <version-tag>  (e.g. v1.0.0)"
  exit 1
fi

echo "Release pipeline is implemented in ZV-057"
echo "To release, push tag: git tag $VERSION && git push origin $VERSION"
