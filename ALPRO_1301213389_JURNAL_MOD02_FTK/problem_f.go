package main

import "fmt"

func main() {
	var nilai, jumlah, n int
	var rata2 float64

	fmt.Scan(&nilai)
	jumlah = 0
	n = 0
	for nilai != -1 {
		n = n + 1
		jumlah = jumlah + nilai
	}
	fmt.Scan(&nilai)
	if n == 0 {
		rata2 = 0.0
	} else {
		rata2 = float64(jumlah) / float64(n)

	}
	fmt.Println(rata2)
}

// Muhammad Fatih Yumna LL 1301213389 IF-45-05
