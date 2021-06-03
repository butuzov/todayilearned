// Example from the go101 article.
package main


// Iteration over the key, so its always in the array.
func f5(s []int) {
	for i := range s {
		_ = s[i]
		_ = s[i:len(s)]      // Range is always in the bounds.
		_ = s[:i+1]          // i+1 wound't be used as it "not include bound" and i is in the slice
	}
}


// Side of slice is determined by len(s)
func f6(s []int) {
	for i := 0; i < len(s); i++ {
		_ = s[i]
		_ = s[i:len(s)]
		_ = s[:i+1]
	}
}

// Same by reverse.
func f7(s []int) {
	for i := len(s) - 1; i >= 0; i-- {
		_ = s[i]
		_ = s[i:len(s)]
	}
}

// no need to bounds check s we doing it.
func f8(s []int, index int) {
	if index >= 0 && index < len(s) {
		_ = s[index]
		_ = s[index:len(s)]
	}
}

// predefined len of s
func f9(s []int) {
	if len(s) > 2 {
	    _, _, _ = s[0], s[1], s[2]
	}
}
