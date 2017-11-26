package main

import (
	"fmt"
	"math/big"
)

func main() {
	x := factorial(1000)
	res := big.NewInt(1)
	for n := range multiply(x) {
		res.Mul(res, big.NewInt(int64(n)))
	}
	fmt.Println("Result:", res)
}

func factorial(n int) chan []int {
	out := make(chan []int)
	go func() {
		for i := n; i > 0; i -= 2 {
			x := make([]int, 0)
			x = append(x, i)
			x = append(x, i-1)
			out <- x
		}
		close(out)
	}()
	return out
}

func multiply(c chan []int) chan int {
	out := make(chan int)
	go func() {
		var result int
		for n := range c {
			if n[1] == 0 {
				result = 1
			} else {
				result = n[0] * n[1]
			}
			out <- result
		}
		close(out)
	}()
	return out
}
