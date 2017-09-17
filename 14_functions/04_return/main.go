package main

import "fmt"

func main() {
	fmt.Println(greet("Habib ", "Fikri "))
}

func greet(fName, lName string) string {
	return fmt.Sprint(fName, lName)
}