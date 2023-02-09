package main

import "fmt"

func main() {
	var n int
	var a1, a2, a3 int
	var b1, b2, b3 int
	var hsl1, hsl2 int

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a1, &a2, &a3, &b1, &b2, &b3)
		hsl1 = hsl1 + a1 + a2 + a3
		hsl2 = hsl2 + b1 + b2 + b3
	}
	fmt.Println("Petinju A:", hsl1, "dan petinju B:", hsl2)
	if hsl1 > hsl2 {
		fmt.Print("Pemenang adalah Petinju A")
	} else if hsl1 < hsl2 {
		fmt.Print("Pemenang adalah Petinju B")
	} else {
		fmt.Print("Draw")
	}
}

// Muhammad Fatih Yumna LL 1301213389
