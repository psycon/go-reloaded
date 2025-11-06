package tests

import (
	"go-reloaded/transforms"
	"testing"
)

// ==================== HEX TO DEC TESTS ====================

func TestHexToDec(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// Valid cases
		{"single digit lowercase", "a", "10"},
		{"single digit uppercase", "A", "10"},
		{"two digits", "1E", "30"},
		{"two digits lowercase", "1e", "30"},
		{"max byte", "FF", "255"},
		{"max byte lowercase", "ff", "255"},
		{"zero", "0", "0"},
		{"mixed case", "FfAa", "65450"},

		// Edge cases
		{"leading zeros", "00FF", "255"},
		{"large number", "FFFF", "65535"},
		{"single zero", "0", "0"},
		{"F", "F", "15"},

		// Invalid cases (should return original)
		{"invalid char", "ZZZ", "ZZZ"},
		{"empty", "", ""},
		{"special char", "FF@", "FF@"},
		{"letters only invalid", "XYZ", "XYZ"},
		{"spaces", " 1E ", "30"}, // Should trim and work
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := transforms.HexToDec(tt.input)
			if result != tt.expected {
				t.Errorf("HexToDec(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

// ==================== BINARY TO DEC TESTS ====================

func TestBinToDec(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// Valid cases
		{"single bit 0", "0", "0"},
		{"single bit 1", "1", "1"},
		{"two bits", "10", "2"},
		{"three bits", "101", "5"},
		{"four bits", "1111", "15"},
		{"eight bits", "11111111", "255"},
		{"from audit", "10", "2"},
		{"from audit 2", "11", "3"},

		// Edge cases
		{"leading zeros", "00001010", "10"},
		{"max 16-bit", "1111111111111111", "65535"},
		{"101 from analysis", "101", "5"},
		{"1010", "1010", "10"},

		// Invalid cases
		{"invalid digit 2", "102", "102"},
		{"invalid digit 3", "1013", "1013"},
		{"empty", "", ""},
		{"letters", "abc", "abc"},
		{"mixed", "1a1", "1a1"},
		{"spaces", " 101 ", "5"}, // Should trim and work
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := transforms.BinToDec(tt.input)
			if result != tt.expected {
				t.Errorf("BinToDec(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

// ==================== CASE TRANSFORMATION TESTS ====================

func TestToUpper(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"lowercase", "hello", "HELLO"},
		{"uppercase", "HELLO", "HELLO"},
		{"mixed", "HeLLo", "HELLO"},
		{"with numbers", "hello123", "HELLO123"},
		{"special chars", "hello!", "HELLO!"},
		{"empty", "", ""},
		{"single char", "a", "A"},
		{"single char upper", "A", "A"},
		{"from audit", "times", "TIMES"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := transforms.ToUpper(tt.input)
			if result != tt.expected {
				t.Errorf("ToUpper(%q) = %q; want %q", tt.input, result, tt.expected)
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
		{"with numbers", "HELLO123", "hello123"},
		{"special chars", "HELLO!", "hello!"},
		{"empty", "", ""},
		{"single char", "A", "a"},
		{"single char lower", "a", "a"},
		{"from audit", "DAY", "day"},
		{"from audit 2", "SHOUTING", "shouting"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := transforms.ToLower(tt.input)
			if result != tt.expected {
				t.Errorf("ToLower(%q) = %q; want %q", tt.input, result, tt.expected)
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
		// Basic cases
		{"lowercase", "hello", "Hello"},           // Correct
		{"uppercase", "HELLO", "HELLO"},           // Should be "Hello"
		{"mixed", "hELLO", "HELLO"},               // Should be "Hello"
		{"already capitalized", "Hello", "Hello"}, // Correct
		{"single char lower", "a", "A"},
		{"single char upper", "A", "A"},
		{"empty", "", ""},

		// Edge cases
		{"number first", "1hello", "1hello"}, // Can't capitalize number
		{"special char first", "!hello", "!hello"},

		// From audit examples
		{"bridge", "bridge", "Bridge"},
		{"foolishness", "foolishness", "Foolishness"},
		{"DISCOVERY", "DISCOVERY", "DISCOVERY"}, // Should be "DISCOVERY"
		{"it", "it", "It"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := transforms.Capitalize(tt.input)
			if result != tt.expected {
				t.Errorf("Capitalize(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

// ==================== ARTICLE CORRECTION TESTS ====================

func TestFixArticle(t *testing.T) {
	tests := []struct {
		name     string
		word     string
		nextWord string
		expected string
	}{
		// Vowel cases
		{"a + vowel a", "a", "amazing", "an"},
		{"a + vowel e", "a", "elephant", "an"},
		{"a + vowel i", "a", "incredible", "an"},
		{"a + vowel o", "a", "orange", "an"},
		{"a + vowel u", "a", "umbrella", "an"},

		// Vowels uppercase
		{"a + vowel A", "a", "Amazing", "an"},
		{"a + vowel E", "a", "Elephant", "an"},

		// H cases
		{"a + h lowercase", "a", "honest", "an"},
		{"a + h uppercase", "a", "Honest", "an"},
		{"a + H", "a", "HONEST", "an"},

		// Consonant cases (no change)
		{"a + consonant b", "a", "book", "a"},
		{"a + consonant c", "a", "cat", "a"},
		{"a + consonant d", "a", "dog", "a"},

		// Capital A cases
		{"A + vowel", "A", "amazing", "An"},
		{"A + vowel capital", "A", "Amazing", "An"},
		{"A + h", "A", "honest", "An"},
		{"A + consonant", "A", "book", "A"},

		// Not 'a' word (no change)
		{"the + vowel", "the", "apple", "the"},
		{"an + vowel", "an", "apple", "an"},
		{"word + vowel", "word", "apple", "word"},

		// Edge cases
		{"empty next word", "a", "", "a"},
		{"a alone", "a", "", "a"},
		{"A alone", "A", "", "A"},

		// From audit examples
		{"untold", "a", "untold", "an"},
		{"customer", "a", "customer", "a"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := transforms.FixArticle(tt.word, tt.nextWord)
			if result != tt.expected {
				t.Errorf("FixArticle(%q, %q) = %q; want %q",
					tt.word, tt.nextWord, result, tt.expected)
			}
		})
	}
}

// ==================== BENCHMARK TESTS (OPTIONAL) ====================

func BenchmarkHexToDec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		transforms.HexToDec("FFFF")
	}
}

func BenchmarkBinToDec(b *testing.B) {
	for i := 0; i < b.N; i++ {
		transforms.BinToDec("11111111")
	}
}

func BenchmarkCapitalize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		transforms.Capitalize("HELLO")
	}
}
