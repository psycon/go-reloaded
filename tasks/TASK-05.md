# TASK-05: Main Entry Point

**Status:** ⬜ TODO  
**Dependencies:** TASK-04 (FSM Processor)  
**Estimated Time:** 15 minutes  

---

## 📋 PROMPT — FULL 4-STEP FLOW (execute sequentially)

---

### **STEP 1: Analyze & Confirm**

**Context:**
Create the main entry point that reads input file, processes with FSM, and writes output file.

**Requirements:**
1. Update `main.go` with complete implementation
2. Handle command-line arguments (input and output files)
3. Read input file
4. Process using FSM processor
5. Write output file
6. Error handling for file I/O

**Command Line Usage:**
```bash
go run . input.txt output.txt
```

**Architecture Reference:**
This is the entry point that orchestrates file I/O and calls FSM processor.

**Acceptance Criteria:**
- [ ] Validates command-line arguments (exactly 2 args)
- [ ] Reads input file with error handling
- [ ] Creates FSM processor instance
- [ ] Calls processor.Process() with input text
- [ ] Writes result to output file
- [ ] Proper error messages for failures
- [ ] Success message on completion

**Questions for Human:**
1. Should we display processing time? (Default: no, keep simple)
2. Should we support stdin/stdout? (Default: no, files only)
3. File permissions for output? (Default: 0644)

**AI Response:** 
"I understand the requirements. I will implement the main entry point with file I/O and FSM integration. Waiting for confirmation to proceed to Step 2."

---

### **STEP 2: Generate the Tests**

**Task:** Create end-to-end tests using actual files.

**Test File:** `tests/main_test.go`

**Test Coverage:**

1. **File I/O Tests:**
   - Read sample file
   - Process content
   - Write output
   - Verify output matches expected

2. **Error Handling Tests:**
   - Missing arguments
   - Non-existent input file
   - Invalid permissions
   - Empty input file

**Test Structure:**
```go
package tests

import (
    "os"
    "testing"
)

func TestMainWithValidFiles(t *testing.T) {
    // Create test input file
    input := "it (cap) was great"
    err := os.WriteFile("test_input.txt", []byte(input), 0644)
    if err != nil {
        t.Fatal(err)
    }
    defer os.Remove("test_input.txt")
    
    // Run main (would need to refactor main for testing)
    // For now, test the components separately
    
    // Verify output file created
    // Read and verify content
}

func TestMissingArguments(t *testing.T) {
    // Test that program exits with error when args missing
    // This tests the validation logic
}

func TestNonExistentInputFile(t *testing.T) {
    // Test error handling for missing input file
}
```

**Note:** Testing `main()` directly is tricky. We'll test the logic components and create sample files for manual testing.

**AI Deliverable:** Complete `tests/main_test.go` with basic end-to-end test

---

### **STEP 3: Generate the Code**

**Task:** Implement complete main.go

**File to Update: main.go**

```go
package main

import (
	"fmt"
	"os"
	"text-editor/fsm"
)

func main() {
	// Check command line arguments
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input_file> <output_file>")
		fmt.Println("Example: go run . input.txt output.txt")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Read input file
	input, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading input file '%s': %v\n", inputFile, err)
		os.Exit(1)
	}

	// Check if input is empty
	if len(input) == 0 {
		fmt.Printf("Warning: Input file '%s' is empty\n", inputFile)
	}

	// Process text using FSM
	processor := fsm.NewProcessor()
	result := processor.Process(string(input))

	// Write output file
	err = os.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Printf("Error writing output file '%s': %v\n", outputFile, err)
		os.Exit(1)
	}

	// Success message
	fmt.Printf("✓ Processing complete: %s → %s\n", inputFile, outputFile)
}
```

**Implementation Guidelines:**
- Use `os.Args` for command-line arguments
- Use `os.ReadFile()` for input (simple, reads entire file)
- Use `os.WriteFile()` for output (creates or overwrites)
- Proper error messages with file names
- Exit codes: 0 for success, 1 for errors

**AI Deliverable:** Complete `main.go` implementation

---

### **STEP 4: QA & Mark Complete**

**Task:** Verify main entry point works end-to-end.

**QA Checklist:**

1. **Create Test Files:**
   ```bash
   # Create sample input
   echo "it (cap) was the best of times (up)" > sample.txt
   ```

2. **Build and Run:**
   ```bash
   # Build executable
   go build -o text-editor .
   
   # Should create 'text-editor' executable (or text-editor.exe on Windows)
   ```

3. **Test Valid Input:**
   ```bash
   # Run with sample file
   ./text-editor sample.txt result.txt
   
   # Should output: ✓ Processing complete: sample.txt → result.txt
   
   # Verify result
   cat result.txt
   # Expected: "It was the best of TIMES"
   ```

4. **Test Error Cases:**
   ```bash
   # No arguments
   ./text-editor
   # Expected: Usage message + exit
   
   # Non-existent file
   ./text-editor missing.txt output.txt
   # Expected: Error message about reading file
   
   # One argument only
   ./text-editor input.txt
   # Expected: Usage message + exit
   ```

5. **Test All Audit Examples:**
   ```bash
   # Test 1: Case transformations
   echo "it (cap) was the worst of times (up)" > test1.txt
   ./text-editor test1.txt out1.txt
   cat out1.txt
   # Expected: "It was the worst of TIMES"
   
   # Test 2: Hex and Bin
   echo "Simply add 42 (hex) and 10 (bin)" > test2.txt
   ./text-editor test2.txt out2.txt
   cat out2.txt
   # Expected: "Simply add 66 and 2"
   
   # Test 3: Article correction
   echo "There is a untold story" > test3.txt
   ./text-editor test3.txt out3.txt
   cat out3.txt
   # Expected: "There is an untold story"
   ```

6. **Integration Test:**
   ```bash
   # Run all tests
   go test ./... -v
   
   # All tests should pass
   ```

**Update Progress:**
- Update `AGENTS.md` progress table: TASK-05 → ✅ COMPLETE
- Git commit:
  ```bash
  git add main.go tests/main_test.go
  git commit -m "feat: complete TASK-05 - main entry point with file I/O"
  ```

**Final Verification:**
```bash
# The program should now be fully functional!
go run . input.txt output.txt
```

**AI Final Response:**
"TASK-05 completed successfully. Main entry point implemented and tested. The text editor is now fully functional! Ready to proceed to testing tasks (TASK-06, TASK-07, TASK-08)."

---

## 📊 Success Metrics

- ✅ Main entry point implemented
- ✅ File I/O working correctly
- ✅ Error handling functional
- ✅ Command-line argument validation
- ✅ End-to-end flow working
- ✅ All audit examples pass

---

## 🔗 Related Tasks

- **Previous:** TASK-04 (FSM Processor)
- **Next:** TASK-06 (Unit Tests)
- **Milestone:** 🎉 Core implementation complete!

---

*Task created: October 26, 2025*