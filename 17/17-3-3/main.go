package main

import (
	"fmt"
	"sync"
)

const step int = 1
const iterationAmount int = 1000

func main() {
	var counter int = 0
	var mutex = sync.Mutex{}
	var c = sync.NewCond(&mutex)

	increment := func() {
		mutex.Lock()
		counter += step
		if counter == iterationAmount {
			c.Signal()
		}
		mutex.Unlock()
	}

	for i := 1; i <= iterationAmount; i++ {
		go increment()
	}

	mutex.Lock()
	for counter < iterationAmount {
		c.Wait()
	}
	mutex.Unlock()

	fmt.Println(counter)
}
