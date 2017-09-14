package main

import "fmt"

const celsiusToReamur float64 = 0.8

func main() {
	var celsius float64
	fmt.Print("Enter celsius swam: ")
	fmt.Scan(&celsius)
	reamur := celsius * celsiusToReamur
	fmt.Println(celsius, " celsius is ", reamur, " reamur.")
}
