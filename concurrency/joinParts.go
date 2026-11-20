package main

import (
	"errors"
	"slices"
)

type ResultPart struct {
	Index  int
	Values []int
}

func JoinParts(in <-chan ResultPart, n int) ([]int, error) {
	vals := make([][]int, n)
	var err error
	err = nil
	res := []int{}
	lenIn := 0
	for v := range in {
		lenIn++
		if v.Index > n-1 || v.Index < 0 {
		err = errors.New("Неверный индекс")
		res = nil
		}
		if err == nil{
			if slices.Equal(vals[v.Index], v.Values) && vals[v.Index] != nil {
				err = errors.New("повтор Id")
				res = nil
			}else{
				vals[v.Index] = v.Values
			}
		}
	
	}
	if lenIn != n{
		err = errors.New("не хватает частей")
		res = nil
	}
	if err == nil{
		for _, v := range vals{
			res = append(res, v...)
		}
	}
	return res, err
}