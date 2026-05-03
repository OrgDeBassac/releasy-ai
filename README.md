# 🚀 Releasy

A remote-first CLI tool for managing semantic releases and `REL-` branches with precision and safety.

## 🌟 Overview

**Releasy** is designed to be the trigger for your GitHub-based release workflows. It provides a bridge between your local development environment and remote CI/CD by synchronizing with `origin`, calculating the next semantic version, and pushing the necessary tags and branches to trigger your production builds.

## ✨ Key Features

- **Remote-First Workflow**: Always fetches from `origin` to ensure you are working with the latest repository state.
- **Visual Release Mapping**: See exactly where your `REL-` branches stand in relation to semantic tags.
- **Automated SemVer**: Intelligent calculation of Major, Minor, and Patch increments.
- **Safety Guaranteed**: Includes dirty working directory checks and confirmation prompts before any mutation.
- **Trigger-Ready**: Automatically pushes new tags and release branches to `origin` to kick off GitHub Actions or other CI/CD pipelines.

## 🛠 Tech Stack

- **Language**: [Go (Golang)](https://go.dev/)
- **CLI Framework**: [Cobra](https://github.com/spf13/cobra)
- **Git Interaction**: Native `git` binary calls via `os/exec`

## 🚀 Getting Started

### Installation (Coming Soon)

```bash
# Clone the repository
git clone https://github.com/username/releasy-ai.git

# Build the binary
go build -o releasy
```

### Usage

#### Visualize Release State
List all `REL-` branches and their associated semantic tags.
```bash
releasy list
```

#### Trigger a New Release
Calculate and apply the next semantic version.
```bash
# Create a patch release (e.g., v1.0.1 -> v1.0.2)
releasy release patch

# Create a minor release and a new REL- branch (e.g., v1.1.0 and REL-1.1)
releasy release minor

# Create a major release (e.g., v2.0.0 and REL-2.0)
releasy release major
```

## 🛡 Safety Checks

Releasy performs several checks before performing any Git actions:
1. **Fetch Sync**: Ensures your local tags/branches are up-to-date with `origin`.
2. **Clean Workspace**: Verifies there are no uncommitted changes.
3. **Collision Detection**: Prevents creating tags that already exist on the remote.
4. **User Confirmation**: Displays the exact version and branch to be created before proceeding.

---

*Built with precision for streamlined engineering workflows.*
