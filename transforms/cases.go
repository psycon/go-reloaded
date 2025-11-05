package transforms

import (
	"strings"
	"unicode"
)

// ToUpper converts string to uppercase
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// ToLower converts string to lowercase
func ToLower(s string) string {
	return strings.ToLower(s)
}

// Capitalize capitalizes first letter and lowercases the rest
func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}

	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])

	return string(runes)
}
