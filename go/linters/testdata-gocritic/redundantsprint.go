package testdata_gocritic

// redundantSprint

import "fmt"

func _() {
	var x saa
	fmt.Println(fmt.Sprint(x))
}

type saa struct{}

func (s saa) String() string {
	return "saa{}"
}
