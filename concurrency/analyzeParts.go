package main

import "sync"

type result struct {
	index int
	value int
}

func AnalyzeParts(parts [][][]int, analyze func([][]int) int) []int {
	if len(parts) == 0 {
		return []int{}
	}

	ch := make(chan result, len(parts))
	var wg sync.WaitGroup

	for i, p := range parts {
		wg.Add(1)
		go func(idx int, part [][]int) {
			defer wg.Done()
			ch <- result{index: idx, value: analyze(part)}
		}(i, p)
	}

	wg.Wait()
	close(ch)

	res := make([]int, len(parts))
	for r := range ch {
		res[r.index] = r.value
	}

	return res
}