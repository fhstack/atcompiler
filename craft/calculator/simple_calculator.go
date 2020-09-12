package calculator

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/l-f-h/atcompiler/craft/ast"
	"github.com/l-f-h/atcompiler/craft/lexer"
	"github.com/l-f-h/atcompiler/craft/token"
)

// SimpleCalculator 一个计算器，其语法规则如下:
// add -> mul (+mul)*
// mul -> pri | pri * mul
// 递归项在右边，因此是右结合性，而我们需要左结合，故结合性有一定问题
type SimpleCalculator struct{}

func NewSimpleCalculator() *SimpleCalculator {
	return &SimpleCalculator{}
}

// Parse 解析脚本 并返回根节点
func (c *SimpleCalculator) parse(code string) (ast.ASTNode, error) {
	lexerIns := lexer.NewSimpleLexer()
	tokens := lexerIns.Tokenize([]byte(code))

	if rootNode, err := c.program(tokens); err != nil {
		return nil, err
	} else {
		return rootNode, nil
	}
}

// Evaluate 执行脚本，并打印输出AST和求值过程
func (c *SimpleCalculator) Evaluate(script string) {
	tree, err := c.parse(script)
	if err != nil {
		fmt.Printf("语法解析错误: %v\n", err)
		return
	}

	fmt.Println("语法树")
	dumpAST(tree, "")
	fmt.Println("计算过程")
	c.evaluate(tree, "")
}

// evaluate 对某个ast节点求值，并打印出过程
func (c *SimpleCalculator) evaluate(node ast.ASTNode, indent string) int64 {
	var res int64
	fmt.Printf("%sCalculating: %s\n", indent, node.GetType())
	switch node.GetType() {
	case ast.ASTNodeType_Program:
		// 实际上program节点永远只有一个子节点
		for _, childNode := range node.GetChildren() {
			res = c.evaluate(childNode, indent+"\t")
			break
		}
	case ast.ASTNodeType_Additive:
		if len(node.GetChildren()) != 2 {
			panic(errors.New("additive expecting two child node"))
		}

		child1 := node.GetChildren()[0]
		value1 := c.evaluate(child1, indent+"\t")
		child2 := node.GetChildren()[1]
		value2 := c.evaluate(child2, indent+"\t")

		switch node.GetText() {
		case "+":
			return value1 + value2
		case "-":
			return value1 - value2
		default:
			panic(errors.New("expecting + or -"))
		}

	case ast.ASTNodeType_Multiplicative:
		if len(node.GetChildren()) != 2 {
			panic(errors.New("multiplicative expecting two child node"))
		}

		child1 := node.GetChildren()[0]
		value1 := c.evaluate(child1, indent+"\t")
		child2 := node.GetChildren()[1]
		value2 := c.evaluate(child2, indent+"\t")

		switch node.GetText() {
		case "*":
			return value1 * value2
		case "/":
			return value1 / value2
		default:
			panic(errors.New("expecting * or /"))
		}

	case ast.ASTNodeType_IntLiteral:
		if val, err := strconv.ParseInt(node.GetText(), 10, 64); err != nil {
			panic(err)
		} else {
			return val
		}

	default:
	}
	fmt.Printf("%sResult: %d\n", indent, res)
	return res
}

// program 语法解析：根结点开始
func (c *SimpleCalculator) program(tokens token.TokenReader) (*ast.SimpleASTNode, error) {
	root := ast.NewSimpleASTNode("Calculator", ast.ASTNodeType_Program)

	child, err := c.additive(tokens)
	if err != nil {
		return nil, err
	}

	if child != nil {
		root.AddChild(child)
	}
	return root, nil
}

