package main

import "fmt"

func FPB(a, b int) int {

	for b != 0 {
		i := b
		b = a % b
		a = i
	}
	return a
}

func main() {
	var a, b, KPK int
	fmt.Scan(&a, &b)

	KPK = a * b / FPB(a, b)
	fmt.Print(KPK)
}

// Muhammad Fatih Yumna LL 1301213389
