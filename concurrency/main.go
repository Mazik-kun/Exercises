package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		ch <- 4
		ch <- 1
		close(ch)
	}()

	fmt.Println(ReceiveIn(ch, 3))
	// fmt.Println(SumStream(ch))
	// fmt.Println(MinStream(ch))
	fmt.Println(ReadExactlyTwo(ch))

	nums := []int{1,4,6,5,3,7}
	ch1 :=PrepareNumbers(nums)
	fmt.Println(ReceiveIn(ch1, 3))


	ch2 := BufferedSource(nums, 4)

	var received []int
	for v := range ch2 {
		received = append(received, v)
	}
	
	fmt.Println("Получено из канала:", received)

	ch3 := make(chan int)

	nums3 := []int{1,2,4,5}
	go func(){
		for v := range ch3{
			fmt.Println("получено:", v)
		}
	}()
	SendAll(nums3, ch3)
	close(ch3)

	fmt.Println("разделитель")

	ch4 := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			ch4 <- i 
			time.Sleep(1000 * time.Millisecond)
		}
		close(ch4)
	}()
	// Drain(ch4) 
	fmt.Println(FirstK(ch4, 3))
	
	fmt.Println("разделитель")

	left := []int{1, 2, 1, 3, 1}
	right := []int{1, 4, 1, 5}
	target := 1
	cl, cr := CountInBoth(left, right, target)
	fmt.Println(cl, cr)
	fmt.Println("разделитель")

	errA, errB := RunChecks(
		func() error { return nil },
		func() error { return nil },
	)
	fmt.Println(errA, errB)
	fmt.Println("разделитель")

	newCh := StartValue(func() int {
		time.Sleep(1 * time.Second) 
		return 42
	})

	for v := range newCh {
		fmt.Println(v) 
	}
	fmt.Println("разделитель")
	ch5 := make(chan int, 3)
	ch5 <- 1
	ch5 <- 2
	ch5 <- 3
	close(ch5) 

	done := StartConsumer(ch5, func(v int) {
		time.Sleep(1 * time.Second)
		fmt.Println("Обработано:", v*10)
	})

	<-done
	fmt.Println("завершено")
	
	newDone := StartCompletion(func(){
		time.Sleep(1 * time.Second)
	})

	<-newDone
	fmt.Println("new завершено")
	<-newDone
	fmt.Println("new завершено 2")

}