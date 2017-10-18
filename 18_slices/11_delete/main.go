package main

import "fmt"

func main() {
	hariKerja := []string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat"}
	hariLibur := []string{"Sabtu", "Minggu"}

	hariLembur := append(hariKerja, hariLibur...)
	fmt.Println("Hari kerja lembur:", hariLembur)

	hariCuti := append(hariKerja[1:])
	fmt.Println(hariCuti)
}
