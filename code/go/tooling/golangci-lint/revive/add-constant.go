// file-header-check
package testdata

func foo(a float64, b string, c int, d string) {
	a = 1.0 // ignore
	b = "ignore"
	c = 2              // ignore
	println("lit", 12) // MATCH /avoid magic numbers like '12', create a named constant for it/
	if a == 12.50 {    // MATCH /avoid magic numbers like '12.50', create a named constant for it/
		if b == "lit" {
			d = "lit" // MATCH /string literal "lit" appears, at least, 3 times, create a named constant for it/
		}
		for i := 0; i < 1; i++ {
			println("lit")
		}
	}
}
