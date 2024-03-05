// PENJELASAN : Looping pada map di Go dilakukan dengan menggunakan perulangan for bersamaan dengan range. range digunakan untuk mengambil pasangan kunci-nilai dari map saat melakukan iterasi.

package main

import "fmt"

func main() {
	var person = map[string]string{
		"name":    "Herdiyan Adam Putra",
		"age":     "28",
		"address": "Jalan Jati Kelapa No 31",
	}

	for key, value := range person {
		fmt.Println(key+": ", value)
	}
}

/*
Ketika kita ingin melakukan loop terhadap sebuah map, maka kitadapat menggunakan range loop. Contohnya seperti pada gambarpertama di sebelah kanan.Kita bisa dengan mudah mendapatkan key dan value nya.Jika kita jalankan pada terminal kita, maka hasilnya akan seperti padagambar kedua.



name:  Herdiyan Adam Putra
age:  28
address:  Jalan Jati Kelapa No 31
*/
