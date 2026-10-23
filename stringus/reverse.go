package main

func Reversy(s string) string{
	origin:= []rune(s)
    leng := len(origin)-1
    nigiro:= []rune(s)
    for i:=0; i<=leng; i++{
        nigiro[i] = origin[leng-i]
    }
    return string(nigiro)
}
