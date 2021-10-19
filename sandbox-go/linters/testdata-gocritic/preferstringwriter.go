package testdata_gocritic

import (
	"bytes"
	"io"
)

func writerstring(w io.Writer, s string) {
	w.Write([]byte(s))
}

func _() {
	var b bytes.Buffer
	writerstring(&b, "foobar")
}
