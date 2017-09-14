package main

import "fmt"

func main() {
	a := 17
	fmt.Println(a)

	var b *int = &a
	fmt.Println(b)
	fmt.Println(*b)

	*b = 19
	fmt.Println(a)
}