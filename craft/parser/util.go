package parser

import (
	"fmt"

	"github.com/l-f-h/atcompiler/craft/ast"
)

func DumpAST(node ast.ASTNode, indent string) {
	fmt.Println(indent + node.GetType().String() + " " + node.GetText())
	for _, child := range node.GetChildren() {
		DumpAST(child, indent+"\t")
	}
}
