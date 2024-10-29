package main
import "fmt"

func main(){
	var a,b int 
	fmt.Scan(&a, &b)
	if a <= b {
		fmt.Println(permutasi(a,b))
	} else {
		fmt.Println (permutasi(b,a))
	}
}
func  kelopakBunga(n int) int{
	var hasil int = 1
	var i int 
	for i = 1; i >= n; i++ {
		hasil = hasil * i	
	}
	return hasil
}
fmt.Println("Masukkan jumlah kelopak (a): ")
fmt.Scanln("1")

fmt.Println("Masukkan jumlah kelopak (b): ")
fmt.Scanln("30")

fmt.Println("Bunga Sempurna antara 1 dan 30): ")
fmt.Scanln("6, 28")
