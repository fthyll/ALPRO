package main

import "fmt"

func hitDig(bil int) {
	var n1, hasil int
	hasil = 0
	for bil != 0 {
		n1 = bil % 10
		hasil = hasil + n1
		bil = bil / 10
	}
	fmt.Println(hasil)
}

func main() {
	var x int
	fmt.Scan(&x)
	hitDig(x)
}
