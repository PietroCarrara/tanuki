package lexer

import (
	"math"
	"sort"
	"strings"
)

var keys map[string]Token = map[string]Token{}

var separators []string

func init() {

	sort.Slice(separators, func(i, j int) bool {
		return len(separators[i]) > len(separators[j])
	})

	var i TokenValue
	for i = 0; i < OPERATORS_LENGTH; i++ {
		keys[OperatorStr(i)] = Token{Type: OPERATOR, Value: i}
		separators = append(separators, OperatorStr(i))
	}
	for i = 0; i < SEPARATORS_LENGTH; i++ {
		keys[SeparatorStr(i)] = Token{Type: SEPARATOR, Value: i}
		separators = append(separators, SeparatorStr(i))
	}
	for i = 0; i < KEYWORDS_LENGTH; i++ {
		keys[KeywordStr(i)] = Token{Type: KEYWORD, Value: i}
	}
}

func Lex(code string) []Token {

	units := lex(code)

	tokens := []Token{}

	isOnString := false
	str := ""
	for _, unit := range units {
		token := Token{}

		if isOnString {
			str += unit
			if strings.HasSuffix(unit, "\"") {
				isOnString = false

				token.Type = LITERAL
				token.Name = str
				str = ""

				tokens = append(tokens, token)
			}
			continue
		}

		if val, ok := keys[unit]; ok {
			token = val
			token.Name = unit

			tokens = append(tokens, token)

			continue
		}

		if strings.HasPrefix(unit, "\"") {
			isOnString = true
			str += unit
			continue
		}

		if IsValidLiteral(unit) {
			token.Type = LITERAL
			token.Name = unit

			tokens = append(tokens, token)

			continue
		}

		token.Type = IDENTIFIER
		token.Name = unit

		tokens = append(tokens, token)
	}

	return tokens
}

func lex(code string) []string {

	result := []string{}

	nextIndex, endIndex := nextSeparator(code)

	// if we're done with the separators
	if nextIndex < 0 {
		if len(code) > 0 {
			return []string{code}
		}
		return []string{}
	}

	prev := code[:nextIndex]
	sep := code[nextIndex:endIndex]
	code = code[endIndex:]

	if len(prev) > 0 {
		result = append(result, prev)
	}

	result = append(result, sep)

	return append(result, lex(code)...)
}

// Receives a code string and returns the 0-based
// index of the next separator. (begin, end (both inclusive))
// Returns -1 if there is none
func nextSeparator(code string) (int, int) {

	sepIndex := -1
	smallestIndex := math.MaxInt32
	for i, sep := range separators {
		index := strings.Index(code, sep)

		if index >= 0 && index < smallestIndex {
			sepIndex = i
			smallestIndex = index
		}
	}

	if sepIndex < 0 {
		return -1, 0
	}

	return smallestIndex, smallestIndex + len(separators[sepIndex])

}
