package main

func StringSize(s string)(bytes int, runes int){
	bytes = len([]byte(s))
	runes = len([]rune(s))
	return bytes, runes
}