package testdata_gocritic

import "strings"

// Example for the rule: argOrder

func _() bool {
	return strings.HasPrefix("#", "long long string")
}
