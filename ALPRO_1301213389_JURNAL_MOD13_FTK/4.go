package main

import "fmt"

const NMAX = 1000

type arrInt [NMAX]int

func main() {
	var A arrInt
	var i, N int

	fmt.Scan(&N)
	for i = 0; i < N; i++ {
		fmt.Scan(&A[i])
	}
	fmt.Println(sum(A, N))

}
func sum(T arrInt, n int) int {
	return sumArray(T, 0, n)
}
func sumArray(T arrInt, i, n int) int {
	if i == n {
		return 0
	} else {
		return T[i] + sumArray(T, i+1, n)
	}
}
