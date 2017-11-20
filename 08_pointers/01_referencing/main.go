package main

import "fmt"

func main() {
	a := 17

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a

	fmt.Println(b)
	fmt.Printf("%T \n", b)
}
