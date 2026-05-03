## Why

Managing releases in a repository with `REL-` branches and semantic versioned tags manually is error-prone and tedious. A CLI tool will provide a streamlined way to visualize the current release state and safely perform semantic version increments.

## What Changes

- A new CLI tool `releasy` (or similar name) will be introduced.
- Capability to list and visualize `REL-` branches alongside their associated semantic version tags, synchronized with the `origin` remote.
- Capability to trigger the creation of a new semantic version release (Major, Minor, or Patch) by creating and pushing tags/branches to `origin`, thereby triggering GitHub workflows.

## Capabilities

### New Capabilities
- `release-visualization`: Visualize the relationship between release branches and semantic tags.
- `release-management`: Create new semantic releases (Major, Minor, Patch) based on current repository state.

### Modified Capabilities
<!-- No existing capabilities found in openspec/specs/ -->

## Impact

- New CLI utility in the codebase.
- Potential integration with Git for branch and tag manipulation.
- No impact on existing production code as this is a new tool.
