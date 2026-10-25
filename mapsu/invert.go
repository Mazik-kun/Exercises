package main

import "errors"

func Invert(m map[string]int) (map[int]string, error) {
	inverted := make(map[int]string)
	for key, v := range m {
		if _, ok := inverted[m[key]]; ok {
			return nil, errors.New("да было такое уже")
		}
		inverted[v] = key
	}
	return inverted, nil
}