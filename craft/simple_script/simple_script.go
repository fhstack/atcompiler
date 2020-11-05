package main

import (
	"fmt"
	"strconv"

	"github.com/l-f-h/atcompiler/craft/ast"
)

type SimpleScript struct {
	variables map[string]int
	verbose   bool
}

func NewSimpleScript(verbose bool) *SimpleScript {
	return &SimpleScript{
		variables: make(map[string]int),
		verbose:   verbose,
	}
}

func (s *SimpleScript) Evaluate(node ast.ASTNode, indent string) (int, error) {
	res := 0
	switch node.GetType() {
	case ast.ASTNodeType_Program:
		// program 下只会有一个节点
		for _, child := range node.GetChildren() {
			val, err := s.Evaluate(child, indent+"\t")
			if err != nil {
				return 0, err
			}
			res = val
			break
		}
	case ast.ASTNodeType_Additive, ast.ASTNodeType_Multiplicative:
		l, err := s.Evaluate(node.GetChildren()[0], indent+"\t")
		if err != nil {
			return 0, err
		}
		r, err := s.Evaluate(node.GetChildren()[1], indent+"\t")
		if err != nil {
			return 0, err
		}
		if node.GetText() == "+" {
			return l + r, nil
		} else if node.GetText() == "-" {
			return l - r, nil
		} else if node.GetText() == "*" {
			return l * r, nil
		} else if node.GetText() == "/" {
			return l / r, nil
		}
		return 0, fmt.Errorf("illegal add or mul op: %s", node.GetText())
	case ast.ASTNodeType_IntDeclaration, ast.ASTNodeType_AssignmentStmt:
		if node.GetType() == ast.ASTNodeType_AssignmentStmt {
			if _, exist := s.variables[node.GetText()]; !exist {
				return 0, fmt.Errorf("variable %s not declare", node.GetText())
			}
		}

		val, err := s.Evaluate(node.GetChildren()[0], indent+"\t")
		if err != nil {
			return 0, err
		}
		res = val
		s.variables[node.GetText()] = val
	case ast.ASTNodeType_IntLiteral:
		return strconv.Atoi(node.GetText())
	case ast.ASTNodeType_Identifier:
		val, exist := s.variables[node.GetText()]
		if !exist {
			return 0, fmt.Errorf("indentifier variable %s not declare", node.GetText())
		}
		return val, nil
	default:

	}

	if s.verbose {
		fmt.Printf("%sresult: %d\n", indent, res)
	} else {
		// 顶层语句
		if node.GetType() == ast.ASTNodeType_IntDeclaration || node.GetType() == ast.ASTNodeType_AssignmentStmt {
			fmt.Printf("%s result: %d\n", node.GetText(), res)
		} else {
			if node.GetType() != ast.ASTNodeType_Program {
				fmt.Println(res)
			}
		}
	}

	return res, nil
}
