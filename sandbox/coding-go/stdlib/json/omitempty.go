package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	Struct := struct {
		A string
		B string  `json:"b"`
		C string  `json:"e"`
		D string  `json:",omitempty"`
		F *string `json:",omitempty"`
		i string  `json:"r"`
	}{
		i: "i",
	}

	b, _ := json.Marshal(&Struct)
	fmt.Println(string(b))
}
