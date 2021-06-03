package testdataabstract

import (
	"context"
	ctxpkg "context"
	"fmt"
)

func ReadFromContext(context context.Context) context.Context {
	return ctxpkg.WithValue(context, "key", "val")
}

func main() {
	a := 1
	if true {
		a := 2
		fmt.Println(a)
	}
	fmt.Println(a)
}
