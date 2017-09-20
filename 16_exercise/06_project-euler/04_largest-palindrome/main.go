// Largest palindrome product
// Problem 4
// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
//
// Find the largest palindrome made from the product of two 3-digit numbers.

package main

import (
	"github.com/harkce/golang-training/02_package/stringutil"
	"fmt"
	"strconv"
)

func isPalindrome(n int) bool {
	return strconv.Itoa(n) == stringutil.Reverse(strconv.Itoa(n))
}

func main() {
	var max int
	a := 999
	for a > 99 {
		b := 999
		for b >= a {
			prod := b * a
			if isPalindrome(prod) && prod > max {
				max = prod
			}
			b--
		}
		a--
	}
	fmt.Println(max)
}