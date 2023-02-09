package main

import "fmt"

func main() {
	var berat int
	var beratKg, sisaG, biayaKG, sisabiayaG, totBiaya int

	fmt.Print("Berat parsel (gram) : ")
	fmt.Scan(&berat)

	beratKg = berat / 1000
	sisaG = berat % 1000

	biayaKG = beratKg * 10000

	if sisaG >= 500 {
		sisabiayaG = sisaG * 5
	} else {
		sisabiayaG = sisaG * 15
	}

	totBiaya = biayaKG

	if !(beratKg >= 10) {
		totBiaya += sisabiayaG
	}

	fmt.Printf("Detail berat : %d kg + %d gr\n", beratKg, sisaG)
	fmt.Printf("Detail biaya : Rp. %d + Rp. %d\n", biayaKG, sisabiayaG)
	fmt.Printf("Total biaya : Rp. %d\n", totBiaya)
}

// Muhammad Fatih Yumna Lajuwirdi Lirrahman - IF 45 05 - 1301213389
