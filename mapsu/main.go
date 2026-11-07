package main

import "fmt"

func main() {
	sli := []int{5, 2, 9, 4, 0, 11, 1, 1, 1, 1, -10, 12, 54, -11}
	fmt.Println(CountInts(sli))

	words := []string{"f","g","qqq","tt","t","t","f"}
	fmt.Println(LastIndex(words))

	fmt.Println(Unique(sli))

	fmt.Println(ContainsAll(words, words))

	testMap := map[string]int{
		"banana": 2,
		"apple":  1,
		"cherry": 3,
		"pea": 3,
	}
	fmt.Println(SortedKeys(testMap))
	fmt.Println(SumValues(testMap))
	fmt.Println(Invert(testMap))

	testMap2 := map[string]int{
		"banana": 5,
		"applee":  4,
		"cherry": 0,
		"pea": 2,
	}
	
	fmt.Println(MergeCounts(testMap,testMap2))

	DeleteZeroValues(testMap2)

	fmt.Println(testMap2)

	testMap3 := map[string][]int{
		"first":  {1, 2, 3},
		"second": {4, 5},
	}
	fmt.Println(Clone(testMap3))

	mapA := map[string]int{"apple": 1, "banana": 2}
	mapB := map[string]int{"banana": 2, "apple": 1}

	fmt.Println(Equal(mapA, mapB))

	fmt.Println(Difference([]int{1, 2, 3, 4}, []int{2, 4}))

	fmt.Println(Intersection(
		[]string{"apple", "banana", "cherry", "apple"}, 
		[]string{"cherry", "banana", "date"},
	)) 
	fmt.Println(SymmetricDifference([]int{3}, []int{3, 3})) 

	fmt.Println(UniqueFold([]string{"Apple", "apple", "Banana", "banana", "Cherry"}))

	fmt.Println(MostFrequent([]int{1, 1, 2, 2, 3, 3, 3})) 

	fmt.Println(FirstUniqueRune("бадаак"))

	fmt.Println(AreAnagrams("Listen", "Silent"))

	fmt.Println(WordFrequency("Hello, hello! World. 123"))

	fmt.Println(TopKWords("a b b c c c d d d d", 2))

	fmt.Println(GroupByParity([]int{5, 2, 8, 1, 4}))

	fmt.Println(GroupByLength([]string{"я", "ты", "он", "она", "они"}))

	fmt.Println(GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))

	fmt.Println(MapValues(map[string]int{"a": 1, "b": 2}, func(x int) int { return x * 10 }))

	fmt.Println(FilterMap(map[string]int{"a": 1, "b": 2, "c": 3}, func(k string, v int) bool { return v > 1 }))
}