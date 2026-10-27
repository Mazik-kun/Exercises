package main

import "maps"

func MergeCounts(a, b map[string]int) map[string]int {
	sum := make(map[string]int)
	maps.Copy(sum, a)
	for k, v := range b {
		sum[k] += v
	}
	return sum
}