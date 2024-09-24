package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 1; i <= 5; i++ {
		go func(i int) {
			defer wg.Done()
			for j := 1; j <= 10; j++ {
				fmt.Printf("Горутина %v отработала \n", i)
			}
		}(i)
	}
	wg.Wait()
}
