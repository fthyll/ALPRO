package main

import "fmt"

// deklarasi variabel global
const mark = '.'
const space = '_' // satu spasi
var pita string
var indeks int
var CC byte
var EOP bool // true apabila CC == mark

var LKata int  // panjang kata terakhir
var LTotal int // total panjang kata
var NKata int  // jumlah kata

func main() {
	fmt.Scan(&pita)
	LTotal = 0
	NKata = 0
	// start kata
	startKata()
	for LKata != 0 {
		// hitung LTotal
		LTotal += LKata
		// hitung NKata
		NKata++
		// adv kata
		advKata()
	}
	if NKata != 0 {
		fmt.Printf("%.2f", float64(LTotal)/float64(NKata))
	} else {
		fmt.Println("Pita Kosong")
	}
}

func START() {
	/* I.S. pita telah terpasang pada mesinF.S. indeks bernilai 0, CC berisi karakter pertama pada pita, EOP bernilai true apabila CC adalah mark */
	indeks = 0
	CC = pita[indeks]
	EOP = CC == mark
}

func ADV() {
	/* I.S. EOP bernilai falseF.S. indeks bertambah 1, CC berisi karakter berikutnya dari CC saat ini, dan EOP bernilai true apabila CC adalah mark */
	indeks++
	CC = pita[indeks]
	EOP = CC == mark
}

func ignoreBlank() {
	/* I.S. CC bernilai apapunF.S. CC != spasi, tetapi karakter pertama kata atau EOP == truecatatan: Jumlah spasi antar kata mungkin bisa lebih dari satu, sehingga harus diabaikan spasi tersebut. */
	for CC == space {
		ADV()
	}
}

func hitungPanjang() {
	/* I.S. CC != spasi, yaitu karakter pertama kata atau markF.S. Lkata berisi panjang kata terakhir yang didapatkan, CC == spasi atau EOP == true */
	LKata = 0
	for CC != space && !EOP {
		LKata++
		ADV()
	}
}

func startKata() {
	/* I.S. pita telah terpasang pada mesinF.S. LKata == 0 karena pita kosong atau LKata berisi panjang kata pertama pada pita */
	START()
	ignoreBlank()
	hitungPanjang()
}

func advKata() {
	/* I.S. CC == spasi atau EOP == trueF.S. LKata == 0 atau Lkata != 0 yaitu panjang kata yang terakhir didapatkan */
	ignoreBlank()
	hitungPanjang()
}
