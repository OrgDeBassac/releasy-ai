Design: Treat default branch as virtual release branch

Overview
- Releases are represented by annotated tags following vMAJOR.MINOR.PATCH (e.g., v1.2.3) pushed to the repository.
- The default branch (main) is considered a first-class place where releases may be tagged. No implicit REL- branch is required for each release.

Key components
1. release-cli
   - Update release CLI behavior:
     - When run on default branch and tagging, do not auto-create a REL- branch.
     - Provide explicit flag `--create-rel` to force creating a REL- branch when needed.
     - Detect tag-based releases and operate on the tag (use GITHUB_REF / git tag detection) when running in CI.
2. CI (tags.yml)
   - Already triggers on v* tags and builds per-OS artifacts.
   - Ensure publish job creates/uses GitHub Release for the pushed tag and uploads assets.
3. Branch policy and docs
   - Update CONTRIBUTING.md and release docs to explain that main is the virtual release branch and how to create REL- branches manually.

Behavior details
- Normal release flow:
  1. Dev runs `release-cli release <patch|minor|major>` on main (interactive or CI).
  2. release-cli computes next version and, unless `--create-rel` is set, creates and pushes the tag only.
  3. CI triggers on tag push, builds artifacts, and publishes release assets.
- Long-lived maintenance flow:
  - Use `release-cli --create-rel` or `git checkout -b REL-X.Y` to create a release branch when backports or long-running stability maintenance is required.

Backward compatibility
- Existing REL- branches remain supported.
- If current automation relied on REL- branches always existing, adapt scripts to prefer tags or check for tag presence on main.

Security & permissions
- Ensure tag creation is restricted or controlled by CI or required roles if necessary.

Telemetry
- Optional: Emit logs in release-cli when creating tags vs. branches to ease audit.

Implementation notes
- Keep changes minimal in release-cli: add a new flag and tag-detection code paths.
- Update docs and CI triggers; no database or API changes required.