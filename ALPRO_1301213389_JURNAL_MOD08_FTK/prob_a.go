package main

import "fmt"

func tracking(a, b int) int {
	if a%2 != 0 && b%2 != 0 {
		return 1
	} else if a%2 == 0 && b%2 == 0 {
		return 0
	} else {
		return 2
	}
}

func main() {
	var a, b, hasil int
	fmt.Scanln(&a, &b)
	for tracking(a, b) != 0 {
		if tracking(a, b) == 1 {
			hasil++
		}
		fmt.Scanln(&a, &b)
	}
	fmt.Print(hasil)
}

// Muhammad Fatih Yumna LL 1301213389
