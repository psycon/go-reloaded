package tests

import (
	"go-reloaded/formatters"
	"testing"
)

// ==================== PUNCTUATION TESTS ====================

func TestFormatPunctuation(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// Single punctuation marks
		{"single period", ".", "."},
		{"single comma", ",", ","},
		{"single exclamation", "!", "!"},
		{"single question", "?", "?"},
		{"colon", ":", ":"},
		{"semicolon", ";", ";"},

		// Punctuation groups
		{"ellipsis", "...", "..."},
		{"question exclaim", "!?", "!?"},
		{"exclaim question", "?!", "?!"},
		{"multiple exclaim", "!!!", "!!!"},
		{"double question", "??", "??"},

		// Edge cases
		{"empty", "", ""},
		{"complex group", "...!?", "...!?"},

		// Note: FormatPunctuation just returns the string as-is
		// The FSM handles spacing logic
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatters.FormatPunctuation(tt.input)
			if result != tt.expected {
				t.Errorf("FormatPunctuation(%q) = %q; want %q",
					tt.input, result, tt.expected)
			}
		})
	}
}

// ==================== QUOTE TESTS ====================

func TestFormatQuote(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		// Single word quotes
		{"single word", []string{"awesome"}, "'awesome'"},
		{"single word caps", []string{"HELLO"}, "'HELLO'"},
		{"single word lower", []string{"hello"}, "'hello'"},

		// Multiple word quotes
		{"two words", []string{"hello", "world"}, "'hello world'"},
		{"three words", []string{"I", "am", "great"}, "'I am great'"},
		{"four words", []string{"This", "is", "a", "test"}, "'This is a test'"},
		{"many words", []string{"I", "am", "the", "most", "well-known"}, "'I am the most well-known'"},

		// From audit examples
		{"awesome from audit", []string{"awesome"}, "'awesome'"},
		{"long quote from audit", []string{"I", "am", "the", "most", "well-known", "homosexual", "in", "the", "world"},
			"'I am the most well-known homosexual in the world'"},

		// Edge cases
		{"empty slice", []string{}, "''"},
		{"word with punctuation", []string{"hello!"}, "'hello!'"},
		{"multiple with punctuation", []string{"hello", "world!"}, "'hello world!'"},
		{"word with comma", []string{"hello,"}, "'hello,'"},

		// Special cases
		{"numbers", []string{"123"}, "'123'"},
		{"mixed", []string{"hello", "123", "world"}, "'hello 123 world'"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatters.FormatQuote(tt.input)
			if result != tt.expected {
				t.Errorf("FormatQuote(%v) = %q; want %q",
					tt.input, result, tt.expected)
			}
		})
	}
}

// ==================== EDGE CASE TESTS ====================

func TestFormatQuote_EmptyStrings(t *testing.T) {
	// Test with slice containing empty strings
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{"single empty string", []string{""}, "''"},
		{"two empty strings", []string{"", ""}, "' '"},
		{"empty between words", []string{"hello", "", "world"}, "'hello  world'"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatters.FormatQuote(tt.input)
			if result != tt.expected {
				t.Errorf("FormatQuote(%v) = %q; want %q",
					tt.input, result, tt.expected)
			}
		})
	}
}

// ==================== BENCHMARK TESTS ====================

func BenchmarkFormatQuoteSingle(b *testing.B) {
	words := []string{"awesome"}
	for i := 0; i < b.N; i++ {
		formatters.FormatQuote(words)
	}
}

func BenchmarkFormatQuoteMultiple(b *testing.B) {
	words := []string{"I", "am", "the", "most", "well-known", "person"}
	for i := 0; i < b.N; i++ {
		formatters.FormatQuote(words)
	}
}
