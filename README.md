# Text Editor Project

## Overview

A text processing tool that reads an input file, applies transformations and formatting rules, and writes the result to an output file using **FSM (Finite State Machine) architecture**.

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

- **Number base conversions**: `(hex)`, `(bin)`
- **Case transformations**: `(up)`, `(low)`, `(cap)`
- **Batch operations**: `(up, N)` - apply to N previous words
- **Smart punctuation**: Automatic spacing and grouping
- **Quote handling**: Single and multiple word quotes
- **Article correction**: `a` → `an` before vowels/h

---

## Project Structure

```
.
├── README.md              # Project overview (this file)
├── docs/
│   └── ANALYSIS.md        # Architecture analysis & test cases
├── assets/
│   └── fsm flow diagram.png
├── main.go                # Entry point
├── fsm/                   # FSM state machine (orchestration)
│   ├── fsm.go
│   ├── states.go
│   └── transitions.go
├── transforms/            # Transformation logic (pure functions)
│   ├── numbers.go         # hex/bin conversions
│   ├── cases.go           # case transformations
│   └── article.go         # a/an correction
├── formatters/            # Formatting logic (pure functions)
│   ├── punctuation.go
│   └── quotes.go
└── tests/                 # Unit & integration tests
    └── *_test.go
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

# Run with coverage
go test -cover ./...

# Run specific test suite
go test ./tests/transforms_test.go
```

For detailed test cases, see `docs/ANALYSIS.md`.

---

## Architecture Highlight

This project uses **FSM (Finite State Machine)** architecture for:
- ✅ Single-pass processing (O(n) efficiency)
- ✅ Context-aware transformations
- ✅ Memory efficiency
- ✅ Industry-standard approach for text parsing

For detailed architecture analysis, see `docs/ANALYSIS.md`.

---

## Author

Text Editor Project - FSM Implementation

---

## License

Educational project for learning purposes.