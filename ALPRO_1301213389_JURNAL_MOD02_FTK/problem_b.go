package main

import "fmt"

func main() {

	var r, luas, phi float64

	phi = 3.14

	fmt.Print("Masukkan panjang jari-jari lingkaran: ")
	fmt.Scanln(&r)

	luas = phi * r * r

	fmt.Println("Luas lingkaran adalah ", luas)

}

// Muhammad Fatih Yumna LL 1301213389 IF-45-05
