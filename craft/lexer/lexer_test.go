package lexer

import (
	"testing"
)

func TestLexer(t *testing.T) {
	sl := NewSimpleLexer()

	// r := sl.Tokenize([]byte("int a = 123;"))
	// dumpToken(r)
	// fmt.Println("**********************************")
	// r = sl.Tokenize([]byte("inta a = 123;"))
	// dumpToken(r)
	// fmt.Println("**********************************")
	// r = sl.Tokenize([]byte("in a = 123;"))
	// dumpToken(r)
	// fmt.Println("**********************************")
	// r = sl.Tokenize([]byte("a >= 123;"))
	// dumpToken(r)
	// fmt.Println("**********************************")
	// r = sl.Tokenize([]byte("a > 123;"))
	// dumpToken(r)
	r := sl.Tokenize([]byte("1+2*3"))
	dumpToken(r)
}
