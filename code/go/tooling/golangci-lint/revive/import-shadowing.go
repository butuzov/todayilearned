// file-header-check
package testdata

import (
	ast "go/ast"
)

const str = "" // MATCH /The name 'str' shadows an import name/

type myAst struct {
	ast *ast.GenDecl
}

func (ast myAst) foo() {} // MATCH /The name 'ast' shadows an import name/
