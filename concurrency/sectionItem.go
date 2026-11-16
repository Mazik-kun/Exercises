package main

import "errors"

type SectionItem struct {
	Value int
	End   bool
}

func ReadSections(in <-chan SectionItem) ([][]int, error) {
	sliceSlice := [][]int{}
	curSlice := []int{}
	for v := range in {
		if !v.End {
			curSlice = append(curSlice, v.Value)
		} else {
			sliceSlice = append(sliceSlice, curSlice)
			curSlice = []int{}
		}
	}
	sliceSlice = append(sliceSlice, curSlice)
	if len(sliceSlice[len(sliceSlice)-1]) != 0 {
		return nil, errors.New("не закрыт последний раздел")
	}
	return sliceSlice[:len(sliceSlice)-1], nil
}
