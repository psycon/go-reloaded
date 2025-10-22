# Text Editor Project

## Overview

This project implements a text processing tool that reads an input file, applies a series of transformations and formatting rules, and writes the result to an output file.

The program recognizes special modifiers within the text (e.g., `(hex)`, `(up)`, `(cap)`) and applies the corresponding transformations to previous words. Additionally, it automatically corrects punctuation, spacing around punctuation marks, and handles special cases such as quotes and article correction (a/an).

---

## Features

✅ **Number Base Conversions**
- Hexadecimal to Decimal: `1E (hex)` → `30`
- Binary to Decimal: `10 (bin)` → `2`

✅ **Case Transformations**
- Uppercase: `go (up)` → `GO`
- Lowercase: `SHOUTING (low)` → `shouting`
- Capitalize: `bridge (cap)` → `Bridge`

✅ **Batch Transformations**
- Apply to N words: `exciting (up, 2)` → `SO EXCITING`

✅ **Smart Punctuation**
- Basic spacing: `there ,and` → `there, and`
- Punctuation groups: `. . .` → `...`

✅ **Quote Handling**
- Single word: `' awesome '` → `'awesome'`
- Multiple words: `' I am great '` → `'I am great'`

✅ **Article Correction**
- a → an: `a amazing` → `an amazing`
- Before h: `a honest` → `an honest`

---

## Usage

```bash
go run . input.txt output.txt
```

### Example

**Input file (sample.txt):**
```
it (cap) was the best of times, it was the worst of times (up)
```

**Command:**
```bash
go run . sample.txt result.txt
```

**Output file (result.txt):**
```
It was the best of times, it was the worst of TIMES
```

---

## Transformation Rules

### 1. Number Base Conversions

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

### 2. Case Transformations

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

### 3. Batch Transformations

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

### 4. Punctuation Rules

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

### 5. Quote Handling: `'`

#### Single Word
Quotes stick to the left and right of the word.

**Example:**
- `"I am: ' awesome '"` → `"I am: 'awesome'"`

#### Multiple Words
Quotes stick to the first and last word.

**Example:**
- `"' I am the best '"` → `"'I am the best'"`

---

### 6. Article Correction: a → an

The article "a" becomes "an" if the next word starts with a vowel (a, e, i, o, u) or 'h'.

**Examples:**
- `"a amazing"` → `"an amazing"`
- `"a honest"` → `"an honest"`
- `"a book"` → `"a book"` (no change)

---

## Testing

The project includes comprehensive test cases covering:
- Individual rule validation
- Rule combination scenarios
- Edge cases and tricky scenarios
- Real-world text with multiple rules

For detailed test cases and expected outputs, see the **Golden Test Set** in `ANALYSIS.md`.

---

## Project Structure

```
.
├── README.md           # Project overview (this file)
├── ANALYSIS.md         # Architecture analysis and test cases
├── main.go             # Entry point
├── fsm/                # FSM implementation
├── transforms/         # Transformation logic
├── formatters/         # Formatting rules
└── tests/              # Unit tests
```

---

## Documentation

- **README.md** (this file): Project overview and usage guide
- **ANALYSIS.md**: Detailed architecture analysis, design decisions, and comprehensive test suite

---

## Author

Text Editor Project - FSM Architecture Implementation

---

## License

Educational project for learning purposes.