package main

import "fmt"

func main() {

	/*		Penulisan
		1. fmt.Println():
		Mencetak teks dengan menambahkan baris baru di akhir.
		
		2. fmt.Printf():
		Digunakan untuk mencetak teks dengan format khusus, di sini %d untuk integer dan %s untuk string.
		
		3. fmt.Sprint():
		Mengembalikan string yang digabungkan tanpa mencetak langsung ke output.
		
		4. fmt.Scan():
		Membaca input dari pengguna dan menyimpannya dalam variabel yang sesuai.
	*/
	
	/*	1	*/
	// Hasil: "Hello, World!" diikuti dengan baris baru.
	fmt.Println("Hello, World!") // Output: Hello, World!

	// Menggunakan fmt.Printf untuk mencetak teks dengan format tertentu.
	// "%d" adalah placeholder untuk integer. "\n" menambahkan baris baru di akhir.
	fmt.Printf("My age is %d\n", 25) // Output: My age is 25

	// Menggunakan fmt.Sprint untuk menggabungkan string tanpa mencetak langsung.
	// Fungsi Sprint mengembalikan string yang diformat dan disimpan dalam variabel message.
	message := fmt.Sprint("Hello, ", "Go!")
	// Mencetak isi dari variabel message
	fmt.Println(message) // Output: Hello, Go!

	// Menggunakan fmt.Scan untuk membaca input dari pengguna.
	// Variabel name akan menyimpan input dari pengguna.
	var name string
	fmt.Print("Masukkan nama: ") // Memberi petunjuk kepada pengguna untuk memasukkan nama
	fmt.Scan(&name) // Menunggu input dari pengguna dan menyimpannya di variabel 'name'

	// Menggunakan fmt.Printf untuk menampilkan nama yang dimasukkan oleh pengguna.
	// "%s" adalah placeholder untuk string.
	fmt.Printf("Hello, %s\n", name)
}
