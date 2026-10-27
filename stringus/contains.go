package main



func Contains(s, sub string) bool {
	if len(s) < len(sub) {
		return false
	}
	if sub == "" {
		return true
	}
	for i := 0; i <= len(s)-len(sub); i++ {
		lens := len(sub)
		if rune(s[i]) == rune(sub[0]) {
			for j := 0; j < len(sub); j++ {
				lens--
				if !(rune(s[i+j]) == rune(sub[j])) {
					break
				}
				if lens == 0 {
					return true
				}
			}
		}
	}
	return false
}