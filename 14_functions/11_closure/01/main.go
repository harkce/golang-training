package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "The quick brown fox jumps over the lazy dog"
		fmt.Println(y)
	}
	//fmt.Println(y)
}
