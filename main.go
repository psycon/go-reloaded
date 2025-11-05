package main

import (
	"fmt"
	"go-reloaded/fsm"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input_file> <output_file>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	input, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		os.Exit(1)
	}

	processor := fsm.NewProcessor()
	result := processor.Process(string(input))

	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Printf("Error writing output file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✓ Success: %s → %s\n", inputFile, outputFile)
}
