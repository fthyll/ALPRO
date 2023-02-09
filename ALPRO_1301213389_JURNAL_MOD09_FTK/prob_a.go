package main

import "fmt"

type text [2022]rume

func main() {
	var text text
	var s rune
	var i, len int
	var p bool

	len = 0
	fmt.Scanf("%c", &s)
	for s != '.' {
		teks[len] = s
		len++
		fmt.Scanf("%c", &s)
	}
	p = true
	for i = 0; i < len && p; i++ {
		p = teks[i] == teks[len-i-1]
	}
	fmt.Println(p)
}
// Muhammad Fatih Yumna IF-45-05 1301213389