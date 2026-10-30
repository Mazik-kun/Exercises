package main

import "errors"

func Splice(s []int, start, deleteCount int, values []int) ([]int, error) {
	sNew, err := ReplaceRange(s, start, start+deleteCount, values)
	if err != nil {
		return nil, errors.New("неверный диапазон")
	}
	return sNew, nil
}