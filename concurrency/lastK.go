package main

func LastK(in <-chan int, k int) []int {
	last := make([]int, k)
	c := 0
	for v := range in {
		if k>0{
			last[c%k] = v
		}
		c++
	}
	if k>0{
		last = append(last[c%k:], last[0:c%k]...)
	}
	if c < k {
		last = last[k-c:]
	}
	return last
}