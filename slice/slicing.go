package main

import "fmt"

func main() {
	var fruits1 = []string{"Herdiyan", "Adam", "Putra", "Pasti", "Sukses"}

	var fruits2 = fruits1[1:4]
	fmt.Printf("%#v\n", fruits2)

	var fruits3 = fruits1[0:]
	fmt.Printf("%#v\n", fruits3)

	var fruits4 = fruits1[:3]
	fmt.Printf("%v\n", fruits4)

	var fruits5 = fruits1[:] // sama dengan fruits1[:len(fruits1)]
	fmt.Printf("%#v\n", fruits5)
}

/*
[]string{"Adam", "Putra", "Pasti"}
[]string{"Herdiyan", "Adam", "Putra", "Pasti", "Sukses"}
[Herdiyan Adam Putra]
[]string{"Herdiyan", "Adam", "Putra", "Pasti", "Sukses"}


PENJELASAN II :
-Variable fruits2 ingin mendapatkan element dari fruits1 dari index kesatuhingga index ke tiga, maka dari itu cara penulisannya adalahfruits[1:4].-Variable fruits3 ingin mendapatkan element dari fruits1 dari index kenol hingga index terakhirnya, maka dari itu cara penulisannya adalahfruits[0:] yang dimana keterangan stop boleh dihilangkan jika inginmendapatkan hingga index terakhir.- Variable fruits4 ingin mendapatkan element dari fruits1 dari index kenol hingga index kedua,maka dari itu cara penulisannya adalahfruits[:3] yang dimana keterangan start nya boleh dihilangkan ketikaingin mendapatkan element dari index ke nol.


*/

/*
Slice - Sesi 2 Slice (Slicing)Ada cara lain lagi agar kita dapat mendapatkan element-element darisebuah slice dan kita juga bisa menentukan element dari index keberapa yang ingin kita dapatkan. Caranya adalah denganmenggunakan slicing.Contohnya seperti pada gambar disebelah kanan. Cara penulisandalam melakukan slicing adalah sama dengan [start:stop]. Start samadengan awal index yang ingin kita akses dan stop berarti indexakhirnya. Perhatikan hasil dari syntax pada gambar pertama di gambarkedua

*/
