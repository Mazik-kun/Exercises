package main

func Equal(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}
	for key, v := range a {
		v2, ok := b[key]
		if !ok || v != v2 {
			return false
		}
	}
	return true
}