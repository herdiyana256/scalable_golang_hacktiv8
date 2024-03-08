package main

import "fmt"

func main() {

	var isPalindrome = func(str string) bool {
		var temp string = ""

		for i := len(str) - 1; i >= 0; i-- {
			temp += string(byte(str[i]))
		}

		return temp == str
	}("katak") // true

	fmt.Println(isPalindrome)
}

/*

Closure ( IIFE )


IIFE atau singkatan dari immediately-invoked function expression merupakan sebuah closure yang dapat langsung tereksekusi ketika pertama kali dideklarasikan. Contohnya seperti pada gambar pertama di sebelah kanan.Variable isPalindrome merupakan sebuah penampung dari closure IIFE yang digunakan untuk mengetahui apakah parameter yang diberikan merupakan sebuah kalimat palindrome atau bukan.Jika kita ingin membuat suatu closure menjadi IIFE, maka kita perlu langsung memanggil closure tersebut secara langsung pada saat dideklarasikan. Perlu diingat bahwa kita tidak perlu lagi memanggil closure IIFE dengan tanda kurung () karena closure IIFE tereksekusi pada saat dideklarasikan.Jika kita jalankan pada terminal kita, maka hasilnya akan seperti pada gambar kedua di sebelah kanan.


PENJELASAN CODE :
Dalam kode yang Anda berikan, Anda menggunakan string "katak" sebagai argumen untuk fungsi isPalindrome. String "katak" adalah sebuah palindrom, artinya jika dibalik urutannya, hasilnya tetap sama.

Namun, dalam proses pembalikan string di dalam fungsi isPalindrome, Anda menggunakan string(byte(str[i])) untuk mendapatkan karakter pada indeks ke-i dari string str dan kemudian mengonversinya menjadi string sebelum menambahkannya ke variabel temp.

Perubahan ini sebenarnya tidak mempengaruhi hasil akhir, tetapi dapat memperjelas cara kerja kode. Dengan menggunakan string(byte(str[i])), Anda secara eksplisit mengonversi byte dari karakter ke dalam string. Namun, karena karakter dalam string di Go diwakili oleh rune, Anda sebenarnya bisa langsung menambahkan karakter ke temp tanpa perlu konversi.

Berikut adalah kode yang memperjelas bagaimana Anda bisa langsung menambahkan karakter ke temp tanpa konversi:

*/
