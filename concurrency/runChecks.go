package main

import "sync"

func RunChecks(first, second func() error) (error, error) {
	var err1, err2 error
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		err1 = first()
	}()
	go func() {
		defer wg.Done()
		err2 = second()
	}()

	wg.Wait()

	return err1, err2
}