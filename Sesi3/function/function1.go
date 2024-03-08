package main

import "fmt"

func main() {

	greet("Herdiyan", "Jalan Jati Kelapa No.31")

}
func greet(name, address string) {
	fmt.Println("Hello there! My name is", name)
	fmt.Println("I live in", address)
	/*
			Hello there! My name is Herdiyan
		I live in Jalan Jati Kelapa No.31
	*/
}

/*

Ketika function yang kita buat menerima lebih dari 1 parameter namun semua tipe datanya sama, maka kita tidak perlu menuliskan seluruh tipe datanya secara satu per satu.
Contohnya seperti pada gambar pertama disebelah kanan. Function greet yang kita buat menerima 2 parameter yang sama - sama memiliki tipe data yang sama yatu tipe data string.
Maka dari itu kita tidak perlu menuliskan satu per satu tipe datanya .package function

Ketika kita jalankan pada terminal, maka hasilnya akan seperti pada gambar kedua


PENJELASAN CODE :
1.Package main adalah package khusus yang menandakan bahwa file tersebut merupakan program utama yang dapat dieksekusi.
2.Di dalam fungsi main(), kita memanggil fungsi greet() dengan memberikan dua argumen, yaitu nama ("Herdiyan") dan alamat ("Jalan Jati Kelapa No.31").
3.Fungsi greet() kemudian didefinisikan di bawah fungsi main(). Fungsi ini memiliki dua parameter, yaitu name dan address, keduanya bertipe string.
4.Di dalam fungsi greet(), pesan sapaan dicetak menggunakan fmt.Println() dengan menyertakan nama dan alamat yang diberikan.
5.Output yang dihasilkan adalah pesan sapaan yang mencakup nama dan alamat yang telah diberikan.






*/
