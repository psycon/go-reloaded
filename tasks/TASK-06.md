# TASK-06: Comprehensive Unit Tests

**Status:** ⬜ TODO  
**Dependencies:** TASK-02, TASK-03 (Transforms & Formatters)  
**Estimated Time:** 30 minutes  

---

## 📋 PROMPT — FULL 4-STEP FLOW (execute sequentially)

---

### **STEP 1: Analyze & Confirm**

**Context:**
Expand unit tests to achieve comprehensive coverage of all transformation and formatting functions.

**Requirements:**
1. Enhance existing `tests/transforms_test.go` with edge cases
2. Enhance existing `tests/formatters_test.go` with edge cases
3. Add test cases for error conditions
4. Achieve 100% code coverage for pure functions
5. Test boundary conditions

**Architecture Reference:**
See `docs/ANALYSIS.md` section 4 (Golden Test Set) for test case inspiration.

**Test Categories:**

**Transforms Tests:**
- Valid inputs (happy path)
- Invalid inputs (error handling)
- Edge cases (empty, zero, max values)
- Boundary conditions

**Formatters Tests:**
- Single vs multiple elements
- Empty inputs
- Special characters
- Edge cases

**Acceptance Criteria:**
- [ ] 30+ test cases for transforms
- [ ] 15+ test cases for formatters
- [ ] All edge cases covered
- [ ] Test coverage > 95%
- [ ] Clear test names and documentation
- [ ] Table-driven tests for readability

**Questions for Human:**
1. Should we test performance/benchmarks? (Default: no, focus on correctness)
2. Include negative test cases? (Default: yes, test error handling)

**AI Response:** 
"I understand the requirements. I will expand unit tests with comprehensive coverage including edge cases. Waiting for confirmation to proceed to Step 2."

---

### **STEP 2: Generate the Tests**

**Task:** Create comprehensive unit test suite.

**Files to Enhance:**

#### **tests/transforms_test.go (Enhanced)**

```go
package tests

import (
	"testing"
	"text-editor/transforms"
)

// Hex to Dec Tests
func TestHexToDec(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// Valid cases
		{"single digit", "A", "10"},
		{"two digits", "1E", "30"},
		{"max byte", "FF", "255"},
		{"zero", "0", "0"},
		{"lowercase", "ff", "255"},
		{"mixed case", "FfAa", "65450"},
		
		// Edge cases
		{"leading zeros", "00FF", "255"},
		{"large number", "FFFF", "65535"},
		
		// Invalid cases (should return original)
		{"invalid char", "ZZZ", "ZZZ"},
		{"empty", "", ""},
		{"special char", "FF@", "FF@"},
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
		// Valid cases
		{"single bit", "1", "1"},
		{"two bits", "10", "2"},
		{"four bits", "1111", "15"},
		{"eight bits", "11111111", "255"},
		{"zero", "0", "0"},
		
		// Edge cases
		{"leading zeros", "00001010", "10"},
		{"max 16-bit", "1111111111111111", "65535"},
		
		// Invalid cases
		{"invalid digit", "102", "102"},
		{"empty", "", ""},
		{"letters", "abc", "abc"},
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
		{"with numbers", "hello123", "HELLO123"},
		{"special chars", "hello!", "HELLO!"},
		{"empty", "", ""},
		{"single char", "a", "A"},
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
		{"with numbers", "HELLO123", "hello123"},
		{"special chars", "HELLO!", "hello!"},
		{"empty", "", ""},
		{"single char", "A", "a"},
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
		{"single char", "a", "A"},
		{"empty", "", ""},
		{"number first", "1hello", "1hello"},
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
		// Vowel cases
		{"a + vowel a", "a", "amazing", "an"},
		{"a + vowel e", "a", "elephant", "an"},
		{"a + vowel i", "a", "incredible", "an"},
		{"a + vowel o", "a", "orange", "an"},
		{"a + vowel u", "a", "umbrella", "an"},
		
		// H case
		{"a + h", "a", "honest", "an"},
		{"a + H capital", "a", "Honest", "an"},
		
		// Consonant cases
		{"a + consonant", "a", "book", "a"},
		{"a + consonant b", "a", "banana", "a"},
		
		// Capital A
		{"A + vowel", "A", "amazing", "An"},
		{"A + consonant", "A", "book", "A"},
		
		// Not 'a' word
		{"the + vowel", "the", "apple", "the"},
		{"an + vowel", "an", "apple", "an"},
		
		// Edge cases
		{"empty next", "a", "", "a"},
		{"a alone", "a", "", "a"},
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
```

