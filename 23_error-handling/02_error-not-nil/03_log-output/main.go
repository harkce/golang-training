package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	lg, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(lg)
}

func main() {
	_, err := os.Open("nothing.txt")
	if err != nil {
		log.Println(err)
	}
}
