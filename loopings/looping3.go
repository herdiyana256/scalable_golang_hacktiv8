package main

import "fmt"

func main() {
	var i = 0

	for {
		fmt.Println("Angka", i)

		i++
		if i == 3 {
			break
		}
	}
}

/*
Angka 0
Angka 1
Angka 2


Penjelasan :

Loopings (third way of looping)Cara ketiga dalam melakukan looping pada bahasa Go adalah dengan melakukan looping tanpa memberikan kondisiapapun. Contohnya seperti pada gambar pertama dibawah. Looping pada gambar dibawah memakai bantuan keywordbreak yang dimana dengan menggunakan keyword ini, maka dapat menghentikan suatu proses looping.



Pada looping tersebut, terdapat suatu kondisi yang dimana jika nilai pada variable i sudah memiliki nilai dengan angkasama dengan 3, maka looping keyword break akan terpanggil dan proses looping akan berhenti. Jika kita jalankan padaterminal, maka hasilnya akan seperti pada gambar kedua dibawah.

*/
