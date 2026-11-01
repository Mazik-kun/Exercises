package main

func FirstUniqueRune(s string) (rune, bool) {
	seen := make(map[rune]bool)
	for _, r := range s {
		if _, ok := seen[r]; ok {
			seen[r] = false
		} else {
			seen[r] = true
		}
	}
	for _, r := range s {
		if seen[r] {
			return r, true
		}
	}
	return 0, false
}