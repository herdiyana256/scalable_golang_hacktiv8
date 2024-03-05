/*
Looping (first way of looping)
Looping atau perulangan merupakan suatu proses berulang yang dimana proses tersebut akan berhenti jika memenuhi suatu kondisi.
Bahasa Go hanya memiliki 1 looping yaitu looping dengan menggunakan keyword for atau yang kita kenal dengan sebutan for loop




Jikaika kita lihat pada gambar pertama dibawah, ini merupakan cara pertama agar kita dapat melakukan looping pada bahasaGo. Looping tersebut akan bekerja selama variable i memiliki nilai kurang dari angka 3.Jika kita jalankan maka hasilnya akan seperti pada gambar kedua dibawah. Perlu diingat disini bahwa kita tidak boleh lupauntuk menambahkan variable  i seperti pada gambar pertama yang dimana i++ sama saja dengan i = i + 1. Jika variable itidak ditambahkan maka akan menimbulkan infinite loop atau proses perulangan yang tidak akan berhenti.

*/

package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println("Angka", i)
	}
}

/*
Angka 0
Angka 1
Angka 2

*/
