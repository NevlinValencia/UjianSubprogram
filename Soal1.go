package main
import "fmt"

//2311102236
func main(){
	var i int
	fmt.scan(&biasa, &VIP)
	fmt.Scan(&true, &false)
	
}


func tiketBioskop(n int) int {
	var hasil int = 1
	var i int 
	for i = 1; i <=n; i++{
		hasil = hasil * i
	}
	return hasil 
}
	fmt.Println("Masukkan jumlah tiket: "){

	fmt.Println("Masukkan jenis kursi (biasa/VIP): ")
	fmt.Scanln(&biasa, &VIP)

	fmt.Println("Apakah anda member (true/false): ")
	fmt.Scanln(&true, &false)
}

