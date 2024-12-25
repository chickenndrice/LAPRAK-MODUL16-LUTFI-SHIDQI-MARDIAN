package main

import "fmt"
func main() {
	var total, hitung, bil float64
	for {
		fmt.Scan(&bil)
		if bil == 9999 {
			break
		}
		total += bil
		hitung++
	}
	if hitung > 0 {
		fmt.Printf("Rata-rata dari bilangan yang diinput: %.2f\n", total/hitung)
	} else {
		fmt.Println("Tidak ada bilangan yang diinput untuk dihitung.")
	}
}