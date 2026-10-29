package main

import "unicode"

func WordCount(s string) int {
	sNew := NormalizeString(s)
	c := 1
	for _, v := range sNew {
		if unicode.IsSpace(v){
			c++
		}
	}
	return c
}