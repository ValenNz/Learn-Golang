package main

import(
	"fmt"
)

func main()  {
	fmt.Println("\n=== Type Data Convertation ===")
	// Deklarasi variable dengan tipe data int32
	var nilai32 int32 = 32768

	// Konversi dari int32 ke int64
	var nilai64 int64 = int64(nilai32)

	// Konversi dari int32 ke int16
	var nilai16 int16 = int16(nilai32)

	// Menampilkan hasil konversi
	fmt.Println("Nilai 32-bit:", nilai32)
	fmt.Println("Nilai 64-bit:", nilai64)
	fmt.Println("Nilai 16-bit:", nilai16)

	// Mengkonversi string ke byte dan sebaliknya
	var name = "Nuevalen Refitra"
	var e = name[0]         // Mengambil karakter pertama sebagai byte
	var eString = string(e) // Mengkonversi byte kembali ke string

	// Menampilkan hasil konversi
	fmt.Println("Nama:", name)
	fmt.Println("Karakter Pertama:", eString)

	/*		Keterangan:
	*	Deklarasi Variable: Di sini, kita mendeklarasikan nilai32 sebagai int32 dan kemudian mengkonversinya ke int64 dan int16.
	*	Konversi Byte ke String: Kita juga menunjukkan bagaimana cara mengambil karakter pertama dari string dan mengkonversinya dari byte kembali ke string.
	*	Menampilkan Hasil: Menggunakan fmt.Println untuk menampilkan nilai dari variable yang telah dikonversi.
	 */
}