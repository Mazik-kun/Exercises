package main

func IndexesOf(s []int, target int) []int{
	idxes := []int{}
	for i:=0;i<len(s);i++{
		if s[i] == target{
			idxes = append(idxes, i)
		}
	}
	return idxes
}