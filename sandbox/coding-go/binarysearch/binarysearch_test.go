package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	desc string // Test Case name
	nums []int  // Slice we looing in (sorted)
	look int    // Value we looking for
	indx int    // Expected index
}{
	{desc: "Empty Array", nums: []int{}, look: 1, indx: -1},
	{desc: "Simple Array: Not Exists", nums: []int{10}, look: 1, indx: -1},
	{desc: "Simple Array: Exists", nums: []int{10}, look: 10, indx: 0},
	{desc: "Array: Not found", nums: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, look: 10, indx: -1},
	{desc: "Array:  9", nums: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, look: 9, indx: 9},
	{desc: "Array:  8", nums: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, look: 8, indx: 8},
	{desc: "Array:  7", nums: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, look: 7, indx: 7},
	{desc: "Array:  6", nums: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, look: 6, indx: 6},
	{desc: "Array:  5", nums: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, look: 5, indx: 5},
	{desc: "Array:  4", nums: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, look: 4, indx: 4},
	{desc: "Array:  3", nums: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, look: 3, indx: 3},
	{desc: "Array:  2", nums: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, look: 2, indx: 2},
	{desc: "Array:  1", nums: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, look: 1, indx: 1},
	{desc: "Array:  0", nums: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, look: 0, indx: 0},
	{desc: "Array: -1", nums: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, look: -1, indx: -1},
}

func TestBinarySearchSimple(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.indx, BinarySearchSimple(tt.nums, tt.look))
		})
	}
}

func TestBinarySearchShort(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.desc, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.indx, BinarySearchShort(tt.nums, tt.look))
		})
	}
}
