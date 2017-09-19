package main

import "fmt"

func main() {
	name := "Habib"
	changeMe(name)
	fmt.Println(name)
}

func changeMe(z string) {
	z = "Cingkleung"
}