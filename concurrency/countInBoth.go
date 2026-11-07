package main

import "sync"

func CountInBoth(left, right []int, target int) (int, int) {
	var countLeft, countRight int
	var wg sync.WaitGroup
	wg.Add(2)

	go func(){
		defer wg.Done()
		for _,v := range left{
			if v == target{
				countLeft++
			}
		}
	}()
	go func(){
		defer wg.Done()
		for _,v := range right{
			if v == target{
				countRight++
			}
		}
	}()

	wg.Wait()

	return countLeft, countRight
}