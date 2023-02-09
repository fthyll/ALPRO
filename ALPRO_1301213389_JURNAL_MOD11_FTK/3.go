package main

import "fmt"

const NMAX = 1000000

type partai struct {
	nama, suara int
}

type tabPartai [NMAX]partai

func main() {
	var p tabPartai
	var n, x, idx int

	n = 0
	fmt.Scan(&x)
	for x != -1 {
		idx = posisi(p, n, x)
		if idx == -1 {
			p[n].nama = x
			p[n].suara = 1
			n++
		} else {
			p[idx].suara++
		}
		fmt.Scan(&x)
	}
	var pass, k int
	var temp partai
	for pass = 1; pass <= n-1; pass++ {
		k = pass
		temp = p[k]
		for k > 0 && temp.suara > p[k-1].suara {
			p[k] = p[k-1]
			k--
		}
		p[k] = temp
	}
	for k = 0; k < n; k++ {
		fmt.Printf("%v(%v) ", p[k].nama, p[k].suara)
	}
}

func posisi(t tabPartai, n int, nama int) int {
	var i, ketemu int
	i = 0
	ketemu = -1
	for i < n && ketemu == -1 {
		if t[i].nama == nama {
			ketemu = i
		}
		i++
	}
	return ketemu
}
