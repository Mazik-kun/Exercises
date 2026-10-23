package main

func ReplaceAll(s []int,old, new int){
	for i:=0; i<len(s);i++{
		if s[i] == old{
			s[i] = new
		}
	}
}