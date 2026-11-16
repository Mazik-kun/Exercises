package main

import "sync"

func FindFirst(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	mid := len(nums) / 2
	leftIdx, rightIdx := -1, -1
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i, v := range nums[:mid] {
			if v == target {
				leftIdx = i
				break
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i, v := range nums[mid:] {
			if v == target {
				rightIdx = mid + i
				break
			}
		}
	}()

	wg.Wait()

	if leftIdx == -1 {
		return rightIdx
	}
	if rightIdx == -1 {
		return leftIdx
	}
	if leftIdx < rightIdx {
		return leftIdx
	}
	return rightIdx
}