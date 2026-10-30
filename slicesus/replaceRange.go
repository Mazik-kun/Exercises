package main

import "errors"

func ReplaceRange(s []int, left, right int, values []int) ([]int, error) {
	sNew, err := DeleteRange(s, left, right)
	if err != nil {
		return nil, errors.New("неверный диапазон")
	}
	sNew,err = InsertMany(sNew, left, values)
	if err != nil {
		return nil, errors.New("неверный диапазон")
	}
	return sNew, nil
}