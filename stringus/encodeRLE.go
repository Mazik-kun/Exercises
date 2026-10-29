package main

import (
	"strconv"
)

func EncodeRLE(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	c := 1
	result := make([]rune, 0, len(runes)*3)
	last := runes[0]
	for i := 1; i < len(runes); i++ {
		if runes[i] != last {
			result = append(result, []rune(strconv.Itoa(c) + ":" + string(last) + ";")...)
			last = runes[i]
			c = 1
		} else {
			c++
		}

	}
	result = append(result, []rune(strconv.Itoa(c) + ":" + string(last) + ";")...)
	return string(result)
}