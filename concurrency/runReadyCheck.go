package main

import (
	"sync"
)

type GatedCheck struct {
	Ready <-chan struct{}
	Check func() error
}

func RunReadyChecks(checks []GatedCheck) []error {
	var wg sync.WaitGroup
	
	errs := make([]error, len(checks))

	for i, item := range checks {
		wg.Add(1)
		go func(idx int, check GatedCheck) {
			defer wg.Done()
			<-check.Ready 
			errs[idx] = check.Check() 
		}(i, item)
	}
	wg.Wait()
	
	return errs
}