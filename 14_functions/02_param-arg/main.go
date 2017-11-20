package main

import "fmt"

func main() {
	greet("Habib")
	greet("Yasni")
}

func greet(name string) {
	fmt.Println("Hello", name)
}
