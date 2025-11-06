package tests

import (
	"go-reloaded/fsm"
	"os"
	"testing"
)

// TEST 12: Comprehensive Paragraph (THE GOLDEN TEST)
func TestGolden_ComprehensiveParagraph(t *testing.T) {
	input := `it (cap) was a amazing DAY (low) ! the sun was shining and the temperature reached 1F (hex) degrees . I went to the store , bought 11 (bin) apples and A (up) orange . the shopkeeper said : ' you are a honest customer ' . when i got HOME (low, 2) , i realized that 101 (bin) plus A (hex) equals F (hex) ! what a DISCOVERY (cap) ... i could not BELIEVE IT (low, 2) ! ? this was the best day EVER (cap, 2) .`

	expected := `It was an amazing day! the sun was shining and the temperature reached 31 degrees. I went to the store, bought 3 apples and An orange. the shopkeeper said: 'you are an honest customer'. when i got home, i realized that 5 plus 10 equals 15! what a DISCOVERY... i could not believe it!? this was the best Day EVER.`

	processor := fsm.NewProcessor()
	result := processor.Process(input)

	if result != expected {
		t.Errorf(`
=== GOLDEN TEST FAILED ===`)
		t.Errorf(`
Input:
%s`, input)
		t.Errorf(`
Expected:
%s`, expected)
		t.Errorf(`
Got:
%s`, result)
		t.Errorf(`
=========================`)

		// Show detailed diff
		t.Log(`
Detailed comparison:`)
		minLen := len(expected)
		if len(result) < minLen {
			minLen = len(result)
		}

		for i := 0; i < minLen; i++ {
			if expected[i] != result[i] {
				t.Logf("First difference at position %d:", i)
				t.Logf("Expected: %q", expected[i:min(i+20, len(expected))])
				t.Logf("Got:      %q", result[i:min(i+20, len(result))])
				break
			}
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Additional Complex Scenarios

func TestGolden_TechnicalDocument(t *testing.T) {
	input := `The system uses A (cap) advanced FSM (up) architecture . It processes FF (hex) tokens per second and handles 101010 (bin) concurrent requests . Performance metrics are ... impressive ! the throughput increased BY (up, 1) A (cap) FACTOR (low) of 10 (bin) .`

	expected := `The system uses An advanced FSM architecture. It processes 255 tokens per second and handles 42 concurrent requests. Performance metrics are... impressive! the throughput increased BY A factor of 2.`

	processor := fsm.NewProcessor()
	result := processor.Process(input)

	if result != expected {
		t.Errorf(`
Technical Document Test Failed`)
		t.Errorf(`Expected: %s`, expected)
		t.Errorf(`Got:      %s`, result)
	}
}

func TestGolden_CreativeWriting(t *testing.T) {
	input := `' once upon a time (cap, 4) ' , there lived a amazing DRAGON (low) in A (up) ENCHANTED (cap) forest . The dragon loved to count in hexadecimal : 1A (hex) , 2B (hex) , 3C (hex) ... what a PECULIAR (cap) creature ! ? Indeed , it was a EXTRAORDINARY (low) tale .`

	expected := `'Once Upon A Time', there lived an amazing dragon in An ENCHANTED forest. The dragon loved to count in hexadecimal: 26, 43, 60... what a PECULIAR creature!? Indeed, it was an extraordinary tale.`

	processor := fsm.NewProcessor()
	result := processor.Process(input)

	if result != expected {
		t.Errorf(`
Creative Writing Test Failed`)
		t.Errorf(`Expected: %s`, expected)
		t.Errorf(`Got:      %s`, result)
	}
}

// File-Based Golden Test

func TestGolden_FileProcessing(t *testing.T) {
	// Create temporary test file
	input := `it (cap) was the best of times (up, 4) , it was the worst of times . It was A (cap) TALE (low) OF (low) TWO (low) CITIES (low) .`

	inputFile := "test_golden_input.txt"
	outputFile := "test_golden_output.txt"
	expectedOutput := `It was THE BEST OF TIMES, it was the worst of times. It was A tale of two cities.`

	// Write input file
	err := os.WriteFile(inputFile, []byte(input), 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(inputFile)

	// Process
	processor := fsm.NewProcessor()
	result := processor.Process(input)

	// Write output
	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(outputFile)

	// Verify
	if result != expectedOutput {
		t.Errorf(`
File Processing Test Failed`)
		t.Errorf(`Expected: %s`, expectedOutput)
		t.Errorf(`Got:      %s`, result)
	}
}

// Regression Test (Catch Previous Bugs)

func TestGolden_RegressionSuite(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "double modifier application",
			input:    "word (up) (low)", // Second modifier should be treated as text
			expected: "WORD (low)",
		},
		{
			name:     "quote with punctuation inside",
			input:    "He said ' hello ! ' to everyone",
			expected: "He said 'hello!' to everyone",
		},
		{
			name:     "multiple a/an in sequence",
			input:    "a apple a orange a banana",
			expected: "an apple an orange a banana",
		},
		{
			name:     "batch modifier exceeds buffer",
			input:    "word (up, 10)",
			expected: "WORD", // Should apply to all available words
		},
		{
			name:     "zero in hex and bin",
			input:    "The value 0 (hex) equals 0 (bin)",
			expected: "The value 0 equals 0",
		},
	}

	processor := fsm.NewProcessor()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := processor.Process(tt.input)
			if result != tt.expected {
				t.Errorf(`
Regression: %s`, tt.name)
				t.Errorf(`Expected: %s`, tt.expected)
				t.Errorf(`Got:      %s`, result)
			}
		})
	}
}
