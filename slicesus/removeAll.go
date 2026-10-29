package main

func RemoveAll(s []int, target int) []int {
	sNew := make([]int, len(s))
	count := 0
	for _, v := range s {
		if v != target {
			sNew[count] = v
			count++
		}
	}
	return sNew[:count]
}