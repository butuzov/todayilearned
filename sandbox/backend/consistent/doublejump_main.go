package main

import (
	"fmt"

	"github.com/edwingeng/doublejump"
)

func main() {
	h := doublejump.NewHash()
	for i := 0; i < 10; i++ {
		h.Add(fmt.Sprintf("node%d", i))
	}

	fmt.Println(h.Len())
	fmt.Println(h.LooseLen())

	fmt.Println(h.Get(1000))
	// fmt.Println(h.Get("foobar"))
	fmt.Println(h.Get(2000))
	fmt.Println(h.Get(3000))
	fmt.Println(h.All())
}
