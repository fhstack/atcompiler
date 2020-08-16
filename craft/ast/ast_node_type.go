package ast

type ASTNodeType int8

const (
	// 根节点 程序入口
	ASTNodeType_Program ASTNodeType = iota

	ASTNodeType_IntDeclaration // 整型变量声明
	ASTNodeType_ExpressionStmt // 表达式语句 即表达式后面跟个分号
	ASTNodeType_AssignmentStmt // 赋值语句

	ASTNodeType_Primary        // 基础表达式
	ASTNodeType_Multiplicative // 乘法表达式
	ASTNodeType_Additive       // 加法表达式

	ASTNodeType_Identifier // 标识符
	ASTNodeType_IntLiteral // 整型字面量
)

func (t ASTNodeType) String() string {
	switch t {
	case ASTNodeType_Program:
		return "Program"
	case ASTNodeType_IntDeclaration:
		return "IntDeclaration"
	case ASTNodeType_ExpressionStmt:
		return "ExpressionStmt"
	case ASTNodeType_AssignmentStmt:
		return "AssignmentStmt"
	case ASTNodeType_Primary:
		return "Primary"
	case ASTNodeType_Multiplicative:
		return "Multiplicative"
	case ASTNodeType_Additive:
		return "Additive"
	case ASTNodeType_Identifier:
		return "Identifier"
	case ASTNodeType_IntLiteral:
		return "IntLiteral"
	default:
		return "UNKNOWN"
	}
}
