# TASK-02: Transforms Module (Hex/Bin/Case Conversions)

**Status:** ⬜ TODO  
**Dependencies:** TASK-01 (Project Setup)  
**Estimated Time:** 30 minutes  

---

## 📋 PROMPT — FULL 4-STEP FLOW (execute sequentially)

Copy this entire section into your AI tool and execute steps 1-4 in order.

---

### **STEP 1: Analyze & Confirm**

**Context:**
You are building the `transforms/` package for a text editor that applies transformations to words.

**Requirements:**
1. Create three files in `transforms/` directory:
   - `numbers.go` - Hex and Binary to Decimal conversions
   - `cases.go` - Uppercase, Lowercase, Capitalize functions
   - `article.go` - Fix "a" to "an" before vowels/h

2. All functions must be **pure functions** (no side effects)
3. Handle edge cases (invalid input returns original string)
4. Follow Go best practices

**Architecture Reference:**
See `docs/ANALYSIS.md` section 2 for transformation rules.

**Acceptance Criteria:**
- [ ] `HexToDec(hex string) string` - converts hex to decimal
- [ ] `BinToDec(bin string) string` - converts binary to decimal
- [ ] `ToUpper(s string) string` - converts to UPPERCASE
- [ ] `ToLower(s string) string` - converts to lowercase
- [ ] `Capitalize(s string) string` - capitalizes first letter only
- [ ] `FixArticle(word, nextWord string) string` - a → an logic
- [ ] All functions handle invalid input gracefully

**Questions for Human:**
1. Should invalid hex/bin return original string or empty string? (Default: original)
2. Should Capitalize force lowercase rest of word? (Default: no, keep as-is)

**AI Response:** 
"I understand the requirements. I will create three files with pure transformation functions. Waiting for confirmation to proceed to Step 2."

---

### **STEP 2: Generate the Tests**

**Task:** Create comprehensive unit tests for all transformation functions.

**Test File:** `tests/transforms_test.go`

**Test Coverage:**

1. **HexToDec Tests:**
   - Valid hex: "1E" → "30", "FF" → "255", "A" → "10"
   - Edge case: "0" → "0"
   - Invalid: "ZZZ" → "ZZZ" (returns original)

2. **BinToDec Tests:**
   - Valid binary: "10" → "2", "1111" → "15", "101" → "5"
   - Edge case: "0" → "0"
   - Invalid: "102" → "102" (returns original)

3. **Case Transformation Tests:**
   - ToUpper: "hello" → "HELLO"
   - ToLower: "WORLD" → "world"
   - Capitalize: "hello" → "Hello", "WORLD" → "WORLD" (keeps rest)

4. **Article Correction Tests:**
   - "a" + "amazing" → "an"
   - "a" + "honest" → "an"
   - "a" + "book" → "a"
   - "A" + "apple" → "An"

**Test Structure:**
```go
package tests

import (
    "testing"
    "text-editor/transforms"
)

func TestHexToDec(t *testing.T) {
    tests := []struct {
        input    string
        expected string
    }{
        // Test cases here
    }
    
    for _, tt := range tests {
        result := transforms.HexToDec(tt.input)
        if result != tt.expected {
            t.Errorf("HexToDec(%s) = %s; want %s", tt.input, result, tt.expected)
        }
    }
}

// Similar for other functions...
```

**AI Deliverable:** Complete `tests/transforms_test.go` file

---

### **STEP 3: Generate the Code**

**Task:** Implement all transformation functions to pass the tests.

**Files to Create:**

#### **1. transforms/numbers.go**
```go
package transforms

import (
    "strconv"
    "strings"
)

// HexToDec converts hexadecimal string to decimal string
func HexToDec(hex string) string {
    // Implementation here
}

// BinToDec converts binary string to decimal string
func BinToDec(bin string) string {
    // Implementation here
}
```

#### **2. transforms/cases.go**
```go
package transforms

import (
    "strings"
    "unicode"
)

// ToUpper converts string to uppercase
func ToUpper(s string) string {
    // Implementation here
}

// ToLower converts string to lowercase
func ToLower(s string) string {
    // Implementation here
}

// Capitalize capitalizes first letter only
func Capitalize(s string) string {
    // Implementation here
}
```

#### **3. transforms/article.go**
```go
package transforms

import (
    "strings"
    "unicode"
)

// FixArticle converts "a" to "an" if next word starts with vowel or h
func FixArticle(word, nextWord string) string {
    // Implementation here
}

func isVowel(ch rune) bool {
    // Helper function
}
```

**Implementation Guidelines:**
- Use `strconv.ParseInt()` for hex/bin conversion
- Handle errors gracefully (return original on error)
- Use `unicode` package for case operations
- Keep functions simple and readable

**AI Deliverable:** Complete implementation of all three files

---

### **STEP 4: QA & Mark Complete**

**Task:** Verify implementation and update documentation.

**QA Checklist:**

1. **Run Tests:**
   ```bash
   go test ./tests/transforms_test.go -v
   ```
   - [ ] All tests pass
   - [ ] No compiler errors
   - [ ] No warnings

2. **Code Quality:**
   - [ ] Functions are documented with comments
   - [ ] Error handling is proper
   - [ ] Code follows Go conventions
   - [ ] No magic numbers or hardcoded values

3. **Edge Cases:**
   - [ ] Empty string input handled
   - [ ] Invalid hex/bin handled
   - [ ] Nil/empty nextWord in FixArticle handled

4. **Integration Check:**
   ```go
   // Quick manual test
   fmt.Println(transforms.HexToDec("FF"))      // Should print: 255
   fmt.Println(transforms.BinToDec("1010"))    // Should print: 10
   fmt.Println(transforms.Capitalize("hello")) // Should print: Hello
   ```

**Update Progress:**
- Update `AGENTS.md` progress table: TASK-02 → ✅ COMPLETE
- Git commit:
  ```bash
  git add transforms/ tests/transforms_test.go
  git commit -m "feat: complete TASK-02 - transforms module with tests"
  ```

**AI Final Response:**
"TASK-02 completed successfully. All tests pass. Ready to proceed to TASK-03."

---

## 📊 Success Metrics

- ✅ All 6 transformation functions implemented
- ✅ 15+ test cases passing
- ✅ 100% test coverage for transforms package
- ✅ Code follows Go best practices
- ✅ No edge case failures

---

## 🔗 Related Tasks

- **Previous:** TASK-01 (Project Setup)
- **Next:** TASK-03 (Formatters Module)
- **Depends On:** None (independent module)

---

*Task created: October 26, 2025*