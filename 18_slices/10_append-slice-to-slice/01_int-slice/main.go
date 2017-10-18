package main

import "fmt"

func main() {
	sliceOne := []int{1, 1, 1, 1, 1}
	otherSlice := []int{2, 2, 2, 2}

	sliceOne = append(sliceOne, otherSlice...)
	fmt.Println(sliceOne)
}
