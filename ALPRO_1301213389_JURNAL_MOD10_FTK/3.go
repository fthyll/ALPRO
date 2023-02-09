package main

import "fmt"

const NMAX = 1000000

var data [NMAX]int

func main() {
	var n, k, hasil int
	fmt.Scan(&n, &k)
	isiArray(n)
	hasil = posisi(n, k)
	if hasil == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(hasil)
	}
}

func isiArray(n int) {
	var i int
	for i < n && i < NMAX {
		fmt.Scan(&data[i])
		i++
	}
}

func posisi(n, k int) int {
	var i int
	i = 0
	for i < n && data[i] != k {
		i++
	}
	if i < n {
		return i
	} else {

		return -1
	}
}

// Muhammad Fatih Yumna LL IF4505 1301213389
