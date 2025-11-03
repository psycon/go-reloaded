# TASK-04: FSM Processor Core

**Status:** ⬜ TODO  
**Dependencies:** TASK-02 (Transforms), TASK-03 (Formatters)  
**Estimated Time:** 60 minutes  

---

## 📋 PROMPT — FULL 4-STEP FLOW (execute sequentially)

---

### **STEP 1: Analyze & Confirm**

**Context:**
Create the core FSM processor that orchestrates text processing using transforms and formatters.

**Requirements:**
1. Create `fsm/processor.go` - Main FSM implementation
2. Implement state machine with token-by-token processing
3. Use transforms and formatters packages (pure functions)
4. Handle all rules: modifiers, punctuation, quotes, a/an
5. Single-pass processing (O(n) complexity)

**Architecture Reference:**
See `docs/ANALYSIS.md` section 3 for FSM architecture and state flow.

**FSM States:**
- `ReadingWord` - Collecting word tokens
- `ProcessingModifier` - Applying modifiers
- `InQuote` - Tracking quote context

**Core Responsibilities:**
1. **Tokenization** - Split input into tokens (words, punctuation, quotes, modifiers)
2. **State Management** - Track current state and context
3. **Word Buffering** - Keep last N words for batch modifiers
4. **Quote Handling** - Track words inside quotes
5. **Orchestration** - Call transforms/formatters at right times

**Acceptance Criteria:**
- [ ] `Processor` struct with state variables
- [ ] `NewProcessor() *Processor` - Constructor
- [ ] `Process(input string) string` - Main processing function
- [ ] `tokenize(input string) []string` - Tokenizer
- [ ] Helper functions: `isModifier()`, `isPunctuation()`, etc.
- [ ] Handles all transformation rules correctly
- [ ] Single-pass processing
- [ ] Memory efficient (no full text copies)

**Questions for Human:**
1. Should tokenizer split on all whitespace or preserve newlines? (Default: all whitespace)
2. Should processor handle streaming input? (Default: no, process entire string)
3. Max word buffer size for batch modifiers? (Default: unlimited, use slice)

**AI Response:** 
"I understand the requirements. I will create the FSM processor with state management and token processing. This is the most complex component. Waiting for confirmation to proceed to Step 2."

---

### **STEP 2: Generate the Tests**

**Task:** Create integration tests for FSM processor.

**Test File:** `tests/fsm_test.go`

**Test Coverage:**

1. **Basic Transformations:**
   - Single modifier: "word (cap)" → "Word"
   - Hex conversion: "1E (hex)" → "30"
   - Binary conversion: "10 (bin)" → "2"
   - Case: "GO (low)" → "go"

2. **Batch Modifiers:**
   - "(up, 2)": "this is great (up, 2)" → "this IS GREAT"
   - "(cap, 3)": "age of wisdom (cap, 3)" → "Age Of Wisdom"

3. **Punctuation:**
   - Basic: "word ," → "word,"
   - Groups: ". . ." → "..."
   - Multiple: "! ?" → "!?"

4. **Quotes:**
   - Single: "' word '" → "'word'"
   - Multiple: "' hello world '" → "'hello world'"

5. **Article Correction:**
   - "a amazing" → "an amazing"
   - "a honest" → "an honest"
   - "a book" → "a book"

6. **Complex Integration:**
   - From audit examples (Test 1-6 from ANALYSIS.md)

**Test Structure:**
```go
package tests

import (
    "testing"
    "text-editor/fsm"
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

// More tests following the pattern from ANALYSIS.md Golden Test Set
```

**AI Deliverable:** Complete `tests/fsm_test.go` with 10+ test cases

---

### **STEP 3: Generate the Code**

**Task:** Implement the FSM processor.

**File to Create: fsm/processor.go**

