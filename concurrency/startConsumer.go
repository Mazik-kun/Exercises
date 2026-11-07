package main

func StartConsumer(in <-chan int, consume func(int)) <-chan struct{}{
	done := make(chan struct{})
	go func() {
		for v := range in {
			consume(v)
		}
		
		close(done) 
	}()
	return done
}