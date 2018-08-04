package lexer

const (
	ADD TokenValue = iota
	SUB
	MUL
	DIV
	AND
	OR
	ASSIGN
	DOT
	SML
	GTR
	SML_EQ
	GTR_EQ
	EQUAL
	INCR
	DECR
	// How many operators we have
	OPERATORS_LENGTH
)

func OperatorStr(v TokenValue) string {
	switch v {
	case ADD:
		return "+"
	case SUB:
		return "-"
	case MUL:
		return "*"
	case DIV:
		return "/"
	case AND:
		return "&&"
	case OR:
		return "||"
	case ASSIGN:
		return "="
	case DOT:
		return "."
	case SML:
		return "<"
	case GTR:
		return ">"
	case SML_EQ:
		return "<="
	case GTR_EQ:
		return ">="
	case EQUAL:
		return "=="
	case INCR:
		return "++"
	case DECR:
		return "--"
	default:
		panic("Operator not recognized!")
	}
}
