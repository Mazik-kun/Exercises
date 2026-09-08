package main

import (

)

func main() {
	nums := []int{5,4,-1,7,8}
	println(maxSubArray(nums))
}
func maxSubArray(nums []int) int {
	maxSArrSum := 0
	currentSum := 0
    for i:= 0; i < len(nums); i++{
		if currentSum + nums[i] > 0{
			currentSum = currentSum + nums[i]
		} else{
			currentSum = 0
		}
		if maxSArrSum < currentSum{
			maxSArrSum = currentSum
		}
	}
return maxSArrSum
}
