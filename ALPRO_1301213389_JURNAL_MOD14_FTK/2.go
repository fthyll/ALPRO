package main

import "fmt"

// deklarasi variabel global
var pita string
var CC byte
var EOP bool
var indeks int

func START() {
	/* I.S. pita telah berisi oleh sekumpulan karakter yang diakhiri oleh hashtag '.'.F.S.  CC  berisi  karakter  pertama  dari  pita,  EOP  bernilai  true  apabila  CC adalah '.', false untuk sebaliknya*/
	indeks = 0
	CC = pita[indeks]
	EOP = CC == '.'
}

func ADV() {
	/* I.S. EOP bernilai falseF.S. CC berisi karakter berikutnya dari CC pada pita karakter saat ini, EOP bernilai true apabila CC adalah '.', false untuk sebaliknya*/
	indeks++
	CC = pita[indeks]
	EOP = CC == '.'
}

func main() {
	var sum int
	var status = 1
	fmt.Scan(&pita)
	START()
	for !EOP {
		if status == 1 {
			if CC == 'F' {
				status = 2
			}
		} else if status == 2 {
			if CC == 'I' {
				status = 3
			} else {
				status = 1
			}
		} else if status == 3 {
			if CC == 'F' {
				sum++
			}
			status = 1
		}
		ADV()
	}
	fmt.Println(sum)
}
