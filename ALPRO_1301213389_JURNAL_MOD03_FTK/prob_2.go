package main

import "fmt"

func main() {
	var x, y, z string
	var hasil bool
	fmt.Scanln(&x, &y, &z)
	hasil = x == y || x == z || y == z
	fmt.Println(hasil)
}

// Muhammad Fatih Yumna Lajuwirdi Lirrahman 1301213389 IF4505
