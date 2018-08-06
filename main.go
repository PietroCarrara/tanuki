package main

import (
	"fmt"
	"github.com/PietroCarrara/tanuki/lexer"
	"github.com/PietroCarrara/tanuki/parser"
)

func main() {

	tokens := lexer.Lex(`int i = sum(1+2, 3+4, 5)
while (i < 100) {
	print(i)
	i = i + 1
}
`)

	program := parser.Parse(tokens)

	fmt.Println(program)
}