```go
package fsm

import (
    "fmt"
    "strings"
    "text-editor/formatters"
    "text-editor/transforms"
)

type State int

const (
    ReadingWord State = iota
    ProcessingModifier
    InQuote
)

type Processor struct {
    tokens      []string        // All input tokens
    pos         int             // Current position in tokens
    output      strings.Builder // Output buffer
    wordBuffer  []string        // Recent words (for batch modifiers)
    inQuote     bool            // Are we inside quotes?
    quoteWords  []string        // Words collected inside quotes
}

func NewProcessor() *Processor {
    return &Processor{
        wordBuffer: make([]string, 0),
        quoteWords: make([]string, 0),
    }
}

func (p *Processor) Process(input string) string {
    p.tokens = tokenize(input)
    p.pos = 0
    
    for p.pos < len(p.tokens) {
        token := p.tokens[p.pos]
        
        // Skip empty tokens
        if token == "" {
            p.pos++
            continue
        }
        
        // Handle quotes
        if token == "'" {
            p.handleQuote()
            p.pos++
            continue
        }
        
        // Check for modifiers
        if isModifier(token) {
            p.handleModifier(token)
            p.pos++
            continue
        }
        
        // Check for punctuation
        if isPunctuation(token) {
            p.handlePunctuation()
            continue // handlePunctuation advances pos
        }
        
        // Regular word
        if p.inQuote {
            p.quoteWords = append(p.quoteWords, token)
        } else {
            p.wordBuffer = append(p.wordBuffer, token)
        }
        
        p.pos++
    }
    
    // Flush remaining words
    p.flushBuffer()
    
    return strings.TrimSpace(p.output.String())
}

func (p *Processor) handleModifier(modifier string) {
    if len(p.wordBuffer) == 0 {
        return
    }
    
    modType, count := parseModifier(modifier)
    
    switch modType {
    case "hex":
        idx := len(p.wordBuffer) - 1
        p.wordBuffer[idx] = transforms.HexToDec(p.wordBuffer[idx])
    case "bin":
        idx := len(p.wordBuffer) - 1
        p.wordBuffer[idx] = transforms.BinToDec(p.wordBuffer[idx])
    case "up":
        p.applyCase(transforms.ToUpper, count)
    case "low":
        p.applyCase(transforms.ToLower, count)
    case "cap":
        p.applyCase(transforms.Capitalize, count)
    }
}

func (p *Processor) applyCase(fn func(string) string, count int) {
    if count == 0 {
        count = 1
    }
    
    start := len(p.wordBuffer) - count
    if start < 0 {
        start = 0
    }
    
    for i := start; i < len(p.wordBuffer); i++ {
        p.wordBuffer[i] = fn(p.wordBuffer[i])
    }
}

func (p *Processor) handlePunctuation() {
    // Collect consecutive punctuation
    group := ""
    for p.pos < len(p.tokens) && isPunctuation(p.tokens[p.pos]) {
        group += p.tokens[p.pos]
        p.pos++
    }
    
    // Flush words before punctuation
    p.flushBuffer()
    
    // Add punctuation (sticks to previous word)
    p.output.WriteString(group)
}

func (p *Processor) handleQuote() {
    if !p.inQuote {
        // Opening quote
        p.flushBuffer()
        p.inQuote = true
        p.quoteWords = make([]string, 0)
    } else {
        // Closing quote
        quoted := formatters.FormatQuote(p.quoteWords)
        
        if p.output.Len() > 0 && !endsWithSpace(p.output.String()) {
            p.output.WriteString(" ")
        }
        p.output.WriteString(quoted)
        
        p.inQuote = false
        p.quoteWords = make([]string, 0)
    }
}

func (p *Processor) flushBuffer() {
    for i := 0; i < len(p.wordBuffer); i++ {
        word := p.wordBuffer[i]
        
        // Check a/an rule
        if i < len(p.wordBuffer)-1 {
            nextWord := p.wordBuffer[i+1]
            word = transforms.FixArticle(word, nextWord)
        }
        
        // Add spacing
        if p.output.Len() > 0 {
            lastChar := p.output.String()[p.output.Len()-1]
            if lastChar != ' ' && !isPunctuationChar(byte(lastChar)) {
                p.output.WriteString(" ")
            } else if isPunctuationChar(byte(lastChar)) {
                p.output.WriteString(" ")
            }
        }
        
        p.output.WriteString(word)
    }
    
    p.wordBuffer = make([]string, 0)
}

// Tokenize splits input into tokens
func tokenize(input string) []string {
    var tokens []string
    var current strings.Builder
    
    for _, ch := range input {
        switch ch {
        case ' ', '\t', '\n', '\r':
            if current.Len() > 0 {
                tokens = append(tokens, current.String())
                current.Reset()
            }
        case '\'':
            if current.Len() > 0 {
                tokens = append(tokens, current.String())
                current.Reset()
            }
            tokens = append(tokens, string(ch))
        case '.', ',', '!', '?', ':', ';':
            if current.Len() > 0 {
                tokens = append(tokens, current.String())
                current.Reset()
            }
            tokens = append(tokens, string(ch))
        case '(', ')':
            current.WriteRune(ch)
        default:
            current.WriteRune(ch)
        }
    }
    
    if current.Len() > 0 {
        tokens = append(tokens, current.String())
    }
    
    return tokens
}

func isModifier(token string) bool {
    if !strings.HasPrefix(token, "(") || !strings.HasSuffix(token, ")") {
        return false
    }
    
    content := strings.TrimPrefix(strings.TrimSuffix(token, ")"), "(")
    parts := strings.Split(content, ",")
    
    if len(parts) == 0 {
        return false
    }
    
    base := strings.TrimSpace(parts[0])
    return base == "hex" || base == "bin" || base == "up" || base == "low" || base == "cap"
}

func parseModifier(token string) (string, int) {
    content := strings.TrimPrefix(strings.TrimSuffix(token, ")"), "(")
    parts := strings.Split(content, ",")
    
    modType := strings.TrimSpace(parts[0])
    count := 0
    
    if len(parts) > 1 {
        fmt.Sscanf(strings.TrimSpace(parts[1]), "%d", &count)
    }
    
    return modType, count
}

func isPunctuation(token string) bool {
    return len(token) == 1 && isPunctuationChar(token[0])
}

func isPunctuationChar(ch byte) bool {
    return ch == '.' || ch == ',' || ch == '!' || 
           ch == '?' || ch == ':' || ch == ';'
}

func endsWithSpace(s string) bool {
    return len(s) > 0 && s[len(s)-1] == ' '
}
```

