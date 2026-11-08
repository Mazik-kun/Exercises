package main

import "errors"

func ReadFrame(in <-chan int) ([]int, error) {
	var err error
	var res []int

	k, ok := <-in
	if !ok {
		return nil, errors.New("missing header")
	}

	if k < 0 || k > 1024 {
		err = errors.New("invalid length")
	}

	for i := 0; i < k; i++ {
		v, ok := <-in
		if !ok {
			if err == nil {
				err = errors.New("insufficient elements")
			}
			break
		}
		res = append(res, v)
	}
	
	for range in {
		if err == nil {
			err = errors.New("excess elements")
		}
	}

	if err != nil {
		return nil, err
	}
	return res, nil
}