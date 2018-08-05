package lexer

const (
	FOR TokenValue = iota
	WHILE
	VAR
	INT
	KEYWORDS_LENGTH
)

func KeywordStr(v TokenValue) string {
	switch v {
	case FOR:
		return "for"
	case WHILE:
		return "while"
	case VAR:
		return "var"
	case INT:
		return "int"
	default:
		panic("Keyword not recognized!")
	}
}