**Implementation Guidelines:**
- State machine processes tokens sequentially
- Word buffer keeps recent words for batch operations
- Quote state tracks words inside quotes
- Flush buffer applies a/an rule and outputs words
- Tokenizer splits on whitespace and punctuation

**AI Deliverable:** Complete `fsm/processor.go` implementation

---

### **STEP 4: QA & Mark Complete**

**Task:** Verify FSM implementation with comprehensive testing.

**QA Checklist:**

1. **Run FSM Tests:**
   ```bash
   go test ./tests/fsm_test.go -v
   ```
   - [ ] All basic transformation tests pass
   - [ ] Batch modifier tests pass
   - [ ] Punctuation tests pass
   - [ ] Quote tests pass
   - [ ] Article correction tests pass

2. **Run Integration Tests:**
   ```bash
   go test ./tests/... -v
   ```
   - [ ] All unit tests still pass (transforms, formatters)
   - [ ] FSM integrates correctly with pure functions

3. **Code Quality:**
   - [ ] Functions are well-documented
   - [ ] State transitions are clear
   - [ ] No memory leaks (proper buffer resets)
   - [ ] Error handling for edge cases

4. **Manual Testing:**
   ```go
   processor := fsm.NewProcessor()
   
   // Test from audit examples
   result := processor.Process("it (cap) was the worst of times (up)")
   fmt.Println(result)
   // Expected: "It was the worst of TIMES"
   ```

5. **Performance Check:**
   - [ ] Single pass through tokens (no re-reading)
   - [ ] O(n) time complexity
   - [ ] Memory efficient (no full text copies)

**Update Progress:**
- Update `AGENTS.md` progress table: TASK-04 → ✅ COMPLETE
- Git commit:
  ```bash
  git add fsm/ tests/fsm_test.go
  git commit -m "feat: complete TASK-04 - FSM processor core with tests"
  ```

**AI Final Response:**
"TASK-04 completed successfully. FSM processor implemented and all tests pass. Ready to proceed to TASK-05."

---

## 📊 Success Metrics

- ✅ FSM processor fully implemented
- ✅ 10+ integration tests passing
- ✅ All transformation rules working
- ✅ Single-pass processing achieved
- ✅ Memory efficient implementation

---

## 🔗 Related Tasks

- **Previous:** TASK-02 (Transforms), TASK-03 (Formatters)
- **Next:** TASK-05 (Main Entry Point)
- **Depends On:** TASK-02, TASK-03

---

*Task created: October 26, 2025*