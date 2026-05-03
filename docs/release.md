# Release process (tag-driven)

This repository treats the default branch (main) as a virtual release branch. Releases are performed by creating annotated tags following the pattern `vMAJOR.MINOR.PATCH` (for example, `v1.2.3`) and pushing them to the remote.

Key points:

- Use `release-cli` to compute the next version and create tags.
- By default, `release-cli` does NOT create `REL-` branches. To create a `REL-` branch use `--create-rel`.
- CI runs on tag pushes (see .github/workflows/tags.yml) and will build artifacts and attach them to the GitHub Release for that tag.

Examples:

- Dry run locally:

  ./release-cli --dry-run release patch

- Create a tag interactively:

  ./release-cli release patch

- Create a REL- branch when needed:

  ./release-cli --create-rel release minor

Notes:
- If you need a long-lived maintenance branch, create a `REL-X.Y` branch manually or pass `--create-rel`.
