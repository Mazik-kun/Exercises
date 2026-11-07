package main

import "errors"

func DeleteRange(s []int, left, right int) ([]int, error) {
	if left < 0 || right < left || right > len(s){
		return nil, errors.New("неверный диапазон")
	}
	if left == right { 
		return s, nil 
	}
	sNew := make([]int, 0, len(s))
	sNew = append(sNew, s[:left]...)
	sNew = append(sNew, s[right:]...)
	return sNew, nil
}