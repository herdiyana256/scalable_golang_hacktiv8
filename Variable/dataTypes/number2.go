package main

import "fmt"

func main() {
	var decimalNumber float32 = 3.63

	fmt.Printf("decimal Number : %f\n", decimalNumber)
	fmt.Printf("decimal Number: %.3f\n", decimalNumber)
}

/*
decimal Number : 3.630000
decimal Number: 3.630
*/

/*
PENJELASAN :

Number (float)

Jika kita lihat pada gambar kedua, format angka desimal yang di cetak di terminal menghasilkan format angka yang berbeda walaupun masih berasal dari satu variable yang sama .const
Pada dasarya verb %f digunakan untuk meformat angka desimal atau tipe data float menjadi 6 digit angka desimal.const

Jika kita ingin mengecilkan digit desimalnya, maka kita dapat menggunakan format verb seperti (%.nf), yang dimana huruf n dapat kita ganti dengan jumlah digit yang ingin kita hasilnya.const
sama seperti contoh diatas.
*/
