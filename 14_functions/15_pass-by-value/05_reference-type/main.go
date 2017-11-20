package main

import "fmt"

func main() {
	m := make(map[string]int)
	changeMe(m)
	fmt.Println(m["Habib"])
}

func changeMe(z map[string]int) {
	z["Habib"] = 17
}
