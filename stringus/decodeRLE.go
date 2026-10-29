package main

import (
	"errors"
	"unicode"
)

func DecodeRLE(encoded string) (string, error) {
	encRune := []rune(encoded)
	if len(encRune) % 4 != 0{
		return "", errors.New("неверный формат")
	}
	result := []rune{}
	for i := range encRune {
		syb := encRune[i]
		if i%4 == 0 && unicode.IsDigit(syb){
			for j := 0; j < int(syb - '0');j++{
				result = append(result, encRune[i+2])
			}
		} else if i%4 == 1 && syb != ':'{
			return "", errors.New("неверный формат")
		}else if i%4 == 3 && syb != ';'{
			return "", errors.New("неверный формат")
		}else if i%4==0{
			return "", errors.New("неверный формат")
		}
	}
	return string(result), nil
}