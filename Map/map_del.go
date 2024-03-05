// MAP (Deleting value)

package main

import "fmt"

func main() {

	var person = map[string]string{
		"name":    "Herdiyan",
		"age":     "28",
		"address": "Jalan Jati Kelapa No.31",
	}

	fmt.Println("Before deleting:", person)

	delete(person, "age")

	fmt.Println("After deleting:", person)
}

/*
Before deleting: map[address:Jalan Jati Kelapa No.31 age:28 name:Herdiyan]
After deleting: map[address:Jalan Jati Kelapa No.31 name:Herdiyan]


Kita juga dapat menghapus value dari sebuah map dengan caramenggunakan fungsi delete. Contohnya seperti pada gambar disebelah kanan. Argumen pertama yang kita berikan pada fungsidelete adalah map yang ingin kita hapus atau bisa juga variabletempat map disimpan. Lalu argumen kedua adalah key yangmenyimpan value nya.Jika kita jalankan pada terminal, maka hasilnya akan seperti padagambar kedua.Bisa kita perhatikan pada gambar kedua, setelah key age dihapusmaka key dan value nya sudah tidak akan ada lagi pada map nya.

*/
