package parser

import (
	"fmt"

	"github.com/l-f-h/atcompiler/craft/ast"
	"github.com/l-f-h/atcompiler/craft/lexer"
	"github.com/l-f-h/atcompiler/craft/token"
)

type SimpleParser struct {
}

func NewSimpleParser() *SimpleParser {
	return &SimpleParser{}
}

func (p *SimpleParser) Parse(script string) (ast.ASTNode, error) {
	tokens := lexer.NewSimpleLexer().Tokenize([]byte(script))
	return p.prog(tokens)
}

func (p *SimpleParser) DumpAST(node ast.ASTNode, indent string) {
	fmt.Println(indent + node.GetType().String() + " " + node.GetText())
	for _, child := range node.GetChildren() {
		p.DumpAST(child, indent+"\t")
	}
}

func (p *SimpleParser) prog(tokens token.TokenReader) (*ast.SimpleASTNode, error) {
	return nil, nil
}

func (p *SimpleParser) expStatement(tokens token.TokenReader) (*ast.SimpleASTNode, error) {

	return nil, nil
}

func (p *SimpleParser) assignmentStatement(tokens token.TokenReader) (*ast.SimpleASTNode, error) {

	return nil, nil
}

func (p *SimpleParser) intDeclare(tokens token.TokenReader) (*ast.SimpleASTNode, error) {

	return nil, nil
}

func (p *SimpleParser) additive(tokens token.TokenReader) (*ast.SimpleASTNode, error) {

	return nil, nil
}

func (p *SimpleParser) multiplicative(tokens token.TokenReader) (*ast.SimpleASTNode, error) {

	return nil, nil
}

func (p *SimpleParser) primary(tokens token.TokenReader) (*ast.SimpleASTNode, error) {

	return nil, nil
}
