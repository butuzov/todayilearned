// file-header-check
package testdata

func f280_3() (err error) {
	_ = 1
	return // MATCH /avoid using bare returns, please add return expressions/
}
