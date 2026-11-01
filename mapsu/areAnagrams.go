package main

import "unicode"
func AreAnagrams(a, b string) bool {
	mapA := make(map[rune]int)
	mapB := make(map[rune]int)
	for _, r := range a {
		if !unicode.IsSpace(r) {
			mapA[unicode.ToLower(r)]++
		}
	}
	for _, r := range b {
		if !unicode.IsSpace(r) {
			mapB[unicode.ToLower(r)]++
		}
	}
	if len(mapA) != len(mapB) {
		return false
	}
	for k, v := range mapA {
		if mapB[k] != v {
			return false
		}
	}
	return true
}