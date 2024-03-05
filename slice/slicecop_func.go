package main

import "fmt"

func main() {

	var fruits1 = []string{"apple", "banana", "mango"}

	var fruits2 = []string{"durian", "pipeapple", "starfruit"}

	nn := copy(fruits1, fruits2)
	// The length of the slice is not changed. It only copies the elements from one slice to another.

	fmt.Println("Fruits1 =>", fruits1)
	fmt.Println("Fruits2 =>", fruits2)
	fmt.Println("Copied elements =>", nn)
}

/*
Slice (copy function)Pada kasus kita kali ini, variable fruits1 ingin meng-copyseluruh element yang ada pada variable fruits2.Argumen pertama yang diterima oleh fungsi copy  adalahdestinasi atau slice yang ingin meng-copy, lalu argumentkedua adalah source/sumber dari slice yang ingin di copy.Fungsi copy juga akan mengembalikan jumlah element yangberhasil ter-copy.Ketika kita jalankan pada terminal kita, maka hasilnya akanseperti pada gambar kedua disebelah kanan. Element padafruits1 sudah ter-replace oleh fruits2, dan terdapat 3 elementyang berhasil ter-copy.


*/
