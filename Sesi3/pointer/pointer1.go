/*
PointerPointer merupakan sebuah variable spesial yang digunakan untuk menyimpan alamat memori variable lainnya. Sebagaicontoh sebuah variabel bertipe integer memiliki nilai 4, maka yang dimaksud pointer adalah alamat memori dimana nilai 4disimpan, bukan nilai 4 itu sendiri.Variabel-variabel yang memiliki reference atau alamat memori yang sama, saling berhubungan satu sama lain dan nilainyapasti sama. Ketika ada perubahan nilai, maka akan memberikan efek kepada variabel lain (yang alamat memorinya sama) yaitunilainya juga ikut berubah.Cara membuat suatu variable menjadi sebuah pointer cukup mudah. Caranya adalah dengan menggunakan tanda asterisk *sebelum penulisan tipe datanya. Contohnya seperti 2 variable pada gambar dibawah ini.

*/

package main () {
	func main () {
		var firstNumber  * int
		var secondNumber * int

		_, _ = firstNumber, secondNumber
	}
}