package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Шаг наращивания счётчика
const step int64 = 1

// Конечное значение счётчика
const endCounterValue int64 = 1000

// Количество горутин
const goroutineCount = 10

func main() {

	var counter int64 = 0
	var wg sync.WaitGroup

	increment := func(iterations int) {
		defer wg.Done()
		for i := 0; i < iterations; i++ {
			atomic.AddInt64(&counter, step)
		}
	}

	// Количество инкрементов на одну горутину
	iterationsPerGoroutine := int(endCounterValue / goroutineCount)

	for i := 0; i < goroutineCount; i++ {
		wg.Add(1)
		go increment(iterationsPerGoroutine)
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	// Печатаем результат, надеясь, что будет 1000
	fmt.Println(counter)
}
