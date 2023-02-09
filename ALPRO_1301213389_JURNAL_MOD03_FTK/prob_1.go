package main

import "fmt"

func main() {
	var x1, x2, x3, x4, x5, x6, y1, y2, y3, y4, y5, y6, a, b, c, d, e, f int
	var jar1, jar2, jar3 float64

	fmt.Scanln(&x1, &y1, &x2, &y2, &a, &b, jar1) //inputan 1
	fmt.Scanln(&x3, &y3, &x4, &y4, &c, &d, jar2) //inputan 2
	fmt.Scanln(&x5, &y5, &x6, &y6, &e, &f, jar3) //inputan 3
	a = (x1 - x2)
	b = (y1 - y2)
	c = (x3 - x4)
	d = (y3 - y4)
	e = (x5 - x6)
	f = (y5 - y6)
	jar1 = float64(b) / float64(a)
	jar2 = float64(d) / float64(c)
	jar3 = float64(f) / float64(e)
	fmt.Println(jar1)
	fmt.Println(jar2)
	fmt.Println(jar3)

}

// Muhammad Fatih Yumna Lajuwirdi Lirrahman 1301213389 IF4505
