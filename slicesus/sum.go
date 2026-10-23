package main

func Sum(s []int)(int){
	sum := 0
	if len(s) != 0{
		for i:=0; i<len(s); i++{
			sum += s[i]
		}
	} 
	return sum
}