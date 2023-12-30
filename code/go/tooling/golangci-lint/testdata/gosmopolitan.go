package testdata

import "fmt"

func _() {
	fmt.Println("hello world")
	fmt.Println("你好，世界") // want `string literal contains rune in Han script`
	fmt.Println("こんにちは、セカイ")
}
