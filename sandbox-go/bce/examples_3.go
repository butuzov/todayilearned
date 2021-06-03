// Example from the go101 article.
package main

import "math/rand"


// slice defined in visible scope, so its defined that len(s) == cap(s)
func fa() {
	s := []int{0, 1, 2, 3, 4, 5, 6}
	index := rand.Intn(7)
	_ = s[:index] // line 9: bounds check
	_ = s[index:] // line 10: bounds check eliminated!
}

func fb(s []int, i int) {
	_ = s[:i] // line 14: bounds check
	_ = s[i:] // line 15: bounds check, not smart enough?
}

func fc() {
	s := []int{0, 1, 2, 3, 4, 5, 6}
	s = s[:4]
	i := rand.Intn(7)
	_ = s[:i] // line 22: bounds check
	_ = s[i:] // line 23: bounds check, not smart enough?
}