// intDeclare 整型变量声明语句解析 如 int a; int a = 2 + 3;
func (c *SimpleCalculator) intDeclare(tokens token.TokenReader) (*ast.SimpleASTNode, error) {
	var node *ast.SimpleASTNode
	curToken := tokens.Peek()
	if curToken != nil && curToken.GetType() == token.Int {
		tokens.Read() // 消耗掉int关键字

		curToken = tokens.Peek()
		if curToken != nil && curToken.GetType() == token.Indentifier {
			tokens.Read() // 消耗掉标识符
			node = ast.NewSimpleASTNode(curToken.GetText(), ast.ASTNodeType_IntDeclaration)

			curToken = tokens.Peek()
			// 该分支可有可无
			if curToken != nil && curToken.GetType() == token.Assignment {
				tokens.Read() // 消耗掉等号
				if child, err := c.additive(tokens); err != nil {
					return nil, err
				} else if child != nil {
					node.AddChild(child)
				} else {
					return nil, errors.New("intDeclare expecting expressing")
				}

			}
		} else {
			return nil, errors.New("intDeclare expecting Int key word")
		}
	}

	if node != nil {
		curToken = tokens.Peek()
		if curToken != nil && curToken.GetType() == token.Semicolon {
			tokens.Read()
		} else {
			return nil, errors.New("intDeclare expecting ';' in the end")
		}
	}

	return node, nil
}

// additive 语法解析：加法表达式  add -> mul (+mul)*
func (c *SimpleCalculator) additive(tokens token.TokenReader) (*ast.SimpleASTNode, error) {
	child1, err := c.multiplicative(tokens)
	if err != nil {
		return nil, err
	}
	node := child1

	if child1 != nil {
		for {
			nextToken := tokens.Peek()
			if nextToken == nil {
				break
			}
			if nextType := nextToken.GetType(); nextType == token.Plus || nextType == token.Minus {
				curToken := tokens.Read()
				child2, err := c.multiplicative(tokens)
				if err != nil {
					return nil, errors.New("SimpleCalculator additive parse error")
				}
				node = ast.NewSimpleASTNode(curToken.GetText(), ast.ASTNodeType_Additive)
				node.AddChild(child1)
				node.AddChild(child2)
				child1 = node // 注意新节点之后也是作为其他的子节点
			} else {
				break
			}
		}
	}

	return node, nil
}

// multiplicative 语法解析：乘法表达式 mul -> primary | primary * mul
func (c *SimpleCalculator) multiplicative(tokens token.TokenReader) (*ast.SimpleASTNode, error) {
	child1, err := c.primary(tokens)
	if err != nil {
		return nil, err
	}
	node := child1

	if child1 != nil && tokens.Peek() != nil {
		if nextType := tokens.Peek().GetType(); nextType == token.Star || nextType == token.Slash {
			curToken := tokens.Read()

			child2, err := c.multiplicative(tokens)
			if err != nil {
				return nil, err
			}
			if child2 == nil {
				return nil, errors.New("SimpleCalculator multiplicative parse error")
			}
			node = ast.NewSimpleASTNode(curToken.GetText(), ast.ASTNodeType_Multiplicative)
			node.AddChild(child1)
			node.AddChild(child2)
		}
	}

	return node, nil
}

// primary 语法解析：基础表达式
func (c *SimpleCalculator) primary(tokens token.TokenReader) (*ast.SimpleASTNode, error) {
	curToken := tokens.Peek()
	if curToken == nil {
		return nil, nil
	}
	// primary节点简化一下 因为primary永远只有一个 因此直接返回子节点 不构造单独的primary节点
	if curToken.GetType() == token.IntLiteral {
		tokens.Read()
		return ast.NewSimpleASTNode(curToken.GetText(), ast.ASTNodeType_IntLiteral), nil
	} else if curToken.GetType() == token.Indentifier {
		tokens.Read()
		return ast.NewSimpleASTNode(curToken.GetText(), ast.ASTNodeType_Identifier), nil
	} else if curToken.GetType() == token.LeftParen {

	} else {
		// not support
		return nil, fmt.Errorf("SimpleCalculator primary node not support the token: %v", curToken)
	}

	return nil, nil
}

func dumpAST(node ast.ASTNode, indent string) {
	fmt.Println(indent + node.GetType().String() + " " + node.GetText())
	for _, child := range node.GetChildren() {
		dumpAST(child, indent+"\t")
	}
}
