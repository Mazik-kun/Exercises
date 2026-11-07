package main

func StartValue(compute func() int) <-chan int{
	newCh:= make(chan int, 1)
	go func(){
		comp := compute()
		newCh <- comp
		close(newCh)
	}()
	
	return newCh
}