package main

import "fmt"

func main() {
	var a, b, n, pos, neg int
	pos = 0
	neg = 0
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a, &b)
		if a > 0 {
			pos += a
		} else {
			neg += a
		}

		if b > 0 {
			pos += b
		} else {
			neg += b
		}
	}
	fmt.Println("Negatif:", neg)
	fmt.Println("Positif:", pos)
}

// Muhammad Fatih Yumna LL IF45-05 1301213389
