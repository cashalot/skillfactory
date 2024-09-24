package main

import (
	"fmt"
	"sync"
)

// Структура счётчика
type Counter struct {
	value  int           // Текущее значение счётчика
	goal   int           // Конечное значение счётчика
	mu     sync.Mutex    // Мьютекс для защиты данных
	done   chan struct{} // Канал для завершения работы
	isDone bool          // Флаг завершения
}

// Метод для увеличения значения счётчика
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Проверяем, достигнуто ли конечное значение
	if c.value < c.goal {
		c.value++ // Увеличиваем значение счётчика
		fmt.Println("Counter:", c.value)
	}

	// Проверяем, достигнуто ли конечное значение
	if c.value >= c.goal && !c.isDone {
		c.isDone = true
		close(c.done) // Закрываем канал завершения
	}
}

// Функция для запуска горутин
func (c *Counter) StartWorkers(numGoroutines int) {
	for i := 0; i < numGoroutines; i++ {
		go func() {
			for {
				c.mu.Lock()
				if c.value >= c.goal {
					c.mu.Unlock()
					return
				}
				c.mu.Unlock()

				c.Increment()
			}
		}()
	}
}

func main() {
	// Настраиваем конечное значение счётчика и количество горутин
	finalValue := 100  // Конечное значение счётчика
	numGoroutines := 5 // Количество горутин
	counter := &Counter{
		value: 0,
		goal:  finalValue,
		done:  make(chan struct{}),
	}

	// Запускаем горутины
	counter.StartWorkers(numGoroutines)

	// Ждём завершения
	<-counter.done
	fmt.Println("Final counter value reached:", counter.value)
}
