package main

func Difference(a, b []int) []int {
	seen := make(map[int]bool)
	res := []int{}
	for _, v := range b {
		seen[v] = true
	}
	for _, v := range a {
		if !seen[v] {
			res = append(res, v)
			seen[v] = true
		}
	}
	return res
}