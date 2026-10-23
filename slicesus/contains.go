package main

func Contains(s []int, target int) bool{
	flag:= false
	for i:=0;i<len(s);i++{
		if s[i] == target{
			flag = true
		}
	}
	return flag
}