package main

import "unicode"

func AreAnagrams(a, b string) bool {
	mapa := make(map[rune]int)
	mapb := make(map[rune]int)
	for _, v := range a {
		if unicode.IsLetter(v) || unicode.IsDigit(v) {
			mapa[unicode.ToLower(v)]++
		}
	}
	for _, v := range b {
		if unicode.IsLetter(v) || unicode.IsDigit(v) {
			mapb[unicode.ToLower(v)]++
	}
	if len(mapa) != len(mapb){
		return false
	}
	}
	for key := range mapa {
		if mapa[key] != mapb[key] {
			return false
		}
	}
	return true
}