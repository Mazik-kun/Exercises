package main

func Clone(m map [string][]int) map[string][]int{
	clone := make(map[string][]int)
	for key, v := range m{
		newSlice := []int{}
		for _, v2 := range v{
			newSlice = append(newSlice, v2)
		}
		clone[key] = newSlice
	}
	return clone
}