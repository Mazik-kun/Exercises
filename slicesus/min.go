package main

func Min(s []int)(int, bool){
	if len(s) == 0{
		return 0, false
	}
	min:=s[0]
	for i:=0;i<len(s);i++{
		if s[i]< min{
			min = s[i]
		}
	}
	return min, true
}