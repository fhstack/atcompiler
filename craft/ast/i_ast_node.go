package ast

type ASTNode interface {
	// 获取父节点
	GetParent() ASTNode

	// 获取所有子节点
	GetChildren() []ASTNode

	// AST类型
	GetType() ASTNodeType

	// 文本值
	GetText() string
}
