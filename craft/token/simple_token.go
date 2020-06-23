package token

type SimpleToken struct {
	TokenType TokenType
	Text      string
}

func (t *SimpleToken) GetType() TokenType {
	return t.TokenType
}

func (t *SimpleToken) GetText() string {
	return t.Text
}
