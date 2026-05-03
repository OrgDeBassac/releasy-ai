package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "releasy",
	Short: "Releasy is a CLI tool for managing semantic releases",
	Long: `A remote-first CLI tool for managing semantic releases and REL- branches 
with precision and safety. It helps visualize release state and trigger 
GitHub-based release builds by pushing tags and branches.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// Root flags can be defined here
}
