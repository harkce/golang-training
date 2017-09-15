package main

import "fmt"

func main() {
	switch "Habib" {
	case "Cingkleung":
		fmt.Println("Naon kleung?")
	case "Habib":
		fmt.Println("Naon bib?")
		fallthrough
	case "Fikri":
		fmt.Println("Naon fik?")
		fallthrough
	default:
		fmt.Println("Saha maneh?")
	}
}
