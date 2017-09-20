// Multiples of 3 and 5
// Problem 1
// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
//
// Find the sum of all the multiples of 3 or 5 below 1000.

package main

import "fmt"

func multiple35(n int) bool {
	if n % 3 == 0 || n % 5 == 0 {
		return true
	}
	return false
}

func main() {
	var total int
	for i := 0; i < 1000; i++ {
		if multiple35(i) {
			total += i
		}
	}
	fmt.Println(total)
}