package main

import "fmt"

func main() {
	iSlice := make([]int, 3, 5)

	iSlice[0] = 3
	iSlice[1] = 5
	iSlice[2] = 7
	iSlice = append(iSlice, 9)
	iSlice = append(iSlice, 11)
	iSlice = append(iSlice, 13)
	iSlice = append(iSlice, 15)

	fmt.Println(iSlice)
	fmt.Println(len(iSlice))
	fmt.Println(cap(iSlice))
}
