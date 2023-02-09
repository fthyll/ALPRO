package main

import "fmt"

func main() {
	var n, x, res int
	res = 0
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&x)
		for !(x >= 0 && x <= 9) {
			fmt.Scanln(&x)
		}
		res += x
	}
	fmt.Println(res)
}

// Muhammad Fatih Yumna LL 1301213389 IF-45-05
