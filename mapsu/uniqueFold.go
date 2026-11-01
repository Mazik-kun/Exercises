package main

import "strings"

func UniqueFold(words []string) []string {
	seen := make(map[string]struct{})
	res := []string{}
	for _, w := range words {
		lower := strings.ToLower(w)
		if _, ok := seen[lower]; !ok {
			seen[lower] = struct{}{}
			res = append(res, w) 
		}
	}
	return res
}