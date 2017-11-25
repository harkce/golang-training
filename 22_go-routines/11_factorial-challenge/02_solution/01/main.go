package main

import (
	"fmt"
)

func main() {
	for n := range factorial(25) {
		fmt.Println("Result:", n)
	}
}

func factorial(n int) chan int {
	out := make(chan int)
	go func() {
		result := 1
		for i := n; i > 0; i-- {
			result *= i
		}
		out <- result
		close(out)
	}()
	return out
}
