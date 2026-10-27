package main


func IndexOf(s, sub string) int {
	sSlice := []byte(s)
	subSlice := []byte(sub)
	for i := 0; i <= len(sSlice)-len(subSlice); i++ {
		lens := len(subSlice)
		for j := range subSlice {
			if !(sSlice[i+j] == subSlice[j]) {
				break
			}
			lens--
			if lens == 0 {
				return i
			}
		}
	}
	return -1
}