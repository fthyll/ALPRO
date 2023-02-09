package main

import "fmt"

func main() {
	var n, last, now int
	var res bool
	fmt.Scan(&n)
	res = true
	last = n % 10
	n = n / 10
	for n > 0 && res == true {
		now = n % 10
		n = n / 10
		res = (last-now == 1) || (last-now == -1)
		last = now
	}
	fmt.Print(res)
}

// Muhammad Fatih Yumna LL 1301213389 IF-45-05
