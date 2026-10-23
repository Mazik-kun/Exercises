package main

func IndexOf(s []int, target int) int{
	idx := -1
	for i:=0;i<len(s);i++{
		if s[i] == target{
			idx = i
			return idx
		}
	}
	return idx
}