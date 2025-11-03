# TASK-03: Formatters Module (Punctuation & Quotes)

**Status:** ⬜ TODO  
**Dependencies:** TASK-01 (Project Setup)  
**Estimated Time:** 25 minutes  

---

## 📋 PROMPT — FULL 4-STEP FLOW (execute sequentially)

---

### **STEP 1: Analyze & Confirm**

**Context:**
Create the `formatters/` package that handles punctuation spacing and quote formatting.

**Requirements:**
1. Create two files in `formatters/` directory:
   - `punctuation.go` - Format punctuation spacing and grouping
   - `quotes.go` - Format single quotes around words

2. All functions must be **pure functions** (no side effects)
3. Handle edge cases (empty input, single word, multiple words)
4. Follow Go best practices

**Architecture Reference:**
See `docs/ANALYSIS.md` section 2.4 and 2.5 for formatting rules.

**Rules to Implement:**

**Punctuation Rules:**
- Punctuation sticks to previous word (no space before)
- Space after punctuation
- Groups like `...` or `!?` have no internal spaces

**Quote Rules:**
- Single word: `'awesome'` (quotes stick to word)
- Multiple words: `'hello world'` (quotes stick to first and last word)
- Words inside quotes are space-separated

**Acceptance Criteria:**
- [ ] `FormatPunctuation(punct string) string` - returns formatted punctuation
- [ ] `FormatQuote(words []string) string` - returns quoted text
- [ ] Punctuation groups handled correctly (`.` `.` `.` → `...`)
- [ ] Single and multiple word quotes work
- [ ] Empty input handled gracefully

**Questions for Human:**
1. Should FormatPunctuation handle spacing logic? (Default: returns punct only, spacing handled by FSM)
2. Should FormatQuote validate input? (Default: no, assumes valid input)

**AI Response:** 
"I understand the requirements. I will create two files with pure formatting functions. Waiting for confirmation to proceed to Step 2."

---

### **STEP 2: Generate the Tests**

**Task:** Create comprehensive unit tests for formatting functions.

**Test File:** `tests/formatters_test.go`

**Test Coverage:**

1. **FormatPunctuation Tests:**
   - Single punctuation: "." → "."
   - Punctuation groups: "..." → "..." (already grouped)
   - Mixed: "!?" → "!?"

2. **FormatQuote Tests:**
   - Single word: ["awesome"] → "'awesome'"
   - Multiple words: ["hello", "world"] → "'hello world'"
   - Empty: [] → "''"
   - Three words: ["I", "am", "great"] → "'I am great'"

**Test Structure:**
```go
package tests

import (
    "testing"
    "text-editor/formatters"
)

func TestFormatPunctuation(t *testing.T) {
    tests := []struct {
        input    string
        expected string
    }{
        {".", "."},
        {"...", "..."},
        {"!?", "!?"},
        {",", ","},
    }
    
    for _, tt := range tests {
        result := formatters.FormatPunctuation(tt.input)
        if result != tt.expected {
            t.Errorf("FormatPunctuation(%s) = %s; want %s", tt.input, result, tt.expected)
        }
    }
}

func TestFormatQuote(t *testing.T) {
    tests := []struct {
        input    []string
        expected string
    }{
        {[]string{"awesome"}, "'awesome'"},
        {[]string{"hello", "world"}, "'hello world'"},
        {[]string{}, "''"},
        {[]string{"I", "am", "great"}, "'I am great'"},
    }
    
    for _, tt := range tests {
        result := formatters.FormatQuote(tt.input)
        if result != tt.expected {
            t.Errorf("FormatQuote(%v) = %s; want %s", tt.input, result, tt.expected)
        }
    }
}
```

**AI Deliverable:** Complete `tests/formatters_test.go` file

---

### **STEP 3: Generate the Code**

**Task:** Implement all formatting functions to pass the tests.

**Files to Create:**

#### **1. formatters/punctuation.go**
```go
package formatters

// FormatPunctuation formats punctuation with proper spacing
// Punctuation sticks to previous word, space after
// Note: Spacing is handled by FSM, this just returns the punctuation
func FormatPunctuation(punct string) string {
    // For now, just return as-is since FSM handles spacing
    // Punctuation groups (e.g., "...") are already combined by FSM
    return punct
}
```

#### **2. formatters/quotes.go**
```go
package formatters

import "strings"

// FormatQuote formats words between quotes
// Single word: 'awesome'
// Multiple words: 'I am great'
func FormatQuote(words []string) string {
    if len(words) == 0 {
        return "''"
    }
    
    if len(words) == 1 {
        return "'" + words[0] + "'"
    }
    
    // Multiple words: join with spaces and wrap
    return "'" + strings.Join(words, " ") + "'"
}
```

**Implementation Guidelines:**
- Keep functions simple and focused
- Use `strings.Join()` for combining words
- Handle empty slice edge case
- No validation needed (FSM ensures valid input)

**AI Deliverable:** Complete implementation of both files

---

### **STEP 4: QA & Mark Complete**

**Task:** Verify implementation and update documentation.

**QA Checklist:**

1. **Run Tests:**
   ```bash
   go test ./tests/formatters_test.go -v
   ```
   - [ ] All tests pass
   - [ ] No compiler errors
   - [ ] No warnings

2. **Code Quality:**
   - [ ] Functions are documented with comments
   - [ ] Code follows Go conventions
   - [ ] Simple and readable implementation
   - [ ] Edge cases handled (empty input)

3. **Edge Cases:**
   - [ ] Empty slice returns "''"
   - [ ] Single word works correctly
   - [ ] Multiple words joined with spaces
   - [ ] Quotes properly placed

4. **Integration Check:**
   ```go
   // Quick manual test
   fmt.Println(formatters.FormatQuote([]string{"awesome"}))
   // Output: 'awesome'
   
   fmt.Println(formatters.FormatQuote([]string{"hello", "world"}))
   // Output: 'hello world'
   
   fmt.Println(formatters.FormatPunctuation("..."))
   // Output: ...
   ```

**Update Progress:**
- Update `AGENTS.md` progress table: TASK-03 → ✅ COMPLETE
- Git commit:
  ```bash
  git add formatters/ tests/formatters_test.go
  git commit -m "feat: complete TASK-03 - formatters module with tests"
  ```

**AI Final Response:**
"TASK-03 completed successfully. All tests pass. Ready to proceed to TASK-04."

---

## 📊 Success Metrics

- ✅ Both formatting functions implemented
- ✅ 8+ test cases passing
- ✅ 100% test coverage for formatters package
- ✅ Code follows Go best practices
- ✅ Edge cases handled

---

## 🔗 Related Tasks

- **Previous:** TASK-02 (Transforms Module)
- **Next:** TASK-04 (FSM Processor Core)
- **Depends On:** TASK-01 (Project Setup)

---

*Task created: October 26, 2025*