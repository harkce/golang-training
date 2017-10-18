package main

import "fmt"

func main() {
	roomId := make([]int, 3)
	roomId[0] = 7
	roomId[1] = 42
	roomId[2] = 12
	fmt.Println(roomId[0])
	fmt.Println(roomId[1])
	fmt.Println(roomId[2])

	greeting := make([]string, 3, 5)
	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "dias!"

	fmt.Println(greeting[2])
}
