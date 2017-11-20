// Smallest multiple
// Problem 5
// 2520 is the smallest number that can be divided by each of the numbers
// from 1 to 10 without any remainder.
//
// What is the smallest positive number that is evenly divisible by all of
// the numbers from 1 to 20?

package main

import "fmt"

func small(n int) int {
	x := n
	small := false
	for !small {
		for i := n; i >= 1; i-- {
			if x%i != 0 {
				small = false
				break
			}
			small = true
		}
		if !small {
			x += n
		}
	}
	return x
}

func main() {
	fmt.Println(small(20))
}
