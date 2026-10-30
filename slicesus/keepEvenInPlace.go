package main

func KeepEvenInPlace(s []int) []int {
	skip := 0
	idxes := []int{}
	for i, v := range s {
		if v % 2 == 1 {
			idxes = append(idxes, i-skip)
			skip++
		}
	}
	for _, idx := range idxes {
		for j := idx; j < len(s)-1; j++ {
			s[j] = s[j+1]
		}
	}
	return s[:len(s)-skip]
}