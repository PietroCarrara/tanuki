package parser

import (
	"fmt"
	"github.com/PietroCarrara/tanuki/lexer"
)

type Declaration struct {
	variables []variable
	Type      lexer.Token
}

type variable struct {
	identifier lexer.Token

	initialValue *Expression
}

func (d Declaration) String() string {
	res := fmt.Sprint(d.Type.Name + " ")

	v := d.variables[0]

	res += v.identifier.Name
	if init := v.initialValue; init != nil {
		res += fmt.Sprint("=", *init)
	}

	for _, v = range d.variables[1:] {
		res += ", " + v.identifier.Name
		if init := v.initialValue; init != nil {
			res += fmt.Sprint("=", *init)
		}
	}

	return res
}

func parseDeclaration() Declaration {

	res := Declaration{}

	res.Type = current()

	res.variables = append(res.variables, newVar(next()))
	if isAssign(next()) {
		next()
		tmp := parseExpression()
		res.variables[0].initialValue = &tmp
	}

	index := 1
	for isComma(current()) {
		res.variables = append(res.variables, newVar(next()))
		if isAssign(next()) {
			next()
			tmp := parseExpression()
			res.variables[index].initialValue = &tmp
		}
		index++
	}

	return res
}

func beginsDeclaration(t lexer.Token) bool {
	return t.Type == lexer.KEYWORD && (t.Value == lexer.INT)
}

func newVar(t lexer.Token) variable {
	fmt.Println("New variable in", t)
	return variable{identifier: t}
}

func (v variable) String() string {
	return fmt.Sprintf("%s = %v", v.identifier.Name, *v.initialValue)
}
