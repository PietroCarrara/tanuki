package parser

import (
	"fmt"
	"github.com/PietroCarrara/tanuki/lexer"
)

type While struct {
	Condition Expression
	Body      Statement
}

func (w While) String() string {
	return fmt.Sprint("while ", w.Condition, " {\n ", w.Body, "}")
}

func parseWhile() While {

	res := While{}

	next()

	res.Condition = parseExpression()

	res.Body = parseStatement()

	return res
}

func beginsWhile(t lexer.Token) bool {
	return isWhile(t)
}
