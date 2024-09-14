package main

import "fmt"

func main() {
	fmt.Println("\n=== Type Data Map ===")
	/*
	 * Pada Array atau Slice, untuk mengakses data, kita menggunakan index Number dimulai dari 0
	 * Map adalah tipe data lain yang berisikan kumpulan data yang sama, namun kita bisa menentukan jenis tipe data index yang akan kita gunakan
	 * Sederhananya, Map adalah tipe data kumpulan key-value (kata kunci - nilai), dimana kata kuncinya bersifat unik, tidak boleh sama
	 * Berbeda dengan Array dan Slice, jumlah data yang kita masukkan ke dalam Map boleh sebanyak-banyaknya, asalkan kata kuncinya berbeda,
	 * jika kita gunakan kata kunci sama, maka secara otomatis data sebelumnya akan diganti dengan data baru
	 */

	// Membuat map
	person := map[string]string{
		"name":    "Eko",
		"address": "Subang",
	}
	fmt.Println("Map person:", person)

	/*
		map[key]:
		*	Mengambil data dari map berdasarkan key.
		*	Contoh: person["name"] mengembalikan "Eko".
	*/
	fmt.Println("Name:", person["name"])       // Mengambil data dengan key "name"
	fmt.Println("Address:", person["address"]) // Mengambil data dengan key "address"

	/*
		map[key] = value:
		*	Mengubah atau menambahkan data di map dengan key yang diberikan.
		*	Contoh: person["address"] = "Bandung" mengganti nilai address dari "Subang" menjadi "Bandung".
	*/
	person["address"] = "Bandung"
	fmt.Println("\nSetelah update address:", person)

	/* 	len(map)
	*	Mengembalikan jumlah elemen yang ada dalam map (jumlah pasangan key-value).
	*	Contoh: len(person) mengembalikan 2, karena map person memiliki 2 entri.
	 */
	fmt.Println("Jumlah data di map person:", len(person)) // Mengembalikan jumlah data

	/*
		make(map[TypeKey]TypeValue):
		*	Membuat map baru dengan tipe key dan value yang ditentukan.
		*	Contoh: book := make(map[string]string) membuat map book dengan key dan value bertipe string.
	*/
	fmt.Println("\n=== Function make(map) ===")
	book := make(map[string]string) // Membuat map baru
	book["title"] = "Buku Go-Lang"
	book["author"] = "Eko Kurniawan"
	book["wrong"] = "Ups" // Menambahkan data
	fmt.Println("Map book:", book)

	/*
		delete(map, key):
		*	Menghapus pasangan key-value dari map berdasarkan key yang diberikan.
		*	Contoh: delete(book, "wrong") menghapus entri dengan key "wrong" dari map book.
	*/
	fmt.Println("\n=== Function delete(map, key) ===")
	delete(book, "wrong") // Menghapus data dengan key "wrong"
	fmt.Println("Map book setelah delete:", book)

}
