package main

import "fmt"

func main() {
	if !true {
		fmt.Println("It's impossible to go here!")
	}
	if !false {
		fmt.Println("You can do it!")
	}
}
