<span style="color: #FF0000;"># Analysis/Developer Focused Document - GO-reloaded Project</span>

## Table of Contents
1. [Architecture Comparison](#architecture-comparison)
2. [Why FSM](#why-fsm)
3. [Transformation Rules Reference](#transformation-rules-reference)
4. [Golden Test Set](#golden-test-set)

---

<a name="architecture-comparison"></a>
## <span style="color: #CCFF99;">1. Architecture Comparison</span>

This section analyzes two possible architectural approaches for implementing the text editor: **Pipeline Architecture** and **FSM (Finite State Machine) Architecture**.

<span style="color: #FFD700;">### Pipeline Architecture</span>

**How it works:**
The text passes through a series of independent filters/stages:
1. Stage 1: Hex/Bin conversions
2. Stage 2: Case transformations
3. Stage 3: Punctuation formatting
4. Stage 4: Quote handling
5. Stage 5: A/An correction

**Advantages:**
- Simple and linear logic
- Each stage is independent and testable
- Easy to add new stages
- Separation of concerns
- Straightforward debugging per stage

**Disadvantages:**
- Requires multiple passes through the text (2+ times)
- Each stage creates an intermediate string
- High memory usage for large files
- Difficult to maintain context between stages
- Slower for large inputs
- Performance degrades with text size

---

<span style="color: #FFD700;">### FSM (Finite State Machine) Architecture</span>

**How it works:**
The program is always in a specific "state" and reads the input character-by-character or token-by-token. Depending on the current state and the next input, it changes state and executes actions.

**Core States:**
- `READING_WORD`: Collecting characters of a word
- `WORD_COMPLETE`: Word finished, checking for modifiers
- `IN_QUOTES`: Tracking if we're inside quotes
- `READING_MODIFIER`: Reading (hex), (up), etc.
- `HANDLE_PUNCTUATION`: Applying punctuation rules

**Advantages:**
-  **Single pass** - reads the text only once
-  **Memory efficient** - doesn't create intermediate copies
-  **Context awareness** - the state machine "remembers" where it is (inside quotes, after modifier, etc.)
-  **Natural fit** - text parsing is a classic FSM problem
-  **Faster execution** - O(n) complexity with minimal overhead
-  **Easier debugging** - you know exactly which state you're in
-  **Scalability** - handles large files efficiently
-  **Industry standard** - compilers, parsers, and lexers all use FSM

**Disadvantages:**
- More complex initial design
- Requires careful thinking about state transitions
- Harder to modify after initial implementation
- Steeper learning curve for beginners


<div align="center">

### FSM Architecture Diagram

<img src="../assets/fsm flow diagram.png" alt="FSM Flow Diagram" width="800"/>

</div>


---


<a name="why-fsm"></a>
### <span style="color: #CCFF99;">2. **Why FSM? My personal choice:**</span>

1. **Performance**: For text processing, single-pass FSM is objectively faster
2. **Memory**: For large files, the memory difference is significant
3. **Natural Design**: Text parsing is a natural FSM problem - it mimics how one thinks when reading
4. **Context Handling**: Many rules depend on context (am I inside quotes? did I just see a modifier?). FSM handles this naturally
5. **Industry Standard**: Compilers, parsers, lexers - all use FSM
6. **Learning Value**: More educational and professional approach
7. **Extensibility**: Adding new states is more maintainable than adding pipeline stages
8. **Real-time Processing**: Can process streams without buffering entire input

The **tradeoff** of complexity is worth it for the benefits that FSM offers in this project.

---

**State Flow:**
```
START → READING_WORD → WORD_COMPLETE → CHECK_MODIFIER
                                      ↓
                              [Modifier Found]
                                      ↓
                              APPLY_TRANSFORMATION → OUTPUT
                                                        ↓
                                                   BACK_TO_READ

Alternative paths:
- PUNCTUATION → FORMAT_PUNCT → OUTPUT → BACK_TO_READ
- QUOTE_START → IN_QUOTES → QUOTE_END → OUTPUT → BACK_TO_READ
- SPACE → CHECK_CONTEXT → OUTPUT → BACK_TO_READ
```

**Context Management:**
- **Word Buffer**: Maintains last N words for batch transformations
- **Quote State**: Boolean flag tracking if inside quotes
- **Previous Token**: Remembers last processed token for punctuation rules
- **Modifier Stack**: Stores pending modifiers to apply

---

<span style="color: #FFD700;">### Why Separation of Concerns?</span>

The project structure separates **orchestration** (FSM) from **business logic** (transforms/formatters):

```
fsm/          → Controls state transitions and flow
   ↓ uses
transforms/   → Pure functions: hex/bin conversions, case changes
formatters/   → Pure functions: punctuation, quote formatting
```

**Benefits:**
1. **Reusability**: Transforms can be used independently of FSM
2. **Testability**: Unit test pure functions without FSM overhead
3. **Maintainability**: Changes to logic don't affect state machine
4. **Extensibility**: Easy to add Pipeline architecture using same transforms

---

<span style="color: #FFD700;">### Implementation Considerations</span>

**Data Structures:**
- **Circular buffer** for word history (for `(up, N)` modifiers)
- **Stack** for nested quote handling
- **Enum/Constants** for state definitions
- **Output builder** (string buffer) for efficient concatenation

**Performance Optimization:**
- Pre-compile regex patterns for modifiers
- Use string builder instead of concatenation
- Minimize memory allocations
- Process in chunks for large files

**Error Handling:**
- Invalid modifiers (e.g., `(hex)` on non-hex string)
- Malformed input (unclosed quotes)
- Out-of-range batch numbers (e.g., `(up, 100)` when only 3 words exist)
- File I/O errors

---

<a name="transformation-rules-reference"></a>
## <span style="color: #CCFF99;">3. Transformation Rules Reference</span>

### 3.1 Number Base Conversions

#### `(hex)` - Hexadecimal to Decimal
Converts the previous word from hexadecimal to decimal.

**Examples:**
- `"1E (hex) files"` → `"30 files"`
- `"FF (hex) is max"` → `"255 is max"`

#### `(bin)` - Binary to Decimal
Converts the previous word from binary to decimal.

**Examples:**
- `"10 (bin) years"` → `"2 years"`
- `"1010 (bin) equals"` → `"10 equals"`

---

### 3.2 Case Transformations

#### `(up)` - Uppercase
Converts the previous word to UPPERCASE.

**Example:**
- `"go (up) now"` → `"GO now"`

#### `(low)` - Lowercase
Converts the previous word to lowercase.

**Example:**
- `"SHOUTING (low)"` → `"shouting"`

#### `(cap)` - Capitalize
Capitalizes only the first letter of the previous word.

**Example:**
- `"bridge (cap)"` → `"Bridge"`

---

### 3.3 Batch Transformations

#### `(up, N)` - Uppercase N Words
Converts the N previous words to UPPERCASE.

**Example:**
- `"so exciting (up, 2)"` → `"SO EXCITING"`

#### `(low, N)` - Lowercase N Words
Converts the N previous words to lowercase.

**Example:**
- `"IT WAS THE (low, 3)"` → `"it was the"`

#### `(cap, N)` - Capitalize N Words
Capitalizes the N previous words.

**Example:**
- `"age of foolishness (cap, 3)"` → `"Age Of Foolishness"`

---

### 3.4 Punctuation Rules

#### Basic Punctuation: `. , ! ? : ;`
Sticks to the previous word (no space before), space after.

**Examples:**
- `"there ,and then"` → `"there, and then"`
- `"Hello !"` → `"Hello!"`

#### Punctuation Groups: `...` `!?` etc.
Groups of punctuation marks stay together without internal spaces.

**Examples:**
- `"thinking . . ."` → `"thinking..."`
- `"Really ! ?"` → `"Really!?"`

---

### 3.5 Quote Handling: `'`

#### Single Word
Quotes stick to the left and right of the word.

**Example:**
- `"I am: ' awesome '"` → `"I am: 'awesome'"`

#### Multiple Words
Quotes stick to the first and last word.

**Example:**
- `"' I am the best '"` → `"'I am the best'"`

---

### 3.6 Article Correction: a → an

The article "a" becomes "an" if the next word starts with a vowel (a, e, i, o, u) or 'h'.

**Examples:**
- `"a amazing"` → `"an amazing"`
- `"a honest"` → `"an honest"`
- `"a book"` → `"a book"` (no change)

---

<a name="golden-test-set"></a>
## <span style="color: #CCFF99;">4. Golden Test Set</span>

This section contains comprehensive test cases to validate the text editor implementation.

### Testing Strategy

1. **Isolation Tests** (Tests 1-6): Validate individual rules
2. **Integration Tests** (Tests 7-11): Validate rule combinations
3. **Comprehensive Test** (Test 12): Validate real-world usage with multiple rules interacting
4. **Edge Cases**: Focus on boundaries (zero values, empty modifiers, consecutive punctuation)
5. **Context Sensitivity**: Validate context awareness (quotes, modifiers with numbers, a/an before modified words)

---

## Basic Functional Tests

### <span style="color: #FFD700;">Test 1: Mixed Case Transformations</span>

**Input:**
```
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6)
```

**Expected Output:**
```
It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness
```

**Covers:** `(cap)`, `(up)`, `(cap, N)` with punctuation

---

### <span style="color: #FFD700;">Test 2: Hexadecimal and Binary Conversions</span>

**Input:**
```
Simply add 42 (hex) and 10 (bin) and you will see the result is 68.
```

**Expected Output:**
```
Simply add 66 and 2 and you will see the result is 68.
```

**Covers:** `(hex)`, `(bin)` with punctuation

---

### <span style="color: #FFD700;">Test 3: A to An Correction</span>

**Input:**
```
There is no greater agony than bearing a untold story inside you.
```

**Expected Output:**
```
There is no greater agony than bearing an untold story inside you.
```

**Covers:** a/an rule

---

### <span style="color: #FFD700;">Test 4: Punctuation Spacing</span>

**Input:**
```
Punctuation tests are ... kinda boring ,what do you think ?
```

**Expected Output:**
```
Punctuation tests are... kinda boring, what do you think?
```

**Covers:** Punctuation spacing, groups of punctuation

---

### <span style="color: #FFD700;">Test 5: Quote Handling (Single Word)</span>

**Input:**
```
I am exactly how they describe me: ' awesome '
```

**Expected Output:**
```
I am exactly how they describe me: 'awesome'
```

**Covers:** Single word between quotes

---

### <span style="color: #FFD700;">Test 6: Quote Handling (Multiple Words)</span>

**Input:**
```
As Elton John said: ' I am the most well-known homosexual in the world '
```

**Expected Output:**
```
As Elton John said: 'I am the most well-known homosexual in the world'
```

**Covers:** Multiple words between quotes

---

## Advanced Test Cases (Tricky Scenarios)

### <span style="color: #FFD700;">Test 7: Multiple Modifiers in Sequence</span>

**Input:**
```
the word (up) was then (cap) followed by another (low) transformation and FF (hex) things.
```

**Expected Output:**
```
the WORD was then Followed by another transformation and 255 things.
```

**Covers:** Multiple different modifiers, hex at end

---

### <span style="color: #FFD700;">Test 8: Modifier with Number Affecting Multiple Words + Punctuation</span>

**Input:**
```
it was the BEST OF TIMES (low, 3) ! what a story .
```

**Expected Output:**
```
it was the best of times! what a story.
```

**Covers:** `(low, N)` affecting capitals, multiple punctuation marks

---

### <span style="color: #FFD700;">Test 9: Edge Case - A/An with Punctuation and Quotes</span>

**Input:**
```
She found a apple , a orange and a ' honest ' person .
```

**Expected Output:**
```
She found an apple, an orange and an 'honest' person.
```

**Covers:** Multiple a/an corrections, quotes with single word, punctuation

---

### <span style="color: #FFD700;">Test 10: Complex Punctuation Groups</span>

**Input:**
```
Wait ... what ! ? Really ? ! I can not believe it ! ! !
```

**Expected Output:**
```
Wait... what!? Really?! I can not believe it!!!
```

**Covers:** Multiple punctuation groups, mixed punctuation

---

### <span style="color: #FFD700;">Test 11: Binary/Hex Edge Cases</span>

**Input:**
```
The values 0 (bin) and 0 (hex) are equal , but 1111 (bin) equals F (hex) which is 15 (hex) .
```

**Expected Output:**
```
The values 0 and 0 are equal, but 15 equals 15 which is 15.
```

**Covers:** Zero values, same result from different bases, multiple conversions

---

## Comprehensive Test

### <span style="color: #FFD700;">Test 12: Large Text with Multiple Rules</span>

**Input:**
```
it (cap) was a amazing DAY (low) ! the sun was shining and the temperature reached 1F (hex) degrees . I went to the store , bought 11 (bin) apples and A (up) orange . the shopkeeper said : ' you are a honest customer ' . when i got HOME (low, 2) , i realized that 101 (bin) plus A (hex) equals F (hex) ! what a DISCOVERY (cap) ... i could not BELIEVE IT (low, 2) ! ? this was the best day EVER (cap, 2) .
```

**Expected Output:**
```
It was an amazing day! the sun was shining and the temperature reached 31 degrees. I went to the store, bought 3 apples and AN orange. the shopkeeper said: 'you are an honest customer'. when i got home, i realized that 5 plus 10 equals 15! what a Discovery... i could not believe it!? this was The Best Day Ever.
```

**Covers:**
- `(cap)` at start
- a→an with `(low)` transformed word
- `(hex)` conversion (1F → 31)
- `(bin)` conversion (11 → 3)
- `(up)` on article
- Quote with multiple words
- `(low, N)` affecting previous words
- Multiple a→an corrections (with h vowel)
- Complex punctuation (... and !?)
- `(cap, N)` at end of sentence
- Mix of all rules in complex text

---

<span style="color: #ff0000ff;">## Edge Cases to Consider</span>

### Additional Test Scenarios

**Test 13: Empty Modifiers**
```
Input:  "word () text"
Output: "word () text"  (no change, invalid modifier)
```

**Test 14: Modifier Without Previous Word**
```
Input:  "(cap) hello world"
Output: "(cap) hello world"  (no word to modify)
```

**Test 15: Nested Quotes**
```
Input:  "he said ' she said ' hello ' '"
Output: "he said 'she said' hello '"  (handle multiple quote pairs)
```

**Test 16: Numbers in Words**
```
Input:  "abc123 (hex) value"
Output: "abc123 (hex) value"  (invalid hex, no change)
```

**Test 17: Multiple Spaces**
```
Input:  "word     another"
Output: "word another"  (normalize to single space)
```

---

## Test Execution Guide

### Running Tests

**Unit Tests:**
```bash
go test ./tests/transforms_test.go
go test ./tests/formatters_test.go
```

**Integration Tests:**
```bash
go test ./tests/integration_test.go
```

**Full Test Suite:**
```bash
go test ./...
```

**With Coverage:**
```bash
go test -cover ./...
```

### Expected Test Results

All tests should pass with:
- 0 failures
- 100% rule coverage
- Correct handling of edge cases
- Proper context management

---

## Conclusion

This analysis document provides:
1. Clear problem understanding
2. Justified architecture choice (FSM)
3. Comprehensive test suite
4. Implementation guidelines

The FSM architecture is the optimal choice for this text processing task due to its efficiency, natural fit for parsing problems, and industry-standard approach.

## Author

Constantine E.P.

---

## License

Educational project for learning purposes.