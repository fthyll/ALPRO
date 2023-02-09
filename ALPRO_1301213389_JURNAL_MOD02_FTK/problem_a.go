package main

import "fmt"

func main() {
	var s string
	var a, b int

	fmt.Print("Kata : ")
	fmt.Scan(&s)
	fmt.Print("Nilai A :")
	fmt.Scan(&a)
	fmt.Print("Nilai B :")
	fmt.Scan(&b)
	fmt.Print("Kata = ", s)
	fmt.Print(" Jumlah = ", hasil_penjumlahan(a, b))

}

func hasil_penjumlahan(a int, b int) int {
	return a + b
}

// Muhammad Fatih Yumna LL 1301213389 IF-45-05
