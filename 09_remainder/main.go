package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		fmt.Print(i, " - ")
		if i%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}
}
