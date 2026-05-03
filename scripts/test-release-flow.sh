#!/usr/bin/env bash
set -euo pipefail

# This script exercises the release flow locally in dry-run mode.
# It does NOT push tags. To run a full end-to-end test, create and push a tag
# to a fork or protected repo and let CI run.

echo "Running release-cli dry run for patch"
./release-cli --dry-run release patch

echo "Running unit tests"
go test ./...

echo "Done. To perform a real release, run: ./release-cli release patch and push the created tag to origin"
