package calculator

import (
	"fmt"
	"testing"

	"github.com/l-f-h/atcompiler/craft/lexer"
)

func TestCal(t *testing.T) {
	calculatorIns := NewSimpleCalculator()

	define := "int a = b + 3;"
	lexerIns := lexer.NewSimpleLexer()
	tokens := lexerIns.Tokenize([]byte(define))
	node, err := calculatorIns.intDeclare(tokens)
	if err != nil {
		panic(err)
	}
	dumpAST(node, "")

	script := "2+3*5"
	fmt.Printf("计算 %s 看上去一切正常\n", script)
	calculatorIns.Evaluate(script)

	script = "2+"
	fmt.Printf("计算 %s 应该有语法错误\n", script)
	calculatorIns.Evaluate(script)

	script = "2+3+4"
	fmt.Printf("计算 %s 结合性会出现错误\n", script)
	calculatorIns.Evaluate(script)
}
