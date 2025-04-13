package version

import "github.com/blang/semver"

// CompareVersions compares two semantic version strings.
func CompareVersions(v1, v2 string) int {
	v1s, err1 := semver.ParseTolerant(v1)
	v2s, err2 := semver.ParseTolerant(v2)
	if err1 != nil || err2 != nil {
		return 0
	}
	return v1s.Compare(v2s)
}
