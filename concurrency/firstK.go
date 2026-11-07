package main

func FirstK(in <-chan int, k int) []int{
	sliceK := make([]int, 0, k)
	for v := range in{
		if k > 0{
			sliceK = append(sliceK, v)
			k--
		}	
	}
	return sliceK
}