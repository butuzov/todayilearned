package testdata_gocritic

import "flag"

// Example for the rule: flagDeref

func _() {
	_ = *flag.Bool("b", false, "b docs")
}
