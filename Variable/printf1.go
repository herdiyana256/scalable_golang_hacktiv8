// SALAH   --->  package main

// func main() {

// 	var nama = "Herdiyan"

// 	var age = 28

// 	var address = "Jalan Jati Kelapa No.31"

// 	fmt.printf("Halo nama ku %s, umurku adalah %d, dan aku tinggal di %s", nama, age, address)

// }

package main

import (
	"fmt"
)

func main() {

	var nama = "Herdiyan"

	var age = 28

	var address = "Jalan Jati Kelapa No.31"

	fmt.Printf("Halo nama ku %s, umurku adalah %d, dan aku tinggal di %s", nama, age, address)

}

//Halo nama ku Herdiyan, umurku adalah 28, dan aku tinggal di Jalan Jati Kelapa No.31

/*
penjelasan :


fmt.Printf UsageSebelumnya kita telah menggunakan fungsi fmt.Printf untuk mengetahui tipe data suatu variable dan kali ini akan kita bahas sedikit mengenai fungsi fmt.Printf tersebut. Struktur output yang dihasilkan oleh fungsi fmt.Printf akan bergantung dari flag yang kita pakai. Seperti contohnya jika kita ingin mengetahui tipe data dari suatu variable, maka kita dapat menggunakan verb% T. Contohnya seperti pada gambar pertama di sebelah kanan. Jika kita jalankan maka hasilnya akan seperti pada gambar kedua.Jika kita perhatikan pada gambar pertama, kita juga memakai simbol \n yang dimana simbol ini dapat kita pakai untuk membuat baris baru. Karena pada dasarnya, fmt.Printf tidak akan membuat baris baru tidak seperti fmt.Println yang menambah baris baru secara otomatis.


*/
