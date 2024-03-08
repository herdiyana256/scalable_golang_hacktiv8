package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str1 := "makan"

	str2 := "minum"

	// Menggunakan fungsi utf8.RuneCountInString() untuk menghitung jumlah karakter dalam setiap string
	// Ini penting karena dalam Go, sebuah string dapat berisi karakter-karakter UTF-8, yang bisa lebih dari 1 byte per karakter.
	// utf8.RuneCountInString() menghitung jumlah rune (karakter) dalam string, bukan jumlah byte.
	// Ini memperlakukan string sebagai rangkaian kode Rune (karakter) Unicode, bukan sebagai rangkaian byte.

	fmt.Printf("str1 character length => %d\n", utf8.RuneCountInString(str1))
	fmt.Printf("str2 character length => %d\n", utf8.RuneCountInString(str2))

}

/*


str1 character length => 5
str2 character length => 5

PENJELASAN CODE :
Kode yang Anda berikan adalah program sederhana dalam bahasa pemrograman Go (Golang). Ini menghitung jumlah karakter dalam dua string menggunakan fungsi utf8.RuneCountInString() yang disediakan oleh paket unicode/utf8. Berikut penjelasan lebih detail:

Hasilnya menunjukkan bahwa kedua string str1 dan str2 sama-sama memiliki panjang karakter 5. Karena masing-masing string hanya terdiri dari karakter-karakter ASCII yang merupakan subkumpulan dari UTF-8, sehingga setiap karakter dianggap sebagai satu Rune dalam perhitungan.






*/
