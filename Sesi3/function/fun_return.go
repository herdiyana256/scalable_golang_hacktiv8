// FUNCTION (RETURN)
// Di Go (Golang), fungsi dapat mengembalikan satu atau beberapa nilai sebagai hasil dari pemrosesan yang dilakukan oleh fungsi tersebut. Penggunaan kata kunci return digunakan untuk mengembalikan nilai-nilai tersebut.

package main

import "fmt"

// Fungsi add akan mengembalikan hasil penjumlahan dari dua bilangan bulat yang diberikan
func add(a, b int) int {
	return a + b
}

func main() {
	// Memanggil fungsi add dengan nilai 5 dan 3 sebagai argumen, dan menyimpan hasilnya dalam variabel sum
	sum := add(5, 3)
	fmt.Println("5 + 3 =", sum) // Output: 5 + 3 = 8
}

// 5 + 3 = 8

/*
Pada contoh di atas, fungsi add mengembalikan hasil penjumlahan dari dua parameter a dan b yang bertipe integer. Kata kunci return digunakan di dalam fungsi untuk mengembalikan nilai dari operasi penjumlahan tersebut.

Ada beberapa hal yang perlu diperhatikan terkait penggunaan return dalam Go:

Fungsi yang dideklarasikan dengan tipe pengembalian akan selalu mengembalikan nilai tersebut.
Jika tipe pengembalian ditentukan dalam deklarasi fungsi, Anda tidak dapat mengembalikan nilai dengan tipe data yang berbeda.
Anda dapat mengembalikan beberapa nilai dalam satu pernyataan return. Misalnya, return a, b, c.

*/
