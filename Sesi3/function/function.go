
package main  

import "fmt"


func main () {
		greet("Herdiyan", 28)

		
}

		func greet(name string, age int8)
			fmt.Printf("Hello There! My Name is %s and I'm %d years old", name, age)

/*
Bahasa Go juga memiliki function atau fungsi nya sendiri. Cara menulis sebuah function pada Go cukup mudah. Caranyaadalah dengan menggunakan keyword func lalu diikuti dengan nama function dan parameter yang digunakan. Contohnyaseperti pada gambar pertama dibawah. Kita telah membuat function bernama greet yang menerima  2 parameter dengan tipedata string dan int8. Function greet merupakan sebuah function yang tidak mengembalikan/me-return nilai apapun, melainkanfunction ini hanya bertugas untuk menampilkan text pada terminal kita menggunakan fmt.Printf. Ketika kita jalankan padaterminal kita, maka hasilnya akan seperti pada gambar kedua dibawah.Perlu diingat disini bahwa, kita harus menjelaskan jenis dari tipe data nya pada parameter function jika function yang kitagunakan menerima sebuah parameter.
*/