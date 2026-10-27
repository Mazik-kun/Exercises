package main

func LastRune(s string) (string,bool) {
	if s == "" {
		return  "",false
	}
	str := []rune(s)
	runa :=string(str[len(str)-1])
	return runa, true
}