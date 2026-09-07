package main

import (
	// "fmt"
)

func main(){
	s := ")("
	println(checkValidness(s))
}

func checkValidness(s string)(bool){
	round := 0
	square := 0
	figure := 0
	symb := s[0]
	switch symb{
			case '(':
				round++
			case ')':
				return false
			case '[':
				square++
			case ']':
				return false
			case '{':
				figure++
			case '}':
				return false
		}
	for i:= 0; i<len(s)-1; i++{
		symb = s[i+1]
		switch symb{
			case '(':
				round++
			case ')':
				round--
			case '[':
				square++
			case ']':
				square--
			case '{':
				figure++
			case '}':
				figure--
		}
		
		pair := s[i:i+2]
		println(pair)
		switch pair{
			case "(]":
				return false;
			case "(}":
				return false;
			case "[)":
				return false;
			case "[}":
				return false;
			case "{]":
				return false;
			case "{)":
				return false;
			default:
		}
	}
if !(round == 0 && square == 0 && figure == 0){
	return false}
return true}
