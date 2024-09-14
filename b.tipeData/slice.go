package main

import "fmt"

func main() {
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
	slice1 := array[2:5]                         // Slice dari indeks 2 hingga sebelum indeks 5
	fmt.Println("Slice 1 (array[2:5]):", slice1) // Output: [Maret, April, Mei]

	// 2. array[low:] -> Membuat slice dari indeks low hingga akhir array
	slice2 := array[3:]                         // Slice dari indeks 3 hingga akhir array
	fmt.Println("Slice 2 (array[3:]):", slice2) // Output: [April, Mei, Juni, Juli, Agustus, September, Oktober, November, Desember]

	// 3. array[:high] -> Membuat slice dari awal array hingga indeks sebelum high
	slice3 := array[:4]                         // Slice dari awal hingga sebelum indeks 4
	fmt.Println("Slice 3 (array[:4]):", slice3) // Output: [Januari, Februari, Maret, April]

	// 4. array[:] -> Membuat slice dari seluruh array
	slice4 := array[:]                         // Slice dari seluruh elemen array
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
	fmt.Println("Slice:", daySlice)              // Output: [Selasa, Rabu, Kamis]
	fmt.Println("Panjang Slice:", len(daySlice)) // Output: 3 (karena slice memiliki 3 elemen: Selasa, Rabu, Kamis)

	// 2. cap(slice) -> Mendapatkan kapasitas dari slice
	fmt.Println("Kapasitas Slice:", cap(daySlice)) // Output: 6 (karena kapasitas slice dari indeks 1 sampai akhir array adalah 6)

	// 3. append(slice, data) -> Menambahkan elemen baru ke slice
	daySlice = append(daySlice, "Libur")                          // Menambahkan elemen "Libur" ke slice
	fmt.Println("Slice setelah append:", daySlice)                // Output: [Selasa, Rabu, Kamis, Libur]
	fmt.Println("Panjang Slice setelah append:", len(daySlice))   // Output: 4
	fmt.Println("Kapasitas Slice setelah append:", cap(daySlice)) // Output: 6 (masih menggunakan kapasitas dari array asli)

	// Menambahkan lebih banyak elemen untuk melebihi kapasitas asli
	daySlice = append(daySlice, "Hari Ekstra", "Hari Tambahan")
	fmt.Println("Slice setelah penambahan lebih banyak elemen:", daySlice)
	fmt.Println("Panjang Slice setelah penambahan:", len(daySlice))   // Output: 6
	fmt.Println("Kapasitas Slice setelah penambahan:", cap(daySlice)) // Output: Kapasitas baru yang lebih besar dari sebelumnya, array baru dibuat.

	// Print array asli untuk melihat apakah ada perubahan
	fmt.Println("Array asli:", days) // Output: [Senin, Selasa, Rabu, Kamis, Jumat, Sabtu, Minggu] (tidak berubah)

	// 4. make([]TypeData, length, capacity): Membuat slice baru dengan panjang dan kapasitas yang ditentukan
	newSlice := make([]string, 2, 5) // Slice baru dengan panjang 2 dan kapasitas 5
	newSlice[0] = "Liburan"
	newSlice[1] = "Kerja"
	fmt.Println("Slice baru:", newSlice)
	fmt.Println("Panjang slice baru:", len(newSlice))   // Output: 2
	fmt.Println("Kapasitas slice baru:", cap(newSlice)) // Output: 5

	// 5. copy(destination, source): Menyalin elemen dari slice source ke slice destination
	fromSlice := days[:3]                                     // Slice dari elemen pertama hingga "Rabu"
	toSlice := make([]string, len(fromSlice), cap(fromSlice)) // Membuat slice tujuan dengan panjang dan kapasitas yang sama
	copy(toSlice, fromSlice)
	fmt.Println("Slice sumber:", fromSlice)            // Output: [Senin, Selasa, Rabu]
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
	fmt.Println("fruitSlice2:", fruitSlice2) // Output: ["Ups", "Grape Baru", "Honeydew"]
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
