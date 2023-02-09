package main

import "fmt"

func main() {
	var member string

	var a, b, c, d, e int
	var cashback, diskon float64

	fmt.Scanln(&member, &a, &b, &c, &d, &e)

	cashback = 3.1 * float64(a+b+c)
	diskon = 1.7 * float64(c+d+e)

	if member == "yes" {
		cashback *= 1.15
		diskon *= 1.15
	}

	if cashback > 35 {
		cashback = 35
	}

	if diskon > 50 {
		diskon = 50
	}

	if a&2 == 0 && b%2 == 0 && c%2 == 0 && d%2 == 0 && e%2 == 0 {
		diskon = 0
	} else if a%2 != 0 && b%2 != 0 && c%2 != 0 && d%2 != 0 && e%2 != 0 {
		cashback = 0
	}

	fmt.Println(cashback, diskon)
}

// Muhammad Fatih Yumna Lajuwirdi Lirrahman - IF 45 05 - 1301213389
