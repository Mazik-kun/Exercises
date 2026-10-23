package main

func LastIndex(words []string)(map[string]int){
	last := make(map[string]int)
	for i:=0; i < len(words); i++{
		last[words[i]] = i
	}
	return last
}