// -------------------------------------
// Skip fields in json (tags)
// -------------------------------------
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	raw := []byte(`{"uno":"one", "ein":"one"}`)

	var data struct {
		Uno string `json:"uno"`
		Ein string `json:"-"` // <- this is minus, and it used for skipping.....
	}

	_ = json.Unmarshal(raw, &data)

	fmt.Printf("Unmarshaled - %#v\n", data)

	data.Ein = "odyn"

	fmt.Printf("Ein is fixed - %#v\n", data)

	jsonStr, _ := json.Marshal(data)
	fmt.Println("Marshaled", string(jsonStr))
}
