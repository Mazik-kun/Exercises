package main

import (
	"slices"
)

func main() {
	nums := []int{1, -2, 1}
	println(maxSubArray(nums))
}
func maxSubArray(nums []int) int {
	maxSArrSum := slices.Max(nums)
	currentSum := 0
    for i:= 0; i < len(nums); i++{
		if currentSum + nums[i] > 0{
			currentSum = currentSum + nums[i]
			if maxSArrSum < currentSum{
				maxSArrSum = currentSum
			}
		} else{
			currentSum = 0
		}
		
	}
return maxSArrSum
}
