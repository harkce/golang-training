package main

import "fmt"

func main() {
	i := 10923
	switch {
	case i % 3 == 0 && i % 5 == 0:
		fmt.Println("Bisa dibagi 3 dan 5")
	case i % 3 == 0:
		fmt.Println("Bisa dibagi 3")
	case i % 5 == 0:
		fmt.Println("Bisa dibagi 5")
	default:
		fmt.Println("Tidak bisa dibagi 3 dan 5")
	}
}
