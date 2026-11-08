package main

func WordCounts(in <-chan string) map[string]int{
	mapa := make(map[string]int)
	for key := range in{
		mapa[key]++
	}
	return mapa
}