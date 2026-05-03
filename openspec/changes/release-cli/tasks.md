## 1. Project Setup

- [ ] 1.1 Initialize Go module (`go mod init release-cli`)
- [ ] 1.2 Initialize Cobra CLI structure (`main.go` and `cmd/` directory)
- [ ] 1.3 Implement a wrapper for executing Git commands via `os/exec`
- [ ] 1.4 Implement a `git fetch` utility to sync with `origin`

## 2. Release Visualization

- [ ] 2.1 Implement logic to list all local and remote `REL-*` branches
- [ ] 2.2 Implement logic to find the latest semantic version tag for a given branch
- [ ] 2.3 Implement the `list` command (ensuring it fetches first)

## 3. Release Calculation

- [ ] 3.1 Implement a semantic version parser to handle `vX.Y.Z` formats
- [ ] 3.2 Implement logic to calculate Major, Minor, and Patch increments
- [ ] 3.3 Add a dry-run mode to the release command

## 4. Release Execution

- [ ] 4.1 Implement a repository health check (clean working directory)
- [ ] 4.2 Implement the `release` command with Major/Minor/Patch options
- [ ] 4.3 Implement logic to create and push a new Git tag
- [ ] 4.4 Implement logic to create and push a new `REL-` branch (for Major/Minor)
- [ ] 4.5 Add confirmation prompts before any Git mutation

## 5. Testing and Validation

- [ ] 5.1 Verify visualization with mock Git repository state
- [ ] 5.2 Verify version calculation with various edge cases (e.g., first release, pre-releases)
- [ ] 5.3 Verify that tagging and branching work as expected in a test repo
