package main

func ContainsAll(have, need []string)bool{
	unqiueMap := make(map[string]struct{})
	for _, v := range have{
		unqiueMap[v] = struct{}{}
	}
	for _, v := range need{
		if _, ok := unqiueMap[v] ; !ok{
			return false
		}
	}
	return true
}