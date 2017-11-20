// 10001st prime
// Problem 7
// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see
// that the 6th prime is 13.
//
// What is the 10 001st prime number?

package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n%2 == 0 {
		return false
	}
	d := 3
	for d <= int(math.Sqrt(float64(n))) {
		if n%d == 0 {
			return false
		}
		d += 2
	}
	return true
}

func nthPrime(n int) int {
	prime := 2
	count := 1
	iter := 3
	for count < n {
		if isPrime(iter) {
			prime = iter
			count++
		}
		iter += 2
	}
	return prime
}

func main() {
	fmt.Println(nthPrime(10001))
}
