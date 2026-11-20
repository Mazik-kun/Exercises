package main

import (
	"errors"
	"sync"
)

func CompareSources(left, right func() (string, error)) (string, error) {
	ch := make(chan struct{string; error; bool}, 2)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		str, err := left()
		ch <- struct{string; error; bool}{str, err, true}
		wg.Done()
	}()
	go func() {
		str, err := right()
		ch <- struct{string; error; bool}{str, err, false}
		wg.Done()
	}()
	
	wg.Wait()
	var num1, num2 struct{string; error; bool}
	num1 = <- ch
	num2 = <- ch
	if !num1.bool{
		num1, num2 = num2, num1
	}
	if num2.error != nil{
		if num1.error != nil{
			return "", num1.error
		}else{
			return "", num2.error
		}
	}
	if num1.string != num2.string{
		return "", errors.New("разные результаты")
	}
	return num1.string, nil
	
}