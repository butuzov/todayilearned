package testdata_gocritic

import (
	"fmt"
	"go/ast"
)

// Example for the rule: caseOrder

func _(x interface{}) {
	switch x.(type) {
	case ast.Expr:
		fmt.Println("expr")
	case *ast.BasicLit:
		fmt.Println("basic lit") // Never executed
	}
}
