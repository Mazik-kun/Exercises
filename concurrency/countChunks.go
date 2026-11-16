package main

import "sync"

func CountChunks(chunks [][][]int) map[int]int {
	ch := make(chan map[int]int, len(chunks))
	var wg sync.WaitGroup

	for _, chunk := range chunks {
		wg.Add(1)
		go func(c [][]int) {
			defer wg.Done()
			localMap := make(map[int]int)
			for _, row := range c {
				for _, v := range row {
					localMap[v]++
				}
			}
			ch <- localMap
		}(chunk)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	res := make(map[int]int)
	for localMap := range ch {
		for k, v := range localMap {
			res[k] += v
		}
	}

	return res
}