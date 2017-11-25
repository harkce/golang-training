package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
			fmt.Println("halo")
		}
	}()

	go func() {
		for {
			fmt.Println(<-c)
			fmt.Println("ambil halo")
		}
	}()

	time.Sleep(time.Second)
}
