package main

func IsAscii(s string)(bool){
	bytes := len([]byte(s))
	runes := len([]rune(s))
	if bytes == runes{
		return true
	}
	return false
}