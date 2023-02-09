package main

import "fmt"

type tabGol [2022]int
func inputData(t *tabGol, n *int){
	var menang int
	fmt.Scan(&menang)
	*n = 0
	for menang >= 0 {
		t[*n] = menang
		*n++
		fmt.Scan(&menang)
	}
}

func rataan(t tabGol, n int) float64 {
	var i, jum int
	jum = 0
	for i = 0; i < n; i++ {
		jum += t[i]
	}
	return float64(jum)/float64(n)
}

func main(){
	var t1,t2.t3 tabGol
	var n1, n2, n3 int
	inputData(&t1,&n1)
	inputData(&t2,&n2)
	inputData(&t3,&n3)
	fmt.Println(rataan(t1,n1))
	fmt.Println(rataan(t2,n2))
	fmt.Println(rataan(t3,n3))
}
// Muhammad Fatih Yumna IF-45-05 1301213389