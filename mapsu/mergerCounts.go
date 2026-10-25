package main

func MergeCounts(a, b map[string]int) map[string]int{
	sum := make(map[string]int)
	for key, v := range a{
		if v2, ok := b[key]; ok{
			sum[key] = v + v2
		}
	}
	return sum
}