package main

import "fmt"

func main() {
	
	fmt.Println("\n=== Operasi Boolean ===")
	/*
		Operasi boolean menggunakan operator logika untuk menghasilkan nilai boolean berdasarkan operasi yang dilakukan.
		&&	Dan
		`	
		!	Kebalikan
	*/
	// Contoh Operasi AND (&&)
	var nilaiAkhir = 90
	var absensi = 80

	// Cek kelulusan berdasarkan nilai akhir dan absensi
	var lulusNilaiAkhir bool = nilaiAkhir > 80       // true
	var lulusAbsensi bool = absensi > 80             // false
	var lulus bool = lulusNilaiAkhir && lulusAbsensi // false

	fmt.Println("Apakah siswa lulus?", lulus) // Output: false

	// Contoh Operasi OR (||)
	nilaiAkhir = 70
	absensi = 90

	lulusNilaiAkhir = nilaiAkhir > 80       // false
	lulusAbsensi = absensi > 80             // true
	lulus = lulusNilaiAkhir || lulusAbsensi // true

	fmt.Println("Apakah siswa lulus?", lulus) // Output: true

	// Contoh Operasi NOT (!)
	lulus = false

	// Kebalikan dari nilai lulus
	var tidakLulus bool = !lulus

	fmt.Println("Apakah siswa tidak lulus?", tidakLulus) // Output: true

}
