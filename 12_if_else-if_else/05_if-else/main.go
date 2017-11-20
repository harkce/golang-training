package main

import "fmt"

func main() {
	if false {
		fmt.Println("This never be seen")
	} else {
		fmt.Println("This is true!")
	}
}
