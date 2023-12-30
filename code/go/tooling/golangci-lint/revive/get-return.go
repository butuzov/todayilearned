// file-header-check
package testdata

func getfoo() {
}

func getBar(a, b int) { // MATCH /function 'getBar' seems to be a getter but it does not return any result/
}
