package parser

type Statement interface {
	String() string
}

func parseStatement() Statement {

	// Skip new lines, they don't matter
	for isNewLine(current()) {
		next()
	}

	var res Statement

	if isOpenKey(current()) {
		stmt := composedStatement{}
		next()
		for !isCloseKey(current()) {
			stmt.Body = append(stmt.Body, parseStatement())
			next()
		}

		res = stmt

		next()
	} else if beginsDeclaration(current()) {
		res = parseDeclaration()
	} else if beginsWhile(current()) {
		res = parseWhile()
	} else {
		res = parseExpression()
	}

	return res
}

type composedStatement struct {
	Body []Statement
}

func (c composedStatement) String() string {
	res := ""

	for _, s := range c.Body {
		res += s.String() + "\n"
	}

	return res
}
