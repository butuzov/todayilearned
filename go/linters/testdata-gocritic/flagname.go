package testdata_gocritic

import "flag"

// Example for the rule: flagName

func _() {
	_ = flag.Bool(" foo ", false, "description")
}
