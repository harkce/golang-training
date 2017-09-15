package main

import "fmt"

func main() {
	if false {
		fmt.Println("This never be seen")
	} else if true {
		fmt.Println("This is true!")
	} else  if true {
		fmt.Println("Already true. Doh!")
	} else if false {
		fmt.Println("Fix your problem")
	} else {
		fmt.Println("The last statement")
	}
}