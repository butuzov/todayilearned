package testdata

import "bytes"

func _() {
	var bb bytes.Buffer

	bb.WriteString(string([]byte("foobar")))
}
