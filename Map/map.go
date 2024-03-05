package main

import "fmt"

func main() {

	var person map[string]string // Deklarasi

	person = map[string]string{} // inisialisasi

	person["name"] = "Herdiyan"

	person["age"] = "28"

	person["address"] = "Jalan Jati Kelapa No.31"

	fmt.Println("name:", person["name"])
	fmt.Println("age:", person["age"])
	fmt.Println("address:", person["address"])
}

/*
name: Herdiyan
age: 28
address: Jalan Jati Kelapa No.31


PENJELASAN :
Ketika kita ingin membuat sebuah map, maka kita juga haruslangsung menginisialisasikannya. Contohnya seperti pada gambardisebelah kanan.Kita membuat sebuah variable bernama person yang memiliki tipedata map.Penulisan map[string]string berarti tipe data  key dari map harusstring dan value dari map juga harus string. Lalu jika kita inginmemberikan nilai atau value kepada map nya, maka harus diinisialisasi terlebih dahulu.Bisa kita lihat pada line 12, 13 dan 14, kita memberikan value kepadavariable person dengan key yang berbeda-beda atau unik. Kita jugadapat mendapatkan value dari map dengan cara mengakses key darimap nya seperti pada line 18 - 20.Jika kita jalankan pada terminal kita, maka hasilnya akan seperti padagambar kedua di sebelah kanan.

Kita juga dapat memberikan value beserta dengan key nya denganlangsung. Contohnya seperti pada gambar pertama di sebelahkanan.Jika kita perhatikan pada line 33, kita perlu menambahkan tandakoma , setelah kita menuliskan value nya walaupun di bawah nyasudah tidak ada value dan key baru. Hal ini dikarenakan kita menulismap dengan cara horizontal (ke bawah).Jika kita jalankan pada terminal kita, maka hasilnya akan seperti padagambar kedua.

*/
