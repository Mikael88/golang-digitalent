package main

import (
	"fmt"
	"os"
)

// contoh 
// go run biodata.go 1

type Teman struct {
	Absen int
	Nama string
	Alamat string
	Pekerjaan string
	Alasan string
}

var temanTeman = []Teman{
	{1, "Supri", "Jl. Merdeka No. 1", "Software Engineer", "Ingin belajar Golang"},
	{2, "Joko", "Jl. Proklamasi No. 2", "Data Analyst", "Tertarik dengan fitur concurrency Golang"},
	{3, "Anwar", "Jl. Sudirman No. 3", "Web Developer", "Memiliki pengalaman dengan bahasa pemrograman lain"},
}

func tampilkanData(absen int) {
	for _, teman := range temanTeman {
		if teman.Absen == absen {
			fmt.Println("Nama:", teman.Nama)
			fmt.Println("Alamat:", teman.Alamat)
			fmt.Println("Pekerjaan:", teman.Pekerjaan)
			fmt.Println("Alasan memilih kelas Golang:", teman.Alasan)
			return
		}
	}
	fmt.Println("Teman dengan absen", absen, "tidak ditemukan.")
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