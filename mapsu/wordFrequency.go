package main

import (
	"strings"
	"unicode"
)

func WordFrequency(text string) map[string]int {
	words := strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})
	freq := make(map[string]int)

	for _, word := range words {
		lowerWord := strings.ToLower(word)
		freq[lowerWord]++
	}
	return freq
}