## Context

The project requires a CLI tool to manage releases following a specific branching (`REL-`) and tagging (semver) convention. Currently, this process is manual. The tool needs to be integrated into the existing workspace or exist as a standalone utility within it.

## Goals / Non-Goals

**Goals:**
- Provide a clear visualization of the repository's release state, synchronized with `origin`.
- Automate the calculation and creation of semantic version tags.
- Automate the creation of release branches for Major/Minor versions.
- Ensure safety by validating the current repository state before tagging/branching.
- Trigger remote GitHub workflows by pushing tags and branches to `origin`.

**Non-Goals:**
- Support for branching strategies other than `REL-`.
- Complex merge conflict resolution (user handles Git conflicts).

## Decisions

### 1. Language and Runtime
- **Decision:** Go (Golang).
- **Rationale:** Go provides single-binary distribution with zero external dependencies, fast startup times, and excellent support for building CLIs (using libraries like `cobra`).
- **Alternatives:** Python (originally considered, but requires an interpreter and environment management).

### 2. Git Interaction
- **Decision:** Use `os/exec` to call the `git` binary directly.
- **Rationale:** Minimizes dependencies, ensures compatibility with the user's local Git configuration, and is the most reliable way to interact with Git for these operations.
- **Alternatives:** `go-git` (a pure Go implementation, but can be complex and lacks some native Git features).

### 3. CLI Framework
- **Decision:** Cobra.
- **Rationale:** The industry standard for building Go CLIs, providing robust command handling, flags, and help generation.
- **Alternatives:** Standard library `flag` package (too basic for a multi-command CLI).

### 4. Remote-First Workflow
- **Decision:** Always fetch from and push to `origin`.
- **Rationale:** The local CLI is a gateway to GitHub-based release builds. Ensuring the local state matches `origin` avoids version collisions and ensures workflows are triggered correctly.
- **Implementation:** Execute `git fetch --tags --prune` before any logic. Execute `git push origin <ref>` after local mutations.

### 5. Visualization Format
- **Decision:** Tabular output using the built-in `text/tabwriter` package.
- **Rationale:** Provides high readability and alignment in terminal environments without adding external dependencies.

## Risks / Trade-offs

- [Risk] → **Accidental Tagging:** User might create a tag on the wrong branch.
  - *Mitigation:* The tool SHALL show a confirmation prompt with the calculated version and branch name before performing any Git actions.
- [Risk] → **Dirty Working Directory:** Tagging a dirty repo can lead to inconsistent release states.
  - *Mitigation:* The tool SHALL check for a clean working directory before allowing a release.
- [Risk] → **Tag Name Collisions:** A tag with the same version might already exist.
  - *Mitigation:* The tool SHALL check for existing tags (after fetching) before attempting to create a new one.
- [Risk] → **Push Failures:** Remote push might fail (e.g., due to network or permissions).
  - *Mitigation:* The tool SHALL verify the success of the push command and alert the user if the remote workflow was not successfully triggered.
