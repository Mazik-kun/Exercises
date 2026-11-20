package main

import "sync"

func CheckQuorum(checks []func() bool, required int) bool {
	fact := 0

	var wg sync.WaitGroup

	for _, item := range checks{
		wg.Add(1)
		go func() {
			defer wg.Done()
			if item(){
				fact++
			}
		}()
	}
	wg.Wait()
	if fact >= required {
		return true
	}
	return false
}