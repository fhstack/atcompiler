package ast

// SimpleASTNode implements ASTNode interface
type SimpleASTNode struct {
	parent   *SimpleASTNode
	children []ASTNode
	text     string
	nodeType ASTNodeType
}

func NewSimpleASTNode(text string, t ASTNodeType) *SimpleASTNode {
	return &SimpleASTNode{text: text, nodeType: t}
}

func (n *SimpleASTNode) GetParent() ASTNode {
	return n.parent
}

func (n *SimpleASTNode) GetText() string {
	return n.text
}

func (n *SimpleASTNode) GetChildren() []ASTNode {
	return n.children
}

func (n *SimpleASTNode) GetType() ASTNodeType {
	return n.nodeType
}

func (n *SimpleASTNode) AddChild(c *SimpleASTNode) {
	n.children = append(n.children, c)
	c.parent = n
}
