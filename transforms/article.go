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
	shouldBeAn := false

	// Special cases for 'h' and 'u'
	switch firstChar {
	case 'h':
		if strings.HasPrefix(strings.ToLower(nextWord), "hour") ||
			strings.HasPrefix(strings.ToLower(nextWord), "honest") ||
			strings.HasPrefix(strings.ToLower(nextWord), "honor") ||
			strings.HasPrefix(strings.ToLower(nextWord), "heir") {
			shouldBeAn = true
		}
	case 'u':
		if strings.HasPrefix(strings.ToLower(nextWord), "uni") ||
			strings.HasPrefix(strings.ToLower(nextWord), "eura") ||
			strings.HasPrefix(strings.ToLower(nextWord), "use") {
			shouldBeAn = false
		} else {
			shouldBeAn = true
		}
	default:
		if isVowel(firstChar) {
			shouldBeAn = true
		}
	}

	if shouldBeAn {
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
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' ||
		ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U'
}
