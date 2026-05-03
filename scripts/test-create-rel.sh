#!/usr/bin/env bash
set -euo pipefail
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

# Build binary
mkdir -p "$ROOT/out"
cd "$ROOT"
go build -o "$ROOT/out/releasy-cli" .

tmpdir=$(mktemp -d)
printf "Using temp dir: %s\n" "$tmpdir"
pushd "$tmpdir" >/dev/null

git init
git config user.email "test@example.com"
git config user.name "CI Test"

echo "hello" > README.md
git add README.md
git commit -m "init"

git tag v1.2.3

# copy binary into temp repo
cp "$ROOT/out/releasy-cli" ./releasy-cli
chmod +x ./releasy-cli

# Run create-rel locally (no push)
./releasy-cli create-rel --from-tag v1.2.3 || true

printf "Branches in temp repo:\n"
git branch -a

popd >/dev/null

printf "Done. Inspect %s if needed.\n" "$tmpdir"
