package main

import (
	"fmt"
)

// Fungsi NewMap mengembalikan map[string]string atau nil
func NewMap(name string) map[string]string {

	/*
		Penjelasan Nil:
		Nil di Go-Lang adalah representasi dari nilai kosong atau tidak terdefinisi, serupa dengan null pada bahasa pemrograman lain.
		Nil tidak dapat digunakan di semua tipe data, hanya dapat digunakan pada tipe data tertentu seperti interface, function, map, slice, pointer, dan channel.
		Saat variabel dengan tipe data tertentu dibuat, Go-Lang memberikan nilai default, bukan nil, kecuali pada tipe data yang mendukung nil.
	*/

	// Jika string 'name' kosong, maka mengembalikan nilai nil
	if name == "" {
		return nil // Mengembalikan nilai nil jika nama kosong
	} else {
		// Jika name tidak kosong, mengembalikan map dengan data "name"
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	// Memanggil fungsi NewMap dengan parameter string kosong
	data := NewMap("") // name kosong, seharusnya mengembalikan nil

	// Mengecek apakah data adalah nil
	if data == nil {
		fmt.Println("Data Kosong") // Output: Data Kosong
	} else {
		fmt.Println(data) // Jika tidak nil, cetak isi map
	}
}

