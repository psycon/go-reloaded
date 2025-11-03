package main

import (
	"fmt"
	"os"
)

func main() {
	// TODO: Implement in TASK-05
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input_file> <output_file>")
		os.Exit(1)
	}

	fmt.Println("Text Editor - Coming soon!")
	fmt.Println("Input:", os.Args[1])
	fmt.Println("Output:", os.Args[2])
}
