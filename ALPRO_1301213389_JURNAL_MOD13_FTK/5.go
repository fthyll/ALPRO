package main

import "fmt"

func main() {
	var X, Y int
	fmt.Scan(&X, &Y)
	fmt.Println(power(X, Y))
}
func power(a, b int) int {
	if b == 0 {
		return 1
	} else {
		return a * power(a, b-1)
	}
}
