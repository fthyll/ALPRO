package main

import "fmt"

func main() {
	var t1, t2, t3, n int
	var jumlah int = 0
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&t1, &t2, &t3)
		if t1+t2+t3 >= 2 {
			jumlah = jumlah + 1
		}
	}
	fmt.Print(jumlah)
}

// Muhammad Fatih Yumna LL 1301213389 IF-45-05
