package main

import "fmt"

func main() {
	var nama, namaMin, namaMax string
	var p int
	var nilai, jumlah, rata, min, max float64
	namaMin = ""
	namaMax = ""
	fmt.Scan(&nama)
	for nama != "STOP" {
		fmt.Scan(&p)
		jumlah = 0
		for i := 1; i <= 3*p; i++ {
			fmt.Scan(&nilai)
			jumlah += nilai
		}
		rata = jumlah / float64(3*p)
		if namaMin == "" || namaMin != "" && min > rata {
			min = rata
			namaMin = nama
		}
		if namaMax == "" || namaMax != "" && max < rata {
			max = rata
			namaMax = nama
		}
		fmt.Scan(&nama)
	}
	fmt.Printf("%s %.2f\n", namaMax, max)
	fmt.Printf("%s %.2f\n", namaMin, min)
}

// Muhammad Fatih Yumna LL IF45-05 1301213389
