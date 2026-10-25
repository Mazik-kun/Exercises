package main

import "strings"

func NormalizeString(s string) string {
	sTrim := strings.TrimSpace(s)
	sRune := []rune(sTrim)
	normalized := []rune{}
	flag:= false
	for i:=0; i<len(sRune);i++{
		if sRune[i] == ' ' {
			if !flag{
				normalized = append(normalized, sRune[i])
				flag = true
			}
		}else{
			normalized = append(normalized, sRune[i])
			flag = false
		}
	}
	return string(normalized)
}