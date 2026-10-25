package main

func RemoveRune(s string, target rune)(string){
	removed:= []rune{}
	sRune:=[]rune(s)
	for i:=0; i<len(sRune); i++{
		if sRune[i] != target{
			removed = append(removed, sRune[i])
		}
	}
	return string(removed)
}