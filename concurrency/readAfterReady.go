package main

func ReadAfterReady(ready <-chan struct{}, read func() int) <-chan int {
	out := make(chan int, 1)
	go func() {
		<-ready       
		val := read()
		out <- val     
		close(out)   
	}()
	return out
}