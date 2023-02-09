package main

import "fmt"

func main() {
	var tgl1, tgl2, thn1, thn2, bln2 int
	var hr1, hr2, bln1, month2 string

	fmt.Scan(&hr1, &tgl1, &bln1, &thn1)
	pengambilan(tgl1, angkaBulan(bln1), thn1, hr1, &tgl2, &bln2, &thn2, &hr2)
	month2 = bulanAngka(bln2)
	fmt.Print(hr2, " ", tgl2, " ", month2, " ", thn2)

}

func kabisat(tahun int) bool {
	return tahun%4 == 0 || tahun%400 == 0 && tahun%100 != 0
}

func angkaBulan(bulan string) int {
	if bulan == "januari" {
		return 1
	} else if bulan == "februari" {
		return 2
	} else if bulan == "maret" {
		return 3
	} else if bulan == "april" {
		return 4
	} else if bulan == "mei" {
		return 5
	} else if bulan == "juni" {
		return 6
	} else if bulan == "juli" {
		return 7
	} else if bulan == "agustus" {
		return 8
	} else if bulan == "september" {
		return 9
	} else if bulan == "oktober" {
		return 10
	} else if bulan == "november" {
		return 11
	} else if bulan == "desember" {
		return 12
	} else {
		return 0
	}
}

func bulanAngka(angka int) string {
	switch angka {
	case 1:
		return "januari"
	case 2:
		return "februari"
	case 3:
		return "maret"
	case 4:
		return "april"
	case 5:
		return "mei"
	case 6:
		return "juni"
	case 7:
		return "juli"
	case 8:
		return "agustus"
	case 9:
		return "september"
	case 10:
		return "oktober"
	case 11:
		return "november"
	case 12:
		return "desember"
	default:
		return "invalid"
	}
}

func jumlahHari(bulan int, tahun int) int {
	if bulan == 1 || bulan == 3 || bulan == 5 || bulan == 7 || bulan == 8 || bulan == 10 || bulan == 12 {
		return 31
	} else if bulan == 4 || bulan == 6 || bulan == 9 || bulan == 11 {
		return 30
	} else if bulan == 2 {
		if kabisat(tahun) {
			return 29
		} else {
			return 28
		}
	} else {
		return -1
	}
}

func mencariDurasi(hari1 string, hari2 *string, durasi *int) {
	switch hari1 {
	case "senin":
		*hari2 = "rabu"
		*durasi = 2
	case "selasa":
		*hari2 = "kamis"
		*durasi = 2
	case "rabu":
		*hari2 = "jumat"
		*durasi = 2
	case "kamis":
		*hari2 = "senin"
		*durasi = 4
	case "jumat":
		*hari2 = "selasa"
		*durasi = 4
	}
}

func pengambilan(tanggal1, bulan1, tahun1 int, hari1 string, tanggal2, bulan2, tahun2 *int, hari2 *string) {
	var durasi int
	var totalHari int

	mencariDurasi(hari1, hari2, &durasi)
	*tanggal2 = tanggal1 + durasi
	totalHari = jumlahHari(bulan1, tahun1)

	if *tanggal2 <= totalHari {
		*bulan2 = bulan1
		*tahun2 = tahun1
	} else if *tanggal2 > totalHari {
		*tanggal2 = *tanggal2 - totalHari
		if bulan1 == 12 {
			*bulan2 = 1
			*tahun2 = tahun1 + 1
		} else {
			*bulan2 = *bulan2 + 1
			*tahun2 = tahun1
		}
	}

}

// Muhammad Fatih Yumna LL IF45-05 1301213389
