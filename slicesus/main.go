package main

import "fmt"

func main() {
	sli := []int{5, 2, 9, 4, 0, 11, 1, 1, 1, 1,-10, 12, 54, -11}
	slick:= []int{}

	fmt.Println(Sum(sli))
	fmt.Println(Sum(slick))

	fmt.Println(Min(sli))
	fmt.Println(Min(slick))

	fmt.Println(Max(sli))

	fmt.Println(CountEven(sli))

	target := 1
	
	fmt.Println(Contains(sli,target))
	fmt.Println(CountOccurrences(sli,target))
	fmt.Println(IndexOf(sli,target))
	fmt.Println(IndexesOf(sli,target))

	old:=1
	new:=2
	ReplaceAll(sli,old,new)
	Reverse(sli)
	fmt.Println(sli)
}