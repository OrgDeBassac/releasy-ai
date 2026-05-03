package cmd

import (
	"fmt"
	"os"
	"github.com/OrgDeBassac/releasy-ai/internal/git"
	"github.com/OrgDeBassac/releasy-ai/internal/semver"

	"github.com/spf13/cobra"
)

var dryRun bool

var releaseCmd = &cobra.Command{
	Use:   "release [major|minor|patch]",
	Short: "Create a new semantic release",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		releaseType := args[0]
		
		clean, err := git.IsClean()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error checking repository status: %v\n", err)
			os.Exit(1)
		}
		if !clean {
			fmt.Fprintln(os.Stderr, "Error: working directory is not clean. Please commit or stash changes.")
			os.Exit(1)
		}

		fmt.Println("Fetching latest state from origin...")
		if err := git.Fetch(); err != nil {
			fmt.Printf("Warning: failed to fetch from origin: %v\n", err)
		}

		latestTag, err := git.GetLatestTag("HEAD")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting latest tag: %v\n", err)
			os.Exit(1)
		}

		if latestTag == "" {
			fmt.Println("No existing semver tags found. Starting with v0.0.0")
			latestTag = "v0.0.0"
		}

		v, err := semver.Parse(latestTag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing latest tag %s: %v\n", latestTag, err)
			os.Exit(1)
		}

		var nextV semver.Version
		switch releaseType {
		case "major":
			nextV = v.NextMajor()
		case "minor":
			nextV = v.NextMinor()
		case "patch":
			nextV = v.NextPatch()
		default:
			fmt.Fprintf(os.Stderr, "Invalid release type: %s. Use major, minor, or patch.\n", releaseType)
			os.Exit(1)
		}

		fmt.Printf("Current version: %s\n", latestTag)
		fmt.Printf("Next version:    %s\n", nextV.String())

		if dryRun {
			fmt.Println("Dry run enabled. No changes made.")
			return
		}

		fmt.Printf("\nReady to release %s. Confirm? [y/N]: ", nextV.String())
		var response string
		fmt.Scanln(&response)
		if response != "y" && response != "Y" {
			fmt.Println("Release cancelled.")
			return
		}

		// Create tag
		fmt.Printf("Creating tag %s...\n", nextV.String())
		if err := git.CreateTag(nextV.String()); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating tag: %v\n", err)
			os.Exit(1)
		}

		// Create branch for Major/Minor
		if releaseType == "major" || releaseType == "minor" {
			branchName := fmt.Sprintf("REL-%d.%d", nextV.Major, nextV.Minor)
			fmt.Printf("Creating release branch %s...\n", branchName)
			if err := git.CreateBranch(branchName); err != nil {
				fmt.Printf("Warning: failed to create branch %s: %v (it might already exist)\n", branchName, err)
			} else {
				fmt.Printf("Pushing branch %s to origin...\n", branchName)
				if err := git.Push(branchName); err != nil {
					fmt.Fprintf(os.Stderr, "Error pushing branch: %v\n", err)
					os.Exit(1)
				}
			}
		}

		// Push tag
		fmt.Printf("Pushing tag %s to origin...\n", nextV.String())
		if err := git.Push(nextV.String()); err != nil {
			fmt.Fprintf(os.Stderr, "Error pushing tag: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("\nSuccessfully released %s and pushed to origin!\n", nextV.String())
		fmt.Println("GitHub workflows should now be triggered.")
	},
}

func init() {
	releaseCmd.Flags().BoolVarP(&dryRun, "dry-run", "d", false, "Display next version without applying changes")
	rootCmd.AddCommand(releaseCmd)
}
