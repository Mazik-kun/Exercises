package main

func CollapseRuns(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	result := make([]rune, 0, len(runes))
	result = append(result, runes[0])
	last := runes[0]
	for i := 1; i < len(runes); i++ {
		if runes[i] != last {
			result = append(result, runes[i])
			last = runes[i]
		}
	}	
	return string(result)
}