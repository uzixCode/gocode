package changecase

import (
	"strings"
	"unicode"
)

// ToLower converts a string to lowercase.
func ToLower(s string) string {
	return strings.ToLower(s)
}

// ToUpper converts a string to uppercase.
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// ToTitle converts a string to title case (capitalizes the first letter of each word).
func ToTitle(s string) string {
	return strings.Title(s)
}

// ToCamel converts a string to camelCase.
func ToCamel(s string) string {
	words := splitWords(s)
	for i, word := range words {
		if i == 0 {
			words[i] = strings.ToLower(word)
		} else {
			words[i] = strings.Title(word)
		}
	}
	return strings.Join(words, "")
}

// ToSnake converts a string to snake_case.
func ToSnake(s string) string {
	words := splitWords(s)
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}
	return strings.Join(words, "_")
}

// ToKebab converts a string to kebab-case.
func ToKebab(s string) string {
	words := splitWords(s)
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}
	return strings.Join(words, "-")
}

// ToPascal converts a string to PascalCase.
func ToPascal(s string) string {
	words := splitWords(s)
	for i, word := range words {
		words[i] = strings.Title(word)
	}
	return strings.Join(words, "")
}

// ToConstant converts a string to CONSTANT_CASE.
func ToConstant(s string) string {
	words := splitWords(s)
	for i := range words {
		words[i] = strings.ToUpper(words[i])
	}
	return strings.Join(words, "_")
}

// ToDot converts a string to dot.case.
func ToDot(s string) string {
	words := splitWords(s)
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}
	return strings.Join(words, ".")
}

// Helper function to split a string into words.
func splitWords(s string) []string {
	var words []string
	var currentWord []rune

	for i, r := range s {
		// If it's a space or non-alphabetic character, it's the boundary for a word
		if unicode.IsSpace(r) || (!unicode.IsLetter(r) && !unicode.IsDigit(r)) {
			if len(currentWord) > 0 {
				words = append(words, string(currentWord))
				currentWord = []rune{}
			}
		} else {
			if i == 0 || unicode.IsUpper(r) && len(currentWord) > 0 {
				// start of a new word or new capitalized word
				if len(currentWord) > 0 {
					words = append(words, string(currentWord))
				}
				currentWord = []rune{r}
			} else {
				currentWord = append(currentWord, r)
			}
		}
	}

	if len(currentWord) > 0 {
		words = append(words, string(currentWord))
	}

	return words
}
