package main

func DeleteZeroValues(m map[string]int){
	for key, v := range m{
		if v == 0{
			delete(m, key)
		}
	}
}