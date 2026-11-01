package main

func MostFrequent(nums []int) (value int, count int, ok bool) {
	if len(nums) == 0 {
		return 0, 0, false
	}

	counts := make(map[int]int)
	for _, n := range nums {
		counts[n]++
	}

	maxCount := -1
	bestNum := 0
	for k, v := range counts {
		if v > maxCount {
			maxCount = v
			bestNum = k
		} else if v == maxCount && k < bestNum {
			bestNum = k
		}
	}
	return bestNum, maxCount, true
}