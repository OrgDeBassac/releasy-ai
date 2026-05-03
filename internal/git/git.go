package git

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// Run executes a git command and returns the output or an error
func Run(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("git command failed: %w, stderr: %s", err, stderr.String())
	}

	return strings.TrimSpace(stdout.String()), nil
}

// Fetch syncs with origin remote
func Fetch() error {
	_, err := Run("fetch", "--tags", "--prune", "origin")
	return err
}

// GetReleaseBranches returns all branches starting with REL-
func GetReleaseBranches() ([]string, error) {
	output, err := Run("branch", "-a", "--list", "*REL-*", "--format=%(refname:short)")
	if err != nil {
		return nil, err
	}

	if output == "" {
		return []string{}, nil
	}

	branches := strings.Split(output, "\n")
	uniqueBranches := make(map[string]bool)
	var result []string

	for _, b := range branches {
		name := strings.TrimPrefix(b, "origin/")
		if !uniqueBranches[name] {
			uniqueBranches[name] = true
			result = append(result, name)
		}
	}

	return result, nil
}

// GetDefaultBranch attempts to determine the repository's default branch (e.g., main)
func GetDefaultBranch() (string, error) {
	// Try to resolve origin/HEAD -> origin/main
	out, err := Run("rev-parse", "--abbrev-ref", "origin/HEAD")
	if err == nil && out != "" {
		// out is typically "origin/main"
		return strings.TrimPrefix(out, "origin/"), nil
	}

	// Fallback: try symbolic-ref on refs/remotes/origin/HEAD
	out, err = Run("symbolic-ref", "--short", "refs/remotes/origin/HEAD")
	if err == nil && out != "" {
		return strings.TrimPrefix(out, "origin/"), nil
	}

	// As a last resort, assume 'main'
	return "main", nil
}

// GetLatestTag returns the most recent semver tag reachable from a ref
func GetLatestTag(ref string) (string, error) {
	// git describe --tags --abbrev=0 --match "v*" <ref>
	output, err := Run("describe", "--tags", "--abbrev=0", "--match", "v*", ref)
	if err != nil {
		// If no tag is found, return empty string instead of error
		if strings.Contains(err.Error(), "fatal: No names found") {
			return "", nil
		}
		return "", err
	}
	return output, nil
}

// ListTags returns all tags matching v* sorted newest-first
func ListTags() ([]string, error) {
	output, err := Run("tag", "--list", "v*", "--sort=-v:refname")
	if err != nil {
		return nil, err
	}
	if strings.TrimSpace(output) == "" {
		return []string{}, nil
	}
	parts := strings.Split(strings.TrimSpace(output), "\n")
	return parts, nil
}

// GetTagInfo returns the commit date (ISO 8601) and subject message for a tag
func GetTagInfo(tag string) (string, string, error) {
	// Use git show to get committer date in ISO 8601 and the subject
	out, err := Run("show", "-s", "--format=%cI%n%s", tag)
	if err != nil {
		return "", "", err
	}
	lines := strings.SplitN(out, "\n", 2)
	date := ""
	msg := ""
	if len(lines) >= 1 {
		date = strings.TrimSpace(lines[0])
	}
	if len(lines) == 2 {
		msg = strings.TrimSpace(lines[1])
	}
	return date, msg, nil
}

// IsClean checks if the working directory is clean
func IsClean() (bool, error) {
	output, err := Run("status", "--porcelain")
	if err != nil {
		return false, err
	}
	return output == "", nil
}

// CreateTag creates a new local tag
func CreateTag(version string) error {
	_, err := Run("tag", version)
	return err
}

// Push pushes a ref to origin
func Push(ref string) error {
	_, err := Run("push", "origin", ref)
	return err
}

// CreateBranch creates a new branch
func CreateBranch(name string) error {
	_, err := Run("checkout", "-b", name)
	return err
}
