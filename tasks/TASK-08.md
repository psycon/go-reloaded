# TASK-08: Golden Test Suite (Comprehensive Paragraph Test)

**Status:** ⬜ TODO  
**Dependencies:** TASK-07 (Integration Tests)  
**Estimated Time:** 20 minutes  

---

## 📋 PROMPT — FULL 4-STEP FLOW (execute sequentially)

---

### **STEP 1: Analyze & Confirm**

**Context:**
Implement the comprehensive Test 12 from ANALYSIS.md - a large paragraph that combines ALL rules in realistic text.

**Requirements:**
1. Implement Test 12 (comprehensive paragraph)
2. Create additional complex real-world scenarios
3. Test all rules working together
4. Verify no rule conflicts
5. Golden reference tests (expected output is canon)

**Architecture Reference:**
See `docs/ANALYSIS.md` section 4 - Test 12 (Comprehensive Test Paragraph).

**Test Purpose:**
This is the **"golden test"** - if this passes, the entire system is working correctly. It combines:
- Multiple modifiers (hex, bin, up, low, cap)
- Batch operations ((up, N))
- Punctuation (groups, spacing)
- Quotes (single, multiple words)
- Article correction (a/an)
- Complex interactions

**Acceptance Criteria:**
- [ ] Test 12 implemented exactly as specified
- [ ] 2-3 additional complex scenarios
- [ ] All rules interact correctly
- [ ] No conflicts between rules
- [ ] Tests serve as regression suite

**Questions for Human:**
1. Should we create file-based golden tests? (Default: yes, one example)
2. Include performance benchmarks? (Default: no)

**AI Response:** 
"I understand the requirements. I will implement the comprehensive golden test suite. Waiting for confirmation to proceed to Step 2."

---

### **STEP 2: Generate the Tests**

**Task:** Create golden test suite with comprehensive scenarios.

**Test File:** `tests/golden_test.go`

```go
package tests

import (
	"os"
	"testing"
	"text-editor/fsm"
)

// TEST 12: Comprehensive Paragraph (THE GOLDEN TEST)
func TestGolden_ComprehensiveParagraph(t *testing.T) {
	input := `it (cap) was a amazing DAY (low) ! the sun was shining and the temperature reached 1F (hex) degrees . I went to the store , bought 11 (bin) apples and A (up) orange . the shopkeeper said : ' you are a honest customer ' . when i got HOME (low, 2) , i realized that 101 (bin) plus A (hex) equals F (hex) ! what a DISCOVERY (cap) ... i could not BELIEVE IT (low, 2) ! ? this was the best day EVER (cap, 2) .`
	
	expected := `It was an amazing day! the sun was shining and the temperature reached 31 degrees. I went to the store, bought 3 apples and AN orange. the shopkeeper said: 'you are an honest customer'. when i got home, i realized that 5 plus 10 equals 15! what a Discovery... i could not believe it!? this was The Best Day Ever.`
	
	processor := fsm.NewProcessor()
	result := processor.Process(input)
	
	if result != expected {
		t.Errorf("\n=== GOLDEN TEST FAILED ===")
		t.Errorf("\nInput:\n%s", input)
		t.Errorf("\nExpected:\n%s", expected)
		t.Errorf("\nGot:\n%s", result)
		t.Errorf("\n=========================")
		
		// Show detailed diff
		t.Log("\nDetailed comparison:")
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
	
	expected := `The system uses An Advanced FSM architecture. It processes 255 tokens per second and handles 42 concurrent requests. Performance metrics are... impressive! the throughput increased BY an factor of 2.`
	
	processor := fsm.NewProcessor()
	result := processor.Process(input)
	
	if result != expected {
		t.Errorf("\nTechnical Document Test Failed")
		t.Errorf("Expected: %s", expected)
		t.Errorf("Got:      %s", result)
	}
}

func TestGolden_CreativeWriting(t *testing.T) {
	input := `' once upon a time (cap, 4) ' , there lived a amazing DRAGON (low) in A (up) ENCHANTED (cap) forest . The dragon loved to count in hexadecimal : 1A (hex) , 2B (hex) , 3C (hex) ... what a PECULIAR (cap) creature ! ? Indeed , it was a EXTRAORDINARY (low) tale .`
	
	expected := `'Once Upon A Time', there lived an amazing dragon in AN Enchanted forest. The dragon loved to count in hexadecimal: 26, 43, 60... what a Peculiar creature!? Indeed, it was an extraordinary tale.`
	
	processor := fsm.NewProcessor()
	result := processor.Process(input)
	
	if result != expected {
		t.Errorf("\nCreative Writing Test Failed")
		t.Errorf("Expected: %s", expected)
		t.Errorf("Got:      %s", result)
	}
}

// File-Based Golden Test

func TestGolden_FileProcessing(t *testing.T) {
	// Create temporary test file
	input := `it (cap) was the best of times (up, 4) , it was the worst of times . It was A (cap) TALE (low) OF (low) TWO (low) CITIES (low) .`
	
	inputFile := "test_golden_input.txt"
	outputFile := "test_golden_output.txt"
	expectedOutput := `It was The Best Of Times, it was the worst of times. It was A tale of two cities.`
	
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
		t.Errorf("\nFile Processing Test Failed")
		t.Errorf("Expected: %s", expectedOutput)
		t.Errorf("Got:      %s", result)
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
			input:    "word (up) (low)",
			expected: "word (low)", // Second modifier should be treated as text
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
				t.Errorf("\nRegression: %s", tt.name)
				t.Errorf("Expected: %s", tt.expected)
				t.Errorf("Got:      %s", result)
			}
		})
	}
}
```

