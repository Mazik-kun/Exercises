package main

func IsNonDecreasing(in <-chan int)bool{
	vPre := <- in
	vNow:= 0
	flag := true
	for v := range in{
		vNow = v
		if vNow < vPre{
			flag = false
		}
		vPre = vNow
	}
	return flag
}