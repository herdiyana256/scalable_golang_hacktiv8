package main

import "fmt"

func main() {
	var fruits1 = []string{"apple", "banana", "mango"}

	var fruits2 = []string{"durian", "pineapple", "starfruit"}

	fruits1 = append(fruits1, fruits2...) //appending the elements of

	fmt.Printf("%#v", fruits1)
	// []string{"apple", "banana", "mango", "durian", "pineapple", "starfruit"}

	/*
		Slice (append function with ellipsis)Jika kita ingin memasukkan seluruh element-element padasuatu array ke dalam array lainnya, maka kita dapatmenggunakan tanda ellipsis (...) atau tanda titik tiga berurut.Contohnya seperti pada gambar pertama di sebelah kanan.Terdapat 2 buah variable bernama fruits1 dan fruits2 yangmasing-masing memiliki tipe data slice of string danmenyimpan nama-nama buah.Lalu variable fruits1 mencoba untuk menambahkan seluruhelement yang terdapat pada variable fruits2, dan memakaitanda ellipsis untuk mengambil seluruh elementnya.Jika dijalankan pada terminal, maka hasilnya akan  sepertipada gambar kedua disebelah kanan.


	*/
}
