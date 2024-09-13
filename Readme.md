# Belajar Go-Lang

Go-Lang adalah bahasa pemrograman yang cepat, efisien, dan open-source yang dibuat oleh **Google**. Bahasa ini dirancang untuk memenuhi kebutuhan pemrograman modern seperti pengembangan **microservices** dan sistem terdistribusi. Berikut adalah beberapa poin sejarah penting dari Go-Lang:

## Sejarah Go-Lang

- **Dibuat di Google** menggunakan bahasa pemrograman **C**.
- **Dirilis ke publik** sebagai **open-source** pada tahun **2009**.
- Go-Lang mulai **populer sejak digunakan untuk membuat Docker** pada tahun **2011**.
- Saat ini, banyak **teknologi baru** yang dibangun menggunakan **Go-Lang** seperti:
  - **Kubernetes**
  - **Prometheus**
  - **CockroachDB**
  - Dan lainnya.
- **Go-Lang mulai populer untuk pembuatan Backend API** di sistem **Microservices** karena performa dan skalabilitasnya.

## Mengapa Go-Lang?

1. **Sederhana** dan mudah dipelajari.
2. **Cepat dan efisien**, dengan **garbage collection** dan dukungan paralelisme bawaan.
3. Mendukung **concurrency** yang baik, menggunakan **goroutines** dan **channels**.
4. **Dikembangkan untuk skala besar** – ideal untuk sistem terdistribusi dan **microservices**.
5. **Ekosistem yang berkembang pesat** dengan banyak framework dan tool yang mendukung.

## Instalasi

### Langkah-langkah instalasi Go-Lang:

1. Download Go-Lang dari [https://golang.org/dl/](https://golang.org/dl/).
2. Ikuti instruksi instalasi sesuai dengan sistem operasi kamu.
3. Cek apakah instalasi berhasil dengan menjalankan perintah:

```bash
go version
```

Jika Go-Lang terinstal dengan benar, kamu akan melihat versi Go-Lang yang terpasang.

## membuat Module

1. Masuk kedalam folder yang telah dibuat 
2. ketik syntax dibawah
3. 
```bash
go mod init learn-golang
```

1. Maka go.mod akan terinstal

## Program Pertama: "Hello, World!"

Berikut adalah contoh program pertama menggunakan Go-Lang:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

Untuk menjalankan program di atas, simpan file dengan nama `hello.go`, kemudian jalankan perintah:

```bash
go run hello.go
```

## Proyek dan Teknologi yang Menggunakan Go-Lang

Beberapa proyek besar yang dibangun dengan Go-Lang:

- **Docker**: Platform untuk menjalankan aplikasi di dalam container.
- **Kubernetes**: Sistem orkestrasi container.
- **Prometheus**: Sistem monitoring dan alerting.
- **CockroachDB**: Database SQL yang scalable.

## Langkah Selanjutnya

Jika kamu tertarik untuk mendalami Go-Lang lebih lanjut, berikut beberapa langkah yang bisa diambil:

1. **Belajar Concurrency di Go** – Pelajari cara kerja **goroutines** dan **channels**.
2. **Bangun REST API** dengan menggunakan Go-Lang dan framework seperti **Gin** atau **Echo**.
3. Pelajari **Microservices** dan bangun sistem terdistribusi dengan Go-Lang.

## Sumber Belajar

- [Dokumentasi Resmi Go-Lang](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)

Selamat belajar Go-Lang dan semoga sukses!

# GoLang: Multiple Main Function Issue

GoLang adalah bahasa pemrograman yang menuntut nama fungsi dalam sebuah module atau project bersifat unik. Ini berarti, kita **tidak boleh membuat nama function yang sama** dalam satu module atau project. Salah satu contoh yang sering terjadi adalah ketika kita secara tidak sengaja membuat **multiple `main()` functions** dalam satu project.

## Permasalahan
Jika kita memiliki beberapa file di dalam satu project/module Go, seperti `helloworld.go` dan `sample.go`, dan kedua file tersebut memiliki function `main()`, maka proses build module akan mengalami **error** karena Go mendeteksi adanya **duplikasi function `main()`**.

Sebagai contoh:
- **`helloworld.go`**
  ```go
  package main
  
  import "fmt"
  
  func main() {
      fmt.Println("Hello, World!")
  }
  ```

- **`sample.go`**
  ```go
  package main
  
  import "fmt"
  
  func main() {
      fmt.Println("This is another main function")
  }
  ```

Jika Anda mencoba untuk melakukan build dengan dua file ini, Go akan menghasilkan error:

```
go build . 
# error: function main redeclared in this block
```

## Penyebab
Function `main()` dalam Go bersifat **unik**. Fungsi ini adalah entry point untuk aplikasi Go, sehingga dalam satu module atau aplikasi hanya boleh ada satu function `main()` yang akan dijalankan.

## Solusi
Untuk mengatasi masalah ini, ada beberapa cara yang bisa dilakukan:
1. **Pisahkan Function `main()` ke Dalam Modul Berbeda**: Jika Anda ingin memiliki beberapa file Go dengan fungsi `main()`, Anda harus memisahkannya menjadi **modul terpisah** atau mengubah nama salah satu function tersebut.
   
2. **Gunakan Package yang Berbeda**: Jika Anda tidak memerlukan dua `main()` functions, pertimbangkan untuk menggunakan package selain `main` di file kedua. Misalnya:
   - Ubah `sample.go` menjadi:
     ```go
     package sample
     
     import "fmt"
     
     func SampleFunc() {
         fmt.Println("This is a sample function")
     }
     ```
     
3. **Ubah Nama Function**: Jika kedua fungsi `main()` tidak diperlukan, Anda dapat mengganti nama salah satu fungsi agar tidak terjadi duplikasi.

## Kesimpulan
Dalam GoLang, hanya boleh ada satu `main()` function per module atau project, karena function `main()` adalah entry point dari aplikasi. Untuk mengatasi masalah duplikasi function `main()`, Anda dapat memisahkan modul, menggunakan package berbeda, atau mengubah nama function.

---

File `README.md` ini memberi pengantar tentang Go-Lang, sejarah, cara instalasi, serta program sederhana untuk memulai.
    
```