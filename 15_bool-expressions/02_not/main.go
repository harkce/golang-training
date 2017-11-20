package main

import "fmt"

func main() {
	if !true {
		fmt.Println("Buried")
	}

	if !false {
		fmt.Println("Hi there")
	}
}
