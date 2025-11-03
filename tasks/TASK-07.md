# TASK-07: Integration Tests

**Status:** ⬜ TODO  
**Dependencies:** TASK-04 (FSM Processor), TASK-06 (Unit Tests)  
**Estimated Time:** 30 minutes  

---

## 📋 PROMPT — FULL 4-STEP FLOW (execute sequentially)

---

### **STEP 1: Analyze & Confirm**

**Context:**
Create integration tests that validate the FSM processor works correctly with all rules combined.

**Requirements:**
1. Test FSM processor end-to-end
2. Test combinations of multiple rules
3. Test complex scenarios from audit examples
4. Verify rule interactions work correctly
5. Test edge cases in integrated system

**Architecture Reference:**
See `docs/ANALYSIS.md` section 4 - Golden Test Set (Tests 1-11).

**Test Categories:**
- **Basic Integration**: Single rules working through FSM
- **Rule Combinations**: Multiple rules in same input
- **Complex Scenarios**: Real-world text with many rules
- **Edge Cases**: Boundary conditions in integrated system

**Acceptance Criteria:**
- [ ] 15+ integration test cases
- [ ] Cover all audit examples (Tests 1-6)
- [ ] Test rule interactions
- [ ] Test complex multi-rule scenarios
- [ ] All tests pass consistently

**Questions for Human:**
1. Should we test performance/speed? (Default: no, focus on correctness)
2. Include stress tests (very long input)? (Default: yes, one test)

**AI Response:** 
"I understand the requirements. I will create comprehensive integration tests for FSM processor. Waiting for confirmation to proceed to Step 2."

---

### **STEP 2: Generate the Tests**

**Task:** Create integration test suite based on audit examples.

**Test File:** `tests/integration_test.go` (expand existing)

```go
package tests

import (
	"strings"
	"testing"
	"text-editor/fsm"
)

// BASIC AUDIT EXAMPLES (from ANALYSIS.md)

func TestAudit1_MixedCaseTransformations(t *testing.T) {
	input := "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6)"
	expected := "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness"
	
	processor := fsm.NewProcessor()
	result := processor.Process(input)
	
	if result != expected {
		t.Errorf("\nTest: Mixed Case Transformations")
		t.Errorf("Input:    %s", input)
		t.Errorf("Expected: %s", expected)
		t.Errorf("Got:      %s", result)
	}
}

func TestAudit2_HexBinConversions(t *testing.T) {
	input := "Simply add 42 (hex) and 10 (bin) and you will see the result is 68."
	expected := "Simply add 66 and 2 and you will see the result is 68."
	
	processor := fsm.NewProcessor()
	result := processor.Process(input)
	
	if result != expected {
		t.Errorf("\nTest: Hex and Binary Conversions")
		t.Errorf("Input:    %s", input)
		t.Errorf("Expected: %s", expected)
		t.Errorf("Got:      %s", result)
	}
}

func TestAudit3_ArticleCorrection(t *testing.T) {
	input := "There is no greater agony than bearing a untold story inside you."
	expected := "There is no greater agony than bearing an untold story inside you."
	
	processor := fsm.NewProcessor()
	result := processor.Process(input)
	
	if result != expected {
		t.Errorf("\nTest: Article Correction")
		t.Errorf("Input:    %s", input)
		t.Errorf("Expected: %s", expected)
		t.Errorf("Got:      %s", result)
	}
}

func TestAudit4_PunctuationSpacing(t *testing.T) {
	input := "Punctuation tests are ... kinda boring ,what do you think ?"
	expected := "Punctuation tests are... kinda boring, what do you think?"
	
	processor := fsm.NewProcessor()
	result := processor.Process(input)
	
	if result != expected {
		t.Errorf("\nTest: Punctuation Spacing")
		t.Errorf("Input:    %s", input)
		t.Errorf("Expected: %s", expected)
		t.Errorf("Got:      %s", result)
	}
}

func TestAudit5_QuoteSingleWord(t *testing.T) {
	input := "I am exactly how they describe me: ' awesome '"
	expected := "I am exactly how they describe me: 'awesome'"
	
	processor := fsm.NewProcessor()
	result := processor.Process(input)
	
	if result != expected {
		t.Errorf("\nTest: Quote Handling (Single Word)")
		t.Errorf("Input:    %s", input)
		t.Errorf("Expected: %s", expected)
		t.Errorf("Got:      %s", result)
	}
}

func TestAudit6_QuoteMultipleWords(t *testing.T) {
	input := "As Elton John said: ' I am the most well-known homosexual in the world '"
	expected := "As Elton John said: 'I am the most well-known homosexual in the world'"
	
	processor := fsm.NewProcessor()
	result := processor.Process(input)
	
	if result != expected {
		t.Errorf("\nTest: Quote Handling (Multiple Words)")
		t.Errorf("Input:    %s", input)
		t.Errorf("Expected: %s", expected)
		t.Errorf("Got:      %s", result)
	}
}

// ADVANCED SCENARIOS

func TestAdvanced7_MultipleModifiers(t *testing.T) {
	input := "the word (up) was then (cap) followed by another (low) transformation and FF (hex) things."
	expected := "the WORD was then Followed by another transformation and 255 things."
	
	processor := fsm.NewProcessor()
	result := processor.Process(input)
	
	if result != expected {
		t.Errorf("\nTest: Multiple Modifiers in Sequence")
		t.Errorf("Input:    %s", input)
		t.Errorf("Expected: %s", expected)
		t.Errorf("Got:      %s", result)
	}
}

func TestAdvanced8_BatchWithPunctuation(t *testing.T) {
	input := "it was the BEST OF TIMES (low, 3) ! what a story ."
	expected := "it was the best of times! what a story."
	
	processor := fsm.NewProcessor()
	result := processor.Process(input)
	
	if result != expected {
		t.Errorf("\nTest: Batch Modifier with Punctuation")
		t.Errorf("Input:    %s", input)
		t.Errorf("Expected: %s", expected)
		t.Errorf("Got:      %s", result)
	}
}

func TestAdvanced9_ArticleWithQuotes(t *testing.T) {
	input := "She found a apple , a orange and a ' honest ' person ."
	expected := "She found an apple, an orange and an 'honest' person."
	
	processor := fsm.NewProcessor()
	result := processor.Process(input)
	
	if result != expected {
		t.Errorf("\nTest: A/An with Punctuation and Quotes")
		t.Errorf("Input:    %s", input)
		t.Errorf("Expected: %s", expected)
		t.Errorf("Got:      %s", result)
	}
}

func TestAdvanced10_ComplexPunctuation(t *testing.T) {
	input := "Wait ... what ! ? Really ? ! I can not believe it ! ! !"
	expected := "Wait... what!? Really?! I can not believe it!!!"
	
	processor := fsm.NewProcessor()
	result := processor.Process(input)
	
	if result != expected {
		t.Errorf("\nTest: Complex Punctuation Groups")
		t.Errorf("Input:    %s", input)
		t.Errorf("Expected: %s", expected)
		t.Errorf("Got:      %s", result)
	}
}

func TestAdvanced11_BinaryHexEdgeCases(t *testing.T) {
	input := "The values 0 (bin) and 0 (hex) are equal , but 1111 (bin) equals F (hex) which is 15 (hex) ."
	expected := "The values 0 and 0 are equal, but 15 equals 15 which is 15."
	
	processor := fsm.NewProcessor()
	result := processor.Process(input)
	
	if result != expected {
		t.Errorf("\nTest: Binary/Hex Edge Cases")
		t.Errorf("Input:    %s", input)
		t.Errorf("Expected: %s", expected)
		t.Errorf("Got:      %s", result)
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
	
	// Modifier at start with no previous word - should be ignored or kept as-is
	// Behavior depends on implementation
	t.Logf("Modifier without word: %s", result)
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
```

