package main

func PrepareNumbers(nums []int) <- chan int{
	
	ch1 := make(chan int, len(nums))
	for _,v := range nums{
		ch1 <- v
	}
	close(ch1)
	return ch1
}