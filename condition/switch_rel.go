package main

import "fmt"

func main() {
	var score = 6

	switch {
	case score == 8:
		fmt.Println("perfect")
	case (score < 8) && (score > 3):
		fmt.Println("not bad")
	default:
		{
			fmt.Println("study harder")
			fmt.Println("you need to learn more")
		}

	}
}

// Not bad

/*
	Penjelasan :
	Kita juga dapat menggunakan switch dengan operator perbandingan seperti layaknya pada sebuah kondisional dengan keyword if, else, if dan else.

	Contohnya seperti pada gambar disebelah kanan. Kondisional pada gambar disebelah kanan akan menghasilkan kalimat "not bad" karena variable score memiliki angka 6 dam amgka 6  memang lebih kecil dari angka 8 dan lebih besar daripada angka 3 .

	Lalu kita juga bisa melihat bahwa scope block default pada gambar disebelah kanan menggunakan kurun kurawal.

	hal ini sangat bagus diterapkan jika kita ingin membuat lebih dari satu statement pada scope block default agar synyax kita lebih rapi dan mudah dimaintan.


*/
