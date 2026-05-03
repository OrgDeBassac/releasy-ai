package cmd

import (
	"fmt"
	"os"
	"github.com/OrgDeBassac/releasy-ai/internal/git"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all REL- branches and their latest tags",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Fetching latest state from origin...")
		if err := git.Fetch(); err != nil {
			fmt.Printf("Warning: failed to fetch from origin: %v\n", err)
		}

		branches, err := git.GetReleaseBranches()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting branches: %v\n", err)
			os.Exit(1)
		}

		if len(branches) == 0 {
			fmt.Println("No REL- branches found.")
			return
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
		fmt.Fprintln(w, "BRANCH\tLATEST TAG\tSTATUS")

		for _, branch := range branches {
			tag, err := git.GetLatestTag(branch)
			if err != nil {
				tag = "error"
			}
			if tag == "" {
				tag = "no tags"
			}
			
			// Simple status check for now
			status := "stable"
			
			fmt.Fprintf(w, "%s\t%s\t%s\n", branch, tag, status)
		}
		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
