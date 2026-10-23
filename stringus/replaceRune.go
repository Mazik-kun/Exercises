package main

func ReplaceRune(s string, old, new rune)(string){
	sRune:=[]rune(s)
	for i:=0; i<len(sRune)-1; i++{
		if sRune[i] == old{
			sRune[i] = new
		}
	}
	return string(sRune)
}