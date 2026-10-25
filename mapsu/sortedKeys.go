package main

import "sort"

func SortedKeys(m map[string]int) []string {
	mapToString := []string{}
	for key := range m {
		mapToString = append(mapToString, key)
	}
	sort.Strings(mapToString)
	return mapToString
}