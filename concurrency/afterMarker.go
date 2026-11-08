package main

func AfterMarker(in <-chan string, marker string) []string {
	res := []string{}
	found := false

	for v := range in {
		if found {
			res = append(res, v)
		} else if v == marker {
			found = true
		}
	}

	return res
}