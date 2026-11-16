package main

func AllRunesUnique(s string) bool{
	mapa := make(map[rune]struct{})
	for _, v := range s{
		_, ok := mapa[v] 
		if ok{
			return false
		}
		mapa[v] = struct{}{}
	}
	return true
}