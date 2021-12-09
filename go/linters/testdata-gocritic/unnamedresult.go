package testdata_gocritic

import "math"

// Example for the rule: unnamedResult

func _() (float64, float64) {
	return math.Pi, math.Phi
}
