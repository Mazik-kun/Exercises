package main

func GroupByLength(words []string) map[int][]string {
	res := make(map[int][]string)
	for _, w := range words {
		length := len([]rune(w))
		res[length] = append(res[length], w)
	}
	return res
}