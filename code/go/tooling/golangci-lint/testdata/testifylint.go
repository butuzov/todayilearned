package testdata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// OK_with_config

func TestTestifylint(t *testing.T) {
	var (
		predicate bool
		resultInt int
		arr       []string
	)

	assert.Equal(t, predicate, true) // want "bool-compare: use assert\\.True"
	assert.True(t, resultInt == 1)   // want "compares: use assert\\.Equal"
	assert.Equal(t, len(arr), 0)     // want "empty: use assert\\.Empty"
}
