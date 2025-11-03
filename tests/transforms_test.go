package tests

import (
	"go-reloaded/transforms"
	"testing"
)

// Hex to Dec Tests
func TestHexToDec(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"single digit", "A", "10"},
		{"two digits", "1E", "30"},
		{"max byte", "FF", "255"},
		{"zero", "0", "0"},
		{"lowercase", "ff", "255"},
		{"large number", "FFFF", "65535"},
		{"invalid char", "ZZZ", "ZZZ"},
		{"empty", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := transforms.HexToDec(tt.input)
			if result != tt.expected {
				t.Errorf("HexToDec(%s) = %s; want %s", tt.input, result, tt.expected)
			}
		})
	}
}

// Binary to Dec Tests
func TestBinToDec(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"single bit", "1", "1"},
		{"two bits", "10", "2"},
		{"four bits", "1111", "15"},
		{"zero", "0", "0"},
		{"eight bits", "11111111", "255"},
		{"invalid digit", "102", "102"},
		{"empty", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := transforms.BinToDec(tt.input)
			if result != tt.expected {
				t.Errorf("BinToDec(%s) = %s; want %s", tt.input, result, tt.expected)
			}
		})
	}
}

// Case Transformation Tests
func TestToUpper(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"lowercase", "hello", "HELLO"},
		{"uppercase", "HELLO", "HELLO"},
		{"mixed", "HeLLo", "HELLO"},
		{"empty", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := transforms.ToUpper(tt.input)
			if result != tt.expected {
				t.Errorf("ToUpper(%s) = %s; want %s", tt.input, result, tt.expected)
			}
		})
	}
}

func TestToLower(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"uppercase", "HELLO", "hello"},
		{"lowercase", "hello", "hello"},
		{"mixed", "HeLLo", "hello"},
		{"empty", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := transforms.ToLower(tt.input)
			if result != tt.expected {
				t.Errorf("ToLower(%s) = %s; want %s", tt.input, result, tt.expected)
			}
		})
	}
}

func TestCapitalize(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"lowercase", "hello", "Hello"},
		{"uppercase", "HELLO", "HELLO"},
		{"mixed", "hELLO", "HELLO"},
		{"empty", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := transforms.Capitalize(tt.input)
			if result != tt.expected {
				t.Errorf("Capitalize(%s) = %s; want %s", tt.input, result, tt.expected)
			}
		})
	}
}

// Article Correction Tests
func TestFixArticle(t *testing.T) {
	tests := []struct {
		name     string
		word     string
		nextWord string
		expected string
	}{
		{"a + vowel a", "a", "amazing", "an"},
		{"a + vowel e", "a", "elephant", "an"},
		{"a + vowel i", "a", "incredible", "an"},
		{"a + vowel o", "a", "orange", "an"},
		{"a + vowel u", "a", "umbrella", "an"},
		{"a + h", "a", "honest", "an"},
		{"a + consonant", "a", "book", "a"},
		{"A + vowel", "A", "amazing", "An"},
		{"A + consonant", "A", "book", "A"},
		{"the + vowel", "the", "apple", "the"},
		{"empty next", "a", "", "a"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := transforms.FixArticle(tt.word, tt.nextWord)
			if result != tt.expected {
				t.Errorf("FixArticle(%s, %s) = %s; want %s",
					tt.word, tt.nextWord, result, tt.expected)
			}
		})
	}
}
