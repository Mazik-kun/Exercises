package main

func Reverse(s []int){
	for i:=0; i < len(s) / 2; i++{
		s[i], s[len(s)-i-1] = s[len(s)-i-1],s[i]
	}
}