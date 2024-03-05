package main

import "fmt"

func main() {

	var fruits1 = []string{"apple", "banana", "mango", "durian", "pineapple"}

	fruits1 = append(fruits1[:3], "buah")

	fmt.Printf("%#v\n", fruits1)
	//[]string{"apple", "banana", "mango", "buah"}

}

/*
Slice - Sesi 2 Slice (Combining slicing and append)Kita juga dapat mengkombinasikan fungsi append dengan slicing.Contohnya seperti pada gambar pertama di sebelah kanan. Jika kitaperhatikan, variable fruits1 ingin me-replace index ketiga hinggaseterusnya dengan hanya buah “rambutan” saja. Jika kita jalankanpada terminal maka hasilnya akan seperti pada gambar kedua disebelah kanan.

*/
