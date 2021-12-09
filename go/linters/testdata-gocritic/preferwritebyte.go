package testdata_gocritic

import (
	"bytes"
	"io"
)

func writerbytes(w io.StringWriter, і []byte) {
	w.WriteString(string(s))
}

func _() {
	var b bytes.Buffer
	writerbytes(&b, []byte("foobar"))

	b.WriteString(string([]byte("foobar")))
}
