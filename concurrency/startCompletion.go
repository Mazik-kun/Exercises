package main

func StartCompletion(work func()) <-chan struct{} {
	done := make(chan struct{}) 

	go func() {
		work()       
		close(done)  
	}()

	return done
}