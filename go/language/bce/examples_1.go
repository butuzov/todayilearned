// Example from the go101 article.
package main


// Moving last bound check to be fun as first.

func f1_before(s []int) {
	_ = s[0] // line 8: bounds check
	_ = s[1] // line 9: bounds check
	_ = s[2] // line 10: bounds check
}

func f1_after(s []int) {
	_ = s[2] // line 14: bounds check
	_ = s[1] // line 15: bounds check eliminated!
	_ = s[0] // line 16: bounds check eliminated!
}



// Double checks are ignored.

func f2(s []int, index int) {
	_ = s[index] // line 24: bounds check
	_ = s[index] // line 25: bounds check eliminated!
}

// passing fixed size array, no ned for check.
func f3(a [5]int) {
	_ = a[4] // line 30: bounds check eliminated!
}
