package main

import "fmt"

func main() {
	var first = 89
	var second = 12

	fmt.Printf("Tipe data first: %T\n", first)
	fmt.Printf("Bilangan second: %T\n", second)

}

// Tipe data first : int
// Bilangan second : int

/*

Data Types - Sesi 1 Number (integer)Setelah penjelasan yang kita baca pada halaman sebelumnya,tipe data int merupakan representasi atau sama saja dengan  int32 atau int64 tergantung dari nilainya. Mari kita perhatikan pada gambar pertama disebelah kanan. Kita membuat 2 variable yang dimana variable first mempunyai nilai dengan angka positif dan variable second mempunyai nilai dengan angka negatif. Jika kita jalankan pada terminal kita maka kita bisa melihat bahwa tipe data kedua variable tersebut ada int seperti pada gambar kedua. Ada baiknya jika kita menentukan tipe data integer dengan jenis apa yang ingin kita pakai untuk menghindari boros memori. Seperti pada variable first kita bisa memberikan tipe data uint8 dan pada variable second kita bisa memberikan tipe data int8. Hal ini kita lakukan karena uint8 merepresentasikan angka dengan skala 0 - 255 sedangkan uint8 merepresentasikan angka dengan skala -128 - 127.


*/
