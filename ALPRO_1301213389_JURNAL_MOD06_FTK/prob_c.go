package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))
}

func didalam(cx, cy, r, x, y float64) bool {
	var d float64
	d = jarak(cx, cy, x, y)

	return d <= r
}

func main() {
	var (
		x1, y1, r1     float64
		x2, y2, r2     float64
		x, y           float64
		state1, state2 bool
	)

	fmt.Scan(&x1, &y1, &r1)
	fmt.Scan(&x2, &y2, &r2)
	fmt.Scan(&x, &y)

	state1 = didalam(x1, y1, r1, x, y)
	state2 = didalam(x2, y2, r2, x, y)

	if state1 && state2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if state1 && !state2 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if !state1 && state2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}

//Muahammad Fatih Yumna - IF 45 05 - 1301213389
