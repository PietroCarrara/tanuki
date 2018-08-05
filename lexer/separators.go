package lexer

const (
	O_PAR TokenValue = iota
	C_PAR
	O_KEY
	C_KEY
	NLINE
	SPACE
	SEMCOL
	SLASHES
	SEPARATORS_LENGTH
)

func SeparatorStr(v TokenValue) string {
	switch v {
	case O_PAR:
		return "("
	case C_PAR:
		return ")"
	case O_KEY:
		return "{"
	case C_KEY:
		return "}"
	case NLINE:
		return "\n"
	case SPACE:
		return " "
	case SEMCOL:
		return ";"
	case SLASHES:
		return "//"
	default:
		panic("Separator not recognized!")
	}
}

func (v TokenValue) ReverseSeparator() TokenValue {
	switch v {
	case O_PAR:
		return C_PAR
	case C_PAR:
		return O_PAR
	case O_KEY:
		return C_KEY
	case C_KEY:
		return O_KEY
	}

	return v
}
