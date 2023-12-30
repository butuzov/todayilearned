package main

import (
	"fmt"
	"reflect"
)

// --------------------------------------------
// Some quick example of struct tags (tags).
// -------------------------------------------

func main() {
	data := struct {
		Variable1 string `json:"name,omitempty"`
		Variable2 string `unformated strings are OK!`
	}{
		"dummy1",
		"dummy2",
	}

	t := reflect.TypeOf(data)
	for _, val := range []string{"Variable1", "Variable2", "Variable3"} {
		if field, ok := t.FieldByName(val); ok {
			fmt.Printf("\nField: .%s\n", val)
			fmt.Printf("\tWhole tag : %q\n", field.Tag)
			fmt.Printf("\tValue of 'json': %q\n", field.Tag.Get("json"))
		} else {
			fmt.Printf("\nNo Field: .%s\n", val)
		}
	}
}
