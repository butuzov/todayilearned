# Binary Search

Binary Search allows to reduce complexity of search in sorted array to `O(log n)`.

We don't need to implement it each time, again and again, we can use a simple `sort.Search`

```
func BinarySearch(heystack []int, needle int) int {

	var index = sort.Search(len(heystack), func(i int) bool {
		return heystack[i] >= needle
	})

	if index < len(heystack) && heystack[index] == needle {
		return index
	}

	return -1
}
```
