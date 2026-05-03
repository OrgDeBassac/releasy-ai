package main

import (
	"fmt"
	"github.com/OrgDeBassac/releasy-ai/internal/git"
	"regexp"
	"sort"
)

func main() {
	branches, _ := git.GetReleaseBranches()
	fmt.Println("branches:", branches)
	defaultBranch, _ := git.GetDefaultBranch()
	tags, _ := git.ListTags()
	fmt.Println("tags:", tags)
	groupMap := make(map[string]string)
	re := regexp.MustCompile(`^v(\d+)\.(\d+)\.(\d+)$`)
	for _, t := range tags {
		m := re.FindStringSubmatch(t)
		if m == nil {
			continue
		}
		key := fmt.Sprintf("REL-%s.%s", m[1], m[2])
		if _, ok := groupMap[key]; !ok {
			groupMap[key] = t
		}
	}
	fmt.Println("groupMap:", groupMap)

	// build groups
	type kv struct {
		key   string
		major int
		minor int
		tag   string
	}
	var groups []kv
	for k, t := range groupMap {
		var maj, min int
		fmt.Sscanf(k, "REL-%d.%d", &maj, &min)
		groups = append(groups, kv{key: k, major: maj, minor: min, tag: t})
	}
	sort.Slice(groups, func(i, j int) bool {
		if groups[i].major != groups[j].major {
			return groups[i].major > groups[j].major
		}
		return groups[i].minor > groups[j].minor
	})
	fmt.Println("groups:", groups)

	// defaultLatest
	defaultLatest := ""
	if defaultBranch != "" {
		if lt, _ := git.GetLatestTag(defaultBranch); lt != "" {
			defaultLatest = lt
		}
	}
	fmt.Println("defaultBranch", defaultBranch, "defaultLatest", defaultLatest)

	finalBranches := []string{}
	seen := map[string]bool{}
	if defaultBranch != "" {
		finalBranches = append(finalBranches, defaultBranch)
		seen[defaultBranch] = true
	}
	for _, g := range groups {
		if g.tag == defaultLatest {
			continue
		}
		hasReal := false
		for _, b := range branches {
			if b == g.key {
				hasReal = true
				break
			}
		}
		fmt.Printf("group %s hasReal=%v\n", g.key, hasReal)
		if hasReal {
			continue
		}
		if !seen[g.key] {
			finalBranches = append(finalBranches, g.key)
			seen[g.key] = true
		}
	}
	fmt.Println("finalBranches after virtuals:", finalBranches)
	// add real branches
	sort.Strings(branches)
	for _, b := range branches {
		if !seen[b] {
			finalBranches = append(finalBranches, b)
			seen[b] = true
		}
	}
	fmt.Println("finalBranches final:", finalBranches)
}
