package parser

import (
	"fmt"
	"github.com/PietroCarrara/tanuki/lexer"
)

type Expression interface {
	String() string
}

func parseExpression() Expression {

	var res Expression

	if isLiteral(current()) {
		res = newLiteralExpression(current())
		next()
	} else if isIdentifier(current()) {
		res = newIdentifierExpression(current())
		id := current()
		if isOpenPar(next()) {
			// Function call
			function := newFunctionCallExpression(id)

			// If the function has args...
			if !isClosePar(next()) {
				function.args = append(function.args, parseExpression())
				for isComma(current()) {
					next()
					function.args = append(function.args, parseExpression())
				}
			}
			res = function
		}
	} else if isOpenPar(current()) {
		// Skip parenthesis
		next()
		res = parseExpression()
		// Skip parenthesis
		next()
	}

	if isOperator(current()) {
		op := newOperationExpression(current())

		op.left = res

		next()
		op.right = parseExpression()

		res = op
	}

	return res
}

type literalExpression struct {
	value lexer.Token
}

func newLiteralExpression(t lexer.Token) literalExpression {
	return literalExpression{value: t}
}

func (l literalExpression) String() string {
	return l.value.Name
}

type identifierExpression struct {
	value lexer.Token
}

func newIdentifierExpression(t lexer.Token) identifierExpression {
	return identifierExpression{value: t}
}

func (i identifierExpression) String() string {
	return i.value.Name
}

type functionCallExpression struct {
	identifier lexer.Token
	args       []Expression
}

func newFunctionCallExpression(t lexer.Token) functionCallExpression {
	return functionCallExpression{identifier: t}
}

func (f functionCallExpression) String() string {

	res := f.identifier.Name + "("

	if len(f.args) > 0 {
		res += f.args[0].String()
		for _, arg := range f.args[1:] {
			res += ", " + arg.String()
		}
	}

	return res + ")"
}

type operationExpression struct {
	left, right Expression
	operator    lexer.Token
}

func newOperationExpression(operator lexer.Token) operationExpression {
	return operationExpression{operator: operator}
}

func (o operationExpression) String() string {
	return fmt.Sprint(o.left, o.operator.Name, o.right)
}