**AI Deliverable:** Complete `tests/golden_test.go` with comprehensive tests

---

### **STEP 3: Generate the Code**

**Task:** No new code generation - golden tests validate entire system.

**If Golden Test Fails:**

This is the most important test. If it fails:

1. **Analyze the diff carefully:**
   ```bash
   go test ./tests/golden_test.go -v -run TestGolden_Comprehensive
   ```

2. **Common issues in comprehensive tests:**
   - Word order in batch modifiers
   - Spacing with quotes and punctuation
   - Multiple a/an corrections
   - Modifier priority/order

3. **Debug step-by-step:**
   - Test each sentence individually
   - Isolate the failing rule combination
   - Fix in `fsm/processor.go`
   - Rerun golden test

**AI Deliverable:** Bug fixes if needed, confirmation golden test passes

---

### **STEP 4: QA & Mark Complete**

**Task:** Verify golden test suite passes - this validates the entire project.

**QA Checklist:**

1. **Run Golden Tests:**
   ```bash
   go test ./tests/golden_test.go -v
   ```
   - [ ] Test 12 (Comprehensive) passes ⭐
   - [ ] All additional scenarios pass
   - [ ] File-based test passes
   - [ ] Regression suite passes

2. **Run Full Test Suite:**
   ```bash
   go test ./... -v
   ```
   - [ ] All unit tests pass
   - [ ] All integration tests pass
   - [ ] All golden tests pass
   - [ ] Zero failures

3. **Test Summary:**
   ```bash
   go test ./... -v | grep -E "(PASS|FAIL|RUN)"
   ```
   - [ ] Should show all PASS, no FAIL

4. **Final Validation:**
   ```bash
   # The ultimate test - run the actual program
   echo "it (cap) was a amazing day (up, 3)" > final_test.txt
   go run . final_test.txt final_out.txt
   cat final_out.txt
   # Expected: "It was AN AMAZING DAY"
   ```

**Documentation:**
- [ ] Golden tests documented in ANALYSIS.md
- [ ] Test coverage documented
- [ ] Known limitations documented (if any)

**Update Progress:**
- Update `AGENTS.md` progress table: TASK-08 → ✅ COMPLETE
- Git commit:
  ```bash
  git add tests/golden_test.go
  git commit -m "feat: complete TASK-08 - golden test suite with comprehensive validation"
  ```

**AI Final Response:**
"TASK-08 completed successfully. Golden test suite passes including the comprehensive Test 12. All system functionality validated. Ready to proceed to documentation tasks (TASK-09, TASK-10)."

---

## 📊 Success Metrics

- ✅ Test 12 (comprehensive) passes
- ✅ 3+ complex scenarios tested
- ✅ File-based golden test works
- ✅ Regression suite catches edge cases
- ✅ 100% of audit requirements met

---

## 🔗 Related Tasks

- **Previous:** TASK-07 (Integration Tests)
- **Next:** TASK-09 (README), TASK-10 (ANALYSIS)
- **Milestone:** 🎉 All code and tests complete!

---

*Task created: October 26, 2025*