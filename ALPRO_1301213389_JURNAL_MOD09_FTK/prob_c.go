package main
import "fmt"

type buku struct{
	judul,penulis string
	tahun int
}

type TabBuku [5]buku

func tambahBuku(kardus *TabBuku, atas *int){
	if *atas == 4 {
		fmt.Println("kardus penuh")
	}else{
		*atas++
		fmt.Scan(&kardus[*atas].judul,&kardus[*atas].penulis,&kardus[*atas].tahun)
	}
}

func ambilBuku(kardus*TabBuku,atas *int, ambil *buku){
	if *atas == -1 {
		fmt.Println("kardus kosong")
	}else{
		*ambil = kardus[*atas]
		*atas--
	}
}

func cariBuku(kardus *TabBuku, atas*int, X string){
	var found bool = false
	var ambil buku
	for !found && *atas != -1{
		ambilBuku(kardus,atas,&ambil)
		fmt.Println("Judul buku yang di keluarkan :",ambil,judul)
		found = ambil.judul == X
	}
	if found {
		fmt.Println("KETEMU")
	}else{
		fmt.Println("TIDAK KETEMU")
	}
}

func main()  {
	var kardus TabBuku
	var atas int
	atas = -1
	tambahBuku(&kardus,&atas)
	tambahBuku(&kardus,&atas)
	tambahBuku(&kardus,&atas)
	tambahBuku(&kardus,&atas)
	cariBuku(&kardus,&atas,"C")
	tambahBuku(&kardus,&atas)
	tambahBuku(&kardus,&atas)
	tambahBuku(&kardus,&atas)
	tambahBuku(&kardus,&atas)
	cariBuku(&kardus,&atas, "Sistem Digital")
}
// Muhammad Fatih Yumna IF-45-05 1301213389