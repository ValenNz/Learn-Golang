package main

import "fmt"

func main() {
	fmt.Println("\n=== Operasi Perbandingan ===")
	/*
		Operasi perbandingan adalah operasi yang digunakan untuk membandingkan dua data. Hasil dari operasi ini adalah nilai boolean (true atau false).
		
		>	Lebih Dari
		<	Kurang Dari
		>=	Lebih Dari Sama Dengan
		<=	Kurang Dari Sama Dengan
		==	Sama Dengan
		!=	Tidak Sama Dengan
	*/
	var name1 = "Eko"
	var name2 = "Eko"

	// Menggunakan operator perbandingan ==
	var result bool = name1 == name2

	fmt.Println("Apakah name1 sama dengan name2?", result) // Output: true
}