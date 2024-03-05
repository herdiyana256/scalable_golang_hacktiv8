package main

func main() {

	var firstVariable string

	var name, age, address = "Herdiyan", 28, "Jalan Jati Kelapa No.31 "

}

//contoh I

/*

Bahasa Golang memiliki satu aturan yang cukup strict /ketat, yang dimana tidak ada variable yang boleh menganggur ketika sudah kita buat.
Contoh nya pada gambar pertama disebelah kanan.

Jika kita perhatikan pada gambar peetama, terjadi error pada saat compile time yag ditandai dengan garis merah.constini terjadi karena kita telah mebuat dan mendeklarasikan variavle tetapi variable variable tersebut kita biarkan menganggur atau tidak dipakai. Maka dari itu cara menanggulangi adalah dengan memakai variable underscore seperti pada gambar kedua.const

dengan variable underscore maka kita dapat menghindari error yang akan terjadi jika kita mempunyai suatu variable yang menganggur atau tidak dipakai.
*/

package main


func main () {
	var firstVariable string

	var name, age, address = "Herdiyan Adam", 28, "Jalan Jati Kelapa No.31"

	-,-,-,- = firstVariable, name, age, address
	
	
}