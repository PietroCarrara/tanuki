package lexer

import (
	"strconv"
	"strings"
)

// Checks if a given string is a literal
func IsValidLiteral(s string) bool {

	// If is number
	if _, err := strconv.ParseFloat(s, 64); err == nil {
		return true
	}

	// If is string
	if strings.HasSuffix(s, "\"") && strings.HasPrefix(s, "\"") {
		return true
	}

	// If is char
	if strings.HasSuffix(s, "'") && strings.HasPrefix(s, "'") {
		return true
	}

	return false
}
