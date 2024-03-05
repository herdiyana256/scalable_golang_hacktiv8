package main

import (
	"fmt"
)

func main() {
	balances := [2][3]int{{5, 6, 7}, {8, 9, 10}}

	for _, arr := range balances {
		for _, value := range arr {
			fmt.Printf("%d", value)
		}

		fmt.Println()
	}
}

/*

567
8910




PENJELASAN :
ARRAY (MUTIDIMENSIONAL ARRAY)


Kita juga dapat membuat suatu array multidimensional yangberarti terdapat sebuah array di dalam sebuah array. Contohnyaseperti pada gambar pertama di sebelah kanan.Array yang dikandung oleh variable balances merupakan sebuahmultidimensional array.Tanda [2][3]int  memiliki arti bahwa array terluar nya memilikipanjang sama dengan 2, dan array yang berada di dalamnyamemiliki panjang sama dengan 3 dengan tipe data int. {5,6,7}dan {8,9,10} merupakan array yang berada di dalam danmasing-masing memiliki panjang sama dengan 3 dengan tipedata int. Jika kita jalankan pada terminal kita, maka hasilnya akanseperti pada gambar kedua di sebelah kanan.

*/
