package tests

import (
	"go-reloaded/fsm"
	"testing"
)

func TestBasicCapitalize(t *testing.T) {
	input := "it (cap) was great"
	expected := "It was great"

	processor := fsm.NewProcessor()
	result := processor.Process(input)

	if result != expected {
		t.Errorf("\nInput:    %s\nExpected: %s\nGot:      %s", input, expected, result)
	}
}

func TestHexConversion(t *testing.T) {
	input := "Simply add 42 (hex) and 10 (bin)"
	expected := "Simply add 66 and 2"

	processor := fsm.NewProcessor()
	result := processor.Process(input)

	if result != expected {
		t.Errorf("\nInput:    %s\nExpected: %s\nGot:      %s", input, expected, result)
	}
}

func TestBatchModifier(t *testing.T) {
	input := "this is so exciting (up, 2)"
	expected := "this is SO EXCITING"

	processor := fsm.NewProcessor()
	result := processor.Process(input)

	if result != expected {
		t.Errorf("\nInput:    %s\nExpected: %s\nGot:      %s", input, expected, result)
	}
}

func TestPunctuationSpacing(t *testing.T) {
	input := "Hello , world !"
	expected := "Hello, world!"

	processor := fsm.NewProcessor()
	result := processor.Process(input)

	if result != expected {
		t.Errorf("\nInput:    %s\nExpected: %s\nGot:      %s", input, expected, result)
	}
}

func TestPunctuationGroups(t *testing.T) {
	input := "Wait . . . what ! ?"
	expected := "Wait... what!?"

	processor := fsm.NewProcessor()
	result := processor.Process(input)

	if result != expected {
		t.Errorf("\nInput:    %s\nExpected: %s\nGot:      %s", input, expected, result)
	}
}

func TestQuoteSingleWord(t *testing.T) {
	input := "I am ' awesome '"
	expected := "I am 'awesome'"

	processor := fsm.NewProcessor()
	result := processor.Process(input)

	if result != expected {
		t.Errorf("\nInput:    %s\nExpected: %s\nGot:      %s", input, expected, result)
	}
}

func TestQuoteMultipleWords(t *testing.T) {
	input := "He said ' hello world '"
	expected := "He said 'hello world'"

	processor := fsm.NewProcessor()
	result := processor.Process(input)

	if result != expected {
		t.Errorf("\nInput:    %s\nExpected: %s\nGot:      %s", input, expected, result)
	}
}

func TestArticleCorrection(t *testing.T) {
	input := "There is a untold story"
	expected := "There is an untold story"

	processor := fsm.NewProcessor()
	result := processor.Process(input)

	if result != expected {
		t.Errorf("\nInput:    %s\nExpected: %s\nGot:      %s", input, expected, result)
	}
}

func TestArticleWithH(t *testing.T) {
	input := "She is a honest person"
	expected := "She is an honest person"

	processor := fsm.NewProcessor()
	result := processor.Process(input)

	if result != expected {
		t.Errorf("\nInput:    %s\nExpected: %s\nGot:      %s", input, expected, result)
	}
}

func TestComplexExample(t *testing.T) {
	input := "it (cap) was the worst of times (up, 4) , it was a amazing day !"
	expected := "It was THE WORST OF TIMES, it was an amazing day!"

	processor := fsm.NewProcessor()
	result := processor.Process(input)

	if result != expected {
		t.Errorf("\nInput:    %s\nExpected: %s\nGot:      %s", input, expected, result)
	}
}
