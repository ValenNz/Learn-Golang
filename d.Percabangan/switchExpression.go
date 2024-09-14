package main

import "fmt"

func main() {
	/*
		Switch Expression:
		*	Digunakan untuk melakukan pengecekan pada variabel name.
		*	Pada contoh ini, jika name bernilai "Eko", maka program mencetak "Hello Eko". Jika bernilai "Joko", maka mencetak "Hello Joko". Jika bukan keduanya, program menjalankan blok default.
	*/
	fmt.Println("\n=== Switch Expression ===")
	name := "Eko"
	switch name {
	case "Eko":
		fmt.Println("Hello Eko")
	case "Joko":
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Hi, Boleh Kenalan?")
	}

	/*
		Switch dengan Short Statement:

		*	Seperti halnya If dengan short statement, kita bisa melakukan inisialisasi nilai dalam pernyataan Switch.
		*	Pada contoh ini, panjang string name dihitung, dan berdasarkan hasilnya, akan menentukan apakah panjang nama tersebut lebih dari 5.
	*/
	fmt.Println("\n=== Switch dengan Short Statement ===")
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama Sudah Benar")
	}

	/*
		Switch Tanpa Kondisi:

		*	Pada Switch ini, tidak ada variabel yang diperiksa secara langsung di Switch expression, namun kondisi ditambahkan pada setiap case.
		*	Contoh ini mengevaluasi panjang dari string name, dan berdasarkan kondisinya, program mencetak pesan yang sesuai.
	*/
	fmt.Println("\n=== Switch Tanpa Kondisi ===")
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama Terlalu Panjang")
	case length > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama Sudah Benar")
	}
}
