package lexer

const (
	FOR TokenValue = iota
	VAR
	KEYWORDS_LENGTH
)

func KeywordStr(v TokenValue) string {
	switch v {
	case FOR:
		return "for"
	case VAR:
		return "var"
	default:
		panic("Keyword not recognized!")
	}
}
