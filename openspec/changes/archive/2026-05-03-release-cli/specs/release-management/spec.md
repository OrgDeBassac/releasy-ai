## ADDED Requirements

### Requirement: Sync with Remote
The system SHALL fetch the latest tags and branches from the `origin` remote before performing any visualization or version calculation to ensure consistency.

#### Scenario: Fetch before list
- **WHEN** the user runs the `list` command
- **THEN** the system executes `git fetch --tags --prune` before displaying the output.

### Requirement: Calculate next semantic version
The system SHALL be able to calculate the next semantic version (Major, Minor, or Patch) based on the latest existing tag in the repository.

#### Scenario: Calculate next patch
- **WHEN** the current latest tag is `v1.2.3` and the user requests a `patch` release
- **THEN** the system calculates `v1.2.4` as the next version.

### Requirement: Create release tag
The system SHALL create a new git tag for the calculated semantic version on the current branch.

#### Scenario: Successfully tag a release
- **WHEN** the user confirms a new release for version `v1.2.4`
- **THEN** the system executes `git tag v1.2.4` and provides confirmation.

### Requirement: Create release branch
The system SHALL allow creating a new `REL-` branch (e.g., `REL-1.2`) when a Major or Minor release is initiated.

#### Scenario: New Minor release branch
- **WHEN** the user initiates a Minor release for `v1.3.0`
- **THEN** the system creates a new branch `REL-1.3` if it doesn't already exist.

### Requirement: Push to Origin
The system SHALL push the newly created tag and/or branch to the `origin` remote to trigger the actual release build on GitHub.

#### Scenario: Push tag after creation
- **WHEN** a new tag `v1.2.4` is created locally
- **THEN** the system executes `git push origin v1.2.4` to trigger the remote workflow.
