package main

import (
	"fmt"
)

func main() {

	fmt.Println("=== Variabel ===")
	/*		Variable
		* Variable adalah tempat untuk menyimpan data
		* Variable digunakan agar kita bisa mengakses data yang sama dimanapun kita mau
		* Di Go-Lang Variable hanya bisa menyimpan tipe data yang sama, jika kita ingin menyimpan data yang berbeda-beda jenis, kita harus membuat beberapa variable
		* Untuk membuat variable, kita bisa menggunakan kata kunci var, lalu diikuti dengan nama variable dan tipe datanya
	*/

	/* Inisialisasi Manual */
	var name string = "Valen"
	var age int = 21
	var isStudent bool = true

	fmt.Println("Nama:", name)
	fmt.Println("Usia:", age)
	fmt.Println("Apakah mahasiswa:", isStudent)

	/*
		* var digunakan untuk mendeklarasikan variable.
		* Variable di Go-Lang harus memiliki tipe data yang jelas, seperti string, int, atau bool.
		* Setiap variable hanya bisa menyimpan data dengan tipe yang sama. Jika ingin menyimpan data yang berbeda tipe seperti string, int, atau bool, kita perlu membuat beberapa variable.
	 */

	/* Inisialisasi Otomatis */
	nama := "Valen"
	usia := 21
	pelajar := true

	fmt.Println("Nama:", nama)
	fmt.Println("Usia:", usia)
	fmt.Println("Apakah mahasiswa:", pelajar)

	/*
		Deklarasi dengan :=: Go-Lang akan otomatis menentukan tipe data berdasarkan nilai yang di-assign. Misalnya, name := "Valen" akan otomatis dianggap sebagai tipe string.
	*/

	/* Deklarasi Multiple Variable */
	// Deklarasi beberapa variable dengan inisialisasi otomatis
	name, city := "Valen", "Surabaya"
	age, year := 21, 2024
	isStudent, isActive := true, true

	fmt.Println("Nama:", name)
	fmt.Println("Kota:", city)
	fmt.Println("Usia:", age)
	fmt.Println("Tahun:", year)
	fmt.Println("Apakah mahasiswa:", isStudent)
	fmt.Println("Apakah aktif:", isActive)
	/*
		Inisialisasi Otomatis dengan :=: Kita juga bisa menggunakan operator := untuk mendeklarasikan dan menginisialisasi beberapa variable secara bersamaan tanpa menyebutkan tipe data, karena Go-Lang akan menentukan tipe secara otomatis berdasarkan nilai yang diberikan.
	*/

	fmt.Println("\n=== Constant ===")
	/*		Constant
		Constant adalah nilai yang tidak dapat diubah setelah pertama kali diinisialisasi. Nilai constant bersifat tetap dan tidak bisa dimodifikasi selama program berjalan. Untuk mendeklarasikan constant, kita menggunakan kata kunci const.
	*/

	const firstName string = "Nuevalen" // Deklarasi constant dengan tipe string
	const lastName = "Refitra"          // Deklarasi constant tanpa menyebutkan tipe data

	// Menampilkan constant
	fmt.Println("Nama Depan:", firstName)
	fmt.Println("Nama Belakang:", lastName)

	// Mengubah nilai constant (akan menyebabkan error)
	// firstName = "Tidak Bisa Diubah" // Ini akan menghasilkan error
	// lastName = "Tidak Bisa Diubah"   // Ini juga akan menghasilkan error

	/*		Keterangan:
		* Deklarasi Constant: Untuk mendeklarasikan constant, kita menggunakan const diikuti dengan nama constant, tipe data (opsional jika nilai sudah jelas), dan nilai yang akan disimpan.
		* Nilai Tidak Bisa Diubah: Jika kita mencoba mengubah nilai dari constant setelah deklarasi, Go-Lang akan menghasilkan error, seperti yang terlihat pada komentar dalam kode.
	 */

	/* Deklarasi Multiple Constant */
	const (
		namaPertama string = "Nuevalen" // Deklarasi constant dengan tipe string
		namaTerakhir         = "Refitra"  // Deklarasi constant tanpa menyebutkan tipe data
	)

	// Menampilkan constant
	fmt.Println("Nama Depan:", namaPertama)
	fmt.Println("Nama Belakang:", namaTerakhir)

	// Mengubah nilai constant (akan menyebabkan error)
	// firstName = "Budi" // Ini akan menghasilkan error
	// lastName = "Joko"   // Ini juga akan menghasilkan error

	/*
			Keterangan:
		* Deklarasi Multiple Constant: Dengan menggunakan const diikuti oleh tanda kurung ( dan ), kita dapat mendeklarasikan beberapa constant secara bersamaan.
		* Nilai Tidak Bisa Diubah: Jika kita mencoba untuk mengubah nilai dari constant setelah deklarasi, Go-Lang akan menghasilkan error, seperti yang terlihat pada komentar dalam kode.
	*/
}
