# Go-Reloaded v1.0.0 Release Notes

## Overview
A sophisticated text processing tool built with FSM (Finite State Machine) architecture that applies transformations, formatting rules, and intelligent text corrections.

## Features

### Core Transformations
- **Number Base Conversions**: Hexadecimal and binary to decimal
- **Case Transformations**: Uppercase, lowercase, capitalize
- **Batch Operations**: Apply transformations to N previous words
- **Article Correction**: Automatic a/an correction before vowels and silent 'h'

### Advanced Features
- **Quote Support**: Both single `'` and double `"` quotes with type preservation
- **Special Word Handling**: 
  - Contractions: don't, it's, let's
  - Hyphenated words: well-known, state-of-the-art
  - Slash compounds: a/an, and/or
- **Unicode Support**: Handles accented characters (café, naïve, résumé)
- **Smart Punctuation**: Automatic spacing and grouping
- **Newline Preservation**: Maintains original document structure
- **Punctuation Boundaries**: Semantic boundaries for modifier scope

### Architecture
- **FSM-based**: Single-pass O(n) processing
- **Memory Efficient**: No intermediate string copies
- **Context-Aware**: Tracks state for intelligent transformations
- **Modular Design**: Separated concerns (FSM, transforms, formatters)

## Test Coverage
- ✅ 100+ unit tests
- ✅ Integration tests
- ✅ Golden test suite
- ✅ Edge case tests
- ✅ 30+ stress test scenarios
- ✅ All tests passing

## Known Limitations
- Punctuation marks (`. , ! ? : ;`) act as boundaries for batch modifiers
- Modifiers only affect words in the current buffer (after last punctuation)
- Invalid hex/binary strings remain unchanged

## Usage

```bash
# Basic usage
go run . input.txt output.txt

# Run tests
go test ./...

# Run stress test
go run . stress_test.txt output.txt
```

## Files Included
- Source code in `fsm/`, `transforms/`, `formatters/`
- Comprehensive tests in `tests/`
- Documentation in `docs/`
- Sample inputs: `test.txt`, `input2.txt`, `randomInput.txt`, `stress_test.txt`
- Task breakdown in `tasks/`

## Author
Constantine E.P.

## License
MIT License - See LICENSE file

## Acknowledgments
Built as an educational project demonstrating FSM architecture, Go best practices, and comprehensive testing strategies.
