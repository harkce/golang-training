package main

import "fmt"

func main() {
	iSlice := []int{2, 3, 5, 7, 11, 13, 17}

	for k, v := range iSlice {
		fmt.Println(k, "-", v)
	}
}
