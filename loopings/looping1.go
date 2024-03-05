package main

import "fmt"

func main() {
	var i = 0

	for i < 3 {
		fmt.Println("Angka", i)
		i++
	}
}

/*
Angka 0
Angla 1
Angka 2



PENJELASAN :
Cara kedua dalam melakukan looping pada bahasa Go adalah dengan menyelipkan kondisional seperti padalooping dengan menggunakan keyword while.Contohnya seperti pada gambar pertama dibawah. Looping pada gambar pertama dibawah menggunakankondisi yang dimana jika nilai yang dimiliki oleh variable i masih kurang dari angka 3, maka proses loopingtersebut akan terus berlanjut. Namun jika sudah melebihi dari angka 3, maka proses looping tersebut akanberhenti. Jika kita menjalankan pada terminal kita, maka hasilnya akan seperti pada gambar kedua dibawah.


*/
