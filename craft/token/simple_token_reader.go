package token

// SimpleTokenReader implements SimpleToken interface
type SimpleTokenReader struct {
	tokenList []Token
	curPos    int // 当前读取位置
}

func NewSimpleTokenReader(tokenList []Token) *SimpleTokenReader {
	return &SimpleTokenReader{tokenList: tokenList, curPos: 0}
}

func (r *SimpleTokenReader) Read() Token {
	if r.curPos < len(r.tokenList) {
		defer func() {
			r.curPos++
		}()
		return r.tokenList[r.curPos]
	}
	return nil
}

func (r *SimpleTokenReader) Peek() Token {
	if r.curPos < len(r.tokenList) {
		return r.tokenList[r.curPos]
	}
	return nil
}

func (r *SimpleTokenReader) Unread() {
	if r.curPos > 0 {
		r.curPos--
	}
}

func (r *SimpleTokenReader) GetPosition() int {
	return r.curPos
}

func (r *SimpleTokenReader) SetPosition(pos int) {
	if pos >= 0 && pos < len(r.tokenList) {
		r.curPos = pos
	}
}
