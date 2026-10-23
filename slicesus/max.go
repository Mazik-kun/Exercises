package main

func Max(s []int)(int, bool){
	if len(s) == 0{
		return 0, false
	}
	max:=s[0]
	for i:=0;i<len(s);i++{
		if s[i]> max{
			max = s[i]
		}
	}
	return max, true
}