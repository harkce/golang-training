package main

import "fmt"

func main() {
	if false {
		fmt.Println("This never be seen")
	} else if true {
		fmt.Println("This is true!")
	} else {
		fmt.Println("Already true. Doh!")
	}
}
