/* 
	Di Go, interface error digunakan untuk merepresentasikan kondisi error. Interface ini berisi satu method, yaitu Error(), yang mengembalikan string sebagai pesan error. Berikut adalah penjelasan lengkap tentang cara kerja error handling di Go, termasuk cara membuat error menggunakan library errors, serta cara membuat custom error dengan struct.

	1. Interface Error di Go
		type error interface {
    	Error() string

	Ketika terjadi error, nilai error biasanya dikembalikan sebagai nil jika tidak ada error, atau dengan pesan error jika ada masalah.
}

*/

package main

import (
    "errors"
    "fmt"
)

/* 2. Membuat Error Menggunakan Library errors
   Go memiliki package errors yang menyediakan fungsi untuk membuat error dengan mudah.
   Salah satu fungsi yang paling sering digunakan adalah errors.New(), yang mengembalikan sebuah error dengan pesan tertentu.
*/
// Fungsi untuk melakukan pembagian dengan pengecekan pembagi nol
func Pembagian(nilai int, pembagi int) (int, error) {
    if pembagi == 0 {
        return 0, errors.New("Pembagian dengan NOL") // Membuat error jika pembagi 0
    } else {
        return nilai / pembagi, nil // Tidak ada error
    }
}

/* 3. Membuat Custom Error
   Karena error adalah interface, kita dapat membuat custom error dengan mendefinisikan struct yang mengimplementasikan method Error().
   Ini memungkinkan kita untuk memberikan lebih banyak informasi dalam error.
*/
// Membuat struct validationError yang mengimplementasikan interface error
type validationError struct {
    Message string
}

// Mengimplementasikan method Error() untuk validationError
func (v *validationError) Error() string {
    return v.Message
}

// Membuat struct notFoundError yang juga mengimplementasikan interface error
type notFoundError struct {
    Message string
}

// Mengimplementasikan method Error() untuk notFoundError
func (n *notFoundError) Error() string {
    return n.Message
}

// Fungsi untuk mensimulasikan penyimpanan data dan mengembalikan error
func SaveData(id string, data any) error {
    if id == "" {
        return &validationError{Message: "validation error: id is empty"} // Mengembalikan validation error
    }
    if id != "eko" {
        return &notFoundError{Message: "data not found"} // Mengembalikan not found error
    }
    return nil // Tidak ada error
}

func main() {

    // 2. Contoh penggunaan errors.New()
    hasil, err := Pembagian(100, 10)
    if err == nil {
        fmt.Println("Hasil:", hasil) // Output: Hasil: 10
    } else {
        fmt.Println("Error:", err.Error()) // Menampilkan pesan error jika ada
    }

    // 3. Coba simpan data dengan id kosong
    err = SaveData("", nil) // Gunakan `=` karena `err` sudah dideklarasikan sebelumnya
    if err != nil {
        // Mengecek apakah error adalah validationError
        if validationErr, ok := err.(*validationError); ok {
            fmt.Println("Validation error:", validationErr.Message)
        // Mengecek apakah error adalah notFoundError
        } else if notFoundErr, ok := err.(*notFoundError); ok {
            fmt.Println("Not found error:", notFoundErr.Message)
        // Jika error tidak diketahui
        } else {
            fmt.Println("Unknown error:", err.Error())
        }
    }
}
