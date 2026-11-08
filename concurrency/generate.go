package main

func Generate(n int, next func(int) int) <-chan int {
	ch := make(chan int)

	go func() {
		for i := 0; i < n; i++ {
			ch <- next(i) 
		}
		close(ch) 
	}()

	return ch
}