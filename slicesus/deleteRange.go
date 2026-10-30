package main

func DeleteRange(s []int, left, right int)([]int, error){
	sNew := make([]int, 0, len(s))
	sNew = append(sNew, s[:left]...)
	sNew = append(sNew, s[right:]...)
	return sNew, nil
}