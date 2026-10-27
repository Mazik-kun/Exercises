package main

import (
	"unicode"
	"fmt"
)

func TrimUnicodeSpace(s string) string {
	if s == "" {
		return ""
	}
	sNew := []rune(s)
	fmt.Println(sNew, len(sNew), len(s))
	for i:= 0; i < len(s); i ++{
		if !unicode.IsSpace(sNew[i]) {
			sNew = sNew[i:]
			break
		}
	}
	for i:=len(sNew)-1; i>=0; i--{
		fmt.Println(i, sNew)
		if !unicode.IsSpace(sNew[i]) {
			sNew = sNew[:i+1]
			break
		}
	}
	return string(sNew)
}

// func TrimUnicodeSpace(s string) string {
// 	runes := []rune(s)
// 	first := -1
// 	for i, r := range runes {
// 		if !unicode.IsSpace(r) {
// 			first = i
// 			break
// 		}
// 	}

// 	if first == -1 {
// 		return ""
// 	}

// 	last := len(runes) - 1
// 	for unicode.IsSpace(runes[last]) {
// 		last--
// 	}

// 	return string(runes[first : last+1])
// }