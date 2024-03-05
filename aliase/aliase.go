package main

import (
	"fmt"
)

// Membuat alias untuk tipe data integer
type MyInt int

func main() {
	var num MyInt = 10
	fmt.Println(num) // 10
}

// Penjelasan "Di sini, MyInt adalah alias untuk tipe data int. Anda dapat membuat instance dari MyInt dan menggunakan secara identik dengan tipe data int.

"