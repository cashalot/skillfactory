package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Инициализация генератора случайных чисел
	rand.Seed(time.Now().UnixNano())

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 1; i <= 100; i++ {
			ch1 <- i
			time.Sleep(time.Millisecond * 100)
		}
		close(ch1)
	}()

	go func() {
		for i := 1001; i <= 1050; i++ {
			ch2 <- i
			time.Sleep(time.Millisecond * 200)
		}
		close(ch2)
	}()

	for {
		select {
		case msg, ok := <-ch1:
			if ok {
				fmt.Printf("Received from channel1: %d\n", msg)
			} else {
				ch1 = nil // Убираем закрытый канал
			}
		case msg, ok := <-ch2:
			if ok {
				fmt.Printf("Received from channel2: %d\n", msg)
			} else {
				ch2 = nil // Убираем закрытый канал
			}
		default:
			// Увеличиваем вероятность выбора первого канала
			if rand.Intn(3) == 0 && ch2 != nil {
				continue
			}
		}

		// Если оба канала закрыты, выходим из цикла
		if ch1 == nil && ch2 == nil {
			break
		}
	}
}
