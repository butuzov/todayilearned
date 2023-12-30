// file-header-check
package testdata

func Getbaz(a string, b int) (int, float32, string, string) { // MATCH /unnamed results of the same type may be confusing, consider using named results/
	return b, float32(b), a, a
}
