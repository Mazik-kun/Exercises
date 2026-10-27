package main

func LongestCommonPrefix(items []string) string{
	if len(items) == 0{
		return ""
	}
	c:= len([]rune(items[0]))
	for i := range items{
		if c > len([]rune(items[i])){
			c = len([]rune(items[i]))
		}
	}
	common := []rune{}
	for i:= 0; i < c; i++{
		bykStd := []rune(items[0])[i]
		for _, v := range items{
			
			bykNew := []rune(v)[i]
			if bykNew != bykStd{
				return string(common)
			}
		}
		common = append(common, bykStd)
	}
	return string(common)
}