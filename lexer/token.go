package lexer

import (
	"fmt"
)

type TokenType int
type TokenValue int

type Token struct {
	Type  TokenType
	Value TokenValue
	// The variable name, keyword, literal value...
	Name string
}

const (
	OPERATOR TokenType = iota
	KEYWORD
	SEPARATOR
	LITERAL
	IDENTIFIER
	COMMENT
)

func (t Token) String() string {
	return fmt.Sprintf("(%s,%s)", TypeStr(t.Type), ValueStr(t))
}

func TypeStr(t TokenType) string {
	switch t {
	case OPERATOR:
		return "Operator"
	case KEYWORD:
		return "Keyword"
	case SEPARATOR:
		return "Separator"
	case LITERAL:
		return "Literal"
	case IDENTIFIER:
		return "Identifier"
	case COMMENT:
		return "Comment"
	default:
		panic("TokenType not recognized!")
	}
}

func ValueStr(t Token) string {

	switch t.Type {
	case OPERATOR:
		return OperatorStr(t.Value)
	case KEYWORD:
		return KeywordStr(t.Value)
	case SEPARATOR:
		return SeparatorStr(t.Value)
	case LITERAL:
		return t.Name
	case IDENTIFIER:
		return t.Name
	case COMMENT:
		return t.Name
	default:
		panic("TokenType not recognized!")
	}

}
