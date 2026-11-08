package main

import "sync"

func RunPairThenFinish(first, second, finish func()) <-chan struct{} {
	done := make(chan struct{})
	go func() {
		var wg sync.WaitGroup
		wg.Add(2)

		go func() { defer wg.Done(); first() }()
		go func() { defer wg.Done(); second() }()

		wg.Wait()   // Ждем выполнения обеих функций
		finish()    // Вызываем финишер
		close(done) // Сигнализируем о завершении
	}()
	return done
}