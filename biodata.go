package main

import (
	"fmt"
	"os"
)

type Teman struct {
	Nama string
	Alamat string
	Pekerjaan string
	Alasan string
}

var temanTeman = map[int]Teman{
	1: {"Supri", "Jl. Merdeka No. 1", "Software Engineer", "Ingin belajar Golang"},
	2: {"Joko", "Jl. Proklamasi No. 2", "Data Analyst", "Tertarik dengan fitur concurrency Golang"},
	3: {"Anwar", "Jl. Sudirman No. 3", "Web Developer", "Memiliki pengalaman dengan bahasa pemrograman lain"},
	4: {"Budi", "Jl. Diponegoro No. 4", "UI/UX Designer", "Ingin meningkatkan skill programming"},
	5: {"Dewi", "Jl. Gatot Subroto No. 5", "System Administrator", "Suka dengan keamanan sistem"},
	6: {"Rini", "Jl. Pahlawan No. 6", "Database Administrator", "Memiliki pengalaman dalam pengelolaan database besar"},
	7: {"Ahmad", "Jl. Asia Afrika No. 7", "Network Engineer", "Minat dalam pengembangan jaringan komputer"},
	8: {"Linda", "Jl. Cendrawasih No. 8", "Quality Assurance", "Menyukai pekerjaan yang detail-oriented"},
	9: {"Rudi", "Jl. Cikini No. 9", "DevOps Engineer", "Berfokus pada otomatisasi dan deployment"},
	10: {"Siti", "Jl. Kebon Kacang No. 10", "Full Stack Developer", "Ingin memperdalam pengetahuan tentang pengembangan web"},
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Argumen belum diberikan")
		return
	}

	absen := os.Args[1]

	var absenInt int
	_, err := fmt.Sscanf(absen, "%d", &absenInt)
	if err != nil {
		fmt.Println("Argumen harus berupa angka")
		return
	}

	tampilkanData(absenInt)
}

func tampilkanData(absen int) {
	if absen > len(temanTeman) || absen < 1 {
		fmt.Println("Teman dengan absen", absen, "tidak ditemukan.")
		return
	}
	var dataTeman = temanTeman[absen]
	fmt.Println("Nama:", dataTeman.Nama)
	fmt.Println("Alamat:", dataTeman.Alamat)
	fmt.Println("Pekerjaan:", dataTeman.Pekerjaan)
	fmt.Println("Alasan memilih kelas Golang:", dataTeman.Alasan)
}

