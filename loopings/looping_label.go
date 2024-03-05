package main

import "fmt"

func main() {
outerLoop:
	for i := 0; i < 3; i++ {
		fmt.Println("Perulangan ke - ", i+1)
		for j := 0; j < 3; j++ {
			if i == 2 {
				break outerLoop
			}
			fmt.Print(j, " ")
		}

		fmt.Print("\n")
	}
}

/*
PENJELASAN :
LOOPINGS(LABEL)
PADA LOOPING BERSARANG, KEYWORD BREAK DAN CONTINE AKAN BERLAKU PADA BLOK PERULANGAN DIMANA IA DIGUNAKAN SAJA. ADA CARA AGAR KEDUA KEYWORD INI BISA TERTUJU PADA PERULANGAN TERLUAR ATAU PERLUANGAN TERTENTU,
YAITU DENGAN MEMANFAATKAN TEKNIK PEMBERIAN LABEL.
CONTOHNYA SEPERTI PADA GAMBAR PERTAMA DIBAWAH .


Jika kita perhatikan hasil dari looping pertama pada gambar kedua, seluruh looping pada perulangan ketiga terhenti karenaterdapat sebuah kondisional pada proses looping kedua yang dimana jika variable i sudah memiliki nilai dengan angkasama dengan 2, maka looping pertama atau looping terluar akan dihentikan atau sama saja dengan seluruh proses loopingterhenti.

*/

/*

Perulangan ke -  1
0 1 2
Perulangan ke -  2
0 1 2
Perulangan ke -  3

*/
