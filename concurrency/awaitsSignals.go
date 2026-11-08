package main

func AwaitSignals(signals []<-chan struct{}) {
	for _, ch := range signals {
		<-ch
	}
}