## ADDED Requirements

### Requirement: List REL-branches
The system SHALL list all local and remote branches that follow the `REL-*` naming convention.

#### Scenario: List branches
- **WHEN** the user runs the `list` command
- **THEN** the system displays a list of all branches starting with `REL-`

### Requirement: Show associated tags
The system SHALL identify and display semantic version tags associated with each `REL-` branch (e.g., tags reachable from the branch tip).

#### Scenario: Display tags with branches
- **WHEN** the user runs the `list` command
- **THEN** for each `REL-` branch, the system displays the most recent semantic version tag and any other relevant release tags.

### Requirement: Visualize release state
The system SHALL provide a visual representation (e.g., a table or tree) showing the relationship between `REL-` branches and their tags, highlighting which branches are "ahead" or "behind" in terms of versioning.

#### Scenario: Visual table output
- **WHEN** the user runs the `list` command
- **THEN** the system outputs a formatted table with columns for Branch Name, Latest Tag, and Status.
