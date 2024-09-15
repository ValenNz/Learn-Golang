package main

import "fmt"

// Struct Address
type Address struct {
	City, Province, Country string
}

// Contoh Pass by Value
func PassByValue() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := address1 // Meng-copy address1 ke address2
	address2.City = "Bandung" // Mengubah address2
	fmt.Println(address1) // Output: {Subang Jawa Barat Indonesia} - address1 tidak berubah
	fmt.Println(address2) // Output: {Bandung Jawa Barat Indonesia}
}

// Contoh Pointer dan Operator &
func PointerOperator() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1 // Mengambil alamat address1
	address2.City = "Bandung" // Mengubah city melalui pointer
	fmt.Println(address1) // Output: {Bandung Jawa Barat Indonesia} - address1 berubah
	fmt.Println(*address2) // Output: {Bandung Jawa Barat Indonesia} - dereferensi pointer
}

// Contoh Asterisk Operator (*)
func AsteriskOperator1() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1
	address2.City = "Bandung" // Mengubah city melalui pointer
	address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"} // Mengubah pointer
	fmt.Println(address1) // Output: {Bandung Jawa Barat Indonesia} - address1 tidak berubah
	fmt.Println(*address2) // Output: {Jakarta DKI Jakarta Indonesia}
}

func AsteriskOperator2() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1
	address2.City = "Bandung" // Mengubah city melalui pointer
	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"} // Mengubah data di address1
	fmt.Println(address1) // Output: {Jakarta DKI Jakarta Indonesia} - address1 berubah
	fmt.Println(*address2) // Output: {Jakarta DKI Jakarta Indonesia}
}

// Contoh penggunaan function new
func FunctionNew() {
	alamat1 := new(Address) // Menggunakan function new untuk membuat pointer
	alamat1.Country = "Indonesia"
	fmt.Println(*alamat1) // Output: {   Indonesia} - city dan province kosong
}

// Contoh Pointer di Function
func ChangeAddressToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func PointerInFunction() {
	address := Address{"Subang", "Jawa Barat", ""}
	ChangeAddressToIndonesia(&address) // Menggunakan pointer untuk mengubah data asli
	fmt.Println(address) // Output: {Subang Jawa Barat Indonesia} - berubah
}

// Pointer di Method
type Man struct {
	Name string
}

func (man Man) Married() {
	man.Name = "Mr. " + man.Name // Ini tidak akan mengubah nilai di main
}

func PointerInMethod1() {
	eko := Man{"Eko"}
	eko.Married()
	fmt.Println(eko.Name) // Output: Eko - tidak berubah
}

func (man *Man) MarriedPointer() {
	man.Name = "Mr. " + man.Name // Ini akan mengubah nilai di main
}

func PointerInMethod2() {
	eko := Man{"Eko"}
	eko.MarriedPointer()
	fmt.Println(eko.Name) // Output: Mr. Eko - berubah
}

func main() {
	// Menggunakan berbagai fungsi dan demonstrasi
	PassByValue()
	PointerOperator()
	AsteriskOperator1()
	AsteriskOperator2()
	FunctionNew()
	PointerInFunction()
	PointerInMethod1()
	PointerInMethod2()
}
