package main

func FirstRepeatedRune(s string)(rune, bool){
	mapa := make(map[rune]struct{})
	for _, v := range s{
		_, ok := mapa[v] 
		if ok{
			return v, true
		}
		mapa[v] = struct{}{}
	}
	return 0, false
}