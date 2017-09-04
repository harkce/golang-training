package main

import "fmt"

var x int = 67

func main() {
	fmt.Println(x)
	foo()
}

func foo() {
	fmt.Println(x)
}
