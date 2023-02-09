package main

import "fmt"

func findFactorial(n int, result *int) {
	var j int
	*result = 1
	for j = 1; j <= n; j++ {
		*result *= j
	}
}

func permutasi(n, r int) int {
	var nF, nrF int
	nF = 1
	nrF = 1
	findFactorial(n, &nF)
	findFactorial(n-r, &nrF)
	return nF / nrF
}
func kombinasi(n, r int) int {
	var nF, nrF, rF int
	nF = 1
	nrF = 1
	rF = 1
	findFactorial(n, &nF)
	findFactorial(n-r, &nrF)
	findFactorial(n, &nF)
	findFactorial(r, &rF)
	return nF / (rF * nrF)
}
func main() {
	var a, b, c, d int
	var p1, c1, p2, c2 int
	fmt.Scan(&a, &b, &c, &d)
	p1 = permutasi(a, c)
	c1 = kombinasi(a, c)
	p2 = permutasi(b, d)
	c2 = kombinasi(b, d)
	fmt.Println(p1, c1)
	fmt.Println(p2, c2)

}

//Muahammad Fatih Yumna - IF 45 05 - 1301213389
