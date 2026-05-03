# Create REL branch from tag or virtual REL group

What: Add a command to release-cli (e.g., `release-cli create-rel --from-tag vX.Y.Z` or `release-cli create-rel --from-rel REL-X.Y`) that creates a physical REL-X.Y branch from a given tag or the latest tag in a virtual REL group.

Why: Sometimes a release requires a long-lived branch for backports or maintenance after the tag is created. This command should be idempotent and allow creating a branch remotely and pushing it, ensuring correct base commit.

Scope:
- Add a `create-rel` command to the CLI with options: `--from-tag`, `--from-rel`, `--push`, `--force`.
- Validate inputs: tag exists; virtual REL group exists and has a latest tag.
- Create a local branch named `REL-X.Y` at the specified tag's commit and push to origin when `--push` is specified.
- Return clear error messages for invalid states (branch already exists, tag missing, etc.).

Risks/Notes:
- Must avoid accidental branch overwrites; require `--force` to overwrite existing REL branch.

Outcome: Allows operators to create REL branches reliably post-release.
