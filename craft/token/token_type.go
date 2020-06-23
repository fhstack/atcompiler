package token

type TokenType int8

const (
	Plus  TokenType = iota + 1
	Minus           // -
	Star            //  *
	Slash           // /

	GE // >=
	GT // >
	EQ // ==
	LE // <=
	LT // <

	Semicolon  // ;
	LeftParen  // (
	RightParen // )
	Assignment // =

	If   // if
	Else // else

	Int

	Indentifier // 标识符

	IntLiteral    // 整型字面量
	StringLiteral // 字符串字面量
)
