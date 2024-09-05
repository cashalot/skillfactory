package main

import "fmt"

func main() {
	sum := func(first, second int) int {
		return first + second
	}(5, 2)
	fmt.Println(sum)
}
