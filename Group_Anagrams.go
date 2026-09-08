package main

import (
	"fmt"
	"slices"
)

func main() {
	strs := []string{"gege","egeg","cir","irl","crit", "e"}
	fmt.Printf("%v\n", GroupAnagrams(strs))
}
func GroupAnagrams(strs []string) [][]string{
	data := [][]string{}
	mapa := make(map[string][]string, len(strs))
	for i:= 0; i < len(strs);i++{
		b := []byte(strs[i])
	
	slices.Sort(b)
	key := string(b)
	mapa[key] = append(mapa[key], strs[i])
	}
	for _, value := range mapa {
        data = append(data,value)
    }
	return data
}
