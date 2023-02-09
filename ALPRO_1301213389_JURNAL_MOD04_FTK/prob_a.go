package main

import "fmt"

func main() {
	var a, r int
	var n int
	fmt.Scanln(&n, &a, &r)
	fmt.Print(0)
	for i := 1; i <= n; i++ {
		fmt.Printf(" + %d", a*i*r)

	}

}

// Muhammad Fatih Yumna LL 1301213389 IF-45-05
