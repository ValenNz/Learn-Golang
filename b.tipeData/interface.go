package main

import "fmt"

/*
	Definisi Interface:
	Interface HasName hanya memerlukan satu method, yaitu GetName(),
	yang mengembalikan nilai berupa string. Setiap tipe data (seperti Person
	atau Animal) yang mengimplementasikan method ini secara otomatis
	memenuhi kontrak interface HasName.
*/

// Interface HasName
type HasName interface {
	GetName() string
}

/*
	Fungsi SayHello:
	Fungsi ini menerima parameter apa pun yang mengimplementasikan interface HasName.
	Tidak peduli apakah itu Person atau Animal, selama tipe tersebut memiliki
	method GetName, fungsi ini bisa dijalankan.
*/

// Fungsi SayHello
func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

/*
	Struct dan Implementasi Method:
	Baik Person maupun Animal memiliki field Name dan mengimplementasikan method
	GetName(). Karena keduanya memenuhi kontrak interface HasName, keduanya
	bisa digunakan dalam fungsi SayHello.
*/

// Struct Person
type Person struct {
	Name string
}

// Method GetName untuk Person
func (person Person) GetName() string {
	return person.Name
}

// Struct Animal
type Animal struct {
	Name string
}

// Method GetName untuk Animal
func (animal Animal) GetName() string {
	return animal.Name
}

/*
	Interface Kosong:
	Interface kosong di Go-Lang adalah interface yang tidak memiliki deklarasi method.
	Karena tidak ada deklarasi method, semua tipe data akan secara otomatis
	mengimplementasikan interface ini.
*/

// Fungsi Ups mengembalikan tipe data interface kosong (interface{})
func Ups() interface{} {
	// Mengembalikan nilai dengan tipe data apapun
	return "Ups" // bisa string
}

/*
	Penjelasan Type Assertions:
	Type Assertions adalah kemampuan dalam Go-Lang untuk mengambil nilai dari
	sebuah interface kosong (interface{}) dan mengubahnya menjadi tipe data
	yang diinginkan. Type assertions sering digunakan ketika kita bekerja
	dengan interface kosong, yang bisa berisi tipe data apa saja. Jika type
	assertions digunakan dengan tipe yang salah, program akan panic dan berhenti.
*/

// Fungsi random mengembalikan nilai dengan tipe interface kosong
func random() interface{} {
	return "OK" // Mengembalikan string "OK"
}

func main() {
	// Membuat instance dari struct Person
	person := Person{Name: "Eko"}
	SayHello(person) // Output: Hello Eko

	// Membuat instance dari struct Animal
	animal := Animal{Name: "Kucing"}
	SayHello(animal) // Output: Hello Kucing

	// Variabel kosong akan menyimpan hasil dari fungsi Ups
	kosong := Ups()
	fmt.Println(kosong) // Output: Ups

	// Memanggil fungsi random dan menyimpan hasilnya di variabel result
	result := random()

	// Menggunakan type assertion untuk mengubah result menjadi string
	resultString := result.(string)
	fmt.Println(resultString) // Output: OK

	// Menggunakan type assertion yang salah, mencoba mengubah result menjadi int
	// Ini akan menyebabkan panic karena result bukan integer
	// resultInt := result.(int) // Uncommenting this line will cause panic

	// Menggunakan switch untuk melakukan type assertion secara aman
	switch value := result.(type) {
	case string:
		fmt.Println("String:", value) // Output: String: OK
	case int:
		fmt.Println("Int:", value)
	default:
		fmt.Println("Unknown type")
	}
}
