package main

func GroupByParity(nums []int) map[string][]int {
	res := map[string][]int{
		"even": {},
		"odd":  {},
	}
	for _, v := range nums {
		if v%2 == 0 {
			res["even"] = append(res["even"], v)
		} else {
			res["odd"] = append(res["odd"], v)
		}
	}
	return res
}