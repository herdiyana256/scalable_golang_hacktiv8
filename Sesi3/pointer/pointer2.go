package main

import "fmt"

func main() {
	var firstNumber int = 4              // Mendeklarasikan variabel firstNumber dan menginisialisasinya dengan nilai 4
	var secondNumber *int = &firstNumber // Mendeklarasikan variabel secondNumber sebagai pointer ke int dan mengisinya dengan alamat memori dari firstNumber

	fmt.Println("firstNumber (value) :", firstNumber)           // Mencetak nilai dari variabel firstNumber
	fmt.Println("firstNumber (memory address) :", &firstNumber) // Mencetak alamat memori dari variabel firstNumber

	fmt.Println("secondNumber (value) :", secondNumber)          // Mencetak nilai dari variabel secondNumber (yang merupakan alamat memori dari firstNumber)
	fmt.Println("secondNumber (memory address) :", secondNumber) // Mencetak alamat memori yang disimpan di dalam variabel secondNumber
}

/*
Value of a: 42
Address of a: 0xc00000a0a8
Value of b: 42
Address of b: 0xc00000a0a8


PENJELASAN CODE :
firstNumber adalah variabel bertipe int yang memiliki nilai 4.
secondNumber adalah variabel bertipe pointer ke int. Kita memberikan alamat memori dari firstNumber ke secondNumber menggunakan operator alamat &.
Saat mencetak secondNumber, yang dicetak adalah alamat memori dari firstNumber, karena secondNumber menyimpan alamat memori dari firstNumber.
Saat mencetak *secondNumber, yang dicetak adalah nilai yang ditunjuk oleh pointer secondNumber, yaitu nilai dari firstNumber, dalam hal ini adalah 4.

*/

/*
POINTER
Pointer (Memory address)Perhatikan pada gambar pertama disebelah kanan. Terdapat 2 variable bernama firstNumber dan secondNumber. Variable firstNumber merupakan sebuah variable dengan tipe data int.Kemudian secondNumber merupakan sebuah variable pointer int yang mengandung alamat memori dari firstNumber. Perlu diperhatikan disini bahwa ketika kita ingin memberikan sebuah nilai kepada variable pointer, maka kita perlu menggunakan tanda ampersand & seperti yang kita lihat pada line 7. Kemudian pada line 10, kita mencoba untuk menampilkan alamat memori dari variable firstNumber dengan cara menggunakan tanda ampersand &.
Kemudian pada line 12, kita mencoba untuk menampilkan nilai asli (bukan alamat memori) yang dikandung oleh secondNumberdengan cara memberikan tanda asterisk *. Perlu diingat disini bahwa karena variable secondNumber merupakan sebuah variable pointer yang mengandung alamat memori dari firstNumber, maka ketika kita ingin mendapatkan nilai asli yang dikandung oleh sebuah variable pointer, kita perlu menggunakan tanda asterisk *. Jika kita jalankan pada terminal kita, maka hasilnya akan seperti pada gambar kedua.

Jika kita melihat hasilnya pada gambar kedua di sebelah kanan. Kita dapat menarik kesimpulan bahwa:●Kita dapat mendapatkan alamat memori dari variable biasa dengan menggunakan tanda ampersand &.●Kita juga dapat mendapatkan nilai asli dari variable pointer dengan cara menggunakan tanda asterisk *.


*/
