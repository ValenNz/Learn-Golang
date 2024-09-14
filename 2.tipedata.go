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

	// Menampilkan hasil
	fmt.Println("\n=== Type Data Boolean ===")
	fmt.Println("Nilai isTrue:", true)
	fmt.Println("Nilai isFalse:", false)

	// Deklarasi tipe data string
	var firstName string = "Nuevalen"
	var lastName string = "Alswando"

	// String kosong
	var emptyString string = ""

	// Menggabungkan string dengan operator +
	fullName := firstName + " " + lastName

	// Menampilkan hasil
	fmt.Println("\n=== Type Data	 String ===")
	fmt.Println("Nama Depan:", firstName)
	fmt.Println("Nama Belakang:", lastName)
	fmt.Println("Nama Lengkap:", fullName)
	fmt.Println("String Kosong:", emptyString)

	// Menghitung jumlah karakter
	str := "Hello, Go-Lang!"
	panjang := len(str) // Fungsi len(str) digunakan untuk menghitung jumlah karakter dalam string str.
	fmt.Println("Jumlah karakter:", panjang)

	// Mengambil karakter pada posisi ke-7
	char := str[7]
	fmt.Printf("Karakter di posisi ke-7: %c\n", char) // String di Go adalah array karakter. Untuk mengambil karakter pada posisi tertentu, kita bisa menggunakan indeks seperti str[number]. Misalnya, str[7] akan mengambil karakter ke-7 dari string.

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

	fmt.Println("\n=== Type Data Array ===")
	/* Tipe Data Array
	Array adalah tipe data yang menyimpan kumpulan data dengan tipe yang sama. Saat membuat array, kita perlu menentukan jumlah elemen yang bisa ditampung oleh array tersebut. Daya tampung array tidak bisa bertambah setelah array dibuat.

	Indeks di Array
	Setiap elemen dalam array diakses menggunakan indeks. Indeks dimulai dari 0. Contoh berikut menunjukkan beberapa data dalam array:

	Index	Data
	0		Valen
	1		Refitra
	2		Alswando
	*/

	var names [3]string   // Mendeklarasikan array dengan kapasitas 3
	names[0] = "Valen"    // Mengisi elemen pertama
	names[1] = "Refitra"  // Mengisi elemen kedua
	names[2] = "Alswando" // Mengisi elemen ketiga

	// Menampilkan elemen-elemen array
	fmt.Println(names[0]) // Output: Valen
	fmt.Println(names[1]) // Output: Refitra
	fmt.Println(names[2]) // Output: Alswando

	/*		Membuat Array Langsung
			Di Go, kita juga bisa membuat array secara langsung saat mendeklarasikan variabel. Berikut adalah contohnya:
	*/
	var values = [3]int{
		90,
		80,
		95,
	}

	// Menampilkan seluruh elemen array
	fmt.Println(values) // Output: [90 80 95]

	/*		Fungsi Array
			Beberapa operasi yang dapat dilakukan pada array adalah sebagai berikut:

			Operasi					Keterangan
			len(array)				Untuk mendapatkan panjang array
			array[index]			Mendapat data di posisi indeks
			array[index] = value	Mengubah data di posisi indeks
	*/

	var values2 = [...]int{
		90,
		80,
		95,
	}

	// Menampilkan seluruh elemen array
	fmt.Println(values2)      // Output: [90 80 95]
	fmt.Println(len(values2)) // Output: 3

	// Mengubah elemen pertama dari array
	values2[0] = 100
	fmt.Println(values2) // Output: [100 80 95]

	fmt.Println("\n=== Type Data Slice ===")
	
	/*		Slice dan Array:
		*	Slice adalah potongan dari data Array. Berbeda dengan array yang ukurannya tetap, ukuran slice bisa berubah. Slice adalah referensi ke array, sehingga setiap perubahan pada slice juga dapat mempengaruhi array yang dirujuk oleh slice tersebut.
		*	Slice memiliki tiga atribut penting:
				-	Pointer: Menunjuk elemen pertama dari slice dalam array.
				-	Length (Panjang): Menunjukkan jumlah elemen dalam slice.
				-	Capacity (Kapasitas): Menunjukkan jumlah maksimum elemen yang bisa ditampung slice dari array mulai dari elemen pertama slice.
	*/
	 
	/* 	Membuat Slice dari Array
			*	array[low:high]: Membuat slice dari elemen array mulai dari indeks low hingga indeks sebelum high.
			*	array[low:]: Membuat slice dari indeks low hingga akhir array.
			*	array[:high]: Membuat slice dari awal array hingga indeks sebelum high.
			*	array[:]: Membuat slice dari seluruh array.
	*/
   // Membuat array untuk contoh slice
	array := [...]string{"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}

	// 1. array[low:high] -> Membuat slice dari elemen array mulai dari indeks low hingga indeks sebelum high
	slice1 := array[2:5] // Slice dari indeks 2 hingga sebelum indeks 5
	fmt.Println("Slice 1 (array[2:5]):", slice1) // Output: [Maret, April, Mei]

	// 2. array[low:] -> Membuat slice dari indeks low hingga akhir array
	slice2 := array[3:] // Slice dari indeks 3 hingga akhir array
	fmt.Println("Slice 2 (array[3:]):", slice2) // Output: [April, Mei, Juni, Juli, Agustus, September, Oktober, November, Desember]

	// 3. array[:high] -> Membuat slice dari awal array hingga indeks sebelum high
	slice3 := array[:4] // Slice dari awal hingga sebelum indeks 4
	fmt.Println("Slice 3 (array[:4]):", slice3) // Output: [Januari, Februari, Maret, April]

	// 4. array[:] -> Membuat slice dari seluruh array
	slice4 := array[:] // Slice dari seluruh elemen array
	fmt.Println("Slice 4 (array[:]):", slice4) // Output: [Januari, Februari, Maret, April, Mei, Juni, Juli, Agustus, September, Oktober, November, Desember]


	/*		Operasi pada Slice:
		*	len(slice): Mengembalikan panjang dari slice (jumlah elemen dalam slice).
		*	cap(slice): Mengembalikan kapasitas dari slice (jumlah elemen yang dapat ditampung slice, mulai dari elemen pertama slice hingga elemen terakhir array).
		*	append(slice, data): Menambahkan data ke slice. Jika kapasitas slice sudah penuh, maka akan dibuat array baru dengan kapasitas lebih besar, dan data baru akan ditambahkan di posisi terakhir slice.
	*/
	
    // Membuat array untuk contoh slice
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	// Membuat slice dari array
	daySlice := days[1:4] // Slice dari indeks 1 hingga sebelum 4 -> [Selasa, Rabu, Kamis]

	// 1. len(slice) -> Mendapatkan panjang dari slice
	fmt.Println("Slice:", daySlice)               // Output: [Selasa, Rabu, Kamis]
	fmt.Println("Panjang Slice:", len(daySlice))  // Output: 3 (karena slice memiliki 3 elemen: Selasa, Rabu, Kamis)

	// 2. cap(slice) -> Mendapatkan kapasitas dari slice
	fmt.Println("Kapasitas Slice:", cap(daySlice)) // Output: 6 (karena kapasitas slice dari indeks 1 sampai akhir array adalah 6)

	// 3. append(slice, data) -> Menambahkan elemen baru ke slice
	daySlice = append(daySlice, "Libur") // Menambahkan elemen "Libur" ke slice
	fmt.Println("Slice setelah append:", daySlice) // Output: [Selasa, Rabu, Kamis, Libur]
	fmt.Println("Panjang Slice setelah append:", len(daySlice)) // Output: 4
	fmt.Println("Kapasitas Slice setelah append:", cap(daySlice)) // Output: 6 (masih menggunakan kapasitas dari array asli)

	// Menambahkan lebih banyak elemen untuk melebihi kapasitas asli
	daySlice = append(daySlice, "Hari Ekstra", "Hari Tambahan")
	fmt.Println("Slice setelah penambahan lebih banyak elemen:", daySlice)
	fmt.Println("Panjang Slice setelah penambahan:", len(daySlice)) // Output: 6
	fmt.Println("Kapasitas Slice setelah penambahan:", cap(daySlice)) // Output: Kapasitas baru yang lebih besar dari sebelumnya, array baru dibuat.

	// Print array asli untuk melihat apakah ada perubahan
	fmt.Println("Array asli:", days) // Output: [Senin, Selasa, Rabu, Kamis, Jumat, Sabtu, Minggu] (tidak berubah)

	// 4. make([]TypeData, length, capacity): Membuat slice baru dengan panjang dan kapasitas yang ditentukan
	newSlice := make([]string, 2, 5) // Slice baru dengan panjang 2 dan kapasitas 5
	newSlice[0] = "Liburan"
	newSlice[1] = "Kerja"
	fmt.Println("Slice baru:", newSlice)
	fmt.Println("Panjang slice baru:", len(newSlice)) // Output: 2
	fmt.Println("Kapasitas slice baru:", cap(newSlice)) // Output: 5

	// 5. copy(destination, source): Menyalin elemen dari slice source ke slice destination
	fromSlice := days[:3]   // Slice dari elemen pertama hingga "Rabu"
	toSlice := make([]string, len(fromSlice), cap(fromSlice)) // Membuat slice tujuan dengan panjang dan kapasitas yang sama
	copy(toSlice, fromSlice)
	fmt.Println("Slice sumber:", fromSlice) // Output: [Senin, Selasa, Rabu]
	fmt.Println("Slice tujuan setelah copy:", toSlice) // Output: [Senin, Selasa, Rabu]


	/*	Hati-Hati Saat Membuat Array
		*	Jika kita salah, maka yang kita buat bukanlah Array, melainkan Slice
		*	iniArray adalah array, sedangkan iniSlice adalah slice.
		*	Perbedaan terletak pada cara deklarasi: array memiliki ukuran tetap, sedangkan slice tidak.
	*/
	iniArray := [...]int{10, 20, 30, 40, 50}
	iniSlice := []int{10, 20, 30, 40, 50}
	fmt.Println("\nArray vs Slice:")
	fmt.Println("Array:", iniArray) // Output: [10, 20, 30, 40, 50]
	fmt.Println("Slice:", iniSlice) // Output: [10, 20, 30, 40, 50]

	// Contoh lain dengan append
	fruit := [...]string{"Apple", "Banana", "Cherry", "Date", "Elderberry", "Fig", "Grape"}
	fruitSlice1 := fruit[5:] // Slice dari "Fig" dan "Grape"
	fruitSlice1[0] = "Fig Baru"
	fruitSlice1[1] = "Grape Baru"
		// fruitSlice1[0] diubah menjadi "Fig Baru" dan fruitSlice1[1] menjadi "Grape Baru". Perubahan ini juga mempengaruhi array fruit karena slice mengacu pada data yang sama.
	fmt.Println("\nArray Buah Setelah Diubah:", fruit) 
	// Output: ["Apple", "Banana", "Cherry", "Date", "Elderberry", "Fig Baru", "Grape Baru"]

	fruitSlice2 := append(fruitSlice1, "Honeydew")
		//	append(fruitSlice1, "Honeydew") menambahkan "Honeydew" ke dalam slice. Setelah append, slice bisa menjadi array baru jika kapasitas penuh.
	fruitSlice2[0] = "Ups"
	fmt.Println("fruitSlice2:", fruitSlice2)    // Output: ["Ups", "Grape Baru", "Honeydew"]
	fmt.Println("Array Buah Setelah Append:", fruit) 
	// Array tetap sama karena kapasitas penuh setelah append
	// Output: ["Apple", "Banana", "Cherry", "Date", "Elderberry", "Fig Baru", "Grape Baru"]

	// Contoh dengan array lain (vegetables)
	vegetables := [...]string{"Carrot", "Broccoli", "Peas", "Spinach", "Potato", "Tomato", "Cucumber"}
	vegetableSlice1 := vegetables[4:]
	vegetableSlice1[0] = "Potato Baru"
	vegetableSlice1[1] = "Tomato Baru"
	fmt.Println("\nArray Sayur Setelah Diubah:", vegetables) 
	// Output: ["Carrot", "Broccoli", "Peas", "Spinach", "Potato Baru", "Tomato Baru", "Cucumber"]

	vegetableSlice2 := append(vegetableSlice1, "Zucchini")
	vegetableSlice2[0] = "Oops"
	fmt.Println("vegetableSlice2:", vegetableSlice2) // Output: ["Oops", "Tomato Baru", "Zucchini"]
	fmt.Println("Array Sayur Setelah Append:", vegetables) 
	// Output tetap tidak berubah karena append menciptakan array baru
	// Output: ["Carrot", "Broccoli", "Peas", "Spinach", "Potato Baru", "Tomato Baru", "Cucumber"]


}
