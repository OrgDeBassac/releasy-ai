package cmd

import (
	"fmt"
	"testing"

	"github.com/OrgDeBassac/releasy-ai/internal/semver"
)

func TestDeriveRelNameFromTag(t *testing.T) {
	cases := []struct{ tag, want string }{
		{"v1.2.3", "REL-1.2"},
		{"v0.0.1", "REL-0.0"},
		{"v10.20.30", "REL-10.20"},
	}
	for _, c := range cases {
		v, err := semver.Parse(c.tag)
		if err != nil {
			t.Fatalf("parse %s: %v", c.tag, err)
		}
		got := fmt.Sprintf("REL-%d.%d", v.Major, v.Minor)
		if got != c.want {
			t.Errorf("tag %s: got %s want %s", c.tag, got, c.want)
		}
	}
}
