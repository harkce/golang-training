package main

import "fmt"

func main() {
	iSlice := make([]int, 0, 5)

	fmt.Println("--------------")
	fmt.Println(iSlice)
	fmt.Println(len(iSlice))
	fmt.Println(cap(iSlice))
	fmt.Println("--------------")
	for i := 0; i < 80; i++ {
		iSlice = append(iSlice, i)
		fmt.Println("Len :", len(iSlice), "Cap :", cap(iSlice), "Val :", i)
	}
}
