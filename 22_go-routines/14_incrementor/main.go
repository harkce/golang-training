package main

import (
	"fmt"
)

func main() {
	f := incrementor(2)

	var counter int
	for n := range f {
		counter++
		fmt.Println(n)
	}
	fmt.Println("Final coun:", counter)
}

func incrementor(n int) chan string {
	out := make(chan string)
	done := make(chan bool)
	for i := 0; i < n; i++ {
		go func(i int) {
			for j := 0; j < 20; j++ {
				out <- fmt.Sprintf("Process %d is printing %d", i, j)
			}
			done <- true
		}(i)
	}

	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(out)
	}()
	return out
}
