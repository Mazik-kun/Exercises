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
	normalizeSpace := "a\tb"
	UnicodeSpace:= "		к ак к   к    f		"

	trimUnicodeSpace :=(" абсабс ")
	fmt.Println(TrimUnicodeSpace(trimUnicodeSpace))
	fmt.Println(TrimUnicodeSpace(""))
	fmt.Println(NormalizeString(normalizeSpace))
	
	fmt.Println(NormalizeString(UnicodeSpace))

	fmt.Println(Contains(trimUnicodeSpace, trimUnicodeSpace))

	fmt.Println(IndexOf(trimUnicodeSpace, "абс "))


	fmt.Println(CountOverlapping(" aaaaaaa aaaaa", "aa"))
	fmt.Println(CountOccurrences(" aaaaaaa aaaaa", "aa"))

	fmt.Println(CommonPrefix("hel", "hell"))
	
	fmt.Println(LongestCommonPrefix([]string{"кот", "котенок", "котик"}))
	fmt.Println(LongestCommonPrefix([]string{}))

	fmt.Println(Split("брухкрутожеуткаутрокаакут", "ут"))
	fmt.Println(Join(Split("брухкрутожеуткаутрокаакут", "ут"), "утка"))
	fmt.Println(TrimSet("брухкрутожеуткаутрокаакут", "брут"))
	fmt.Println(CollapseRuns("ааааабаабаабабабабббба"))

	fmt.Println(EncodeRLE("ааааабаабаабабабабббба"))
	fmt.Println(DecodeRLE("5:а;1:б;2:а;1:б;2:а;1:б;1:а;1:б;1:а;1:б;1:а;4:б;1:а;"))
	fmt.Println(WordCount(" aaaaaaa      aa    aaa"))

	fmt.Println(WordFrequency("Hello, hello! World."))
	fmt.Println(FirstUniqueRune("бадаак"))
}