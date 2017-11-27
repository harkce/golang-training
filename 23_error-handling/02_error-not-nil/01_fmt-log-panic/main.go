package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("nothing.txt")
	if err != nil {
		fmt.Println("Error euy", err)
		log.Println("Error mak", err)
		// log.Fatalln("Error cuk", err)
		// panic(err)
	}
	fmt.Println("ok")
}
