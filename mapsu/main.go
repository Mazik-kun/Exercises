package main

import "fmt"

func main() {
	sli := []int{5, 2, 9, 4, 0, 11, 1, 1, 1, 1, -10, 12, 54, -11}
	fmt.Println(CountInts(sli))

	words := []string{"f","g","qqq","tt","t","t","f"}
	fmt.Println(LastIndex(words))

	fmt.Println(Unique(sli))
}