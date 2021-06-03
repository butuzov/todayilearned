package binarysearch

import "sort"

func BinarySearchShort(heystack []int, needle int) int {

	var index = sort.Search(len(heystack), func(i int) bool {
		return heystack[i] >= needle
	})

	if index < len(heystack) && heystack[index] == needle {
		return index
	}

	return -1
}

func BinarySearchSimple(heystack []int, needle int) int {

	if len(heystack) == 0 {
		return -1
	}

	var lo, hi = 0, len(heystack) - 1

	for lo < hi {
		middle := int((hi + lo) / 2)

		switch {
		case heystack[middle] == needle:
			return middle
		case heystack[middle] < needle:
			lo = middle + 1
		case heystack[middle] > needle:
			hi = middle - 1
		}
	}

	if heystack[lo] == needle {
		return lo
	}

	return -1
}
