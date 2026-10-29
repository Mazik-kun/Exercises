package main

import "errors"

func ReverseRange(s []int, left, right int) error {
	if left < 0 || right < 0 || right <= left || left > len(s)-1 || right > len(s)-1 {
		return errors.New("incorrect bounds")
	}
	
	for i := left; i <= left+(right-left)/2; i++ {
		s[i], s[right+left-i] = s[right+left-i], s[i]
	}
	return nil
}