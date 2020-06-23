package token

// TokenReader token流
type TokenReader interface {
	// 返回并读取Token流中下一个Token，如果流已经为空，则返回nil
	Read() Token
	// 返回Token流中下一个Token，如果流已经为空，则返回nil
	Peek() Token
	// Token流回退一步
	Unread()
	// 获取Token流当前的读取位置
	GetPosition() int
	// 设置当前读取位置到pos
	SetPosition(pos int)
}
