package main

func ReceiveIn(in <-chan int, n int)[]int{
	res := make([]int, 0, n)

	for range n{
		val := <-in
		res = append(res, val)
	}
	return res
}