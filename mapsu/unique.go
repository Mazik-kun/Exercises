package main

func Unique(nums []int)[]int{
	res := []int{} 
	seen := make(map[int]struct{})
	for i:=0; i < len(nums); i++{
		if _, ok := seen[nums[i]]; !ok {
		res = append(res, nums[i])
		seen[nums[i]] = struct{}{}
		}		
	}
	return res
}