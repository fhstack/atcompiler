package parser

import (
	"github.com/l-f-h/atcompiler/craft/ast"
)

type Parser interface {
	Parse(script string) (ast.ASTNode, error)
	DumpAST(node ast.ASTNode, indent string)
}
