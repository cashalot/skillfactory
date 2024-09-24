package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Фоновая горутина для вывода текущего времени каждую секунду
	go func() {
		for {
			fmt.Println("Current time:", time.Now().Format("15:04:05"))
			time.Sleep(time.Second)
		}
	}()

	// Главная горутина занимается приёмом сообщений из каналов
	for {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received from channel 1:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received from channel 2:", msg2)
		}
	}
}
