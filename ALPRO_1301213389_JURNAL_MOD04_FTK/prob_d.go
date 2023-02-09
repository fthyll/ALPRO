package main

import "fmt"

func main() {
	var n int
	var topi, tulis, tali, pisau bool
	var has bool
	has = true

	fmt.Scanln(&n)
	for (n >= 1) && (has == true) {
		fmt.Scanln(&topi, &tulis, &tali, &pisau)
		has = topi && tulis && tali && pisau
		n = n - 1
	}
	fmt.Println(has)

}

// Muhammad Fatih Yumna LL 1301213389 IF-45-05
