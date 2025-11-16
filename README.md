<span style="color: #FF0000;"># Auditor/Educational focused Document - GO-reloaded Project</span>

<span style="color: #CCFF99;">## Overview</span>

Go-Reloaded is an intelligent text processing tool that transforms and formats text using a **hybrid FSM-orchestrated pipeline architecture**. The program reads an input file, processes it through a single-pass state machine that routes tokens to pure transformation functions, and writes the formatted result to an output file.

**What it does:**
- Applies modifiers like `(up)`, `(hex)`, `(cap)` to transform preceding words
- Automatically corrects grammar (a → an before vowels)
- Formats punctuation and quotes with proper spacing
- Handles special cases: contractions, hyphenated words, Unicode characters
- Preserves document structure (newlines, quote types)

**How it works:**
The FSM controller maintains context (quote state, word history) while routing tokens to stateless transformation functions. Punctuation marks act as semantic boundaries, ensuring modifiers only affect their intended scope. This hybrid approach combines the context awareness of state machines with the modularity of functional pipelines.

---

## Quick Start

```bash
go run . input.txt output.txt
```

### Example

**Input (sample.txt):**
```
it (cap) was the best of times, it was the worst of times (up)
```

**Command:**
```bash
go run . sample.txt result.txt
```

**Output (result.txt):**
```
It was the best of times, it was the worst of TIMES
```

---

## Key Features

- **Number base conversions**: `(hex)`, `(bin)` - Convert hexadecimal and binary to decimal
- **Case transformations**: `(up)`, `(low)`, `(cap)` - Uppercase, lowercase, capitalize
- **Batch operations**: `(up, N)` - Apply transformations to N previous words
- **Smart punctuation**: Automatic spacing and grouping (`. , ! ? : ;`)
- **Quote handling**: Both single `'` and double `"` quotes with modifier support
- **Article correction**: `a` → `an` before vowels and silent 'h'
- **Special word support**: Contractions (don't, it's), hyphenated (well-known), slash compounds (a/an)
- **Newline preservation**: Maintains original line structure
- **Punctuation boundaries**: Modifiers respect punctuation as semantic boundaries

---

## Project Structure

```
/go-reloaded/
├───.gitignore
├───go.mod
├───LICENSE
├───main.go
├───README.md
├───assets/
│   └───fsm flow diagram.png
├───audit/
├───docs/
│   ├───AGENTS.md
│   ├───ANALYSIS.md
│   ├───ARCHITECTURE_DIAGRAM.md
│   ├───AUTHORS.md
│   └───gh-pages/
│       └───index.html
├───formatters/
│   ├───punctuation.go
│   └───quotes.go
├───fsm/
│   └───processor.go
├───tasks/
│   ├───TASK-01.md
│   ├───TASK-02.md
│   ├───TASK-03.md
│   ├───TASK-04.md
│   ├───TASK-05.md
│   ├───TASK-06.md
│   ├───TASK-07.md
│   ├───TASK-08.md
│   └───TASK-09-10.md
├───tests/
│   ├───stress_test.txt
│   ├───test.txt
│   ├───formatters_test.go
│   ├───fsm_test.go
│   ├───golden_test.go
│   ├───integration_test.go
│   ├───main_test.go
│   └───transforms_test.go
└───transforms/
    ├───article.go
    ├───cases.go
    └───numbers.go
```

---

## Documentation

### 📘 **README.md** (this file)
Quick overview and usage guide.

### 📗 **docs/ANALYSIS.md**
Comprehensive documentation including:
- Problem description
- Architecture comparison (Pipeline vs FSM)
- Design decisions and justification
- Complete test suite (Golden Test Set)
- Implementation guidelines

---

## Testing

```bash
# Run all tests
go test ./...

# Run with verbose output
go test -v ./tests

# Run specific test
go test -v ./tests -run TestSpecificName

# Run stress test
go run . tests/stress_test.txt tests/stress_output.txt

# Build executable
go build -o go-reloaded .
```

**Test Coverage:**
- ✅ 100+ unit tests
- ✅ Integration tests
- ✅ Golden test suite
- ✅ Edge case tests
- ✅ Stress tests with 30+ scenarios

For detailed test cases, see `docs/ANALYSIS.md`.

---

## Architecture Highlight

This project uses a **hybrid FSM-orchestrated pipeline architecture**:

**FSM Controller** (State Management):
-  Single-pass token processing
-  Context tracking (quotes, modifiers, boundaries)
-  State-based routing decisions

**Transformation Pipeline** (Pure Functions):
-  Modular, testable transformations
-  Reusable components
-  Clean separation of concerns

**Benefits:**
-  O(n) time complexity
-  Memory efficient (no intermediate copies)
-  Maintainable and extensible
-  Industry-standard patterns (lexer + transformer)

For detailed architecture analysis and diagrams, see:
- `docs/ANALYSIS.md` - Architecture comparison and design decisions
- `docs/ARCHITECTURE_DIAGRAM.md` - Mermaid flowchart and component breakdown

---

## Author

Constantine E.P.

---

## Advanced Features

### Batch Modifiers with Quotes
Batch modifiers count words across quotes:
```
Input:  one two ' three four ' (up, 4)
Output: ONE TWO 'THREE FOUR'
```

### Punctuation as Boundaries
Punctuation marks (`. , ! ? : ;`) act as semantic boundaries:
```
Input:  one two, three four (up, 3)
Output: one two, THREE FOUR
```
Only words after the comma are affected.

### Quote Type Preservation
Single and double quotes are preserved:
```
Input:  ' single ' and " double "
Output: 'single' and "double"
```

---

## License

MIT License - See LICENSE file for details.