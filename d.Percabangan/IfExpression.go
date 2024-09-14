package main

import "fmt"

func main() {
    /*
		If Expression:
		*	Digunakan untuk menjalankan suatu blok kode ketika kondisi bernilai true.
		*	Pada contoh ini, jika variabel name bernilai "Eko", maka akan mencetak "Hello Eko".
	*/
    fmt.Println("\n=== If Expression ===")
    name := "Eko"
    if name == "Eko" { // Menggunakan operator perbandingan ==
        fmt.Println("Hello Eko")
    }

   	/*
		Else Expression:
		*	Digunakan untuk menangani situasi di mana kondisi dalam if tidak terpenuhi (bernilai false).
		*	Jika name bukan "Eko", maka program akan menjalankan blok kode di dalam else.
	*/
    fmt.Println("\n=== Else Expression ===")
    name = "Budi"
    if name == "Eko" {
        fmt.Println("Hello Eko")
    } else {
        fmt.Println("Hi, Boleh Kenalan?")
    }

    /*
		Else If Expression:
		*	Digunakan jika terdapat lebih dari satu kondisi yang perlu dicek.
		*	Pada contoh ini, jika name tidak bernilai "Eko", tetapi bernilai "Joko", maka program mencetak "Hello Joko".
	*/
    fmt.Println("\n=== Else If Expression ===")
    name = "Joko"
    if name == "Eko" {
        fmt.Println("Hello Eko")
    } else if name == "Joko" {
        fmt.Println("Hello Joko")
    } else {
        fmt.Println("Hi, Boleh Kenalan?")
    }

    /*
		If dengan Short Statement:
		*	Ini adalah fitur yang memungkinkan kita membuat pernyataan pendek sebelum mengevaluasi kondisi.
		*	Pada contoh ini, program pertama menghitung panjang string name dengan len(name) dan kemudian mengevaluasi apakah panjangnya lebih dari 5.
	*/
    fmt.Println("\n=== If dengan Short Statement ===")
    name = "Khannedy"
    if length := len(name); length > 5 {
        fmt.Println("Nama terlalu panjang")
    } else {
        fmt.Println("Nama sudah benar")
    }
}
