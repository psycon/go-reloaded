package tests

import (
	"go-reloaded/formatters"
	"testing"
)

// Punctuation Tests
func TestFormatPunctuation(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"single period", ".", "."},
		{"single comma", ",", ","},
		{"single exclamation", "!", "!"},
		{"single question", "?", "?"},
		{"colon", ":", ":"},
		{"semicolon", ";", ";"},
		{"ellipsis", "...", "..."},
		{"question exclaim", "!?", "!?"},
		{"exclaim question", "?!", "?!"},
		{"multiple exclaim", "!!!", "!!!"},
		{"empty", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatters.FormatPunctuation(tt.input)
			if result != tt.expected {
				t.Errorf("FormatPunctuation(%s) = %s; want %s",
					tt.input, result, tt.expected)
			}
		})
	}
}

// Quote Tests
func TestFormatQuote(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{"single word", []string{"awesome"}, "'awesome'"},
		{"single word caps", []string{"HELLO"}, "'HELLO'"},
		{"two words", []string{"hello", "world"}, "'hello world'"},
		{"three words", []string{"I", "am", "great"}, "'I am great'"},
		{"many words", []string{"This", "is", "a", "test"}, "'This is a test'"},
		{"empty slice", []string{}, "''"},
		{"word with punct", []string{"hello!"}, "'hello!'"},
		{"multiple with punct", []string{"hello", "world!"}, "'hello world!'"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatters.FormatQuote(tt.input)
			if result != tt.expected {
				t.Errorf("FormatQuote(%v) = %s; want %s",
					tt.input, result, tt.expected)
			}
		})
	}
}
