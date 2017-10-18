package main

import "fmt"

func main() {
	hariKerja := []string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat"}
	hariLibur := []string{"Sabtu", "Minggu"}

	hariKerja = append(hariKerja, hariLibur...)
	fmt.Println("Hari kerja lembur:", hariKerja)
}
