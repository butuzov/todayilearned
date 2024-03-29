// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// file-header-check
package testdata

type (
	A string
	B = string
	C int
	D = uintptr
)

func StringTest() {
	var (
		i int
		j rune
		k byte
		l C
		m D
		n = []int{0, 1, 2}
		o struct{ x int }
	)
	const p = 0
	_ = string(i) // MATCH /dubious convertion of an integer into a string, use strconv.Itoa/
	_ = string(j)
	_ = string(k)
	_ = string(p)    // MATCH /dubious convertion of an integer into a string, use strconv.Itoa/
	_ = A(l)         // MATCH /dubious convertion of an integer into a string, use strconv.Itoa/
	_ = B(m)         // MATCH /dubious convertion of an integer into a string, use strconv.Itoa/
	_ = string(n[1]) // MATCH /dubious convertion of an integer into a string, use strconv.Itoa/
	_ = string(o.x)  // MATCH /dubious convertion of an integer into a string, use strconv.Itoa/
}
