package main

import (
	"fmt"
)

func main() {

	// Tipe data Integer
	fmt.Println("=== Contoh Integer ===")
	fmt.Println("Nilai intA (int):", 42)      // Integer default (int)
	fmt.Println("Nilai intB (int8):", 127)    // Integer dengan ukuran 8 bit
	fmt.Println("Nilai intC (int16):", 32767) // Integer dengan ukuran 16 bit
	fmt.Println("Satu =", 1)
	fmt.Println("Dua = ", 2)

	// Tipe data Floating Point
	fmt.Println("\n=== Contoh Floating Point ===")
	fmt.Println("Nilai floatA (float32):", 3.14)  // Floating point dengan ukuran 32 bit
	fmt.Println("Nilai floatB (float64):", 3.141) // Floating point dengan ukuran 64 bit
	fmt.Println("Tiga Koma Lima = ", 3.5)

	// Menampilkan hasil
	fmt.Println("\n=== Contoh Boolean ===")
	fmt.Println("Nilai isTrue:", true)
	fmt.Println("Nilai isFalse:", false)

	// Deklarasi tipe data string
	var firstName string = "Nuevalen"
	var lastName string = "Alswando"

	// String kosong
	var emptyString string = ""

	// Menggabungkan string dengan operator +
	fullName := firstName + " " + lastName

	// Menampilkan hasil
	fmt.Println("\n=== Contoh String ===")
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

	fmt.Println("\n=== konversi Tipe Data ===")
	// Deklarasi variable dengan tipe data int32
	var nilai32 int32 = 32768

	// Konversi dari int32 ke int64
	var nilai64 int64 = int64(nilai32)

	// Konversi dari int32 ke int16
	var nilai16 int16 = int16(nilai32)

	// Menampilkan hasil konversi
	fmt.Println("Nilai 32-bit:", nilai32)
	fmt.Println("Nilai 64-bit:", nilai64)
	fmt.Println("Nilai 16-bit:", nilai16)

	// Mengkonversi string ke byte dan sebaliknya
	var name = "Nuevalen Refitra"
	var e = name[0]         // Mengambil karakter pertama sebagai byte
	var eString = string(e) // Mengkonversi byte kembali ke string

	// Menampilkan hasil konversi
	fmt.Println("Nama:", name)
	fmt.Println("Karakter Pertama:", eString)

	/*		Keterangan:
	*	Deklarasi Variable: Di sini, kita mendeklarasikan nilai32 sebagai int32 dan kemudian mengkonversinya ke int64 dan int16.
	*	Konversi Byte ke String: Kita juga menunjukkan bagaimana cara mengambil karakter pertama dari string dan mengkonversinya dari byte kembali ke string.
	*	Menampilkan Hasil: Menggunakan fmt.Println untuk menampilkan nilai dari variable yang telah dikonversi.
	 */

	fmt.Println("\n=== Type Declarations ===")
	/*
		Type Declarations adalah kemampuan untuk membuat tipe data baru dari tipe data yang sudah ada. Hal ini biasanya digunakan untuk membuat alias terhadap tipe data yang sudah ada, sehingga membuat kode lebih mudah dimengerti dan lebih terstruktur.
	*/

	// Deklarasi tipe data baru bernama NoKTP yang merupakan alias dari string
	type NoKTP string

	// Menggunakan tipe data NoKTP
	var ktpValen NoKTP = "1111111111"

	// Menampilkan nilai dari variable ktpValen
	fmt.Println(ktpValen)

	// Mengkonversi string menjadi tipe NoKTP
	fmt.Println(NoKTP("22222222222"))

	/*		Keterangan:
	*	Type Declarations: Pada kode di atas, kita mendeklarasikan NoKTP sebagai tipe data baru yang merupakan alias dari string.
	*	Penggunaan Tipe Data Baru: Kita kemudian menggunakan NoKTP untuk membuat variable ktpValen dan memberikan nilai yang sesuai.
	*	Menampilkan Hasil: Menggunakan fmt.Println untuk menampilkan nilai dari variable yang telah dideklarasikan sebagai NoKTP.
	 */

	fmt.Println("\n=== Type Array ===")
	/* Tipe Data Array
	Array adalah tipe data yang menyimpan kumpulan data dengan tipe yang sama. Saat membuat array, kita perlu menentukan jumlah elemen yang bisa ditampung oleh array tersebut. Daya tampung array tidak bisa bertambah setelah array dibuat.

	Indeks di Array
	Setiap elemen dalam array diakses menggunakan indeks. Indeks dimulai dari 0. Contoh berikut menunjukkan beberapa data dalam array:

	Index	Data
	0		Valen
	1		Refitra
	2		Alswando
	*/

	var names [3]string    // Mendeklarasikan array dengan kapasitas 3
	names[0] = "Valen"       // Mengisi elemen pertama
	names[1] = "Refitra" // Mengisi elemen kedua
	names[2] = "Alswando"  // Mengisi elemen ketiga

	// Menampilkan elemen-elemen array
	fmt.Println(names[0]) // Output: Valen
	fmt.Println(names[1]) // Output: Refitra
	fmt.Println(names[2]) // Output: Alswando

	/*		Membuat Array Langsung
			Di Go, kita juga bisa membuat array secara langsung saat mendeklarasikan variabel. Berikut adalah contohnya:
	*/
	var values = [3]int{
		90,
		80,
		95,
	}

	// Menampilkan seluruh elemen array
	fmt.Println(values) // Output: [90 80 95]

	/*		Fungsi Array
			Beberapa operasi yang dapat dilakukan pada array adalah sebagai berikut:

			Operasi					Keterangan
			len(array)				Untuk mendapatkan panjang array
			array[index]			Mendapat data di posisi indeks
			array[index] = value	Mengubah data di posisi indeks
	*/

	var values2 = [...]int{
        90,
        80,
        95,
    }

    // Menampilkan seluruh elemen array
    fmt.Println(values2)           // Output: [90 80 95]
    fmt.Println(len(values2))      // Output: 3

    // Mengubah elemen pertama dari array
    values2[0] = 100
    fmt.Println(values2)           // Output: [100 80 95]
}
