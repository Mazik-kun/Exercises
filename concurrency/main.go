package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		ch <- 4
		ch <- 1
		close(ch)
	}()

	// fmt.Println(ReceiveIn(ch, 2))
	// fmt.Println(SumStream(ch))
	fmt.Println(MinStream(ch))
}