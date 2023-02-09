package main

import "fmt"

const (
	kg  = 1000
	pon = 453.592
	ons = 28.3495
)

func main() {
	var gr1, gr2, gr3, gr4, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12 float64

	fmt.Scanln(&gr1)
	fmt.Scanln(&gr2)
	fmt.Scanln(&gr3)
	fmt.Scanln(&gr4)

	a1 = gr1 / kg
	a2 = gr1 / pon
	a3 = gr1 / ons

	a4 = gr2 / kg
	a5 = gr2 / pon
	a6 = gr2 / ons

	a7 = gr3 / kg
	a8 = gr3 / pon
	a9 = gr3 / ons

	a10 = gr4 / kg
	a11 = gr4 / pon
	a12 = gr4 / ons

	fmt.Printf(" %g", a1)
	fmt.Printf(" %g", a2)
	fmt.Printf(" %g  \n", a3)
	fmt.Printf(" %g", a4)
	fmt.Printf(" %g", a5)
	fmt.Printf(" %g  \n", a6)
	fmt.Printf(" %g", a7)
	fmt.Printf(" %g", a8)
	fmt.Printf(" %g  \n", a9)
	fmt.Printf(" %g", a10)
	fmt.Printf(" %g", a11)
	fmt.Printf(" %g", a12)

}

// Muhammad Fatih Yumna Lajuwirdi Lirrahman 1301213389 IF4505
