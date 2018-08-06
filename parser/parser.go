package parser

import (
	"github.com/PietroCarrara/tanuki/lexer"
)

var tokens []lexer.Token
var index = 0

func next() lexer.Token {

	index++

	return current()
}

func Parse(t []lexer.Token) Program {

	tokens = t

	return parseProgram()
}

func current() lexer.Token {
	return tokens[index]
}

func isOpenKey(t lexer.Token) bool {
	return t.Type == lexer.SEPARATOR && t.Value == lexer.O_KEY
}

func isCloseKey(t lexer.Token) bool {
	return t.Type == lexer.SEPARATOR && t.Value == lexer.C_KEY
}

func isOpenPar(t lexer.Token) bool {
	return t.Type == lexer.SEPARATOR && t.Value == lexer.O_PAR
}

func isClosePar(t lexer.Token) bool {
	return t.Type == lexer.SEPARATOR && t.Value == lexer.C_KEY
}

func isNewLine(t lexer.Token) bool {
	return t.Type == lexer.SEPARATOR && t.Value == lexer.NLINE
}

func isEOF(t lexer.Token) bool {
	return t.Type == lexer.SEPARATOR && t.Value == lexer.EOF
}

func isWhile(t lexer.Token) bool {
	return t.Type == lexer.KEYWORD && t.Value == lexer.WHILE
}

func isComma(t lexer.Token) bool {
	return t.Type == lexer.SEPARATOR && t.Value == lexer.COMMA
}

func isAssign(t lexer.Token) bool {
	return t.Type == lexer.OPERATOR && t.Value == lexer.ASSIGN
}

func isLiteral(t lexer.Token) bool {
	return t.Type == lexer.LITERAL
}

func isIdentifier(t lexer.Token) bool {
	return t.Type == lexer.IDENTIFIER
}

func isOperator(t lexer.Token) bool {
	return t.Type == lexer.OPERATOR
}
