package main

import "fmt"

func main() {
	fmt.Println("Hello, World!") // Output: Hello, World! (dengan baris baru)
	// Digunakan untuk mencetak teks atau nilai dengan menambahkan baris baru di akhir.

	fmt.Printf("My age is %d\n", 25) // Output: My age is 25
	// Digunakan untuk mencetak teks dengan format tertentu. Misalnya, menampilkan variabel dalam format yang spesifik.

	message := fmt.Sprint("Hello, ", "Go!")
	fmt.Println(message) // Output: Hello, Go!
	// Mengembalikan string yang diformat tanpa mencetaknya langsung ke layar.

	var name string
	fmt.Scan(&name)
	fmt.Printf("Hello, %s\n", name)
	//  Digunakan untuk membaca input dari pengguna.

}
