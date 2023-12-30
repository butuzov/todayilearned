package testdata

// PROJECT URL: https://godoc.org/golang.org/x/tools/cmd/goimports
import (
	"fmt" // ERROR "File is not `goimports`-ed"
	"github.com/pkg/errors"
)

// OK_with_config

func Bar() {
	fmt.Print("x")
	_ = errors.New("oops")
}
