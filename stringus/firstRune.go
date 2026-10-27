package main

func FirstRune(s string) (string,bool) {
	if s == "" {
		return  "",false
	}
	str := []rune(s)
	runa :=string(str[0])
	return runa,true
}