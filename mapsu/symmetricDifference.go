package main

import (
	"fmt"
	"slices"
)

func SymmetricDifference(a, b []int) []int {
	states := make(map[int]int)
	res := []int{}
	for _, v := range a {
		states[v] = 1
	}
	for _, v := range b {
		if states[v] == 1 {
			states[v] = 2 
		} else if _, ok := states[v]; !ok{
			states[v] = 1
		}
	}

	for k, state := range states {
		fmt.Println(state, "дел")
		if state == 1 {
			res = append(res, k)
		}
	}
	slices.Sort(res) 
	return res
}