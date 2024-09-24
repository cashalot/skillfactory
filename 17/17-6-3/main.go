package main

import (
	"fmt"
)

func main() {
	numChan := make(chan int) // Канал для передачи чисел
	done := make(chan bool)   // Канал для завершения работы

	// Горутина для отправки чисел
	go func() {
		for i := 1; i <= 100; i++ {
			numChan <- i // Отправляем число в канал
		}
		close(numChan) // Закрываем канал после отправки всех чисел
	}()

	// Горутина для приёма чисел и их вывода в консоль
	go func() {
		for num := range numChan { // Принимаем числа из канала
			fmt.Println(num)
		}
		done <- true // Отправляем сигнал завершения работы
	}()

	<-done // Ожидаем завершения приёма чисел
}
