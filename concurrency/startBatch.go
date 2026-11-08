package main

import "sync"

func StartBatch(jobs []func() int) <-chan int {
	out := make(chan int, len(jobs))

	if len(jobs) == 0 {
		close(out)
		return out
	}

	var wg sync.WaitGroup
	wg.Add(len(jobs))

	for _, job := range jobs {
		go func(j func() int) {
			defer wg.Done()
			out <- j()
		}(job)
	}
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}