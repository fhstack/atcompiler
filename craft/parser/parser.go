package parser

import (
	"errors"

	"github.com/l-f-h/atcompiler/craft/ast"
	"github.com/l-f-h/atcompiler/craft/lexer"
	"github.com/l-f-h/atcompiler/craft/token"
)

// SimpleParser 一个简单的语法解析器
// 支持解析如下语法：
// program -> intDeclare | expressionStatement | assignmentStatement
// intDeclare -> 'int' Id ( = additive ) ';'
// assignmentStatement -> Id = addtive ';'
// expressionStatement -> additive ';'
// additive -> multiplicative ( (+ | -) multiplicative)*
// multiplicative -> primary ( (* | /) primary)*
// primary -> IntLiteral | Id | (additive)
type SimpleParser struct {
}

func NewSimpleParser() *SimpleParser {
	return &SimpleParser{}
}

func (p *SimpleParser) Parse(script string) (ast.ASTNode, error) {
	tokens := lexer.NewSimpleLexer().Tokenize([]byte(script))
	return p.prog(tokens)
}

func (p *SimpleParser) prog(tokens token.TokenReader) (*ast.SimpleASTNode, error) {
	root := ast.NewSimpleASTNode("prog", ast.ASTNodeType_Program)
	for tokens.Peek() != nil {
		if node, err := p.intDeclare(tokens); err != nil {
			return nil, err
		} else if node != nil {
			root.AddChild(node)
			continue
		}

		if node, err := p.expStatement(tokens); err != nil {
			return nil, err
		} else if node != nil {
			root.AddChild(node)
			continue
		}

		if node, err := p.assignmentStatement(tokens); err != nil {
			return nil, err
		} else if node != nil {
			root.AddChild(node)
			continue
		}
	}
	return root, nil
}

// 表达式语句 如 a + 5 * 3;
// 我们的语法定义里表达式后面必须有分号
func (p *SimpleParser) expStatement(tokens token.TokenReader) (*ast.SimpleASTNode, error) {
	startPos := tokens.GetPosition()
	node, err := p.additive(tokens)
	if err != nil {
		return nil, err
	}

	if nextToken := tokens.Peek(); nextToken != nil && node != nil {
		if nextToken.GetType() != token.Semicolon {
			tokens.SetPosition(startPos) // 说明不是表达式语句则回溯
			return nil, nil
		} else {
			tokens.Read()
		}
	}

	return node, nil
}

// 赋值语句 如 a = 5 + 3 * b;
func (p *SimpleParser) assignmentStatement(tokens token.TokenReader) (*ast.SimpleASTNode, error) {
	idToken := tokens.Peek()
	if idToken.GetType() != token.Identifier {
		return nil, nil
	}
	var node *ast.SimpleASTNode
	tokens.Read() // 消耗掉标识符
	if nextToken := tokens.Peek(); nextToken.GetType() == token.Assignment {
		tokens.Read() // 消耗掉等号
		additiveNode, err := p.additive(tokens)
		if err != nil {
			return nil, err
		}
		node = ast.NewSimpleASTNode(idToken.GetText(), ast.ASTNodeType_AssignmentStmt)
		node.AddChild(additiveNode)
		// 尝试消耗掉分号
		if nextToken := tokens.Read(); nextToken.GetType() != token.Semicolon {
			// 未匹配到则报错
			return nil, errors.New("assignmentStatement expected semicolon")
		}
	} else {
		tokens.Unread() // 回退标识符
	}

	return node, nil
}

// 整型的定义或者声明 int a = 1 + 3 * 5; int b;
func (p *SimpleParser) intDeclare(tokens token.TokenReader) (*ast.SimpleASTNode, error) {
	curToken := tokens.Peek()
	if curToken.GetType() != token.Int {
		return nil, nil
	}
	tokens.Read() // 消耗掉 int

	idToken := tokens.Read() // 消耗掉标识符
	if idToken.GetType() != token.Identifier {
		return nil, errors.New("intDeclare expected identifier")
	}

	if nextToken := tokens.Read(); nextToken.GetType() != token.Assignment { // 消耗掉 =
		return nil, errors.New("intDeclare expected assignment")
	}

	node := ast.NewSimpleASTNode(idToken.GetText(), ast.ASTNodeType_IntDeclaration)
	additiveNode, err := p.additive(tokens)
	if err != nil {
		return nil, err
	}
	if additiveNode != nil {
		node.AddChild(additiveNode)
	}

	// 消耗掉分号
	if nextToken := tokens.Read(); nextToken.GetType() != token.Semicolon {
		return nil, errors.New("intDeclare expected semicolon")
	}

	return node, nil
}

// 加法表达式 1 + 2 * 5 + a
func (p *SimpleParser) additive(tokens token.TokenReader) (*ast.SimpleASTNode, error) {
	node, err := p.multiplicative(tokens)
	if err != nil {
		return nil, err
	}
	if node == nil {
		return nil, nil
	}

	for {
		nextToken := tokens.Peek()
		if nextToken == nil {
			break
		}

		if nextToken.GetType() != token.Plus && nextToken.GetType() != token.Minus {
			break
		}

		tokens.Read()
		child2, err := p.multiplicative(tokens)
		if err != nil {
			return nil, err
		}
		if child2 == nil {
			return nil, errors.New("expected multiplicative after operator")
		}

		child1 := node
		node = ast.NewSimpleASTNode(nextToken.GetText(), ast.ASTNodeType_Additive)
		node.AddChild(child1)
		node.AddChild(child2)
	}

	return node, nil
}

// 乘法表达式
// 5 * 2 * c
// 5
// 2 / 5
func (p *SimpleParser) multiplicative(tokens token.TokenReader) (*ast.SimpleASTNode, error) {
	node, err := p.primary(tokens)
	if err != nil {
		return nil, err
	}
	if node == nil {
		return nil, nil
	}

	for {
		nextToken := tokens.Peek()
		if nextToken == nil {
			break
		}

		if nextToken.GetType() != token.Star && nextToken.GetType() != token.Slash {
			break
		}
		tokens.Read() // 消耗掉运算符
		child2, err := p.primary(tokens)
		if err != nil {
			return nil, err
		}
		if child2 == nil {
			return nil, errors.New("expected primary after operator")
		}

		child1 := node
		node = ast.NewSimpleASTNode(nextToken.GetText(), ast.ASTNodeType_Multiplicative)
		node.AddChild(child1)
		node.AddChild(child2)
	}

	return node, nil
}

func (p *SimpleParser) primary(tokens token.TokenReader) (*ast.SimpleASTNode, error) {
	curToken := tokens.Peek()
	if curToken.GetType() == token.IntLiteral {
		tokens.Read()
		return ast.NewSimpleASTNode(curToken.GetText(), ast.ASTNodeType_IntLiteral), nil
	} else if curToken.GetType() == token.Identifier {
		tokens.Read()
		return ast.NewSimpleASTNode(curToken.GetText(), ast.ASTNodeType_Identifier), nil
	} else if curToken.GetType() == token.LeftParen {
		tokens.Read()
		node, err := p.additive(tokens)
		if err != nil {
			return nil, err
		}
		if nextToken := tokens.Read(); nextToken.GetType() != token.RightParen {
			return nil, errors.New("primary expected right paren ')'")
		}
		return node, nil
	}
	return nil, nil
}
