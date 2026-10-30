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

	SwapPairs(sli)

	fmt.Println(sli)

	fmt.Println(ReverseRange([]int{10, 20, 30, 40, 50, 60}, 1, 4))

	fmt.Println(RemoveAt([]int{10, 20, 30, 40, 50}, 2))

	fmt.Println(RemoveAtFast([]int{10, 20, 30, 40, 50}, 2)) 

	fmt.Println(RemoveAll([]int{1, 2, 3, 2, 4, 2}, 2)) 

	fmt.Println(RemoveAllInPlace([]int{1, 2, 3, 2, 4, 2}, 2))

	fmt.Println(KeepEvenInPlace([]int{1, 2, 3, 2, 4, 2}))

	zeros:= []int{1,2,0,0,0,3,0,4,0,4,2,45,6,4,4,76,0,9,90,0,0,0,}
	fmt.Println(CompactNonZero(zeros))
	MoveZeroToEnd(zeros)
	fmt.Println(zeros)

	fmt.Println(InsertAt([]int{1,2,3,5,6},4,4))

	fmt.Println(InsertMany([]int{1,2,3,5,6},1,[]int{}))

	fmt.Println(DeleteRange([]int{1,2,3,5,6},1,2))

	fmt.Println(ReplaceRange([]int{1,2,3,5,6},1,5,[]int{0,0,0}))

	fmt.Println(Splice([]int{1,2,3,5,6},1,4,[]int{0,0,0}))
	
	fmt.Println(Clone([]int{1,2,3,5,6}))
}