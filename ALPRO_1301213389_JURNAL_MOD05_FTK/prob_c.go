package main

import "fmt"

const (
	num = 0
	dom = 0
	i   = 0
)

func main() {
	var index string
	var i, num, dom, nilaiIndex, sks, ambilMK int
	var ips float64

	fmt.Scanln(&ambilMK)

	for i < ambilMK {
		fmt.Scanln(&index, &sks)

		if (index == "A" || index == "B" || index == "C" || index == "D" || index == "E") && sks > 0 {
			switch index {
			case "A":
				nilaiIndex = 4
			case "B":
				nilaiIndex = 3
			case "C":
				nilaiIndex = 2
			case "D":
				nilaiIndex = 1
			case "E":
				nilaiIndex = 0
			}

			num += nilaiIndex * sks
			dom += sks

			i++
		}

	}

	ips = float64(num) / float64(dom)

	fmt.Println("Keluaran", ips)
}

// Muhammad Fatih Yumna Lajuwirdi Lirrahman - IF 45 05 - 1301213389
