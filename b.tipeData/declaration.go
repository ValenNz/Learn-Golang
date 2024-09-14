package main

import (
	"fmt"
)

func main() {
	fmt.Println("\n=== Type Data Declarations ===")
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
}
