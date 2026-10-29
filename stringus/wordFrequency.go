package main

import (
	"strings"
	"unicode"
)

func WordFrequency(s string) map[string]int {
	words := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})
	dict := make(map[string]int)
	for _, word := range words {
		lowerWord := strings.ToLower(word)
		dict[lowerWord]++
	}
	return dict
}