// Conditions ( Temporary variable )
// Pada bahasa go, kita dapat membuat suatu variable yang dimana variable tersebut hanya bisa diakses dan digunakan pada scope  scope block dari suatu kondisional.

//Contoh nya seperti gambar dibawah ini :

package main

import "fmt"

func main() {
	var currentYear = 2021

	if age := currentYear - 1996; age < 17 {
		fmt.Println("Kami belum boleh membuat kartu sim")

	} else {
		fmt.Println("Kamu sudah boleh membuat kartu sim")
	}
}

// Output : Kamu sudah boleh membuat kartu sim

/*
Penjelasan :
Jika kita perhatikan pada baris 8, age merupakan sebuah bariable yang terbuat pada scope block kondisional if else nya .

Kita membuat variable age yag dimana kita membuat perkurangan antara variable currentYear dengan angk 1996.
Lalu setekag tanda semicolon; kita langsung membuat kondisinya yang mengatakan jika bariable age lebih kecil dari pada angka 17 .  maka kondisi dibawah akan menghasilkan kalimat yang ada pada blok else kaena variabl age menghasilkan nilai dengan angka 23 .

*/
