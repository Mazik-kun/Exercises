package main

import (
	"slices"
)
func main() {
    str:= "Stroka s glasnymy"
    MappingVowels(str)
}
func MappingVowels(str string) string{
    vowels:= []rune("aeiuoAEUIO")
    reversedVowels:= []rune{}
    rns:= []rune(str)
    var newString string
    var letter rune
    for i:= len(str)-1; i>= 0;i--{
        letter:= rns[i]
        if slices.Contains(vowels, letter){
            reversedVowels = append(reversedVowels, letter)
        }
    }
    for i:= 0; i< len(str);i++{
        letter= rns[i]
        if slices.Contains(vowels, letter){
            newString = newString + string(reversedVowels[0])
            reversedVowels = reversedVowels[1:]
            
        }else{
            newString = newString + string(rns[i])
        }
    }
    return newString
}
