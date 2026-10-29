package main

func SwapPairs(s []int) {
	for i := 0; i < len(s)-len(s)%2; i += 2 {
		s[i], s[i+1] = s[i+1], s[i]
	}
}