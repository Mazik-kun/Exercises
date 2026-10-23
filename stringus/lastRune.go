package main

func LastRune(s string) (bool,string) {
	if s == "" {
		return  false, ""
	}
	str := []rune(s)
	runa :=string(str[len(str)-1])
	return true, runa
}