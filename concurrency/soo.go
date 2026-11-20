package main

import "fmt"

func Soo() {
	a := make(chan int)
	b := make(chan int)

	go func() {
		x := <-a
		b <- x * 2
	}()

	a <- 21
	fmt.Println(<-b)
}
