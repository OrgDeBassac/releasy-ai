package cmd

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/OrgDeBassac/releasy-ai/internal/git"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all REL- branches, virtual REL groups, and their latest tags",
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

		// Determine default branch (do not modify branches slice here)
		defaultBranch, _ := git.GetDefaultBranch()

		// Gather tags and derive virtual REL-MAJOR.MINOR groups
		tags, err := git.ListTags()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to list tags: %v\n", err)
			// fall back to printing branches only
			tags = []string{}
		}

		groupMap := make(map[string]string) // REL-x.y -> latest tag
		re := regexp.MustCompile(`^v(\d+)\.(\d+)\.(\d+)$`)
		for _, t := range tags {
			m := re.FindStringSubmatch(t)
			if m == nil {
				continue
			}
			maj := m[1]
			min := m[2]
			key := fmt.Sprintf("REL-%s.%s", maj, min)
			if _, ok := groupMap[key]; !ok {
				// tags are expected to be sorted newest-first by git; keep first seen
				groupMap[key] = t
			}
		}

		// Build sorted list of virtual REL keys (descending by major,minor)
		type kv struct {
			key   string
			major int
			minor int
			tag   string
		}
		var groups []kv
		for k, tg := range groupMap {
			// k format REL-M.N
			var maj, min int
			fmt.Sscanf(k, "REL-%d.%d", &maj, &min)
			groups = append(groups, kv{key: k, major: maj, minor: min, tag: tg})
		}
		sort.Slice(groups, func(i, j int) bool {
			if groups[i].major != groups[j].major {
				return groups[i].major > groups[j].major
			}
			return groups[i].minor > groups[j].minor
		})

		// Prepare final branch list per requested order:
		// 1) default branch
		// 2) virtual REL groups (sorted by version desc)
		// 3) other branches (alphabetical)
		finalBranches := []string{}
		seen := make(map[string]bool)

		// 1) default branch
		if defaultBranch != "" {
			finalBranches = append(finalBranches, defaultBranch)
			seen[defaultBranch] = true
		}

		// 2) virtual groups
		for _, g := range groups {
			if !seen[g.key] {
				finalBranches = append(finalBranches, g.key)
				seen[g.key] = true
			}
		}

		// 3) remaining real branches sorted alphabetically
		sort.Strings(branches)
		for _, b := range branches {
			if !seen[b] {
				finalBranches = append(finalBranches, b)
				seen[b] = true
			}
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
		fmt.Fprintln(w, "BRANCH\tLATEST TAG\tDATE\tMESSAGE\tSTATUS")

		for _, branch := range finalBranches {
			var tag, date, message, status string
			if t, ok := groupMap[branch]; ok {
				// virtual group
				tag = t
				status = "virtual"
				if d, m, err := git.GetTagInfo(tag); err == nil {
					// date: YYYY-MM-DD
					if idx := strings.Index(d, "T"); idx != -1 {
						date = d[:idx]
					} else {
						date = d
					}
					// truncate message to 50 chars (rune-safe)
					runes := []rune(m)
					if len(runes) > 50 {
						message = string(runes[:50]) + "..."
					} else {
						message = m
					}
				}
			} else {
				// real branch
				t, err := git.GetLatestTag(branch)
				if err != nil {
					tag = "error"
				} else if t == "" {
					tag = "no tags"
				} else {
					tag = t
					if d, m, err := git.GetTagInfo(tag); err == nil {
						if idx := strings.Index(d, "T"); idx != -1 {
							date = d[:idx]
						} else {
							date = d
						}
						runes := []rune(m)
						if len(runes) > 50 {
							message = string(runes[:50]) + "..."
						} else {
							message = m
						}
					}
				}
				status = "stable"
			}

			fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n", branch, tag, date, message, status)
		}
		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
