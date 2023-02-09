package main

import "fmt"

type set [2022]int

func exist(T set, n int, val int) bool {
	/*  mengembalikan  true  apabila  bilangan  val  ada  di dalam  array  T  yang  berisi sejumlah n bilangan bulat */
	status := true
	for i := 0; i < n-1; i++ {
		if T[i] == val {
			status = false
			break
		}
	}
	return status
}

func inputSet(T *set, n *int) {
	/*  I.S. data himpunan telah siap pada piranti masukan
		F.S. array T berisi sejumlah n bilangan bulat yang berasal dari masukan (masukan
	berakhir apabila bilangan ada yang duplikat, atau array penuh)
		Catatan: Panggil function exist di sini untuk membantu pengecekan */
	var b int
	status := true
	for status {
		fmt.Scan(&b)
		T[*n] = b
		*n++
		status = exist(*T, *n, b)
	}
}
func findIntersection(T1, T2 set, n, m int, T3 *set, h *int) {
	/*  I.S.  terdefinisi  himpunan  T1  dan  T2  yang  berisi  sejumlah  n  dan  m  anggota
	himpunan
		F.S. himpunan T3 berisi sejumlah h bilangan bulat yang merupakan irisan dari
	himpunan T1 dan T2
		Catatan: Panggil function exist di sini untuk membantu pengecekan */
	var x, z int
	for x = 0; x < n; x++ {
		for z = 0; z < m; z++ {
			if T2[z] == T1[x] && exist(*T3, n, T2[z]) {
				T3[*h] = T2[z]
				*h++
			}
		}

	}
}
func printSet(T set, n int) {
	/*  I.S. terdefinisi sebuah himpunan T yang berisi sejumlah n bilangan bulat
	F.S. menampilkan isi array T secara horizontal (dipisahkan oleh spasi) */
	for i := 0; i < n; i++ {
		fmt.Print(T[i], " ")
	}
}
func main() {
	var s1, s2, s3 set
	var n1, n2, n3 int
	fmt.Println("Masukkan bilangan:")
	inputSet(&s1, &n1)
	inputSet(&s2, &n2)
	findIntersection(s1, s2, n1, n2, &s3, &n3)
	fmt.Println("Maka Bilangan yang sama: ")
	printSet(s3, n3)
}

// Muhammad Fatih Yumna LL IF4505 1301213389
