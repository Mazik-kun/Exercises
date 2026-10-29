package main

func FirstUniqueRune(s string) (rune, bool) {
	counts := make(map[rune]int)
	for _, r := range s {
		counts[r]++
	}
	for _, r := range s {
		if counts[r] == 1 {
			return r, true
		}
	}
	return 0, false
}