#### **tests/formatters_test.go (Enhanced)**

```go
package tests

import (
	"testing"
	"text-editor/formatters"
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
		
		// Groups
		{"ellipsis", "...", "..."},
		{"question exclaim", "!?", "!?"},
		{"exclaim question", "?!", "?!"},
		{"multiple exclaim", "!!!", "!!!"},
		
		// Edge cases
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
		// Single word
		{"single word", []string{"awesome"}, "'awesome'"},
		{"single word caps", []string{"HELLO"}, "'HELLO'"},
		
		// Multiple words
		{"two words", []string{"hello", "world"}, "'hello world'"},
		{"three words", []string{"I", "am", "great"}, "'I am great'"},
		{"many words", []string{"This", "is", "a", "test"}, "'This is a test'"},
		
		// Edge cases
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
```

**AI Deliverable:** Enhanced test files with 40+ test cases

---

### **STEP 3: Generate the Code**

**Task:** No new code generation - tests should pass with existing implementation.

**Verification Steps:**

1. **Run Enhanced Tests:**
   ```bash
   go test ./tests/transforms_test.go -v
   go test ./tests/formatters_test.go -v
   ```

2. **Check Coverage:**
   ```bash
   go test ./transforms -cover
   go test ./formatters -cover
   
   # Should show high coverage (>95%)
   ```

3. **If Tests Fail:**
   - Identify failing test cases
   - Fix bugs in transforms/formatters implementation
   - Re-run tests until all pass

**AI Deliverable:** Confirmation that all tests pass or list of bugs to fix

---

### **STEP 4: QA & Mark Complete**

**Task:** Verify comprehensive test coverage.

**QA Checklist:**

1. **Run All Unit Tests:**
   ```bash
   go test ./tests/transforms_test.go ./tests/formatters_test.go -v
   ```
   - [ ] All tests pass
   - [ ] No flaky tests (run multiple times)
   - [ ] Clear test output

2. **Coverage Analysis:**
   ```bash
   # Detailed coverage
   go test ./transforms -coverprofile=coverage_transforms.out
   go test ./formatters -coverprofile=coverage_formatters.out
   
   # View coverage
   go tool cover -html=coverage_transforms.out
   go tool cover -html=coverage_formatters.out
   ```
   - [ ] Transforms coverage > 95%
   - [ ] Formatters coverage > 95%
   - [ ] All functions tested

3. **Test Quality:**
   - [ ] Test names are descriptive
   - [ ] Table-driven tests used
   - [ ] Edge cases included
   - [ ] Error cases included
   - [ ] No redundant tests

4. **Documentation:**
   ```bash
   # Test examples should serve as documentation
   go test ./tests/transforms_test.go -v | grep "Test"
   ```

**Update Progress:**
- Update `AGENTS.md` progress table: TASK-06 → ✅ COMPLETE
- Git commit:
  ```bash
  git add tests/
  git commit -m "feat: complete TASK-06 - comprehensive unit tests with 95%+ coverage"
  ```

**AI Final Response:**
"TASK-06 completed successfully. Comprehensive unit tests with 40+ test cases and >95% coverage. Ready to proceed to TASK-07 (Integration Tests)."

---

## 📊 Success Metrics

- ✅ 40+ unit test cases
- ✅ 95%+ code coverage
- ✅ All edge cases covered
- ✅ Table-driven test format
- ✅ Clear test documentation

---

## 🔗 Related Tasks

- **Previous:** TASK-02, TASK-03 (Modules being tested)
- **Next:** TASK-07 (Integration Tests)
- **Parallel:** Can be done alongside implementation tasks

---

*Task created: October 26, 2025*