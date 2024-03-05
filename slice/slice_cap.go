package main

import (
	"fmt"
	"strings"
)

func main() {
	var caption1 = []string{"Herdiyan", "Adam", "Putra", "CEO"}

	fmt.Println("caption1 cap:", cap(caption1)) // 4
	fmt.Println("caption1 len:", len(caption1)) // 4

	fmt.Println(strings.Repeat("#", 20))

	var caption2 = caption1[0:3]

	fmt.Println("caption2 cap:", cap(caption2)) // 4
	fmt.Println("caption2 len:", len(caption2)) // 3

	fmt.Println(strings.Repeat("#", 20))

	var caption3 = caption1[1:]

	fmt.Println("caption3 cap:", cap(caption3)) // 3
	fmt.Println("caption3 len:", len(caption3)) // 3
}

/*

caption1 cap: 4
caption1 len: 4
####################
caption2 cap: 4
caption2 len: 3
####################
caption3 cap: 3
caption3 len: 3



PENJELASAN :
Slice (Cap function)Fungsi cap dapat kita gunakan untuk mengetahui kapasitasdari sebuah array maupun slice.Ketika pertama kali kita membuat suatu slice, panjang dankapasitasnya dipastikan sama, namun dapat berubah seiringdengan slicing yang kita lakukan.Contohnya seperti pada gambar disebelah kanan. Variablefruits1 telah disiapkan di awal dengan panjang 4 dankapasitas 4.Lalu terdapat 2 variable baru bernama fruits2 dan fruits3 yangmelakukan slicing terhadap fruits1 untuk mendapatkan 3element dari fruits1.Lalu kenapa fruits2 memiliki panjang dan kapasitas yangberbeda sedangkan fruits3 panjang dan kapasitasnya samatetapi kapasitasnya berkurang? Mari kita bedah di halamanselanjutnya.


*/
