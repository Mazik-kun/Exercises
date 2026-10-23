package main

func FirstRune(s string) (bool,string) {
	if s == "" {
		return  false, ""
	}
	str := []rune(s)
	runa :=string(str[0])
	return true, runa
}