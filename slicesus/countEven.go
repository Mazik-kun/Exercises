package main

func CountEven(s []int)(int){
	count:=0
	for i:=0;i<len(s);i++{
		if s[i] % 2 == 0{
			count++
		}
	}
	return count
}