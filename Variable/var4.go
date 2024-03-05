package main

import "fmt"

func main() {
	var name = "Herdiyan"
	var age = 28

	fmt.Printf("%T, %T", name, age)
}

//string, int

/*
	Variable without data type

	Kita juga dapat membuat suatu variable dengan tidak memberikan tipe datanya secara ekplisit . Hal ini dikenal dengan nama type inference yang dimana sstem dari golang yang akan menentukan sendiri tipe datanya secara otomatis tergantung daru nilai yang kita berikan pada variable tersebut.
	Contohnya seperti pada gambar disebelah kanan.const


	Jika perhatikan gambar pertama disebelah kanan.const


	Jika kita perhatikan gambardisebelah kanan, kita sekarang menggunakan fungsi fmt.Printf.
	Fungsi fmt.Printf mempunyai kegunaan yang sama dengan fungsi fmt.Prinln.
	Hanya saja struktur output dari fungsi fmt.Printf ditentukan dari flag yang kita berikan. Hal ini akan kita bahas nanti.

	Jika kita coba menjalankan gambar pertama, maka kita akan mendapatkan keterangan tipe data kedua variavle yang kita buat. Variable name akan secara otomatis memiliki tipe data string dan variavle age akan mempunyai tipe data int.const



*/

//===============================================

package main

import "fmt"

func main() {
		name := "Herdiyan"
		age := 28

		fmt.Printf("%T, %T", name, age)
}

/*
PENJELASAN :
Kita juga dapat membuat variable dengan menerapkan teknik short declaration yang dimana kita tidak perlu menggunakan keyword var lagi. Kita bisa dengan mudah nya membuat suatu variable tanpa tipe data seperti pada gambar di sebelah kanan.Jika kita perhatikan, kita menggunakan tanda titik dua (:) sebelum tanda sama dengan (=). Dengan seperti ini maka kita telah menerapkan teknik short declaration yang dimana teknik ini cukup mempermudah kita dalam membuat suatu variable. Tapi perlu diingat bahwa, ketika kita memakai teknik short declaration, maka kita harus langsung memberikan nilai kepada variable tersebut. Jika tidak maka akan terjadi error


*/
