package main

func MapValues(m map[string]int, f func(int) int) map[string]int {
	res := make(map[string]int, len(m))
	for k, v := range m {
		res[k] = f(v)
	}
	return res
}