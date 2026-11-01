package main

import (
	"slices"
)

type Group struct {
	firstIdx int
	words    []string 
}

func GroupAnagrams(words []string) [][]string {
	mapa := make(map[string]Group)
	order := 0
	for _, w := range words {
		b := []rune(w)
		slices.Sort(b)
		key := string(b)

		if g, ok := mapa[key]; ok {
			g.words = append(g.words, w)
			mapa[key] = g
		} else {
			mapa[key] = Group{
				firstIdx: order,
				words:    []string{w},
			}
			order++
		}
	}
	res := make([][]string, order)
	for _, g := range mapa {
		res[g.firstIdx] = g.words
	}
	return res
}