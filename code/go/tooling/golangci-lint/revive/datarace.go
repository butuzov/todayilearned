// file-header-check
package testdata

func datarace() (r int, c byte) {
	for _, p := range []int{1, 2} {
		go func() {
			print(r) // MATCH /potential datarace: return value r is captured (by-reference) in goroutine/
			print(p) // MATCH /datarace: range value p is captured (by-reference) in goroutine/
		}()
	}
	return r, c
}
