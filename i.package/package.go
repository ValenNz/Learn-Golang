package main

import (
	"learn-golang/helper" // Mengimpor package helper
	"learn-golang/database" // Mengimpor package datavase
	_"learn-golang/database" // Mengimpor package tanpa menggunakannya

	"fmt"
)

func main() {
	/*		Package
		Package adalah cara untuk mengorganisir kode program di Go. Dengan menggunakan package, kita dapat merapikan dan memisahkan kode berdasarkan fungsionalitas.
	
	*/
	result := helper.SayHello("Eko") // Memanggil fungsi SayHello dari package helper
	fmt.Println(result) // Output: Hello Eko

	/*		Access Modifier
		Di Go, access modifier ditentukan oleh huruf awal nama variabel atau fungsi:

		Jika dimulai dengan huruf besar, dapat diakses dari package lain.
		Jika dimulai dengan huruf kecil, tidak dapat diakses dari package lain.
	*/

	// result := helper.sayGoodBye("Eko") // Memanggil fungsi sayGoodBye dari package helper
	// fmt.Println(result) // Output: Error karena menggunakan huruf kecil diawal

	/*		Package Initialization
		Fungsi init akan secara otomatis dijalankan ketika package diakses. Ini berguna untuk inisialisasi, seperti membuka koneksi database.
	
	*/
	fmt.Println(database.GetDatabase()) // Output: MySQL


	/*				Blank Identifier
		Jika kita ingin mengimpor package tetapi tidak menggunakannya, kita harus menggunakan blank identifier _. Ini mencegah Go memberikan peringatan tentang package yang tidak terpakai.
	*/

	// Di sini kita tidak menggunakan fungsi dari package database,
	// tetapi init function tetap akan dijalankan.
	fmt.Println("Program berjalan tanpa menggunakan database package")
	
	
}
