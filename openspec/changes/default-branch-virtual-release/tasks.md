Tasks to implement "default branch as virtual release branch"

Todos:
1. Update release-cli behavior
   - id: release-cli-update
   - description: Add `--create-rel` flag and avoid creating REL- branches by default when releasing on default branch. Detect tag-based releases in CI.
   - status: done

2. Update CI workflow
   - id: ci-tags
   - description: Ensure tags.yml runs publish job on tag pushes and uploads artifacts (already implemented). Verify CI runs in GitHub actions on pushed tags.
   - status: pending

3. Update documentation
   - id: docs-update
   - description: Update CONTRIBUTING.md and docs/release.md to explain tag-driven releases and when to create REL- branches.
   - status: pending

4. Integration tests & validation
   - id: test-release-flow
   - description: Test end-to-end flow: run release-cli in dry-run and real mode on main; push tag; ensure CI builds and publishes assets to GitHub Release.
   - status: pending

5. Back-compat checks
   - id: backcompat
   - description: Search scripts/automation that expect REL- branches and update them to handle tag-only releases or check for REL- existence.
   - status: pending

Dependencies:
- release-cli-update depends on review and small code change in cmd/release.go
- ci-tags already updated; verify and adjust if needed

Notes:
- Use issue/PR for release-cli changes, include tests for the new flag and behavior.
- Consider adding a non-interactive mode for CI to create tags automatically when releasing from CI (use env var or flag).

Next steps:
- Confirm task list and permission to implement. Run `/opsx:apply` to start making the code changes and tests.