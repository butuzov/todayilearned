// file-header-check
package testdata

func ur1() {
	ur1() // MATCH /unconditional recursive call/
}
