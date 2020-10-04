package parser

import "testing"

func TestParser(t *testing.T) {
	p := NewSimpleParser()
	script := "int age = 5 + 17;age = 20;age + 10 * 2;"
	// script := "2+3*;" // 测试异常语句
	// script := "2+3+;" // 测试异常语句

	node, err := p.Parse(script)
	if err != nil {
		t.Fatal(err)
	}
	DumpAST(node, "")
}
