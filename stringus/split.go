package main

func Split(s, sep string) []string {
	sSlice := []rune(s)
	subSlice := []rune(sep)
	splited := []string{}
	start := 0
	for i := 0; i <= len(sSlice)-len(subSlice); i++ {
		lens := len(subSlice)
		for j := range subSlice {
			if !(sSlice[i+j] == subSlice[j]) {
				break
			}
			lens--
			if lens == 0 {
				splited = append(splited, string(sSlice[start:i]))
				start = i + len(subSlice) 
			}
		}	
	}
	splited = append(splited, string(sSlice[start:]))
	return splited
}