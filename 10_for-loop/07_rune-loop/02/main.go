package main

import "fmt"

func main() {
	for i := 6000; i <= 6100; i++ {
		fmt.Printf("%v - %v - %v\n", i, string(i), []byte(string(i)))
	}
}
