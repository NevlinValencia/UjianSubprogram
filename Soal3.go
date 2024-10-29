package main

import "fmt"

func sesiPelatihan(hari int, p int, q int) int {
	if hari > 365 {
		return 0
	}
	if hari%p == 0 && hari%q != 0 {
		return 1 + sesiPelatihan(hari+1, p, q)
	}
	return sesiPelatihan(hari+1, p, q)
}

func main() {
	var p, q int

	fmt.Print("Masukkan nilai p (kelipatan sesi pelatihan): ")
	fmt.Scan(&p)
	fmt.Print("Masukkan nilai q (bukan kelipatan): ")
	fmt.Scan(&q)

	totalSesi := sesiPelatihan(1, p, q)

	fmt.Printf("Jumlah sesi pelatihan dalam setahun: %d\n", totalSesi)
}
