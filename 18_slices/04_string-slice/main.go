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

	for i, g := range greetings {
		fmt.Println(i, g)
	}

	for i := 0; i < len(greetings); i++ {
		fmt.Println(i, greetings[i])
	}
}
