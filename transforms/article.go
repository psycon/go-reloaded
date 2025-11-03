package transforms

import (
	"strings"
	"unicode"
)

// FixArticle converts "a" to "an" if next word starts with vowel or h
func FixArticle(word, nextWord string) string {
	if strings.ToLower(word) != "a" {
		return word
	}

	if len(nextWord) == 0 {
		return word
	}

	firstChar := unicode.ToLower([]rune(nextWord)[0])

	// Check if starts with vowel or h
	if isVowel(firstChar) || firstChar == 'h' {
		if word == "a" {
			return "an"
		}
		if word == "A" {
			return "An"
		}
	}

	return word
}

func isVowel(ch rune) bool {
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
}
