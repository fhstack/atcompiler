package calculator

import (
	"fmt"
	"testing"
)

func TestCal(t *testing.T) {
	calculatorIns := NewSimpleCalculator()

	script := "2+3*5"
	fmt.Printf("计算 %s 看上去一切正常\n", script)
	calculatorIns.Evaluate(script)
}
