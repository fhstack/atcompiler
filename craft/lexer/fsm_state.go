package lexer

type FSMState int8

const (
	FSMState_Initial FSMState = iota + 1
	FSMState_Plus             // +
	FSMState_Minus            // -
	FSMState_Star             // *
	FSMState_Slash            // /

	FSMState_GT // >
	FSMState_GE // >=
	FSMState_EQ // ==
	FSMState_LT // <
	FSMState_LE // <=

	FSMState_Assignment // =
	FSMState_Semicolon  // ;
	FSMState_LeftParen  // (
	FSMState_RightParen // )

	FSMState_Int1
	FSMState_Int2
	FSMState_Int3
	FSMState_Int

	FSMState_If1
	FSMState_If2
	FSMState_If

	FSMState_Id // identifier

	FSMState_Else1
	FSMState_Else2
	FSMState_Else3
	FSMState_Else4
	FSMState_Else

	FSMState_IntLiteral
)
