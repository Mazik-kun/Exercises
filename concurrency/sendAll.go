package main

func SendAll(nums []int, out chan<- int){
	for _,v := range nums{
		out <- v
	}
}