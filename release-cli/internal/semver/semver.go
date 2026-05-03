package semver

import (
	"fmt"
	"strconv"
	"strings"
)

type Version struct {
	Major int
	Minor int
	Patch int
}

func Parse(v string) (Version, error) {
	v = strings.TrimPrefix(v, "v")
	parts := strings.Split(v, ".")
	if len(parts) != 3 {
		return Version{}, fmt.Errorf("invalid semver format: %s", v)
	}

	major, err := strconv.Atoi(parts[0])
	if err != nil {
		return Version{}, fmt.Errorf("invalid major version: %s", parts[0])
	}

	minor, err := strconv.Atoi(parts[1])
	if err != nil {
		return Version{}, fmt.Errorf("invalid minor version: %s", parts[1])
	}

	patch, err := strconv.Atoi(parts[2])
	if err != nil {
		return Version{}, fmt.Errorf("invalid patch version: %s", parts[2])
	}

	return Version{Major: major, Minor: minor, Patch: patch}, nil
}

func (v Version) String() string {
	return fmt.Sprintf("v%d.%d.%d", v.Major, v.Minor, v.Patch)
}

func (v Version) NextMajor() Version {
	return Version{Major: v.Major + 1, Minor: 0, Patch: 0}
}

func (v Version) NextMinor() Version {
	return Version{Major: v.Major, Minor: v.Minor + 1, Patch: 0}
}

func (v Version) NextPatch() Version {
	return Version{Major: v.Major, Minor: v.Minor, Patch: v.Patch + 1}
}
