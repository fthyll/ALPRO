package main

import "fmt"

func main() {

	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	if d > a && d > b && d > c {
		fmt.Println("Ada Record Baru")
	} else {
		fmt.Println("Tidak Ada Record Baru")
	}

}

// Muhammad Fatih Yumna LL 1301213389 IF-45-05
