package main

import "fmt"

func max(ns ...int) int {
	var max int
	for _, v := range ns {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	fmt.Println(max(1, 2, 3, 4, 5))
}
