package main

import "errors"

func InsertMany(s []int, i int, values []int) ([]int, error) {
	if i > len(s) || i < 0 {
		return nil, errors.New("неверный индекс")
	}
	sNew := make([]int, 0, len(s)+len(values))
	sNew = append(sNew, s[:i]...)
	sNew = append(sNew, values...)
	sNew = append(sNew, s[i:]...)
	return sNew, nil
}