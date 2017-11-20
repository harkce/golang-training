package main

import "fmt"

func half(n int) (int, bool) {
	return n / 2, n%2 == 0
}

func main() {
	var x int
	fmt.Print("Input a number: ")
	fmt.Scan(&x)
	var a, b = half(x)
	fmt.Printf("half(%d) = %v, %v\n", x, a, b)
}
