// Special Pythagorean triplet
// Problem 9
// A Pythagorean triplet is a set of three natural numbers, a < b < c, for
// which,
//
// a^2 + b^2 = c^2
// For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.
//
// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.

package main

import (
	"fmt"
	"math"
)

func tripletSum(n int) int {
	i := 1
	for i < n {
		j := 1
		for j < n-i {
			k := n - i - j
			if math.Pow(float64(i), 2)+math.Pow(float64(j), 2) == math.Pow(float64(k), 2) {
				return i * j * k
			}
			j++
		}
		i++
	}
	return 0
}

func main() {
	fmt.Println(tripletSum(1000))
}
