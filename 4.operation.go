package main

import "fmt"

func main() {
	fmt.Println("=== Operasi Matematika ===")
	/*
		Operasi Matematika
		Go-Lang mendukung berbagai operasi matematika dasar, seperti penjumlahan, pengurangan, perkalian, pembagian, dan sisa pembagian. Berikut adalah operator yang umum digunakan:

		+ : Penjumlahan
		- : Pengurangan
		* : Perkalian
		/ : Pembagian
		% : Sisa Pembagian
	*/

	// Mendeklarasikan variabel
	var a = 10
	var b = 5

	// Operasi Matematika
	fmt.Println("Operasi Matematika:")
	fmt.Println("Penjumlahan:", a+b)    // Penjumlahan
	fmt.Println("Pengurangan:", a-b)    // Pengurangan
	fmt.Println("Perkalian:", a*b)      // Perkalian
	fmt.Println("Pembagian:", a/b)      // Pembagian
	fmt.Println("Sisa Pembagian:", a%b) // Sisa Pembagian

	fmt.Println("\n=== Augmented Assignments ===")
	/*
		Augmented assignments memungkinkan kita untuk melakukan operasi matematika dan sekaligus mengassign kembali nilai ke variabel dengan cara yang lebih ringkas. Berikut adalah beberapa contoh augmented assignments:

		* a = a + 10 atau a += 10 : Menambahkan 10 ke a
		* a = a - 10 atau a -= 10 : Mengurangi 10 dari a
		* a = a * 10 atau a *= 10 : Mengalikan a dengan 10
		* a = a / 10 atau a /= 10 : Membagi a dengan 10
		* a = a % 10 atau a %= 10 : Mengambil sisa pembagian a dengan 10
	*/

	var i = 10

	// Penjumlahan
	i += 10
	fmt.Println("Setelah Penjumlahan (i += 10):", i)

	// Pengurangan
	i -= 5
	fmt.Println("Setelah Pengurangan (i -= 5):", i)

	// Perkalian
	i *= 2
	fmt.Println("Setelah Perkalian (i *= 2):", i)

	// Pembagian
	i /= 5
	fmt.Println("Setelah Pembagian (i /= 5):", i)

	// Sisa Pembagian
	i %= 3
	fmt.Println("Setelah Sisa Pembagian (i %= 3):", i)

	fmt.Println("\n=== Unary Operator ===")

	// Mendeklarasikan variabel
	var j = 1

	// Unary Operators
	fmt.Println("Sebelum Unary Operators:")
	fmt.Println("Nilai j:", j)

	// Operator Increment
	j++ // j = j + 1
	fmt.Println("Setelah Increment (j++):", j)

	// Operator Increment lagi
	j++ // j = j + 1
	fmt.Println("Setelah Increment lagi (j++):", j)

	// Operator Decrement
	j-- // j = j - 1
	fmt.Println("Setelah Decrement (j--):", j)

	// Menggunakan Negative
	var k = -j
	fmt.Println("Negative dari j:", k)

	// Positive (tidak mengubah nilai)
	var l = +j
	fmt.Println("Positive dari j:", l)

	// Boolean kebalikan
	var m = false
	fmt.Println("Nilai m:", m)
	fmt.Println("Boolean kebalikan (!m):", !m)


	fmt.Println("\n=== Operasi Perbandingan ===")
	/*
		Operasi perbandingan adalah operasi yang digunakan untuk membandingkan dua data. Hasil dari operasi ini adalah nilai boolean (true atau false).
		
		>	Lebih Dari
		<	Kurang Dari
		>=	Lebih Dari Sama Dengan
		<=	Kurang Dari Sama Dengan
		==	Sama Dengan
		!=	Tidak Sama Dengan
	*/
	var name1 = "Eko"
	var name2 = "Eko"

	// Menggunakan operator perbandingan ==
	var result bool = name1 == name2

	fmt.Println("Apakah name1 sama dengan name2?", result) // Output: true

	fmt.Println("\n=== Operasi Boolean ===")
	/*
		Operasi boolean menggunakan operator logika untuk menghasilkan nilai boolean berdasarkan operasi yang dilakukan.
		&&	Dan
		`	
		!	Kebalikan
	*/
	// Contoh Operasi AND (&&)
	var nilaiAkhir = 90
	var absensi = 80

	// Cek kelulusan berdasarkan nilai akhir dan absensi
	var lulusNilaiAkhir bool = nilaiAkhir > 80       // true
	var lulusAbsensi bool = absensi > 80             // false
	var lulus bool = lulusNilaiAkhir && lulusAbsensi // false

	fmt.Println("Apakah siswa lulus?", lulus) // Output: false

	// Contoh Operasi OR (||)
	nilaiAkhir = 70
	absensi = 90

	lulusNilaiAkhir = nilaiAkhir > 80       // false
	lulusAbsensi = absensi > 80             // true
	lulus = lulusNilaiAkhir || lulusAbsensi // true

	fmt.Println("Apakah siswa lulus?", lulus) // Output: true

	// Contoh Operasi NOT (!)
	lulus = false

	// Kebalikan dari nilai lulus
	var tidakLulus bool = !lulus

	fmt.Println("Apakah siswa tidak lulus?", tidakLulus) // Output: true

}
