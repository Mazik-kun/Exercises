package main

func TrimSet(s, cutset string) string {
	runes := []rune(s)
	cutMap := make(map[rune]struct{})
	for _, r := range cutset {
		cutMap[r] = struct{}{}
	}
	start := 0
	for start < len(runes) {
		if _, ok := cutMap[runes[start]]; !ok {
			break
		}
		start++
	}
	if start == len(runes) {
		return ""
	}
	end := len(runes) - 1
	for end >= start {
		if _, ok := cutMap[runes[end]]; !ok {
			break 
		}
		end--
	}
	
	return string(runes[start : end+1])
}