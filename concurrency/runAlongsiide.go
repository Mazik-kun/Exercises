package main


func RunAlongside(local, background func() int) (int, int) {
	backgroundResult := make(chan int)

	go func() {
		backgroundResult <- background()
		close(backgroundResult)
		}()
	
	localResult := local()
	return localResult, <- backgroundResult

}