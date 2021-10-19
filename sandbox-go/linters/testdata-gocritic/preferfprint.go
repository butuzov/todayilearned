package testdata_gocritic

import (
	"fmt"
	"io"
)

func foobar11(w io.Writer) {
	w.Write([]byte(fmt.Sprintf("%x", 10)))
}
