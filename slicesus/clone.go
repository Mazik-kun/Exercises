package main

func Clone(s []int) []int{
	sNew:= make([]int, len(s))
	copy(sNew, s)
	return sNew
}