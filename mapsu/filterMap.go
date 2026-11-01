package main

func FilterMap(m map[string]int, keep func(string, int) bool) map[string]int {
	res := make(map[string]int)
	
	for k, v := range m {
		if keep(k, v) {
			res[k] = v
		}
	}
	
	return res
}