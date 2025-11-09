package formatters

import "strings"

// FormatQuote formats words between quotes
// Single word: 'awesome'
// Multiple words: 'I am great'
func FormatQuote(words []string) string {
	if len(words) == 0 {
		return "''"
	}

	if len(words) == 1 {
		return "'" + words[0] + "'"
	}

	// Multiple words: join with spaces and wrap
	return "'" + strings.Join(words, " ") + "'"
}

// FormatDoubleQuote formats words between double quotes
func FormatDoubleQuote(words []string) string {
	if len(words) == 0 {
		return "\"\""
	}

	if len(words) == 1 {
		return "\"" + words[0] + "\""
	}

	return "\"" + strings.Join(words, " ") + "\""
}
