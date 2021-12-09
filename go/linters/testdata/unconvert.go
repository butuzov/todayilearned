package testdata

import "io"

// PROJECT URL:  https://github.com/mdempsky/unconvert

func _() {
	var vbool bool
	_ = bool(vbool)     //@ unnecessary conversion
	_ = (*bool)(&vbool) //@ unnecessary conversion

	var writer io.Writer
	_ = (io.Writer)(writer)   //@ unnecessary conversion
	_ = (*io.Writer)(&writer) //@ unnecessary conversion
}
