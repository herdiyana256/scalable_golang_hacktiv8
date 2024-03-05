/*
Switch
Selain menggunakan keybword if , else if dan else, kita juga bisa menggunakan keyword switch untuk membuat suatu kondisional, seperti contohnya pada gambar disebelah kanan.

Jika kita perhatikan, kondisional pada gambar disebelah kanan akan menghasilkan kalimat "perfect" karena case pertama sudah terpenuhi.
Jika tidak ada case yang terpenuhi maka kondisional tersebut akan menghasilkan kalimat "not bad" karena keyword default sama dengan keyword  e;se yang dimana akan di eksekusi blocknya jika tidak ada kondisi yang terpenuhi .


Switch pada bahasa Go tidak memerlukan keyword break, Jadi jika suatu case telah terpenuhi maka dia tidak akan berlanjut kepada case berikutnya


*/

// contoh
package main

import "fmt"

func main() {

	var score = 8

	switch score {
	case 8:
		fmt.Println("perfect") // perfect
	case 7:
		fmt.Println("Awesome")
	default:
		fmt.Println("default")
	}
}
