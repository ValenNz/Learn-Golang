package main

import "fmt"

func main() {
	fmt.Println("\n=== Type Data Array ===")
	/* Tipe Data Array
	Array adalah tipe data yang menyimpan kumpulan data dengan tipe yang sama. Saat membuat array, kita perlu menentukan jumlah elemen yang bisa ditampung oleh array tersebut. Daya tampung array tidak bisa bertambah setelah array dibuat.

	Indeks di Array
	Setiap elemen dalam array diakses menggunakan indeks. Indeks dimulai dari 0. Contoh berikut menunjukkan beberapa data dalam array:

	Index	Data
	0		Valen
	1		Refitra
	2		Alswando
	*/

	var names [3]string   // Mendeklarasikan array dengan kapasitas 3
	names[0] = "Valen"    // Mengisi elemen pertama
	names[1] = "Refitra"  // Mengisi elemen kedua
	names[2] = "Alswando" // Mengisi elemen ketiga

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
	fmt.Println(values2)      // Output: [90 80 95]
	fmt.Println(len(values2)) // Output: 3

	// Mengubah elemen pertama dari array
	values2[0] = 100
	fmt.Println(values2) // Output: [100 80 95]
}
