package lexer

import (
	"bytes"
	"fmt"

	"github.com/l-f-h/atcompiler/craft/token"
)

type SimpleLexer struct {
	tokenText *bytes.Buffer      //  临时保存已经解析的token文本
	tokenList []token.Token      // 解析出的token列表
	token     *token.SimpleToken // 当前正在解析的token
}

func NewSimpleLexer() *SimpleLexer {
	return &SimpleLexer{}
}

func (sl *SimpleLexer) Tokenize(code []byte) token.TokenReader {
	sl.tokenText = bytes.NewBuffer(nil)
	sl.token = &token.SimpleToken{}
	sl.tokenList = make([]token.Token, 0, 10)
	state := FSMState_Initial
	var c byte
	for _, c = range code {
		switch state {
		case FSMState_Initial:
			state = sl.initToken(c)
		case FSMState_Plus, FSMState_Minus, FSMState_Star, FSMState_Slash, FSMState_Semicolon, FSMState_LeftParen,
			FSMState_RightParen, FSMState_Assignment, FSMState_EQ, FSMState_LE, FSMState_GE,
			FSMState_Int:
			state = sl.initToken(c)

		case FSMState_Id:
			if isAlpha(c) || isDigit(c) {
				sl.tokenText.WriteByte(c)
			} else {
				state = sl.initToken(c)
			}

		case FSMState_GT:
			if c == '=' {
				sl.tokenText.WriteByte(c)
				sl.token.TokenType = token.GE
				state = FSMState_GE
			} else {
				state = sl.initToken(c)
			}

		case FSMState_Int1:
			if c == 'n' {
				sl.tokenText.WriteByte(c)
				state = FSMState_Int2
			} else if isDigit(c) || isAlpha(c) {
				sl.tokenText.WriteByte(c)
				state = FSMState_Id
			} else {
				state = sl.initToken(c)
			}

		case FSMState_Int2:
			if c == 't' {
				sl.tokenText.WriteByte(c)
				state = FSMState_Int3
			} else if isDigit(c) || isAlpha(c) {
				sl.tokenText.WriteByte(c)
				state = FSMState_Id
			} else {
				state = sl.initToken(c)
			}

		case FSMState_Int3:
			if isBlank(c) {
				state = FSMState_Int
				sl.token.TokenType = token.Int
			} else {
				sl.tokenText.WriteByte(c)
				state = FSMState_Id
			}
		case FSMState_IntLiteral:
			if isDigit(c) {
				sl.tokenText.WriteByte(c)
			} else {
				state = sl.initToken(c)
			}
		}
	}
	if sl.tokenText.Len() > 0 {
		sl.initToken(c)
	}

	return token.NewSimpleTokenReader(sl.tokenList)
}

func (sl *SimpleLexer) initToken(c byte) FSMState {
	if sl.tokenText.Len() != 0 {
		sl.token.Text = sl.tokenText.String()
		sl.tokenList = append(sl.tokenList, sl.token)

		sl.token = &token.SimpleToken{}
		sl.tokenText.Reset()
	}

	var newState FSMState
	if c == '+' {
		sl.token.TokenType = token.Plus
		newState = FSMState_Plus
		sl.tokenText.WriteByte(c)
	} else if c == '-' {
		sl.token.TokenType = token.Minus
		newState = FSMState_Minus
		sl.tokenText.WriteByte(c)
	} else if c == '*' {
		sl.token.TokenType = token.Star
		newState = FSMState_Star
		sl.tokenText.WriteByte(c)
	} else if c == '/' {
		sl.token.TokenType = token.Slash
		newState = FSMState_Slash
		sl.tokenText.WriteByte(c)
	} else if c == '>' {
		sl.token.TokenType = token.GT
		newState = FSMState_GT
		sl.tokenText.WriteByte(c)
	} else if c == '<' {
		sl.token.TokenType = token.LT
		newState = FSMState_LT
		sl.tokenText.WriteByte(c)
	} else if c == '=' {
		sl.token.TokenType = token.Assignment
		newState = FSMState_Assignment
		sl.tokenText.WriteByte(c)
	} else if c == '(' {
		sl.token.TokenType = token.LeftParen
		newState = FSMState_LeftParen
		sl.tokenText.WriteByte(c)
	} else if c == ')' {
		sl.token.TokenType = token.RightParen
		newState = FSMState_RightParen
		sl.tokenText.WriteByte(c)
	} else if c == ';' {
		sl.token.TokenType = token.Semicolon
		newState = FSMState_Semicolon
		sl.tokenText.WriteByte(c)
	} else if isAlpha(c) {
		sl.token.TokenType = token.Identifier
		if c == 'i' {
			newState = FSMState_Int1
		} else {
			sl.token.TokenType = token.Identifier
			newState = FSMState_Id
		}
		sl.tokenText.WriteByte(c)
	} else if isDigit(c) {
		sl.token.TokenType = token.IntLiteral
		newState = FSMState_IntLiteral
		sl.tokenText.WriteByte(c)
	} else {
		newState = FSMState_Initial
	}

	return newState
}

func dumpToken(tokenReader token.TokenReader) {
	for {
		token := tokenReader.Read()
		if token == nil {
			return
		}
		fmt.Printf("%s\t\t%v\n", token.GetText(), token.GetType())
	}
}
