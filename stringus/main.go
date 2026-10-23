package main

import "fmt"

func main() {
	stringSize := "67pyuн"
	b, r := StringSize(stringSize)
	fmt.Printf("Символов: %d, Байт: %d\n", r, b)

	isAscii := "нет"
	fmt.Println(IsAscii(isAscii))

	firstRune := "лол"
	fmt.Println(FirstRune(firstRune))
	
	lastRune := "лол не"
	fmt.Println(LastRune(lastRune))

	reversy := "сатра"
	fmt.Println(Reversy(reversy))

	removeRune := "проверкаaaaaaaaпaпа"
	target:= rune('a')
	fmt.Println(RemoveRune(removeRune, target))

	replaceRune := "ffffqqqff"
	old:= rune('q')
	new:= rune('s')
	fmt.Println(ReplaceRune(replaceRune, old, new))

	isPalindrome:= "не казак"
	fmt.Println(IsPalindrome(isPalindrome))

	UnicodeSpace:= "    к ак к   к    f"
	fmt.Println(TrimUnicodeSpace(UnicodeSpace))

	fmt.Println(NormalizeString(UnicodeSpace))
}