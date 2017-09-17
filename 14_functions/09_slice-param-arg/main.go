package main

import "fmt"

func main() {
	data := []float64{4, 5, 3, 7, 8}
	fmt.Println(average(data))
}

func average(sf []float64) float64 {
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}