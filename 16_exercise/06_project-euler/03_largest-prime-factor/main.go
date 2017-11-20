// Largest prime factor
// Problem 3
// The prime factors of 13195 are 5, 7, 13 and 29.
//
// What is the largest prime factor of the number 600851475143 ?

package main

import (
	"fmt"
	"math"
)

func primeFactor(n int) {
	for n%2 == 0 {
		fmt.Printf("%d ", 2)
		n /= 2
	}

	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		for n%i == 0 {
			fmt.Printf("%d ", i)
			n /= i
		}
	}

	if n > 2 {
		fmt.Printf("%d ", n)
	}
}

func main() {
	primeFactor(3628800)
	fmt.Println()
}
