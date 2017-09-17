package main

import "fmt"

func main() {
	greet("Habib", "Bandung")
	greet("Yasni", "Garut")
}

func greet(name string, city string) {
	fmt.Printf("Hello %s from %s!\n", name, city)
}