package main

import (
	"strings"
	"unicode"
)

func TopKWords(text string, k int) []string {
	words := strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})

	counts := make(map[string]int)
	for _, w := range words {
		counts[strings.ToLower(w)]++
	}
	if k > len(counts) {
		k = len(counts)
	}
	res := make([]string, 0, k)
	for i := 0; i < k; i++ {
		bestWord := ""
		maxCount := -1
		for w, c := range counts {
			if c > maxCount {
				maxCount = c
				bestWord = w
			} else if c == maxCount && w < bestWord {
				bestWord = w
			}
		}
		res = append(res, bestWord)
		counts[bestWord] = -1 
	}
	return res
}