**AI Deliverable:** Complete `tests/integration_test.go` with 15+ tests

---

### **STEP 3: Generate the Code**

**Task:** No new code generation - tests validate existing FSM implementation.

**If Tests Fail:**

1. **Debug failing tests:**
   ```bash
   go test ./tests/integration_test.go -v -run TestAudit1
   ```

2. **Common issues:**
   - Spacing around punctuation
   - Quote handling with multiple words
   - Batch modifier count logic
   - a/an rule not applying in all cases

3. **Fix bugs in `fsm/processor.go`** and re-run tests

**AI Deliverable:** Bug fixes if needed, confirmation all tests pass

---

### **STEP 4: QA & Mark Complete**

**Task:** Verify all integration tests pass.

**QA Checklist:**

1. **Run All Integration Tests:**
   ```bash
   go test ./tests/integration_test.go -v
   ```
   - [ ] All 15+ tests pass
   - [ ] No flaky tests
   - [ ] Clear pass/fail output

2. **Run Full Test Suite:**
   ```bash
   go test ./tests/... -v
   ```
   - [ ] Unit tests still pass
   - [ ] Integration tests pass
   - [ ] No regressions

3. **Test Coverage:**
   ```bash
   go test ./fsm -cover
   ```
   - [ ] FSM coverage > 90%
   - [ ] All code paths tested

4. **Performance Check:**
   ```bash
   go test ./tests/integration_test.go -run TestStress -v
   ```
   - [ ] Completes in reasonable time (< 1 second)
   - [ ] No memory issues

**Update Progress:**
- Update `AGENTS.md` progress table: TASK-07 → ✅ COMPLETE
- Git commit:
  ```bash
  git add tests/integration_test.go
  git commit -m "feat: complete TASK-07 - integration tests for FSM processor"
  ```

**AI Final Response:**
"TASK-07 completed successfully. All integration tests pass including audit examples and edge cases. Ready to proceed to TASK-08 (Golden Test Suite)."

---

## 📊 Success Metrics

- ✅ 15+ integration tests
- ✅ All 6 audit examples pass
- ✅ Advanced scenarios tested
- ✅ Edge cases covered
- ✅ FSM coverage > 90%

---

## 🔗 Related Tasks

- **Previous:** TASK-04 (FSM Processor), TASK-06 (Unit Tests)
- **Next:** TASK-08 (Golden Test Suite)
- **Validates:** Entire system integration

---

*Task created: October 26, 2025*