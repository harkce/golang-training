package main

import "fmt"

func main() {
	greetings := []string{
		"Good morning!",
		"Bonjour!",
		"Dias!",
		"Bongiorno!",
		"Ohayo!",
		"Selamat pagi!",
		"Gutten morgen!",
	}

	fmt.Println("[1:2]", greetings[1:2])
	fmt.Println("[:2]", greetings[:2])
	fmt.Println("[5:]", greetings[5:])
	fmt.Println("[:]", greetings[:])
}
