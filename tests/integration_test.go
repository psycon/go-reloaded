package tests

import (
	"go-reloaded/fsm"
	"strings"
	"testing"
)

// BASIC AUDIT EXAMPLES (from ANALYSIS.md)

func TestAudit1_MixedCaseTransformations(t *testing.T) {
	input := "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6)"
	expected := "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness"

	processor := fsm.NewProcessor()
	result := processor.Process(input)

	if result != expected {
		t.Errorf("Test: Mixed Case Transformations\nInput:    %s\nExpected: %s\nGot:      %s", input, expected, result)
	}
}

func TestAudit2_HexBinConversions(t *testing.T) {
	input := "Simply add 42 (hex) and 10 (bin) and you will see the result is 68."
	expected := "Simply add 66 and 2 and you will see the result is 68."

	processor := fsm.NewProcessor()
	result := processor.Process(input)

	if result != expected {
		t.Errorf("Test: Hex and Binary Conversions\nInput:    %s\nExpected: %s\nGot:      %s", input, expected, result)
	}
}

func TestAudit3_ArticleCorrection(t *testing.T) {
	input := "There is no greater agony than bearing a untold story inside you."
	expected := "There is no greater agony than bearing an untold story inside you."

	processor := fsm.NewProcessor()
	result := processor.Process(input)

	if result != expected {
		t.Errorf("Test: Article Correction\nInput:    %s\nExpected: %s\nGot:      %s", input, expected, result)
	}
}

func TestAudit4_PunctuationSpacing(t *testing.T) {
	input := "Punctuation tests are ... kinda boring ,what do you think ?"
	expected := "Punctuation tests are... kinda boring, what do you think?"

	processor := fsm.NewProcessor()
	result := processor.Process(input)

	if result != expected {
		t.Errorf("Test: Punctuation Spacing\nInput:    %s\nExpected: %s\nGot:      %s", input, expected, result)
	}
}

func TestAudit5_QuoteSingleWord(t *testing.T) {
	input := "I am exactly how they describe me: ' awesome '"
	expected := "I am exactly how they describe me: 'awesome'"

	processor := fsm.NewProcessor()
	result := processor.Process(input)

	if result != expected {
		t.Errorf("Test: Quote Handling (Single Word)\nInput:    %s\nExpected: %s\nGot:      %s", input, expected, result)
	}
}

func TestAudit6_QuoteMultipleWords(t *testing.T) {
	input := "As Elton John said: ' I am the most well-known homosexual in the world '"
	expected := "As Elton John said: 'I am the most well-known homosexual in the world'"

	processor := fsm.NewProcessor()
	result := processor.Process(input)

	if result != expected {
		t.Errorf("Test: Quote Handling (Multiple Words)\nInput:    %s\nExpected: %s\nGot:      %s", input, expected, result)
	}
}

// ADVANCED SCENARIOS

func TestAdvanced7_MultipleModifiers(t *testing.T) {
	input := "the word (up) was then (cap) followed by another (low) transformation and FF (hex) things." // Corrected input
	expected := "the WORD was Then followed by another transformation and 255 things."

	processor := fsm.NewProcessor()
	result := processor.Process(input)

	if result != expected {
		t.Errorf("Test: Multiple Modifiers in Sequence\nInput:    %s\nExpected: %s\nGot:      %s", input, expected, result)
	}
}

func TestAdvanced8_BatchWithPunctuation(t *testing.T) {
	input := "it was the BEST OF TIMES (low, 3) ! what a story ."
	expected := "it was the best of times! what a story."

	processor := fsm.NewProcessor()
	result := processor.Process(input)

	if result != expected {
		t.Errorf("Test: Batch Modifier with Punctuation\nInput:    %s\nExpected: %s\nGot:      %s", input, expected, result)
	}
}

func TestAdvanced9_ArticleWithQuotes(t *testing.T) {
	input := "She found a apple , a orange and a ' honest ' person ."
	expected := "She found an apple, an orange and an 'honest' person."

	processor := fsm.NewProcessor()
	result := processor.Process(input)

	if result != expected {
		t.Errorf("Test: A/An with Punctuation and Quotes\nInput:    %s\nExpected: %s\nGot:      %s", input, expected, result)
	}
}

func TestAdvanced10_ComplexPunctuation(t *testing.T) {
	input := "Wait ... what ! ? Really ? ! I can not believe it ! ! !"
	expected := "Wait... what!? Really?! I can not believe it!!!"

	processor := fsm.NewProcessor()
	result := processor.Process(input)

	if result != expected {
		t.Errorf("Test: Complex Punctuation Groups\nInput:    %s\nExpected: %s\nGot:      %s", input, expected, result)
	}
}

func TestAdvanced11_BinaryHexEdgeCases(t *testing.T) {
	input := "The values 0 (bin) and 0 (hex) are equal , but 1111 (bin) equals F (hex) which is 15 ."
	expected := "The values 0 and 0 are equal, but 15 equals 15 which is 15."

	processor := fsm.NewProcessor()
	result := processor.Process(input)

	if result != expected {
		t.Errorf("Test: Binary/Hex Edge Cases\nInput:    %s\nExpected: %s\nGot:      %s", input, expected, result)
	}
}

// EDGE CASES

func TestEdge_EmptyInput(t *testing.T) {
	input := ""
	expected := ""

	processor := fsm.NewProcessor()
	result := processor.Process(input)

	if result != expected {
		t.Errorf("Empty input should return empty output")
	}
}

func TestEdge_OnlySpaces(t *testing.T) {
	input := "     "
	expected := ""

	processor := fsm.NewProcessor()
	result := processor.Process(input)

	if result != expected {
		t.Errorf("Only spaces should return empty output")
	}
}

func TestEdge_ModifierWithoutWord(t *testing.T) {
	input := "(cap) hello"
	expected := "(cap) hello"

	processor := fsm.NewProcessor()
	result := processor.Process(input)

	if result != expected {
		t.Errorf("Test: Modifier without word\nInput:    %s\nExpected: %s\nGot:      %s", input, expected, result)
	}
}

func TestEdge_InvalidModifier(t *testing.T) {
	input := "word (invalid) test"
	expected := "word (invalid) test"

	processor := fsm.NewProcessor()
	result := processor.Process(input)

	// Invalid modifiers should be treated as regular text
	if result != expected {
		t.Errorf("Invalid modifier should be kept as regular text")
	}
}

// STRESS TEST

func TestStress_LongInput(t *testing.T) {
	// Generate long input with multiple rules
	var builder strings.Builder
	for i := 0; i < 100; i++ {
		builder.WriteString("word (cap) and FF (hex) test ")
	}
	input := builder.String()

	processor := fsm.NewProcessor()
	result := processor.Process(input)

	// Just verify it completes without crashing
	if len(result) == 0 {
		t.Error("Long input produced empty output")
	}

	t.Logf("Processed %d chars successfully", len(input))
}
