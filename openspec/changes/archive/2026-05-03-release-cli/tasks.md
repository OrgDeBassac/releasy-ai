## 1. Project Setup

- [x] 1.1 Initialize Go module (`go mod init release-cli`)
- [x] 1.2 Initialize Cobra CLI structure (`main.go` and `cmd/` directory)
- [x] 1.3 Implement a wrapper for executing Git commands via `os/exec`
- [x] 1.4 Implement a `git fetch` utility to sync with `origin`

## 2. Release Visualization

- [x] 2.1 Implement logic to list all local and remote `REL-*` branches
- [x] 2.2 Implement logic to find the latest semantic version tag for a given branch
- [x] 2.3 Implement the `list` command (ensuring it fetches first)

## 3. Release Calculation

- [x] 3.1 Implement a semantic version parser to handle `vX.Y.Z` formats
- [x] 3.2 Implement logic to calculate Major, Minor, and Patch increments
- [x] 3.3 Add a dry-run mode to the release command

## 4. Release Execution

- [x] 4.1 Implement a repository health check (clean working directory)
- [x] 4.2 Implement the `release` command with Major/Minor/Patch options
- [x] 4.3 Implement logic to create and push a new Git tag
- [x] 4.4 Implement logic to create and push a new `REL-` branch (for Major/Minor)
- [x] 4.5 Add confirmation prompts before any Git mutation

## 5. Testing and Validation

- [x] 5.1 Verify visualization with mock Git repository state
- [x] 5.2 Verify version calculation with various edge cases (e.g., first release, pre-releases)
- [x] 5.3 Verify that tagging and branching work as expected in a test repo
