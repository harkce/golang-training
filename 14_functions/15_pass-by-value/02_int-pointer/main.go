package main

import "fmt"

func main() {
	age := 42
	changeMe(&age)
	fmt.Println(age)
}

func changeMe(z *int) {
	*z = 24
}
