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
	fmt.Println(palindrom(A, N))
}

func palindrom(T arrInt, n int) bool {
	return check(T, 0, n-1, n)
}

func check(T arrInt, i, j, n int) bool {
	if i == n/2 {
		return (T[i] == T[j])
	} else {
		return (T[i] == T[j]) && check(T, i+1, j-1, n)
	}
}
