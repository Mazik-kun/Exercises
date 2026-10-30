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
	fmt.Println(SymmetricDifference([]int{1, 2, 3, 4}, []int{2, 4, 5})) 


}