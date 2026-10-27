package main

func CountOccurrences(s, sub string)(int, error){
	sSlice := []byte(s)
	subSlice := []byte(sub)
	c := 0
	for i := 0; i <= len(sSlice)-len(subSlice); i++ {
		lens := len(subSlice)
		for j := range subSlice {
			if !(sSlice[i+j] == subSlice[j]) {
				break
			}
			lens--
			if lens == 0 {
				c++
				i = i + j
			}
		}
	}
	return c, nil
}