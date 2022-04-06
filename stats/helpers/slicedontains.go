package helpers

import "github.com/byvko-dev/am-core/helpers/slices"

// Check if at least one of the strings in a slice is in the other slice
func SliceContains(a, b []string) bool {
	for _, item := range a {
		if slices.Contains(b, item) > -1 {
			return true
		}
	}
	return false
}
