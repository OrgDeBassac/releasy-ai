#!/usr/bin/env bash
set -euo pipefail

# Build releasy-cli into out/
mkdir -p out

echo "Building releasy-cli into out/releasy-cli"
go build -v -o out/releasy-cli .

echo "Build complete: out/releasy-cli"
