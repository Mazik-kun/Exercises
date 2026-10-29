package main

import "errors"

func RemoveAt(s []int, i int) ([]int, error) {
	if i < 0 || i >= len(s) {
		return s, errors.New("index out of bounds")
	}

	for idx := i; idx < len(s)-1; idx++ {
		s[idx] = s[idx+1]
	}
	return s[:len(s)-1], nil
}