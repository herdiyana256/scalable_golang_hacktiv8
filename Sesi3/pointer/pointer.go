// Di Go (Golang), sebuah pointer adalah variabel yang menyimpan alamat memori dari suatu nilai. Pointer digunakan untuk mengakses dan memanipulasi data secara langsung di lokasi memori yang sebenarnya, bukan melalui salinan nilai tersebut. Hal ini memungkinkan kita untuk menghindari penggunaan memori yang berlebihan dan meningkatkan kinerja program dalam beberapa kasus.

// Contoh penggunaan pointer di Go:

package main

import "fmt"

func main() {
	var a int = 42
	var b *int // mendeklarasikan variabel b sebagai pointer ke tipe data int

	b = &a // mengisi b dengan alamat memori dari a

	fmt.Println("Value of a:", a)
	fmt.Println("Address of a:", &a)
	fmt.Println("Value of b:", *b) // mendapatkan nilai yang ditunjuk oleh pointer b
	fmt.Println("Address of b:", b)
}

/*
PENJELASAN CODE :
a adalah variabel bertipe data int yang menyimpan nilai 42.
b adalah variabel bertipe pointer ke tipe data int. Untuk mendeklarasikan variabel sebagai pointer, kita menggunakan tanda asterisk * di depan tipe data.
&a digunakan untuk mendapatkan alamat memori dari variabel a. Operator & disebut sebagai operator alamat.
Nilai dari variabel b diset ke alamat memori dari variabel a menggunakan operator alamat &. Artinya, b sekarang menunjuk ke lokasi memori di mana nilai a disimpan.
Untuk mendapatkan nilai yang disimpan di lokasi yang ditunjuk oleh pointer b, kita menggunakan operator dereference *. Dalam hal ini, *b menghasilkan nilai yang sama dengan a, yaitu 42.
Alamat memori dari b sendiri dapat dicetak menggunakan b.


OUTPUT
alue of a: 42
Address of a: 0xc00000a0a8
Value of b: 42
Address of b: 0xc00000a0a8

*/
