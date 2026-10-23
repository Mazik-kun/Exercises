package main

import (
	// "fmt"
)

func main(){
	nums := []int{1,2,3,4,}
	println(contains(nums))

}

func contains(nums []int) bool{
	mapa := make(map[int]int)
	for i:=0; i < len(nums); i++{
		mapa[nums[i]] = i
	}

	if len(nums) != len(mapa){
		return true
	}
return false}
