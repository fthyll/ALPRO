package main

import "fmt"

func desimalToBiner(bil int) {
	var biner []int

	for bil != 0 {
		biner = append(biner, bil%2)
		bil = bil / 2
	}
	if len(biner) == 0 {
		fmt.Printf("%d\n", 0)
	} else {
		for i := len(biner) - 1; i >= 0; i-- {
			fmt.Printf("%d", biner[i])
		}
		fmt.Println()
	}
}

func main() {
	var x int
	fmt.Scan(&x)
	desimalToBiner(x)
}
