// file-header-check
package testdata

func UselessBreaks() {
	switch {
	case true:
		break // MATCH /useless break in case clause/
	case false:
		if false {
			break
		}
	}
}
