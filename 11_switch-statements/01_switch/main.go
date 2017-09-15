package main

import "fmt"

func main() {
	switch "Habib" {
	case "Cingkleung":
		fmt.Println("Naon kleung?")
	case "Habib":
		fmt.Println("Naon bib?")
	case "Fikri":
		fmt.Println("Naon fik?")
	default:
		fmt.Println("Saha maneh?")
	}
}
