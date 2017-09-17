package main

import "fmt"

func main() {
	fmt.Println(greet("Habib ", "Fikri "))
}

func greet(fName, lName string) (s string) {
	s = fmt.Sprint(fName, lName)
	return
}