package helper

var Version = "1.23.1" // Dapat diakses dari luar package
var application = "golang" // Tidak dapat diakses dari luar package

// SayHello adalah fungsi yang mengembalikan salam
func SayHello(name string) string {
	return "Hello " + name
}

// sayGoodBye tidak dapat diakses dari luar package
func sayGoodBye(name string) string {
	return "Goodbye " + name
}
