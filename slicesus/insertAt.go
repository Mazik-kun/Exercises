package main

import "errors"

func InsertAt(s []int, i, value int) ([]int, error) {
	if i > len(s) || i < 0 {
		return nil, errors.New("неверный индекс")
	}
	sNew := make([]int, 0, len(s)+1)
	sNew = append(sNew, s[:i]...)
	sNew = append(sNew, value)
	sNew = append(sNew, s[i:]...)
	return sNew, nil
}