package main

import (
	"fmt"
	"github.com/PietroCarrara/tanuki/lexer"
)

func main() {

	tokens := lexer.Lex(`
int i = 0
while (i < 20) {
	print(i)
	i = i + 1
}
`)

	fmt.Println(tokens)

	for _, tok := range tokens {
		fmt.Print(tok.Name)
	}
	fmt.Println()
}
