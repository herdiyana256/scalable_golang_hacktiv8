package main 

import "fmt"

func main() {
	var value = 2 + 2*3
	fmt.Println(value)
}

//8 

/*

Jika kita melihat pada gambar pertama disebelah kanan, kita telahmembuat suatu variable bernama value yang dimana variable valuemenyimpan nilai dari hasil kalkulasi 2 + 2 * 3. Walaupun padabahasa Go, syntax kita akan dibaca dari kiri ke kanan, namunkarena pada kalkulasi tersebut ada simbol * yang dimana simboltersebut merupakan simbol perkalian, maka Go akan melakukanperkalian tersebut terlebih dahulu kemudian baru melakukanpertambahannya.  Kalkulasi pada gambar pertama akanmenghasilkan angka 8

*/





package main 

import "fmt"


func main () {
	var value = (2 + 2) * 3 
	fmt.Println(value)
}
// 12


/*
Lalu jika kita ingin agar pertambahannyadieksekusi terlebih dahulu, maka kita bisa mengelompokkankalkulasi yang ini kita eksekusi terlebih dahulu menggunakansimbol () seperti pada gambar kedua. Jika kita jalankan gambarkedua pada terminal kita, maka akan menghasilkan angka 12.

*/
