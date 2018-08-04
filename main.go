package main

import (
	"fmt"
	"github.com/PietroCarrara/tanuki/lexer"
)

func main() {

	tokens := lexer.Lex(`
for (var i = 0; i <= 20; i++) {
	print(i)
}
var sla = 20;
var myVar = \"my string\";`)

	fmt.Println(tokens)

	for _, tok := range tokens {
		fmt.Print(tok.Name)
	}
	fmt.Println()

}
