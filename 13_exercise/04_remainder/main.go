package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("1st number: ")
	fmt.Scanln(&a)
	fmt.Print("2nd number: ")
	fmt.Scanln(&b)

	if a > b {
		fmt.Printf("%d mod %d = %d\n", a, b, a % b)
	} else if a < b {
		fmt.Printf("%d mod %d = %d\n", b, a, b % a)
	} else {
		fmt.Printf("%d mod %d = %d\n", a, b, a % b)
	}
}