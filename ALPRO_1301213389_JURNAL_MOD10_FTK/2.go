package main

import "fmt"

type calon [20]int

func main() {
	var x calon
	var masuk, sah, pilihan, i int
	masuk = 0
	sah = 0
	fmt.Scan(&pilihan)
	for pilihan != 0 {
		masuk++
		if pilihan >= 1 && pilihan <= 20 {
			sah++
			x[pilihan-1] = x[pilihan-1] + 1
		}
		fmt.Scan(&pilihan)
	}
	fmt.Println("Suara masuk:", masuk)
	fmt.Println("Suara sah:", sah)
	for i = 0; i < 20; i++ {
		if x[i] != 0 {
			fmt.Printf("%v: %v\n", i+1, x[i])
		}
	}
}

// Muhammad Fatih Yumna LL IF4505 1301213389
