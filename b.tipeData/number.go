package main

import (
	"fmt"
)

func main() {
	// Tipe data Integer
	fmt.Println("=== Type Data Integer ===")
	fmt.Println("Nilai intA (int):", 42)      // Integer default (int)
	fmt.Println("Nilai intB (int8):", 127)    // Integer dengan ukuran 8 bit
	fmt.Println("Nilai intC (int16):", 32767) // Integer dengan ukuran 16 bit
	fmt.Println("Satu =", 1)
	fmt.Println("Dua = ", 2)

	// Tipe data Floating Point
	fmt.Println("\n=== Type Data Floating Point ===")
	fmt.Println("Nilai floatA (float32):", 3.14)  // Floating point dengan ukuran 32 bit
	fmt.Println("Nilai floatB (float64):", 3.141) // Floating point dengan ukuran 64 bit
	fmt.Println("Tiga Koma Lima = ", 3.5)
}
