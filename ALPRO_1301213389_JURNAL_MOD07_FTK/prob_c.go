package main

import "fmt"

func pangkat(a, b int) int {
	var res int
	res = 1
	for i := 1; i <= b; i++ {
		res *= a
	}
	return res
}
func konversi(bin int, des *int) {
	var x, i int
	for bin > 0 {
		x = bin % 10
		*des += x * pangkat(2, i)
		i++
		bin = bin / 10
	}
}
func main() {
	var n, res int
	fmt.Scan(&n)
	konversi(n, &res)
	fmt.Println(res)
}

// Muhammad Fatih Yumna LL IF45-05 1301213389
