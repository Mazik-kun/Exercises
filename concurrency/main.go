package main

import (
	"fmt"
	"time"
)

	func source(values ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, value := range values {
			out <- value
		}
	}()

	return out
	}

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
			time.Sleep(10 * time.Millisecond)
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
		time.Sleep(1 * time.Millisecond)
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

	ch6 := make(chan string, 7)
	ch6 <- "apple"
	ch6 <- "Apple"
	ch6 <- ""
	ch6 <- "apple"
	ch6 <- "banana"
	ch6 <- ""
	close(ch6)
	fmt.Println(WordCounts(ch6))

	ch17 := make(chan int, 4)
	ch17 <- 1
	ch17 <- 2
	ch17 <- 5
	ch17 <- 4
	close(ch17)
	fmt.Println(IsNonDecreasing(ch17))

	square := func(x int) int { return x * x }
	ch18 := Generate(3, square)

	for v := range ch18 {
		fmt.Println(v) 
	}

		ch19_1 := make(chan struct{})
	ch19_2 := make(chan struct{})
	ch19_3 := make(chan struct{})
	go func() {
		close(ch19_1)
		close(ch19_2)
		close(ch19_3)
	}()
	AwaitSignals([]<-chan struct{}{ch19_1, ch19_2, ch19_3})
	fmt.Println("AwaitSignals completed")

	fmt.Println("разделитель")

	ch20_ready := make(chan struct{})
	ch20_out := ReadAfterReady(ch20_ready, func() int { return 100 })
	go func() {
		ch20_ready <- struct{}{}
	}()
	fmt.Println(<-ch20_out)

	fmt.Println("разделитель")

	ch21_done := RunPairThenFinish(
		func() { fmt.Println("first") },
		func() { fmt.Println("second") },
		func() { fmt.Println("finish") },
	)
	<-ch21_done

	fmt.Println("разделитель")

	ch22_out := StartBatch([]func() int{
		func() int { return 1 },
		func() int { return 2 },
		func() int { return 3 },
	})
	for v := range ch22_out {
		fmt.Println(v)
	}

	fmt.Println("разделитель")

	ch23_out := StartCalculation(func() (int, error) {
		return 42, nil
	})
	fmt.Println(<-ch23_out)

	fmt.Println("разделитель")

	ch24_out := make(chan int)
	ch24_processed := make(chan struct{})
	go func() {
		val := <-ch24_out
		fmt.Println("Received:", val)
		close(ch24_processed)
	}()
	SendAndWait(55, ch24_out, ch24_processed)

	fmt.Println("разделитель")

	ch25_1 := make(chan int, 3)
	ch25_1 <- 2
	ch25_1 <- 10
	ch25_1 <- 20
	close(ch25_1)
	res25, err25 := ReadFrame(ch25_1)
	fmt.Println(res25, err25)

	ch25_2 := make(chan int, 1)
	ch25_2 <- -1
	close(ch25_2)
	_, err25_2 := ReadFrame(ch25_2)
	fmt.Println(err25_2)

	fmt.Println("разделитель")

	ch26 := make(chan string, 5)
	ch26 <- "a"
	ch26 <- "start"
	ch26 <- "b"
	ch26 <- "start"
	ch26 <- "c"
	close(ch26)
	fmt.Println(AfterMarker(ch26, "start"))

	start := time.Now()
	local, background := RunAlongside(
		func() int {
			fmt.Println("local завершилась")
			return 10
		},
		func() int {
			time.Sleep(2 * time.Second)
			fmt.Println("background завершилась")
			return 20
		},
	)
	fmt.Println("Результаты:", local, background)
	fmt.Println("Время:", time.Since(start))

	pairs, err := PairStreams(
		source(10, 20),
		source(1, 2),
	)
	fmt.Println(pairs, err)
}