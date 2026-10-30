package main

func CompactNonZero(s []int) []int{
	readPnt := 0
	writePnt := 0
	for range s{
		if s[writePnt] != 0 {
			if s[readPnt] !=0{
				s[writePnt] = s[readPnt]
				writePnt++
				readPnt++
			}else{
				readPnt++
			}
		}else{
			if s[readPnt] !=0{
				s[writePnt], s[readPnt] = s[readPnt], s[writePnt]
				writePnt++
				readPnt++
			}else{
				readPnt++
			}
		}
	}
	s = s[0:writePnt]
	return s
}