package main

func CountOccurrences(s []int, target int)(int){
	count:=0
	for i:=0;i<len(s);i++{
		if s[i] == target{
			count++
		}
	}
	return count
}