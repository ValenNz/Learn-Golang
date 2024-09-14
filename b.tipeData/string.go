package main

import (
	"fmt"
)

func main() {

	// Deklarasi tipe data string
	var firstName string = "Nuevalen"
	var lastName string = "Alswando"

	// String kosong
	var emptyString string = ""

	// Menggabungkan string dengan operator +
	fullName := firstName + " " + lastName

	// Menampilkan hasil
	fmt.Println("\n=== Type Data	 String ===")
	fmt.Println("Nama Depan:", firstName)
	fmt.Println("Nama Belakang:", lastName)
	fmt.Println("Nama Lengkap:", fullName)
	fmt.Println("String Kosong:", emptyString)

	// Menghitung jumlah karakter
	str := "Hello, Go-Lang!"
	panjang := len(str) // Fungsi len(str) digunakan untuk menghitung jumlah karakter dalam string str.
	fmt.Println("Jumlah karakter:", panjang)

	// Mengambil karakter pada posisi ke-7
	char := str[7]
	fmt.Printf("Karakter di posisi ke-7: %c\n", char) // String di Go adalah array karakter. Untuk mengambil karakter pada posisi tertentu, kita bisa menggunakan indeks seperti str[number]. Misalnya, str[7] akan mengambil karakter ke-7 dari string.
}
