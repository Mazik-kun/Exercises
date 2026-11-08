package main

func SendAndWait(value int, out chan<- int, processed <-chan struct{}) {
	out <- value  
	<-processed  
}