package cmd

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/OrgDeBassac/releasy-ai/internal/git"
	"github.com/OrgDeBassac/releasy-ai/internal/semver"

	"github.com/spf13/cobra"
)

var fromTag string
var fromRel string
var pushBranch bool
var forceBranch bool

var createRelCmd = &cobra.Command{
	Use:   "create-rel",
	Short: "Create a REL-X.Y branch from a tag or virtual REL group",
	Run: func(cmd *cobra.Command, args []string) {
		if fromTag == "" && fromRel == "" {
			fmt.Fprintln(os.Stderr, "Error: --from-tag or --from-rel is required")
			os.Exit(1)
		}

		// Prefer --from-tag when both provided
		if fromTag != "" && fromRel != "" {
			fmt.Printf("Both --from-tag and --from-rel provided; preferring --from-tag (%s)\n", fromTag)
		}

		tag := fromTag
		if tag == "" {
			// Resolve from virtual REL group
			re := regexp.MustCompile(`REL-(\d+)\.(\d+)`)
			m := re.FindStringSubmatch(fromRel)
			if m == nil {
				// allow passing just X.Y as well
				re2 := regexp.MustCompile(`^(\d+)\.(\d+)$`)
				m = re2.FindStringSubmatch(fromRel)
				if m == nil {
					fmt.Fprintf(os.Stderr, "Invalid --from-rel value: %s. Expect REL-X.Y or X.Y\n", fromRel)
					os.Exit(1)
				}
			}
			maj := m[1]
			min := m[2]
			// find latest tag for this group
			tags, err := git.ListTags()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error listing tags: %v\n", err)
				os.Exit(1)
			}
			pattern := fmt.Sprintf("^v%s\\.%s\\.\\d+$", maj, min)
			r := regexp.MustCompile(pattern)
			found := ""
			for _, t := range tags {
				if r.MatchString(t) {
					found = t
					break
				}
			}
			if found == "" {
				fmt.Fprintf(os.Stderr, "No tags found for virtual REL-%s.%s\n", maj, min)
				os.Exit(1)
			}
			tag = found
		}

		// Validate tag exists
		exists, err := git.TagExists(tag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error checking tag existence: %v\n", err)
			os.Exit(1)
		}
		if !exists {
			fmt.Fprintf(os.Stderr, "Tag %s not found\n", tag)
			os.Exit(1)
		}

		// Derive REL name
		v, err := semver.Parse(tag)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse tag %s: %v\n", tag, err)
			os.Exit(1)
		}
		branchName := fmt.Sprintf("REL-%d.%d", v.Major, v.Minor)

		// Check local/remote existence
		localExists, _ := git.BranchExistsLocal(branchName)
		remoteExists, _ := git.BranchExistsRemote(branchName)

		if localExists || remoteExists {
			if !forceBranch {
				fmt.Fprintf(os.Stderr, "Branch %s already exists. Use --force to recreate.\n", branchName)
				os.Exit(1)
			}
			// attempt delete
			if localExists {
				if err := git.DeleteLocalBranch(branchName); err != nil {
					fmt.Fprintf(os.Stderr, "Warning: failed to delete local branch %s: %v\n", branchName, err)
				}
			}
			if remoteExists {
				if err := git.DeleteRemoteBranch(branchName); err != nil {
					fmt.Fprintf(os.Stderr, "Warning: failed to delete remote branch %s: %v\n", branchName, err)
				}
			}
		}

		// Create branch at tag
		fmt.Printf("Creating branch %s at %s...\n", branchName, tag)
		if err := git.CreateBranchAtTag(branchName, tag); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating branch: %v\n", err)
			os.Exit(1)
		}

		// Optionally push
		if pushBranch {
			fmt.Printf("Pushing branch %s to origin (force=%v)...\n", branchName, forceBranch)
			if err := git.PushWithForce(branchName, forceBranch); err != nil {
				fmt.Fprintf(os.Stderr, "Error pushing branch: %v\n", err)
				os.Exit(1)
			}
		}

		fmt.Printf("Branch %s created and checked out locally.\n", branchName)
		if pushBranch {
			fmt.Println("Branch pushed to origin.")
		}
	},
}

func init() {
	createRelCmd.Flags().StringVar(&fromTag, "from-tag", "", "Tag to create branch from (vX.Y.Z)")
	createRelCmd.Flags().StringVar(&fromRel, "from-rel", "", "Virtual REL group to use (REL-X.Y or X.Y)")
	createRelCmd.Flags().BoolVar(&pushBranch, "push", false, "Push created branch to origin")
	createRelCmd.Flags().BoolVar(&forceBranch, "force", false, "Force recreate branch if it exists")
	rootCmd.AddCommand(createRelCmd)
}
