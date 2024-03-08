/*
Closure (Callback)Callback adalah sebuah closure yang dijadikan sebagai parameter pada sebuah function. Contohnya seperti pada gambar pertama di sebelah kanan.Function findOddNumbers merupakan sebuah function menerima 2 parameter. Parameter pertama memiliki tipe data slice of int dan parameter kedua merupakan sebuah callback yang menerima satu parameter dengan tipe data int dan me-return nilai bool.Function findOddNumbers ini digunakan untuk mencari jumlah angka ganjil yang diberikan pada parameter pertamanya. Lalu pengecekan nya tersebut dilakukan oleh callback nya.Jika callback me-return nilai true, maka variable totalOddNumbers pada line 104 akan ditambahkan. Dan function findOddNumbers akan me-return variable tersebut.Jika kita jalankan pada terminal kita, maka hasilnya akan seperti pada gambar kedua.


Jika kita perhatikan pada penulisan parameter callback pada halaman sebelumnya. Kita dapat menggunakan type alias untuk mempersingkat penulisan parameter callback nya. Contohnya seperti pada gambar di bawah yang telah di highlight.
*/

package main

import "fmt"

func main() {
	var numbers = []int{2, 5, 8, 10, 3, 99, 23}

	var find = findOddNumbers(numbers, func(number int) bool {
		return number%2 != 0
	})

	fmt.Println("Total odd numbers:", find)
}

func findOddNumbers(numbers []int, callback func(int) bool) int {
	var totalOddNumbers int
	for _, v := range numbers {
		if callback(v) {
			totalOddNumbers++
		}
	}

	return totalOddNumbers // Total odd numbers: 4
}

/*
PENJELASAN CODE :

main(): Di sini, kita membuat sebuah slice numbers yang berisi sejumlah angka. Kita kemudian memanggil fungsi findOddNumbers() dengan memberikan slice numbers dan sebuah fungsi callback sebagai argumen. Fungsi callback ini mengecek apakah suatu angka ganjil atau tidak. Hasil dari findOddNumbers() kemudian dicetak.

findOddNumbers(): Ini adalah fungsi yang mencari angka-angka ganjil dalam slice numbers. Fungsi ini menerima dua parameter: slice numbers dan sebuah fungsi callback. Di dalamnya, kita melakukan iterasi melalui setiap angka dalam slice dan memanggil fungsi callback untuk setiap angka. Jika fungsi callback mengembalikan true, itu berarti angka tersebut ganjil, sehingga kita menambahkan satu ke variabel totalOddNumbers. Akhirnya, kita mengembalikan totalOddNumbers, yang merupakan jumlah angka ganjil dalam slice.

Dalam hal ini, fungsi callback adalah fungsi anonim yang didefinisikan langsung sebagai argumen untuk findOddNumbers(). Fungsi callback ini memeriksa apakah suatu angka adalah ganjil atau tidak dengan melakukan operasi modulus (number % 2 != 0). Jika modulus dari angka tersebut tidak sama dengan nol, itu berarti angka tersebut ganjil.

*/
