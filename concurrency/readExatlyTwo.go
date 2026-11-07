package main

import "errors"

func ReadExactlyTwo(in <-chan int) (int, int, error) {
	var a, b int
	count := 0
	for v := range in {
		if count == 0 {
			a = v
		} else if count == 1 {
			b = v
		}
		count++
	}

	if count == 2 {
		return a, b, nil
	}
	return 0, 0, errors.New("not exactly two numbers")
}