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
	print(A, N)
}
func print(T arrInt, n int) {
	// fungsi master (opsional)
	printArray(T, 0, n)
}
func printArray(T arrInt, i, n int) {
	// fungsi rekursif
	if i < n {
		fmt.Printf("%v ", T[i])
		printArray(T, i+1, n)
	}
}
