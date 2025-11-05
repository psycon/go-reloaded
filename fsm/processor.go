package fsm

import (
	"fmt"
	"go-reloaded/formatters"
	"go-reloaded/transforms"
	"strings"
)

type Processor struct {
	//It holds the input data
	tokens []string
	//tracks its progress
	pos int
	//builds the result
	output strings.Builder
	//temp buffer for words
	wordBuffer []string
	//flags T if ' found
	inQuote bool
	//temp buffer for quoted words
	quoteWords []string
}

func NewProcessor() *Processor {
	return &Processor{
		wordBuffer: make([]string, 0),
		quoteWords: make([]string, 0),
	}
}

func (p *Processor) Process(input string) string {
	// RESET STATE - IMPORTANT!
	p.tokens = nil
	p.pos = 0
	p.output.Reset() // Clear previous output
	p.wordBuffer = make([]string, 0)
	p.inQuote = false
	p.quoteWords = make([]string, 0)

	// Now tokenize and process
	p.tokens = tokenize(input)

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
	targetBuffer := &p.wordBuffer
	if p.inQuote {
		targetBuffer = &p.quoteWords
	}

	if len(*targetBuffer) == 0 {
		// No preceding word, treat modifier as regular text
		if p.inQuote {
			p.quoteWords = append(p.quoteWords, modifier)
		} else if len(p.wordBuffer) == 0 {
			p.wordBuffer = append(p.wordBuffer, modifier)
		} else {
			p.wordBuffer = append(p.wordBuffer, modifier)
		}
		return
	}

	modType, count := parseModifier(modifier)

	switch modType {
	case "hex":
		idx := len(*targetBuffer) - 1
		(*targetBuffer)[idx] = transforms.HexToDec((*targetBuffer)[idx])
	case "bin":
		idx := len(*targetBuffer) - 1
		(*targetBuffer)[idx] = transforms.BinToDec((*targetBuffer)[idx])
	case "up":
		p.applyCase(transforms.ToUpper, count, targetBuffer)
	case "low":
		p.applyCase(transforms.ToLower, count, targetBuffer)
	case "cap":
		p.applyCase(transforms.Capitalize, count, targetBuffer)
	}
}

func (p *Processor) applyCase(fn func(string) string, count int, buffer *[]string) {
	if count == 0 {
		count = 1
	}

	start := len(*buffer) - count
	if start < 0 {
		start = 0
	}

	for i := start; i < len(*buffer); i++ {
		(*buffer)[i] = fn((*buffer)[i])
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

	// Add punctuation (sticks to previous word, space after)
	p.output.WriteString(formatters.FormatPunctuation(group))
}

func (p *Processor) handleQuote() {
	if !p.inQuote {
		// If there are words in the buffer before an opening quote,
		// and the last word is punctuation, it's a syntax error in input.
		// But we can handle it gracefully by flushing.
		if len(p.wordBuffer) > 0 && isPunctuation(p.wordBuffer[len(p.wordBuffer)-1]) {
			p.flushBuffer()
		}
		// Opening quote
		p.flushBuffer()
		p.inQuote = true
		p.quoteWords = make([]string, 0)
	} else {
		// Closing quote
		// APPLY ARTICLE CORRECTION TO QUOTED WORDS
		for i := 0; i < len(p.quoteWords); i++ {
			var nextWord string
			if i < len(p.quoteWords)-1 {
				nextWord = p.quoteWords[i+1]
			}
			p.quoteWords[i] = transforms.FixArticle(p.quoteWords[i], nextWord)
		}

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
			var nextWord string
			// Look ahead for the next actual word, skipping over any invalid modifiers treated as text
			for j := i + 1; j < len(p.wordBuffer); j++ {
				nextWord = p.wordBuffer[j]
				break
			}
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
	inParens := false

	for _, ch := range input {
		switch ch {
		case '(':
			current.WriteRune(ch)
			inParens = true
		case ')':
			current.WriteRune(ch)
			inParens = false
			// Complete the token and add it
			if current.Len() > 0 {
				tokens = append(tokens, current.String())
				current.Reset()
			}
		case ' ', '\t', '\n', '\r':
			if inParens {
				// Keep spaces inside parentheses
				current.WriteRune(ch)
			} else {
				// Space outside parentheses - end token
				if current.Len() > 0 {
					tokens = append(tokens, current.String())
					current.Reset()
				}
			}
		case '\'':
			if current.Len() > 0 {
				tokens = append(tokens, current.String())
				current.Reset()
			}
			tokens = append(tokens, string(ch))
		case '.', ',', '!', '?', ':', ';':
			if inParens {
				// Keep punctuation inside parentheses
				current.WriteRune(ch)
			} else {
				// Punctuation outside parentheses - separate token
				if current.Len() > 0 {
					tokens = append(tokens, current.String())
					current.Reset()
				}
				tokens = append(tokens, string(ch))
			}
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
	count := 1 // Default to 1 if no count is specified

	if len(parts) > 1 {
		fmt.Sscanf(strings.TrimSpace(parts[1]), "%d", &count)
	} else if modType == "hex" || modType == "bin" {
		count = 1 // These never have a count, ensure it's 1
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
