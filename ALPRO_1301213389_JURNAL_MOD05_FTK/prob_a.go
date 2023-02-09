package main

import "fmt"

func main() {
	var b int
	var countFactor int

	fmt.Scanln(&b)

	countFactor = 0

	fmt.Print("Faktor : ")
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			countFactor++
			fmt.Print(i, " ")
		}
	}
	fmt.Println()

	fmt.Println("Prima :", countFactor == 2)
}

// Muhammad Fatih Yumna Lajuwirdi Lirrahman - IF 45 05 - 1301213389
