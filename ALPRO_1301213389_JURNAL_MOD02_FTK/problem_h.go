package main

import "fmt"

func main() {
	var r, t, hitungVolume, phi float64

	phi = 3.14

	fmt.Print("Masukan Angaka: ")
	fmt.Scanln(&r, &t)

	hitungVolume = r * r * phi * t

	fmt.Println("Volume ", hitungVolume)
}

// Muhammad Fatih Yumna LL 1301213389 IF-45-05
