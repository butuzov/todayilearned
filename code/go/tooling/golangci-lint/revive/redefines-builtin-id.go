// file-header-check
package testdata

type data struct{}

func (this data) vmethod() {
	nil := true // MATCH /assignment creates a shadow of built-in identifier nil/
	_ = nil
}
