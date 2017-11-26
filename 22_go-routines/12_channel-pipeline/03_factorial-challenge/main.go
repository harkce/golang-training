package main

import "fmt"

func main() {
	fmt.Println(<-factorial(4))
}

func factorial(n int) chan int {
	out := make(chan int)
	go func() {
		result := 1
		for i := n; i > 1; i-- {
			result *= i
		}
		out <- result
		close(out)
	}()
	return out
}
