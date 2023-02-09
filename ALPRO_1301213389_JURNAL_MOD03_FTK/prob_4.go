package main

import "fmt"

func main() {
	var th1, th2, th3, th4 int
	var kab1, kab2, kab3, kab4 bool

	fmt.Scanln(&th1)
	fmt.Scanln(&th2)
	fmt.Scanln(&th3)
	fmt.Scanln(&th4)
	kab1 = th1%400 == 0 || th1%4 == 0 && th1&100 != 0
	kab2 = th2%400 == 0 || th2%4 == 0 && th2&100 != 0
	kab3 = th3%400 == 0 || th3%4 == 0 && th3&100 != 0
	kab4 = th4%400 == 0 || th4%4 == 0 && th4&100 != 0

	fmt.Println(kab1 && kab2 && kab3 && kab4)
}

// Muhammad Fatih Yumna Lajuwirdi Lirrahman 1301213389 IF4505
