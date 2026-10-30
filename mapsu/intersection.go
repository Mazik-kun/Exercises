package main

func Intersection(a, b []string) []string{
	seen := make(map[string]bool)
	res := []string{}
	for _, v := range b {
		seen[v] = true
	}
	for _, v := range a {
		if seen[v] {
			res = append(res, v)
			seen[v] = false
		}
	}
	return res
}