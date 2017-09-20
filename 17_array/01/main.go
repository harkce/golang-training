package main

import "fmt"

func main() {
	var x [58]int
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[17])
	x[17] = 19
	fmt.Println(x[17])
}