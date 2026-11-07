package main

func BufferedSource(nums []int, capacity int) <-chan int {
	ch := make(chan int, capacity) 

	go func() {
		for _, v := range nums {
			ch <- v
		}
		close(ch)
	}()

	return ch
}