package main

func SumStream(in <-chan int) int {
	sum := 0
	for v := range in {
		sum += v
	}
	return sum
}