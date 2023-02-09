package main

import "fmt"

const NMAX = 1000000

type arrInt [NMAX]int

func sorting(T *arrInt, n int) {
	var pass, idx_min, i, temp int
	for pass = 1; pass <= n-1; pass++ {
		idx_min = pass - 1
		for i = pass; i <= n-1; i++ {
			if T[idx_min] > T[i] {
				idx_min = i
			}
		}
		temp = T[idx_min]
		T[idx_min] = T[pass-1]
		T[pass-1] = temp
	}
}
func median(T arrInt, n int) float64 {
	var mid int = n / 2
	if n%2 == 0 {
		return float64(T[mid-1]+T[mid]) / 2.0
	} else {
		return float64(T[mid])
	}
}
func main() {
	var A arrInt
	var x, n int
	fmt.Scan(&x)
	n = 0
	for x != -5313541 && n < NMAX {
		if x == 0 {
			sorting(&A, n)
			fmt.Println(median(A, n))
		} else {
			A[n] = x
			n++
		}
		fmt.Scan(&x)
	}
}

// Muhammad Fatih Yumna LL - 1301213389
