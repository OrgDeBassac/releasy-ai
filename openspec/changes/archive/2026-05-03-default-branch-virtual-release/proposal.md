# Consider default branch a virtual release branch

What: Treat the repository's default branch (main) as a virtual release branch that carries released versions via semantic tags (v*). REL- branches continue to exist but are only created when a true long-lived release branch is needed.

Why: Simplifies workflow for projects that prefer tagging releases on the default branch rather than maintaining per-release branches. Reduces branch clutter, aligns CI (already triggers on v* tags), and makes releases lightweight: a tag identifies a release while the default branch remains the main development line.

Scope:
- Change release process so releasing is primarily tag-driven (vX.Y.Z) on main
- Prevent automatic creation of REL- branches except when explicitly requested
- Ensure tooling (release-cli, CI) recognizes tags and treats default branch as eligible for releases

Risks/Notes:
- Consumers expecting REL- branches for backports will need to create them explicitly
- CI/workflows must be adapted to run on tags (already added in tags.yml)

Outcome: A streamlined release flow using tags on the default branch; REL- branches kept for exceptional long-lived maintenance releases.