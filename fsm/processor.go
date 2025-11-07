package fsm

import (
	"fmt"
	"go-reloaded/formatters"
	"go-reloaded/transforms"
	"regexp"
	"strconv"
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
	quoteWords           []string
	lastProcessedWasWord bool // Tracks if the last token processed was a word (not punctuation, modifier, or quote)
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
	p.lastProcessedWasWord = false // Reset state for new input

	// Now tokenize and process
	p.tokens = tokenize(input)

	for p.pos < len(p.tokens) {
		token := p.tokens[p.pos]
		trimmedToken := strings.TrimSpace(token)

		// Skip empty tokens after trimming
		if trimmedToken == "" {
			p.pos++
			continue
		}

		// Handle quotes
		if trimmedToken == "'" {
			p.flushBuffer() // Flush any pending words before handling quote
			p.handleQuote()
			p.pos++
			// Don't override lastProcessedWasWord here - handleQuote sets it correctly
			continue
		}

		// Check for modifiers (only if last token was a word)
		if isModifier(trimmedToken) {
			targetBuffer := &p.wordBuffer
			if p.inQuote {
				targetBuffer = &p.quoteWords
			}
			// Only apply modifier if:
			// 1. There's a word in the buffer AND
			// 2. The last processed token was a word (not another modifier)
			if len(*targetBuffer) > 0 && p.lastProcessedWasWord {
				p.handleModifier(trimmedToken)
				p.pos++
				p.lastProcessedWasWord = false
				continue
			}
			// Otherwise, treat it as regular text (fall through)
		}
		// Check for punctuation
		if isPunctuation(trimmedToken) {
			p.handlePunctuation()
			continue // handlePunctuation advances pos
		}

		// Regular word
		if p.inQuote {
			p.quoteWords = append(p.quoteWords, trimmedToken)
		} else {
			p.wordBuffer = append(p.wordBuffer, trimmedToken) // Add to buffer
		}
		p.lastProcessedWasWord = true // A word was just processed

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

		modType, count := parseModifier(modifier)

	

		switch modType {

		case "hex":

			idx := len(*targetBuffer) - 1

			if idx >= 0 {

				(*targetBuffer)[idx] = transforms.HexToDec((*targetBuffer)[idx])

			}

		case "bin":

			idx := len(*targetBuffer) - 1

			if idx >= 0 {

				(*targetBuffer)[idx] = transforms.BinToDec((*targetBuffer)[idx])

			}

		case "up":

			p.applyCase(transforms.ToUpper, count, targetBuffer)

		case "low":

			p.applyCase(transforms.ToLower, count, targetBuffer)

		case "cap":

			p.applyCase(transforms.Capitalize, count, targetBuffer)

		}

		p.lastProcessedWasWord = false // A modifier was just processed
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

	if p.inQuote {
		// If inside a quote, attach punctuation to the last word.
		if len(p.quoteWords) > 0 {
			lastIndex := len(p.quoteWords) - 1
			p.quoteWords[lastIndex] += group
		} else {
			p.quoteWords = append(p.quoteWords, group)
		}
		return
	}

	// Flush words before punctuation
	p.flushBuffer()

	// Trim trailing space from output before adding punctuation
	if p.output.Len() > 0 && p.output.String()[p.output.Len()-1] == ' ' {
		currentOutput := p.output.String()
		p.output.Reset()
		p.output.WriteString(currentOutput[:len(currentOutput)-1])
	}

	// Add punctuation (sticks to previous word, space after)
	p.lastProcessedWasWord = false // Punctuation was just processed
	p.output.WriteString(formatters.FormatPunctuation(group))
	p.output.WriteString(" ") // Add space after punctuation
}
func (p *Processor) handleQuote() {
	if !p.inQuote {
		p.flushBuffer()
		p.inQuote = true
		p.lastProcessedWasWord = false
		p.quoteWords = make([]string, 0)
	} else {
		// Apply a/an transformation inside quotes before formatting
		for i := 0; i < len(p.quoteWords)-1; i++ {
			p.quoteWords[i] = transforms.FixArticle(p.quoteWords[i], p.quoteWords[i+1])
		}

		quoted := formatters.FormatQuote(p.quoteWords)

		// Add quoted text to word buffer instead of directly to output
		// This allows modifiers after the quote to transform it
		p.wordBuffer = append(p.wordBuffer, quoted)

		p.inQuote = false
		p.lastProcessedWasWord = true // Treat quoted text as a word
		p.quoteWords = make([]string, 0)
	}
}

func (p *Processor) flushBuffer() {
	for i := 0; i < len(p.wordBuffer); i++ {
		word := p.wordBuffer[i]
		nextWord := ""

		// Check a/an rule
		if i < len(p.wordBuffer)-1 { // If there's a next word in the current buffer
			nextWord = p.wordBuffer[i+1]
		} else { // This is the last word in wordBuffer, need to peek ahead in p.tokens
			// Find the next actual word in the main token stream
			for j := p.pos; j < len(p.tokens); j++ {
				potentialNextToken := p.tokens[j]
				if !isModifier(potentialNextToken) && !isPunctuation(potentialNextToken) && potentialNextToken != "'" && potentialNextToken != "" {
					nextWord = potentialNextToken
					break
				}
			}
		}

		// Apply FixArticle
		if nextWord != "" {
			word = transforms.FixArticle(word, nextWord)
		}

		if p.output.Len() > 0 && !endsWithSpace(p.output.String()) {
			// Ensure there's a space between words, unless it's punctuation
			p.output.WriteString(" ")
		}
		p.output.WriteString(word)
	}

	p.wordBuffer = make([]string, 0)
}

// tokenize function with contraction and hyphenated word support
func tokenize(input string) []string {
	// Updated regex to handle:
	// - Contractions: "Let's", "It's", "don't"
	// - Hyphenated words: "well-known", "state-of-the-art"
	// - Slash compounds: "a/an", "and/or"
	// - Modifiers: (hex), (up, 2)
	// - Punctuation: . , ! ? : ;
	// - Quotes: '
	re := regexp.MustCompile(`(\w+(?:[-'/]\w+)*|[.,!?:;']|\(\s*\w+\s*(?:,\s*\d+\s*)?\))`)
	matches := re.FindAllString(input, -1)

	var tokens []string
	for _, match := range matches {
		trimmedMatch := strings.TrimSpace(match)
		if trimmedMatch != "" {
			tokens = append(tokens, trimmedMatch)
		}
	}
	return tokens
}
func isLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
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

	modType := strings.TrimSpace(parts[0]) // Default to 1 if no count is specified
	count := 1

	if len(parts) > 1 {
		countStr := strings.TrimSpace(parts[1])
		if countStr == "" {
			// If the count string is empty, default to 1 and log a warning
			fmt.Printf("Warning: Empty count in modifier '%s', defaulting to 1\n", token)
			count = 1
		} else {
			var err error
			count, err = strconv.Atoi(countStr)
			if err != nil {
				// Log the error if strconv.Atoi fails, then default to 1
				fmt.Printf("Error parsing count '%s' in modifier '%s': %v, defaulting to 1\n", countStr, token, err)
				count = 1
			}
		}
	} else if modType != "hex" && modType != "bin" && len(parts) == 1 {
		// This is for single-word modifiers like (up), (low), (cap)
	}

	return modType, count
}

// isPunctuation checks if a string consists entirely of punctuation characters.
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
