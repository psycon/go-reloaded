<span style="color: #FF0000;"># Auditor/Educational focused Document - GO-reloaded Project</span>

<span style="color: #CCFF99;">## Overview</span>

This project implements a text processing tool that reads an input file, applies transformations and formatting rules, and writes the result to an output file using **FSM (Finite State Machine) architecture**.
The program recognizes special modifiers within the text and applies the corresponding transformations to previous words. Additionally, it automatically corrects punctuation, spacing around punctuation marks, and handles special cases such as quotes and article correction (a/an).

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
├───AGENTS.md
├───AUTHORS.md
├───go.mod
├───main.go
├───README.md
├───test.txt
├───assets/
│   ├───.gitkeep
│   └───fsm flow diagram.png
├───docs/
│   └───ANALYSIS.md
├───formatters/
│   ├───.gitkeep
│   ├───punctuation.go
│   └───quotes.go
├───fsm/
│   ├───.gitkeep
│   └───processor.go
├───tasks/
│   ├───.gitkeep
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
│   ├───.gitkeep
│   ├───documentation_test.go
│   ├───fsm_test.go
│   ├───formatters_test.go
│   ├───golden_test.go
│   ├───integration_test.go
│   ├───main_test.go
│   └───transforms_test.go
└───transforms/
    ├───.gitkeep
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
go run . stress_test.txt stress_output.txt

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

This project uses **FSM (Finite State Machine)** architecture for:
-  Single-pass processing (O(n) efficiency)
-  Context-aware transformations
-  Memory efficiency
-  Industry-standard approach for text parsing

For detailed architecture analysis, see `docs/ANALYSIS.md`.

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