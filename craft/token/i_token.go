package token

// Token 词法分析的基本单元
type Token interface {
	// Token类型
	GetType() TokenType
	// Token的文本值
	GetText() string
}
