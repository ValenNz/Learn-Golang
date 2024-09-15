package database

var connection string

// init function untuk menginisialisasi koneksi
func init() {
	connection = "MySQL"
}

// GetDatabase mengembalikan string nama database
func GetDatabase() string {
	return connection
}
