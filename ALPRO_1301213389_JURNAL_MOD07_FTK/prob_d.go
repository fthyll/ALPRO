package main

import "fmt"

func main() {
	var turun, pengunjung, pengunjung_prev, total int
	var valid bool
	i := 0
	for turun < 3 {
		i++
		fmt.Printf("Hari %d : ", i)
		fmt.Scanln(&pengunjung)
		valid = pengunjung >= 0 && pengunjung <= 100

		for !valid {
			fmt.Printf("Hari %d : ", i)
			fmt.Scanln(&pengunjung)
			valid = pengunjung > 0 && pengunjung <= 100
		}
		total += pengunjung

		if pengunjung_prev > pengunjung {
			turun += 1
		} else {
			turun = 0
		}

		pengunjung_prev = pengunjung
	}
	fmt.Printf("Museum buka selama %d hari \n", i)
	fmt.Printf("Rata-rata pengunjung %.2f orang", float64(total)/float64(i))
}

// Muhammad Fatih Yumna LL IF45-05 1301213389
