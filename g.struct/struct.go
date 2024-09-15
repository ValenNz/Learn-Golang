package main

import "fmt"

/*		
	Struct
	Struct adalah template data yang digunakan untuk menggabungkan beberapa tipe data menjadi satu kesatuan.
	Pada contoh di atas, struct `Customer` memiliki tiga field: `Name`, `Address`, dan `Age`.
*/

// Mendefinisikan struct dengan nama Customer
// Struct ini memiliki tiga field: Name, Address, dan Age
type Customer struct {
	Name, Address string // Field Name dan Address bertipe string
	Age           int    // Field Age bertipe int
}

/*		
	Struct Method
	Struct adalah tipe data seperti tipe data lainnya dan bisa digunakan sebagai parameter pada function.
	Jika kita ingin menambahkan method ke dalam struct, kita bisa membuat function yang terasosiasi dengan struct.
	Method adalah function yang didefinisikan untuk tipe data tertentu (dalam hal ini, struct).
*/
func (customer Customer) sayHello() {
	// Method sayHello akan mencetak pesan sapaan menggunakan field Name dari struct Customer
	fmt.Println("Hello, My Name is", customer.Name)
}

func main() {
	/*		
		Membuat Data Struct
		Struct adalah template data atau prototype data. Struct tidak bisa langsung digunakan,
		kita harus membuat data atau objek dari struct yang telah kita buat.
	*/
	
	// Membuat objek struct bernama 'eko' dari struct Customer
	var eko Customer
	eko.Name = "Eko Kurniawan"   // Mengisi field Name
	eko.Address = "Indonesia"    // Mengisi field Address
	eko.Age = 30                 // Mengisi field Age

	// Menampilkan data 'eko' ke terminal
	fmt.Println(eko)

	/*		
		Struct Literals
		Ada berbagai cara untuk membuat data dari struct, salah satunya adalah menggunakan struct literals.
	*/
	
	// Membuat objek 'joko' menggunakan struct literals dengan menyebutkan nama field
	joko := Customer{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     30,
	}
	// Menampilkan data 'joko' ke terminal
	fmt.Println(joko)

	// Membuat objek 'budi' menggunakan struct literals tanpa menyebutkan nama field
	// Data diisi secara berurutan sesuai urutan field di struct Customer
	budi := Customer{"Budi", "Indonesia", 30}
	// Menampilkan data 'budi' ke terminal
	fmt.Println(budi)

	// Membuat objek 'rully' menggunakan struct literals, hanya mengisi field Name
	// Field Address dan Age akan mendapatkan nilai default (string kosong dan 0)
	rully := Customer{Name: "Rully"}
	// Memanggil method sayHello pada objek rully
	rully.sayHello()
}
