package main

import "errors"

func RemoveAtFast(s []int, i int) ([]int, error) {
	if i < 0 || i >= len(s) {
		return s, errors.New("index out of bounds")
	}
	s[i] = s[len(s)-1]
	
	return s[:len(s)-1], nil
}