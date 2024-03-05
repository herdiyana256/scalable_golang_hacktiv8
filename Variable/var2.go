package main

import "fmt"

func main() {
	var name string
	var age int = 23

	name = "Herdiyan Adam Putra"
	age = 28

	fmt.Println("Ini adalah namanya ==>", name)
	fmt.Println("Ini adalah umurnya ==>", age)
}

/*

VARIABLE WITH DATA TYPE
Kita juga bisa me-reassign suatu variavle dengan suatu nilai asalkan kita memberikan nilai dengan tipe ata yang sama dengan tipe data yang sudah dimiliki oleh varible tersebut.
Contoh seperti pada gambar sebelah kanan

Jika kita perhatikan pada gambar pertama, kita tidak memberikan nilai apapun pada variavle name pada pertama kali kita membuat variable tersebut. Hal ini dapat dilakukan asalkan kita sudah memberikan tipe datanya.
Lalu kita me-reassign variable namae dan age dengan suatu nilai baru  dan tidak ada eror yang terjadi. Hal ini dikarenakan kita masih memberikan nilai dengan tipe data tang sama seperti pertama kali variable name dan age dibuat.

*/
