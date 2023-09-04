package casenames

import (
	"unicode"
)

func caseNames(input string) []string {
	var words []string
	start := 0
	length := 0

	for i, char := range input {
		if char == ' ' || char == '_' {
			// Treat spaces and underscores as word separators.
			if length != 0 {
				words = append(words, input[start:start+length])
			}
			length = 0
			start = i + 1
		} else if i > 0 && unicode.IsUpper(rune(input[i])) && !unicode.IsUpper(rune(input[i-1])) {
			// catches pascal and camel case beginnings
			if length != 0 {
				words = append(words, input[start:start+length])
			}

			start = i
			length = 1
		} else if unicode.IsUpper(rune(input[i])) && unicode.IsLower(rune(input[i+1])) && i < len(input) {
			// ignores ABBR (abbreviations) from splitting into multiple single letters
			words = append(words, input[start:start+length])
			start = i
			length = 1
		} else {
			length++
		}
	}

	// Append the last word.
	if length != 0 {
		words = append(words, input[start:start+length])
	}

	return words
}
