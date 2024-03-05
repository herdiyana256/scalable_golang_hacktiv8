package main

func main() {
	// Deklarasi variable dengan tipe data uint8

	var a uint8 = 10
	var b byte // byte adalah aliasa dari tipe data uint8

	b = a // no error, karena byte memiliki tipe data yang sama dengan uint8

	b = a // no error. karena byte memiliki tipe data yang sama dengan uint8
	_ = b
}

/*
Perhatikan pada contoh diatas mengenai penggunaan aliase:

●Variable a memiliki tipe data uint8, sedangka varible b memiliki tipe data byte.
●Ketika b di-reassign dengan a, maka tidak akan terjadi error karena byte merupakan tipe data alias daritipe data uint8.

*/
