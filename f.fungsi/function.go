package main

import "fmt"

/*		Function tanpa parameter
sayHello() adalah function sederhana yang tidak memerlukan input. Ketika dipanggil, ia akan mencetak "Hello" ke layar. Function ini digunakan untuk mendemonstrasikan dasar pembuatan function.
*/
func sayHello() {
	fmt.Println("Hello")
}

/*		Function dengan parameter
sayHelloTo(firstName, lastName) menerima dua parameter bertipe string. Saat dipanggil, ia mencetak salam dengan nama lengkap yang diberikan. Ini menunjukkan bagaimana kita bisa menggunakan parameter untuk menerima input ke dalam function.
*/
func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

/*		Function dengan return value
getHello(name) mengembalikan sebuah string yang merupakan sambutan untuk nama yang diberikan. Dengan menggunakan kata kunci return, kita bisa mengirimkan kembali nilai dari dalam function.

*/
func getHello(name string) string {
	return "Hello " + name
}

/*		Function dengan multiple return value
getFullName() mengembalikan dua nilai, yaitu firstName dan lastName. Ini menunjukkan bahwa sebuah function dapat mengembalikan lebih dari satu nilai sekaligus, yang sangat berguna dalam banyak situasi.
*/
func getFullName() (string, string) {
	return "Eko", "Khannedy"
}

/*		Function dengan Named Return Values:
Di sini, kita mendeklarasikan nama variabel untuk return value di dalam signature function. Hal ini memungkinkan kita untuk langsung mengisi nilai dan mengembalikannya tanpa menyebutkan return secara eksplisit dengan variabel.
*/
func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Eko"
	middleName = "Kurniawan"
	lastName = "Khannedy"
	return
}

/*		Variadic Function:
sumAll(numbers ...int) dapat menerima banyak argumen integer. Ini sangat berguna ketika jumlah argumen tidak diketahui sebelumnya. Kita bisa menggunakan ellipsis (...) untuk mendefinisikannya.
*/
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

/*		Function sebagai Value: (disimpan dalam variabel)
Di sini, kita menyimpan function getGoodBye ke dalam variabel goodbye. Kita kemudian dapat memanggil function tersebut melalui variabel, menunjukkan bahwa function adalah tipe data di Go.
*/
func getGoodBye(name string) string {
	return "Good Bye " + name
}

/*		Function sebagai Parameter:
sayHelloWithFilter(name, filter) menerima sebuah function sebagai parameter kedua yang digunakan untuk memfilter input sebelum mencetaknya. Ini menunjukkan fleksibilitas function di Go.
*/
func sayHelloWithFilter(name string, filter func(string) string) {
	fmt.Println("Hello", filter(name))
}

// Contoh filter function
func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

/*		Recursive function untuk menghitung factorial
factorialRecursive(value) adalah contoh function yang memanggil dirinya sendiri untuk menghitung faktorial dari value. Rekursi bisa lebih efisien dan mudah dibaca untuk beberapa jenis masalah.
*/
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

/*		Contoh anonymous function dan recursive function
Dalam contoh ini, kita mendeklarasikan function tanpa nama dan menyimpannya dalam variabel blacklist. Ini berguna ketika kita tidak perlu mendefinisikan function yang sama di banyak tempat.
*/
type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You Are Blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

/*		Closure
	Closure adalah kemampuan sebuah function untuk mengakses variabel yang dideklarasikan di luar function tersebut. Dalam contoh ini, increment adalah closure yang mengakses dan memodifikasi counter yang dideklarasikan di closureExample. Ketika kita memanggil increment dua kali, counter meningkat setiap kali dan mencetak nilainya di akhir.

*/
func closureExample() {
	counter := 0
	increment := func() {
		fmt.Println("Increment")
		counter++ // Mengakses variabel 'counter' dari scope luar
	}
	increment() // Memanggil function increment
	increment() // Memanggil lagi
	fmt.Println("Counter:", counter) // Menampilkan nilai counter
}

/*		Defer
	defer digunakan untuk menunda eksekusi suatu function sampai function yang memanggilnya selesai. Dalam contoh ini, logging() akan dipanggil setelah runApplication() selesai, terlepas dari apakah ada error atau tidak.
*/
func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	defer logging() // Menjadwalkan function logging untuk dieksekusi setelah runApplication selesai
	fmt.Println("Run Application")
}

/*		Panic
	panic digunakan untuk menghentikan program ketika terjadi kondisi yang tidak dapat ditangani, seperti error kritis. Ketika panic dipanggil, semua defer function yang telah dideklarasikan akan tetap dieksekusi sebelum program benar-benar berhenti. Dalam contoh ini, jika error bernilai true, program akan memicu panic.
*/
func endApp() {
	fmt.Println("End App")
}

func runApp(error bool) {
	defer endApp() // Menjadwalkan function endApp untuk dieksekusi setelah runApp selesai
	if error {
		panic("ERROR") // Menghentikan program dan memicu panic
	}
	fmt.Println("Application running smoothly")
}

/*		Recover
recover digunakan untuk menangkap panic yang terjadi dan melanjutkan eksekusi program. Jika panic terjadi, kita bisa menggunakan recover untuk menangkap pesan panic dan mencegah program berhenti secara tiba-tiba. Dalam contoh ini, runApp(true) akan memicu panic, tetapi kita menangkapnya dengan recover, sehingga program tetap berjalan.
*/
func runAppWithRecover(error bool) {
	defer endApp() // Menjadwalkan function endApp untuk dieksekusi setelah runAppWithRecover selesai
	if error {
		panic("ERROR") // Menghentikan program dan memicu panic
	}
	fmt.Println("Application running smoothly")
}


func main() {
	// Memanggil function tanpa parameter
	sayHello()

	// Memanggil function dengan parameter
	sayHelloTo("Eko", "Khannedy")

	// Memanggil function dengan return value
	result := getHello("Eko")
	fmt.Println(result)

	// Memanggil function dengan multiple return value
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	// Mengabaikan salah satu return value
	firstNameOnly, _ := getFullName()
	fmt.Println(firstNameOnly)

	// Memanggil function dengan named return values
	first, middle, last := getCompleteName()
	fmt.Println(first, middle, last)

	// Memanggil variadic function
	total := sumAll(10, 10, 10, 10)
	fmt.Println("Total:", total)

	// Menyimpan function dalam variable
	goodbye := getGoodBye
	fmt.Println(goodbye("Eko"))

	// Menggunakan function sebagai parameter
	sayHelloWithFilter("Eko", spamFilter)

	// Menggunakan recursive function
	fmt.Println("Factorial of 5:", factorialRecursive(5))

	// Contoh anonymous function
	blacklist := func(name string) bool {
		return name == "Anjing"
	}
	registerUser("Eko", blacklist)
	registerUser("Anjing", func(name string) bool {
		return name == "Anjing"
	})


	// Menjalankan contoh Closure
	fmt.Println("=== Closure Example ===")
	closureExample()

	// Menjalankan contoh Defer
	fmt.Println("\n=== Defer Example ===")
	runApplication()

	// Menjalankan contoh Panic
	fmt.Println("\n=== Panic Example ===")
	defer func() {
		if r := recover(); r != nil { // Menangkap panic jika terjadi
			fmt.Println("Recover from panic:", r)
		}
	}()
	runApp(true) // Memicu panic

	// Menjalankan contoh Recover
	fmt.Println("\n=== Recover Example ===")
	runAppWithRecover(true) // Ini akan memicu panic dan mengeksekusi endApp
}
