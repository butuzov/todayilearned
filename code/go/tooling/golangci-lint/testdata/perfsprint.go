package testdata

import "fmt"

func _() {
	var s string
	var err error
	var i int
	var i64 uint64
	var b []byte
	var ui uint

	fmt.Sprintf("%s", s) // want "fmt.Sprintf can be replaced with just using the string"
	fmt.Sprint(s)        // want "fmt.Sprint can be replaced with just using the string"
	fmt.Sprintf("%s", err)
	fmt.Sprint(err)
	fmt.Sprintf("%t", b)           // want "fmt.Sprintf can be replaced with faster strconv.FormatBool"
	fmt.Sprint(b)                  // want "fmt.Sprint can be replaced with faster strconv.FormatBool"
	fmt.Sprintf("%d", i)           // want "fmt.Sprintf can be replaced with faster strconv.Itoa"
	fmt.Sprint(i)                  // want "fmt.Sprint can be replaced with faster strconv.Itoa"
	fmt.Sprintf("%d", i64)         // want "fmt.Sprintf can be replaced with faster strconv.FormatInt"
	fmt.Sprint(i64)                // want "fmt.Sprint can be replaced with faster strconv.FormatInt"
	fmt.Sprintf("%d", ui)          // want "fmt.Sprintf can be replaced with faster strconv.FormatUint"
	fmt.Sprint(ui)                 // want "fmt.Sprint can be replaced with faster strconv.FormatUint"
	fmt.Sprintf("%x", []byte{'a'}) // want "fmt.Sprintf can be replaced with faster hex.EncodeToString"
	fmt.Errorf("hello")            // want "fmt.Errorf can be replaced with errors.New"
	fmt.Sprintf("Hello %s", s)     // want "fmt.Sprintf can be replaced with string addition"
}
