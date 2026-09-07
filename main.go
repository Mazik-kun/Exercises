package main

import (
	// "fmt"
)

func main(){
	
	nums := []int{11,44,5123,53,76,67,9,21,10}
	target := 31
	index1, index2,count := TwoSum(nums,target)
	println(index1,index2,count)
	
}

// func TwoSum(nums []int, target int)(i,j,c int){
// 	c = 0
// 	for i := 0; i < len(nums) - 1; i++ {
// 		for j := i+1; j < len(nums); j++{
// 			c++
// 			if (nums[i] + nums[j]) == target{
// 				return i, j, c
// 			}
// 		}
// 	}
// return 0, 0, c}

func TwoSum(nums []int, target int)(i,j,c int){
	seen := make(map[int]int)
	count := 0
	for i, num := range nums{
		count ++ 
		complement := target - num
		
		if idx, found := seen[complement]; found{
			return idx,i,count
		}
		seen[num] = i
	}
	return 0,0,count
}
