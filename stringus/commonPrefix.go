package main

func CommonPrefix(a, b string)(string){
	c := []rune{}
	if len(b) > len(a){
		a,b = b,a
	}
	for i := range(b){
		if !(rune(a[i]) == rune(b[i])){
			break
		}
		c = append(c, rune(a[i]))
	}
	return string(c)
}