package main

import "fmt"

func main() {
    fmt.Println("=== Contoh For Loop dengan Kondisi ===")
	/*For Loop dengan Kondisi:
		*	Pada contoh pertama, for counter <= 10, perulangan akan berjalan selama nilai counter kurang dari atau sama dengan 10.
		*	Di dalam loop, nilai counter dicetak dan kemudian ditambah 1 pada setiap iterasi.
	*/
    // Inisialisasi variabel counter dengan nilai awal 1
    counter := 1

    // Perulangan akan terus berjalan selama counter <= 10
    for counter <= 10 {
        // Menampilkan nilai counter saat ini
        fmt.Println("Perulangan ke", counter)

        // Increment counter (menambahkan 1 setiap iterasi)
        counter++
    }

    fmt.Println("\n=== Contoh For Loop dengan Init dan Post Statement ===")
	/*		For dengan Init dan Post Statement:

		*	Di contoh ini, kita bisa mendeklarasikan variabel (init statement), menambahkan kondisi (counter <= 10), dan melakukan operasi setelah setiap iterasi (post statement).
		*	Ini adalah cara yang lebih umum digunakan untuk menulis loop yang memiliki inisialisasi, kondisi, dan peningkatan variabel dalam satu baris.
	*/
    for counter := 1; counter <= 10; counter++ {
        fmt.Println("Perulangan ke", counter)
    }
	// For loop dengan init statement (counter := 1), kondisi (counter <= 10),
    // dan post statement (counter++)

    fmt.Println("\n=== Contoh For Loop dengan Range ===")
	/*		For Loop dengan Range:
		*	range digunakan untuk mengiterasi elemen-elemen dari sebuah slice, array, atau map.
		*	Pada contoh ini, kita mengiterasi slice names yang berisi beberapa nama. Setiap iterasi memberikan index dan nilai dari elemen.
	*/
    // Membuat slice of string yang berisi nama-nama
    names := []string{"Eko", "Kurniawan", "Khannedy"}
    
    for index, name := range names {
        fmt.Println("Index", index, "=", name)
    }
	// Menggunakan for dengan range untuk iterasi slice
    // index adalah posisi elemen dalam slice, dan name adalah nilai elemen tersebut


    fmt.Println("\n=== Contoh For Loop dengan Range tanpa Index ===")
	/*		For Loop dengan Range tanpa Index:
		*	Terkadang, kita hanya membutuhkan nilai elemen, tanpa peduli dengan index. Kita bisa mengabaikan index dengan menggunakan _.
	*/
    for _, name := range names {
        fmt.Println("Nama:", name)
    }
	// Jika hanya perlu nilai tanpa index, kita bisa mengabaikan index dengan menggunakan _

}
