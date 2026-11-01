package main

func MinStream(in <-chan int) (int, bool) {
	first, ok := <-in
	if !ok{
		return 0, false
	}
	min := first
	for v := range in {
		if min > v{
			min = v
		}
	}
	return min, true
}